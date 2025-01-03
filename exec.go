package steamcmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"

	vdf "github.com/frantjc/go-encoding-vdf"
	xerrors "github.com/frantjc/x/errors"
	"github.com/frantjc/x/slice"
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
	stdin  io.WriteCloser
	stdout io.ReadCloser
	err    error
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

		errC <- xerrors.Ignore(func() error {
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
				}
			}
		}(), io.EOF)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errC:
		return err
	}
}

func (p *Prompt) Run(ctx context.Context, commands ...Command) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, command := range commands {
		if p.err != nil {
			return p.err
		}

		if err := command.Check(p.flags); err != nil {
			return errors.Join(err, p.err)
		}

		args, err := command.Args()
		if err != nil {
			return errors.Join(err, p.err)
		}

		if _, err := fmt.Fprintln(p.stdin, xslice.Map(args, func(arg string, _ int) any {
			return arg
		})...); err != nil {
			return errors.Join(err, p.err)
		}

		if err = p.readOutput(ctx); err != nil {
			return errors.Join(err, p.err)
		}

		if err := command.Modify(p.flags); err != nil {
			return errors.Join(err, p.err)
		}
	}

	return p.err
}

func (p *Prompt) Close(ctx context.Context) error {
	return errors.Join(
		p.Run(ctx, Quit),
		p.stdin.Close(),
		p.stdout.Close(),
		p.err,
	)
}

func (c Path) Run(ctx context.Context, commands ...Command) (*Prompt, error) {
	var (
		arg   = []string{}
		flags = &Flags{}
	)

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
		flags:  flags,
		stdin:  stdin,
		stdout: stdout,
	}

	go func() {
		if err = cmd.Run(); err != nil {
			p.err = errors.Join(p.err, err)
		}
	}()

	if err = p.readOutput(ctx); err != nil {
		return nil, errors.Join(err, p.err)
	}

	return p, p.err
}
