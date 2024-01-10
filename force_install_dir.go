package steamcmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type forceInstallDir string

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

	if err := os.MkdirAll(a, 0o644); err != nil {
		return nil, err
	}

	return []string{"force_install_dir", a}, nil
}

func (c forceInstallDir) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c forceInstallDir) modify(_ *promptFlags) error {
	return nil
}
