//go:build linux

package steamcmd

import (
	"io"
	"net/url"
	"os/exec"
)

var (
	DownloadURL = func() *url.URL {
		u, err := url.Parse("https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz")
		if err != nil {
			panic(err)
		}

		return u
	}()
	DefaultPlatformType = PlatformTypeLinux
	steamcmdBinaryPath  = "linux32/steamcmd"
)

func pipes(cmd *exec.Cmd) (io.ReadCloser, io.ReadCloser, error) {
	stdoutr, stdoutw := io.Pipe()
	cmd.Stdout = stdoutw
	stderrr, stderrw := io.Pipe()
	cmd.Stderr = stderrw
	return stdoutr, stderrr, nil
}
