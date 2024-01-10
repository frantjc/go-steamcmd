package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type AppUpdate struct {
	AppID        string
	Beta         string
	BetaPassword string
	Validate     bool
}

func (*AppUpdate) check(flags *promptFlags) error {
	if !flags.loggedIn {
		return fmt.Errorf("cannot app_update before login")
	}

	return nil
}

func (c *AppUpdate) args() ([]string, error) {
	if c == nil || c.AppID == "" {
		return nil, fmt.Errorf("app_update requires app ID")
	}

	args := []string{"app_update", c.AppID}

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

func (c *AppUpdate) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c *AppUpdate) modify(_ *promptFlags) error {
	return nil
}
