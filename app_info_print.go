package steamcmd

import (
	"fmt"
)

type AppInfoPrint int

var _ Command = AppInfoPrint(0)

func (c AppInfoPrint) String() string {
	return fmt.Sprintf("%d", c)
}

func (c AppInfoPrint) Check(flags *Flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("cannot app_info_print before login")
	}

	return nil
}

func (c AppInfoPrint) Args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_print requires app ID")
	}

	return []string{"app_info_print", c.String()}, nil
}

func (c AppInfoPrint) Modify(_ *Flags) error {
	return nil
}
