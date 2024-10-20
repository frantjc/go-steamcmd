//go:build darwin

package steamcmd

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/frantjc/go-steamcmd/internal"
	xtar "github.com/frantjc/x/archive/tar"
)

var DownloadURL = func() *url.URL {
	u, err := url.Parse("https://steamcdn-a.akamaihd.net/client/installer/steamcmd_osx.tar.gz")
	if err != nil {
		panic(err)
	}

	return u
}()

func New(ctx context.Context) (Command, error) {
	entrypoint := "steamcmd.sh"
	if bash, err := exec.LookPath("bash"); err != nil || bash == "" {
		entrypoint = "steamcmd"
	}

	if bin, err := exec.LookPath(entrypoint); errors.Is(err, exec.ErrDot) || err == nil {
		return Command(bin), nil
	} else if fi, err := os.Stat(filepath.Join(internal.Cache, entrypoint)); err != nil || fi.IsDir() {
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

		if err = xtar.Extract(tar.NewReader(r), internal.Cache); err != nil {
			return Command(""), err
		}
	}

	return Command(filepath.Join(internal.Cache, entrypoint)), nil
}
