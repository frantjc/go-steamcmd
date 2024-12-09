//go:build darwin

package steamcmd

import (
	"io"
	"net/url"
	"os/exec"
)

var (
	DownloadURL = func() *url.URL {
		u, err := url.Parse("https://steamcdn-a.akamaihd.net/client/installer/steamcmd_osx.tar.gz")
		if err != nil {
			panic(err)
		}

		return u
	}()
	DefaultPlatformType = PlatformTypeMacOS
	steamcmdBinaryPath  = "steamcmd"
)

func pipes(cmd *exec.Cmd) (io.ReadCloser, io.ReadCloser, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	return stdout, stderr, nil
}
