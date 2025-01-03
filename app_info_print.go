package steamcmd

import (
	"fmt"
)

type AppInfoPrint int

var _ Command = AppInfoPrint(0)

func (c AppInfoPrint) String() string {
	return fmt.Sprintf("%d", c)
}

func (c AppInfoPrint) check(flags *flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("cannot app_info_print before login")
	}

	return nil
}

func (c AppInfoPrint) args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_print requires app ID")
	}

	return []string{"app_info_print", c.String()}, nil
}

func (c AppInfoPrint) modify(_ *flags) error {
	return nil
}
