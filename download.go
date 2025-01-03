//go:build darwin || linux

package steamcmd

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/frantjc/go-steamcmd/internal/cache"
	xtar "github.com/frantjc/x/archive/tar"
)

func Download(ctx context.Context) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, DownloadURL.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func New(ctx context.Context) (Path, error) {
	entrypoint := "steamcmd.sh"
	if bash, err := exec.LookPath("bash"); err != nil || bash == "" {
		entrypoint = steamcmdBinaryPath
	}

	if bin, err := exec.LookPath(entrypoint); errors.Is(err, exec.ErrDot) || err == nil {
		return Path(bin), nil
	} else if _, err := os.Stat(filepath.Join(cache.Dir, entrypoint)); errors.Is(err, os.ErrNotExist) {
		rc, err := Download(ctx)
		if err != nil {
			return "", err
		}
		defer rc.Close()

		r, err := gzip.NewReader(rc)
		if err != nil {
			return "", fmt.Errorf("steamcmd: ungzipping download tarball: %w", err)
		}
		defer r.Close()

		if err = xtar.Extract(tar.NewReader(r), cache.Dir); err != nil {
			return "", fmt.Errorf("steamcmd: extracting download tarball: %w", err)
		}
	} else if err != nil {
		return "", err
	}

	return Path(filepath.Join(cache.Dir, entrypoint)), nil
}
