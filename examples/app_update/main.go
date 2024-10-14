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

	fmt.Println(tmp)

	if err := steamcmd.Run(ctx,
		steamcmd.ForceInstallDir(tmp),
		&steamcmd.Login{},
		&steamcmd.AppUpdate{
			AppID: appID,
		},
	); err != nil {
		panic(err)
	}
}
