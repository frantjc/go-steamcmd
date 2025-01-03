package steamcmd

import (
	"fmt"
	"path/filepath"
)

type ForceInstallDir string

var _ Command = ForceInstallDir("")

func (c ForceInstallDir) String() string {
	return string(c)
}

func (ForceInstallDir) check(flags *flags) error {
	if flags.LoggedIn {
		return fmt.Errorf("cannot force_install_dir after login")
	}

	return nil
}

func (c ForceInstallDir) args() ([]string, error) {
	if c == "" {
		return nil, fmt.Errorf("empty force_install_dir")
	}

	a, err := filepath.Abs(c.String())
	if err != nil {
		return nil, err
	}

	return []string{"force_install_dir", a}, nil
}

func (c ForceInstallDir) modify(_ *flags) error {
	return nil
}
