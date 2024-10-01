package steamcmd

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
)

type Command string

func (c Command) String() string {
	return string(c)
}

func (c Command) Start(ctx context.Context) (Prompt, error) {
	//nolint:gosec
	cmd := exec.CommandContext(ctx, c.String())

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	p := &prompt{&promptFlags{}, stdin, stdout, sync.Mutex{}, nil}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go func() {
		p.err = cmd.Wait()
		if p.err == nil {
			p.err = fmt.Errorf("steamcmd exited")
		}
	}()

	return p, p.Run(ctx)
}

func (c Command) Run(ctx context.Context, cmds ...Cmd) error {
	args := []string{}

	for _, cmd := range cmds {
		if a, err := cmd.args(); err != nil {
			return err
		} else if len(a) > 0 {
			a[0] = "+" + a[0]
			args = append(args, a...)
		}
	}

	//nolint:gosec
	return exec.CommandContext(ctx, c.String(), args...).Run()
}
