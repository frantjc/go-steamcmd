package steamcmd

import (
	"context"
	"fmt"
)

type login struct {
	Username       string
	Password       string
	SteamGuardCode string
}

func (c *login) check(_ *promptFlags) error {
	return nil
}

func (c *login) args() ([]string, error) {
	if c == nil || c.Username == "" || c.Username == "anonymous" {
		return []string{"login", "anonymous"}, nil
	}

	args := []string{"login", c.Username}

	switch {
	case c.Password != "":
		args = append(args, c.Password)

		if c.SteamGuardCode != "" {
			args = append(args, c.SteamGuardCode)
		}
	case c.SteamGuardCode != "":
		return nil, fmt.Errorf("specified Steam Guard code without password")
	default:
		return nil, fmt.Errorf("non-anonymous username given without a password or Steam Guard code")
	}

	return args, nil
}

func (c *login) readOutput(ctx context.Context, p *Prompt) error {
	return readOutput(ctx, p, 0)
}

func (*login) modify(flags *promptFlags) error {
	if flags == nil {
		flags = &promptFlags{}
	}

	flags.loggedIn = true

	return nil
}
