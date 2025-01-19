package steamcmd

import (
	"fmt"
)

type AppUpdate struct {
	AppID        int
	Beta         string
	BetaPassword string
	Validate     bool
}

var _ Command = new(AppUpdate)

func (AppUpdate) check(flags *flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("cannot app_update before login")
	}

	return nil
}

func (c AppUpdate) args() ([]string, error) {
	if c.AppID == 0 {
		return nil, fmt.Errorf("app_update requires app ID")
	}

	args := []string{"app_update", fmt.Sprint(c.AppID)}

	if c.Beta != "" {
		args = append(args, "-beta", c.Beta)
	}

	if c.BetaPassword != "" {
		args = append(args, "-betapassword", c.BetaPassword)
	}

	if c.Validate {
		args = append(args, "validate")
	}

	return args, nil
}

func (c AppUpdate) modify(_ *flags) error {
	return nil
}
