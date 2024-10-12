package steamcmd

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/adrg/xdg"
	xtar "github.com/frantjc/x/archive/tar"
)

func Start(ctx context.Context) (Prompt, error) {
	c, err := New(ctx)
	if err != nil {
		return nil, err
	}

	return c.Start(ctx)
}

func Run(ctx context.Context, cmds ...Cmd) error {
	c, err := New(ctx)
	if err != nil {
		return err
	}

	return c.Run(ctx, cmds...)
}

func New(ctx context.Context) (Command, error) {
	var (
		cache      = filepath.Join(xdg.CacheHome, "steamcmd")
		entrypoint = "steamcmd.sh"
	)
	if bash, err := exec.LookPath("bash"); err != nil || bash == "" {
		entrypoint = "steamcmd"
	}

	if bin, err := exec.LookPath(entrypoint); !errors.Is(err, exec.ErrDot) && err == nil {
		return Command(bin), nil
	} else if fi, err := os.Stat(filepath.Join(cache, entrypoint)); err != nil || fi.IsDir() {
		rc, err := Download(ctx)
		if err != nil {
			return Command(""), err
		}
		defer rc.Close()

		r, err := gzip.NewReader(rc)
		if err != nil {
			return Command(""), err
		}
		defer r.Close()

		if err = xtar.Extract(tar.NewReader(r), cache); err != nil {
			return Command(""), err
		}
	}

	return Command(filepath.Join(cache, entrypoint)), nil
}
