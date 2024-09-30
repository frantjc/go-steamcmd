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
	defer prompt.Quit()

	if err = prompt.Login(ctx, &steamcmd.Login{}); err != nil {
		panic(err)
	}

	// `steamcmd` tries to do this automatically during app_info_print,
	// but it often just hangs.
	if err = prompt.AppInfoRequest(ctx, steamcmd.AppInfoRequest(appID)); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 9)

	appInfo, err := prompt.AppInfoPrint(ctx, steamcmd.AppInfoPrint(appID))
	if err != nil {
		panic(err)
	}

	enc.SetIndent("", "  ")

	enc.Encode(appInfo)
}
