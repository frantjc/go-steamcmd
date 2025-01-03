package steamcmd

import (
	"context"
	"os"

	"github.com/frantjc/go-steamcmd/internal/cache"
)

func NewPrompt(ctx context.Context, cmds ...Command) (*Prompt, error) {
	c, err := New(ctx)
	if err != nil {
		return nil, err
	}

	return c.Run(ctx, cmds...)
}

func Clean() error {
	return os.RemoveAll(cache.Dir)
}
