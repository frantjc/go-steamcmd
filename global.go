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
