package steamcmd

import (
	"context"
	"fmt"
)

type appUpdate struct {
	AppID        int
	Beta         string
	BetaPassword string
	Validate     bool
}

var _ cmd = new(appUpdate)

func (*appUpdate) check(flags *promptFlags) error {
	if !flags.loggedIn {
		return fmt.Errorf("cannot app_update before login")
	}

	return nil
}

func (c *appUpdate) args() ([]string, error) {
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

func (c *appUpdate) readOutput(ctx context.Context, p *Prompt) error {
	return readOutput(ctx, p, 0)
}

func (c *appUpdate) modify(_ *promptFlags) error {
	return nil
}
