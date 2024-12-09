//go:build linux

package steamcmd

import "net/url"

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
