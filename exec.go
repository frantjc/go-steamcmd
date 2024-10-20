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

	return p, p.run(ctx, base)
}
