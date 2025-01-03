package steamcmd

import (
	"fmt"
)

type AppInfoRequest int

var _ Command = AppInfoRequest(0)

func (c AppInfoRequest) String() string {
	return fmt.Sprintf("%d", c)
}

func (c AppInfoRequest) check(flags *flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("cannot app_info_request before login")
	}

	return nil
}

func (c AppInfoRequest) args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_request requires app ID")
	}

	return []string{"app_info_request", c.String()}, nil
}

func (c AppInfoRequest) modify(_ *flags) error {
	return nil
}
