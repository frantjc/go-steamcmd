package steamcmd

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
)

type forceInstallDir string

var _ cmd = forceInstallDir("")

func (c forceInstallDir) String() string {
	return string(c)
}

func (forceInstallDir) Check(flags *promptFlags) error {
	if flags.loggedIn {
		return fmt.Errorf("cannot force_install_dir after login")
	}

	return nil
}

func (c forceInstallDir) Args() ([]string, error) {
	if c == "" {
		return nil, fmt.Errorf("empty force_install_dir")
	}

	a, err := filepath.Abs(c.String())
	if err != nil {
		return nil, err
	}

	return []string{"force_install_dir", a}, nil
}

func (c forceInstallDir) ReadOutput(ctx context.Context, r io.Reader) error {
	return readOutput(ctx, r, 0)
}

func (c forceInstallDir) Modify(_ *promptFlags) error {
	return nil
}
