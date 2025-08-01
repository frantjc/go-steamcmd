package steamcmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	vdf "github.com/frantjc/go-encoding-vdf"
	xerrors "github.com/frantjc/x/errors"
	xslices "github.com/frantjc/x/slices"
)

type Path string

func (c Path) String() string {
	return string(c)
}

type Flags struct {
	LoggedIn bool
}

type Command interface {
	Check(*Flags) error
	Args() ([]string, error)
	Modify(*Flags) error
}

type Prompt struct {
	flags  *Flags
	stdin  io.Writer
	stdout io.Reader
	cmd    *exec.Cmd
	mu     sync.Mutex
}

var (
	promptBytes                  = []byte("Steam>")
	errBytes                     = []byte("ERROR! ")
	failedBytes                  = []byte("FAILED ")
	requestingAppInfoPrefixBytes = []byte("No app info for AppID ")
	requestingAppInfoSuffixBytes = []byte(" found, requesting...")
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
				} else if p.stdin != nil && bytes.Contains(q, requestingAppInfoPrefixBytes) && bytes.Contains(q, requestingAppInfoSuffixBytes) {
					// Handle a special case where steamcmd requests the app info because
					// it does not have it and then hangs without more input.
					_, appIDBytes, _ := bytes.Cut(q, requestingAppInfoPrefixBytes)
					appIDBytes, _, _ = bytes.Cut(appIDBytes, requestingAppInfoSuffixBytes)
					appID, err := strconv.Atoi(string(bytes.TrimSpace(appIDBytes)))
					if err != nil {
						return &CommandError{
							Msg:    "parsing appID to rerequest app info",
							Err:    err,
							Output: q,
						}
					}

					if _, err := fmt.Fprintln(p.stdin, "app_info_print", fmt.Sprint(appID)); err != nil {
						return &CommandError{
							Msg:    "rerequesting app info",
							Err:    err,
							Output: q,
						}
					}
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

					// If this is non-nil, then we are running as a prompt
					// and need to exit so that the next command can get ran.
					// If it is nil, then we are running as a script and should
					// keep reading the output as the next command is already
					// queued.
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
		if err := command.Check(p.flags); err != nil {
			return err
		}

		args, err := command.Args()
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintln(p.stdin, xslices.Map(args, func(arg string, _ int) any {
			return arg
		})...); err != nil {
			return err
		}

		if err = p.readOutput(ctx); err != nil {
			return err
		}

		if err := command.Modify(p.flags); err != nil {
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

func Args(flags *Flags, commands ...Command) ([]string, error) {
	arg := []string{}

	if flags == nil {
		flags = &Flags{}
	}

	for _, command := range commands {
		if err := command.Check(flags); err != nil {
			return nil, err
		}

		args, err := command.Args()
		if err != nil {
			return nil, err
		}

		if len(args) > 0 {
			args[0] = fmt.Sprintf("+%s", args[0])
		}

		arg = append(arg, args...)

		if err := command.Modify(flags); err != nil {
			return nil, err
		}
	}

	return arg, nil
}

func (c Path) Start(ctx context.Context, commands ...Command) (*Prompt, error) {
	flags := &Flags{}

	arg, err := Args(flags, commands...)
	if err != nil {
		return nil, err
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
	flags := &Flags{}

	arg, err := Args(flags, append(commands, Quit)...)
	if err != nil {
		return err
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
