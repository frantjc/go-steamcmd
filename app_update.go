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

func (AppUpdate) Check(flags *Flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("cannot app_update before login")
	}

	return nil
}

const (
	DefaultBranch = "public"
)

func (c AppUpdate) Args() ([]string, error) {
	if c.AppID == 0 {
		return nil, fmt.Errorf("app_update requires app ID")
	}

	args := []string{"app_update", fmt.Sprint(c.AppID)}

	if c.Beta != "" && c.Beta != DefaultBranch {
		if c.BetaPassword == "" {
			return nil, fmt.Errorf("app_update -betapassword is required with -beta")
		}

		args = append(args, "-beta", c.Beta, "-betapassword", c.BetaPassword)
	} else if c.BetaPassword != "" {
		return nil, fmt.Errorf("app_update -betapassword is prohibited without -beta")
	}

	if c.Validate {
		args = append(args, "validate")
	}

	return args, nil
}

func (c AppUpdate) Modify(_ *Flags) error {
	return nil
}
