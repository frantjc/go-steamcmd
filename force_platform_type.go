package steamcmd

import (
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

type ForcePlatformType PlatformType

var _ Command = ForcePlatformType("")

func (c ForcePlatformType) String() string {
	return string(c)
}

func (c ForcePlatformType) Check(_ *Flags) error {
	return nil
}

func (c ForcePlatformType) Args() ([]string, error) {
	if c == "" {
		return nil, fmt.Errorf("empty PlatformType")
	}

	return []string{"@sSteamCmdForcePlatformType", c.String()}, nil
}

func (c ForcePlatformType) Modify(_ *Flags) error {
	return nil
}
