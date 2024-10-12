package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/adrg/xdg"
	"github.com/frantjc/go-steamcmd"

	xtar "github.com/frantjc/x/archive/tar"
	"github.com/spf13/cobra"
)

func NewEntrypoint() *cobra.Command {
	var (
		download, clean bool
		cmd             = &cobra.Command{
			Use:           "steamcmd",
			SilenceErrors: true,
			SilenceUsage:  true,
			RunE: func(cmd *cobra.Command, args []string) error {
				var (
					cache = filepath.Join(xdg.CacheHome, "steamcmd")
					bin   = filepath.Join(cache, "steamcmd.sh")
				)

				if clean {
					return os.RemoveAll(cache)
				}

				if !download {
					if path, err := exec.LookPath("steamcmd.sh"); errors.Is(err, exec.ErrDot) || err == nil {
						bin = path
					} else {
						if fi, err := os.Stat(bin); err != nil || fi.IsDir() {
							download = true
						}
					}
				}

				if download {
					if err := os.RemoveAll(cache); err != nil {
						return err
					}

					if err := os.MkdirAll(cache, 0o775); err != nil {
						return err
					}

					rc, err := steamcmd.Download(cmd.Context())
					if err != nil {
						return err
					}
					defer rc.Close()

					r, err := gzip.NewReader(rc)
					if err != nil {
						return err
					}
					defer r.Close()

					if err = xtar.Extract(tar.NewReader(r), cache); err != nil {
						return err
					}
				}

				sub := exec.CommandContext(cmd.Context(), bin, args...)
				sub.Stdin = os.Stdin
				sub.Stdout = os.Stdout
				sub.Stderr = os.Stderr

				return sub.Run()
			},
		}
	)

	cmd.SetVersionTemplate("{{ .Name }}{{ .Version }} " + runtime.Version() + "\n")
	cmd.PersistentFlags().BoolVar(&clean, "clean", false, "delete steamcmd download and exit")
	cmd.PersistentFlags().BoolVar(&download, "download", false, "force redownload of steamcmd")

	return cmd
}
