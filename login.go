package steamcmd

import (
	"fmt"
)

type Login struct {
	Username       string
	Password       string
	SteamGuardCode string
}

func (c Login) Check(_ *Flags) error {
	return nil
}

func (c Login) Args() ([]string, error) {
	if c.Username == "" || c.Username == "anonymous" {
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

func (Login) Modify(flags *Flags) error {
	flags.LoggedIn = true
	return nil
}
