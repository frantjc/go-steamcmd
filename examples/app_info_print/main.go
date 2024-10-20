package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/frantjc/go-steamcmd"
)

func main() {
	var (
		ctx   = context.Background()
		appID = 896660
		enc   = json.NewEncoder(os.Stdout)
	)

	prompt, err := steamcmd.Start(ctx)
	if err != nil {
		panic(err)
	}
	defer prompt.Quit(ctx)

	if err = prompt.Login(ctx); err != nil {
		panic(err)
	}

	// `steamcmd` tries to do this automatically during app_info_print,
	// but it often just hangs.
	if err = prompt.AppInfoRequest(ctx, appID); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 9)

	appInfo, err := prompt.AppInfoPrint(ctx, appID)
	if err != nil {
		panic(err)
	}

	enc.SetIndent("", "  ")

	enc.Encode(appInfo)
}
