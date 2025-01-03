package steamcmd_test

import (
	"context"
	"testing"
	"time"

	"github.com/frantjc/go-steamcmd"
)

const (
	AppID = 730
)

func TestAppInfoPrint(t *testing.T) {
	ctx, stop := context.WithTimeout(context.TODO(), time.Second*33)
	defer stop()

	if _, err := steamcmd.NewPrompt(ctx,
		steamcmd.ForcePlatformType(steamcmd.PlatformTypeLinux),
		steamcmd.Login{},
		steamcmd.AppInfoRequest(AppID),
		steamcmd.AppInfoPrint(AppID),
		steamcmd.Quit,
	); err != nil {
		t.Error(err)
		t.FailNow()
	}

	appInfo, found := steamcmd.GetAppInfo(AppID)
	if !found {
		t.Error("did not get app info for app ID", AppID)
		t.FailNow()
	}

	if appInfo.Common.GameID != AppID {
		t.Error("got wrong app ID", appInfo.Common.GameID)
		t.FailNow()
	}
}

func TestPrompt(t *testing.T) {
	ctx, stop := context.WithTimeout(context.TODO(), time.Second*33)
	defer stop()

	prompt, err := steamcmd.NewPrompt(ctx,
		steamcmd.ForcePlatformType(steamcmd.PlatformTypeLinux),
		steamcmd.Login{},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer func() {
		if err := prompt.Close(ctx); err != nil {
			t.Error(err)
			t.FailNow()
		}
	}()

	if err = prompt.Run(ctx, steamcmd.AppInfoRequest(AppID)); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
