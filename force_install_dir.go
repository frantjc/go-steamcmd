package steamcmd

import (
	"context"
	"fmt"
	"path/filepath"
)

type forceInstallDir string

var _ cmd = forceInstallDir("")

func (c forceInstallDir) String() string {
	return string(c)
}

func (forceInstallDir) check(flags *promptFlags) error {
	if flags.loggedIn {
		return fmt.Errorf("cannot force_install_dir after login")
	}

	return nil
}

func (c forceInstallDir) args() ([]string, error) {
	if c == "" {
		return nil, fmt.Errorf("empty force_install_dir")
	}

	a, err := filepath.Abs(c.String())
	if err != nil {
		return nil, err
	}

	return []string{"force_install_dir", a}, nil
}

func (c forceInstallDir) readOutput(ctx context.Context, p *Prompt) error {
	return readOutput(ctx, p, 0)
}

func (c forceInstallDir) modify(_ *promptFlags) error {
	return nil
}
