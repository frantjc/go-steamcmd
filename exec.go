package steamcmd

import (
	"context"
	"io"
	"os/exec"
)

type Command string

func (c Command) String() string {
	return string(c)
}

func (c Command) Start(ctx context.Context) (*Prompt, error) {
	var (
		//nolint:gosec
		cmd = exec.CommandContext(ctx, c.String())
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	p := &Prompt{
		flags:  &promptFlags{},
		stdin:  stdin,
		stdout: stdout,
	}

	go func() {
		defer stderr.Close()

		if err = func() error {
			if _, err = io.Copy(io.Discard, stderr); err != nil {
				return err
			}

			return stderr.Close()
		}(); err != nil {
			p.err = err
		}
	}()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go func() {
		p.err = cmd.Wait()
	}()

	if err = readOutput(ctx, p, 0); err != nil {
		return nil, err
	}

	return p, p.err
}
