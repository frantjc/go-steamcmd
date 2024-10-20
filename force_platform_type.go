package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type forcePlatformType PlatformType

var _ cmd = forcePlatformType("")

func (c forcePlatformType) String() string {
	return string(c)
}

func (c forcePlatformType) check(_ *promptFlags) error {
	return nil
}

func (c forcePlatformType) args() ([]string, error) {
	if c == "" {
		return nil, fmt.Errorf("empty PlatformType")
	}

	return []string{"@sSteamCmdForcePlatformType", c.String()}, nil
}

func (c forcePlatformType) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c forcePlatformType) modify(_ *promptFlags) error {
	return nil
}
