package steamcmd

import (
	"context"
	"os"

	"github.com/frantjc/go-steamcmd/internal/cache"
)

func Start(ctx context.Context, cmds ...Command) (*Prompt, error) {
	c, err := New(ctx)
	if err != nil {
		return nil, err
	}

	return c.Start(ctx, cmds...)
}

func Run(ctx context.Context, cmds ...Command) error {
	c, err := New(ctx)
	if err != nil {
		return err
	}

	return c.Run(ctx, cmds...)
}

func Clean() error {
	return os.RemoveAll(cache.Dir)
}
