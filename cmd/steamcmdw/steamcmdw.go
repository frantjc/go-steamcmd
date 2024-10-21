package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/frantjc/go-steamcmd"
	"github.com/frantjc/go-steamcmd/internal"
	"github.com/spf13/cobra"
)

func NewEntrypoint() *cobra.Command {
	var (
		clean bool
		cmd   = &cobra.Command{
			Use:           "steamcmd",
			SilenceErrors: true,
			SilenceUsage:  true,
			RunE: func(cmd *cobra.Command, args []string) error {
				if clean {
					return os.RemoveAll(internal.Cache)
				}

				c, err := steamcmd.New(cmd.Context())
				if err != nil {
					return err
				}

				//nolint:gosec
				sub := exec.CommandContext(cmd.Context(), c.String(), args...)
				sub.Stdin = os.Stdin
				sub.Stdout = os.Stdout
				sub.Stderr = os.Stderr

				return sub.Run()
			},
		}
	)

	cmd.SetVersionTemplate("{{ .Name }}{{ .Version }} " + runtime.Version() + "\n")
	cmd.PersistentFlags().BoolVar(&clean, "clean", false, "delete steamcmd download and exit")

	return cmd
}
