package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type ForcePlatformType PlatformType

var _ Cmd = ForcePlatformType("")

func (c ForcePlatformType) String() string {
	return string(c)
}

func (c ForcePlatformType) check(_ *promptFlags) error {
	return nil
}

func (c ForcePlatformType) args() ([]string, error) {
	if c == "" {
		return nil, fmt.Errorf("empty PlatformType")
	}

	return []string{"@sSteamCmdForcePlatformType", c.String()}, nil
}

func (c ForcePlatformType) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c ForcePlatformType) modify(_ *promptFlags) error {
	return nil
}
