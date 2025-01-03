package steamcmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	vdf "github.com/frantjc/go-encoding-vdf"
	xerrors "github.com/frantjc/x/errors"
	xslice "github.com/frantjc/x/slice"
)

type Path string

func (c Path) String() string {
	return string(c)
}

type flags struct {
	LoggedIn bool
}

type Command interface {
	check(*flags) error
	args() ([]string, error)
	modify(*flags) error
}

type Prompt struct {
	flags  *flags
	stdin  io.Writer
	stdout io.Reader
	cmd    *exec.Cmd
	mu     sync.Mutex
}

var (
	promptBytes = []byte("Steam>")
	errBytes    = []byte("ERROR! ")
	failedBytes = []byte("FAILED ")
)

func (p *Prompt) readOutput(ctx context.Context) error {
	errC := make(chan error, 1)

	go func() {
		defer close(errC)

		errC <- func() error {
			buf := new(bytes.Buffer)

			for {
				var b [512]byte

				n, err := p.stdout.Read(b[:])
				if err != nil {
					return err
				}

				if _, err := buf.Write(b[:n]); err != nil {
					return err
				}

				q := buf.Bytes()
				if _, msgB, found := bytes.Cut(q, errBytes); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					return &CommandError{
						Msg:    strings.ToLower(string(msgB)),
						Output: q,
					}
				} else if _, msgB, found := bytes.Cut(q, failedBytes); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					return &CommandError{
						Msg:    strings.ToLower(string(msgB)),
						Output: q,
					}
				} else if bytes.Contains(q, promptBytes) {
					return nil
				} else if i := bytes.Index(q, []byte("{")); i >= 0 {
					appInfo := &AppInfo{}

					// On Linux, steamcmd needs a kick to make it
					// finish printing the output of app_info_print
					// if we're running as a prompt.
					if runtime.GOOS == "linux" && p.stdin != nil {
						cctx, cancel := context.WithCancel(ctx)
						defer cancel()

						go func() {
							for {
								select {
								case <-cctx.Done():
									return
								case <-time.NewTimer(time.Second).C:
									_, _ = fmt.Fprintln(p.stdin)
								}
							}
						}()
					}

					if err := vdf.NewDecoder(
						io.MultiReader(
							bytes.NewReader(q[i:]),
							p.stdout,
						),
					).Decode(appInfo); err != nil {
						return &CommandError{
							Msg:    "decoding vdf",
							Err:    err,
							Output: q,
						}
					}

					appInfos[appInfo.Common.GameID] = *appInfo

					// If this is nil, then we're running as a prompt and need to exit so that
					// the next command can get ran. If it's not nil, then we're running as a
					// script and should keep reading the output as the next command is already
					// queued up.
					if p.stdin != nil {
						return nil
					}
				}
			}
		}()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errC:
		return err
	}
}

func (p *Prompt) Run(ctx context.Context, commands ...Command) error {
	if p.stdin == nil {
		return fmt.Errorf("prompt is closed")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	for _, command := range commands {
		if err := command.check(p.flags); err != nil {
			return err
		}

		args, err := command.args()
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintln(p.stdin, xslice.Map(args, func(arg string, _ int) any {
			return arg
		})...); err != nil {
			return err
		}

		if err = p.readOutput(ctx); err != nil {
			return err
		}

		if err := command.modify(p.flags); err != nil {
			return err
		}
	}

	return nil
}

func (p *Prompt) Close() error {
	if _, err := fmt.Fprintln(p.stdin, "quit"); err != nil {
		return err
	}

	if err := p.cmd.Wait(); err != nil {
		return err
	}

	p.cmd = nil
	p.flags = nil
	p.stdin = nil
	p.stdout = nil

	return nil
}

func (c Path) Start(ctx context.Context, commands ...Command) (*Prompt, error) {
	var (
		arg   = []string{}
		flags = &flags{}
	)

	for _, command := range commands {
		if err := command.check(flags); err != nil {
			return nil, err
		}

		args, err := command.args()
		if err != nil {
			return nil, err
		}

		if len(args) > 0 {
			args[0] = fmt.Sprintf("+%s", args[0])
		}

		arg = append(arg, args...)

		if err := command.modify(flags); err != nil {
			return nil, err
		}
	}

	//nolint:gosec
	cmd := exec.CommandContext(ctx, c.String(), arg...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	p := &Prompt{
		stdin:  stdin,
		stdout: stdout,
		flags:  flags,
		cmd:    cmd,
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if err := p.readOutput(ctx); err != nil {
		return nil, err
	}

	return p, nil
}

func (c Path) Run(ctx context.Context, commands ...Command) error {
	var (
		arg   = []string{}
		flags = &flags{}
	)

	for _, command := range append(commands, quit) {
		if err := command.check(flags); err != nil {
			return err
		}

		args, err := command.args()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			args[0] = fmt.Sprintf("+%s", args[0])
		}

		arg = append(arg, args...)

		if err := command.modify(flags); err != nil {
			return err
		}
	}

	var (
		//nolint:gosec
		cmd    = exec.CommandContext(ctx, c.String(), arg...)
		stdout = new(bytes.Buffer)
	)
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	if err := (&Prompt{stdout: stdout}).readOutput(ctx); xerrors.Ignore(err, io.EOF) != nil {
		return err
	}

	return nil
}
