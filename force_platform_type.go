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

func (c forcePlatformType) Check(_ *promptFlags) error {
	return nil
}

func (c forcePlatformType) Args() ([]string, error) {
	if c == "" {
		return nil, fmt.Errorf("empty PlatformType")
	}

	return []string{"@sSteamCmdForcePlatformType", c.String()}, nil
}

func (c forcePlatformType) ReadOutput(ctx context.Context, r io.Reader) error {
	return readOutput(ctx, r, 0)
}

func (c forcePlatformType) Modify(_ *promptFlags) error {
	return nil
}
