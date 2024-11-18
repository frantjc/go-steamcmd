package main

import (
	"context"
	"os"

	"github.com/frantjc/go-steamcmd"
)

func main() {
	var (
		ctx   = context.Background()
		appID = 896660
	)

	prompt, err := steamcmd.Start(ctx)
	if err != nil {
		panic(err)
	}
	defer prompt.Close(ctx)

	if err := prompt.ForceInstallDir(ctx, os.Args[1]); err != nil {
		panic(err)
	}

	if err := prompt.ForcePlatformType(ctx, steamcmd.PlatformTypeLinux); err != nil {
		panic(err)
	}

	if err := prompt.Login(ctx); err != nil {
		panic(err)
	}

	if err := prompt.AppUpdate(ctx, appID); err != nil {
		panic(err)
	}
}
