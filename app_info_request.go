package steamcmd

import (
	"fmt"
)

type AppInfoRequest int

var _ Command = AppInfoRequest(0)

func (c AppInfoRequest) String() string {
	return fmt.Sprintf("%d", c)
}

func (c AppInfoRequest) Check(flags *Flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("cannot app_info_request before login")
	}

	return nil
}

func (c AppInfoRequest) Args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_request requires app ID")
	}

	return []string{"app_info_request", c.String()}, nil
}

func (c AppInfoRequest) Modify(_ *Flags) error {
	return nil
}
