package steamcmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ForceInstallDir string

var _ Cmd = ForceInstallDir("")

func (c ForceInstallDir) String() string {
	return string(c)
}

func (ForceInstallDir) check(flags *promptFlags) error {
	if flags.loggedIn {
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

	if err := os.MkdirAll(a, 0o644); err != nil {
		return nil, err
	}

	return []string{"force_install_dir", a}, nil
}

func (c ForceInstallDir) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c ForceInstallDir) modify(_ *promptFlags) error {
	return nil
}
