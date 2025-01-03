package steamcmd_test

import (
	"context"
	"testing"

	"github.com/frantjc/go-steamcmd"
)

// CS2, Core Keeper (server), Valheim (server).
var AppIDs = []int{730, 1963720, 896660}

func TestAppInfoPrint(t *testing.T) {
	ctx := context.TODO()

	for _, appID := range AppIDs {
		if err := steamcmd.Run(ctx,
			steamcmd.Login{},
			steamcmd.AppInfoPrint(appID),
		); err != nil {
			t.Error(err)
			t.FailNow()
		}

		appInfo, found := steamcmd.GetAppInfo(appID)
		if !found {
			t.Error("did not get app info for app ID", appID)
			t.FailNow()
		}

		if appInfo.Common.GameID != appID {
			t.Error("got wrong app ID", appInfo.Common.GameID)
			t.FailNow()
		}
	}
}

func TestPrompt(t *testing.T) {
	ctx := context.TODO()

	for _, appID := range AppIDs {
		prompt, err := steamcmd.Start(ctx,
			steamcmd.Login{},
		)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if err = prompt.Run(ctx, steamcmd.AppInfoRequest(appID)); err != nil {
			t.Error(err)
			t.FailNow()
		}

		if err := prompt.Close(); err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
}
