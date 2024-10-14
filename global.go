package steamcmd

import (
	"context"
)

func Start(ctx context.Context) (Prompt, error) {
	c, err := New(ctx)
	if err != nil {
		return nil, err
	}

	return c.Start(ctx)
}

func Run(ctx context.Context, cmds ...Cmd) error {
	c, err := New(ctx)
	if err != nil {
		return err
	}

	return c.Run(ctx, cmds...)
}
