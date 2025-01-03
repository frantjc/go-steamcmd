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

func (*AppUpdate) Check(flags *Flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("cannot app_update before login")
	}

	return nil
}

func (c *AppUpdate) Args() ([]string, error) {
	if c == nil || c.AppID == 0 {
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

func (c *AppUpdate) Modify(_ *Flags) error {
	return nil
}
