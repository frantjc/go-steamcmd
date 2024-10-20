package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/frantjc/go-steamcmd"
	"github.com/frantjc/go-steamcmd/internal"
)

func main() {
	var (
		ctx   = context.Background()
		appID = 896660
		tmp   = filepath.Join(internal.Cache, fmt.Sprint(appID))
	)

	prompt, err := steamcmd.Start(ctx)
	if err != nil {
		panic(err)
	}

	if err := prompt.ForceInstallDir(ctx, tmp); err != nil {
		panic(err)
	}

	if err := prompt.Login(ctx); err != nil {
		panic(err)
	}

	if err := prompt.AppUpdate(ctx, appID); err != nil {
		panic(err)
	}
}
