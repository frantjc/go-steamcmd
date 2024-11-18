package steamcmd

import (
	"context"
	"os/exec"
)

type Command string

func (c Command) String() string {
	return string(c)
}

func (c Command) Start(ctx context.Context) (*Prompt, error) {
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

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return &Prompt{
		flags:  &promptFlags{},
		stdin:  stdin,
		stdout: stdout,
	}, readOutput(ctx, stdout, 0)
}
