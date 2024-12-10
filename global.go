package steamcmd

import (
	"context"
	"os"

	"github.com/frantjc/go-steamcmd/internal/cache"
)

func Start(ctx context.Context) (*Prompt, error) {
	c, err := New(ctx)
	if err != nil {
		return nil, err
	}

	return c.Start(ctx)
}

func Clean() error {
	return os.RemoveAll(cache.Dir)
}
