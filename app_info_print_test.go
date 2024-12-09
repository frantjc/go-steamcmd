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
	ctx, stop := context.WithTimeout(context.TODO(), time.Minute)
	defer stop()

	prompt, err := steamcmd.Start(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer prompt.Close(ctx)

	if err = prompt.ForcePlatformType(ctx, steamcmd.PlatformTypeLinux); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err = prompt.Login(ctx); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err = prompt.AppInfoRequest(ctx, AppID); err != nil {
		t.Error(err)
		t.FailNow()
	}

	appInfo, err := prompt.AppInfoPrint(ctx, AppID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if appInfo.Common.GameID != AppID {
		t.Error("got wrong app ID", appInfo.Common.GameID)
		t.FailNow()
	}
}
