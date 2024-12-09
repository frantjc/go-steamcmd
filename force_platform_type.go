package steamcmd

import (
	"context"
	"fmt"
)

type PlatformType string

func (t PlatformType) String() string {
	return string(t)
}

var (
	PlatformTypeWindows PlatformType = "windows"
	PlatformTypeLinux   PlatformType = "linux"
	PlatformTypeMacOS   PlatformType = "macos"
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

func (c forcePlatformType) readOutput(ctx context.Context, p *Prompt) error {
	return readOutput(ctx, p, 0)
}

func (c forcePlatformType) modify(_ *promptFlags) error {
	return nil
}
