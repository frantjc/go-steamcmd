package steamcmd

import (
	"context"
)

func Start(ctx context.Context) (Prompt, error) {
	c, err := New(ctx)
	if err != nil {
		return nil, err
	}

	return c.Start(ctx)
}

func Run(ctx context.Context, cmds ...Cmd) error {
	c, err := New(ctx)
	if err != nil {
		return err
	}

	return c.Run(ctx, cmds...)
}

func GetAppInfo(ctx context.Context, appID int) (*AppInfo, error) {
	if appInfo, ok := appInfos[appID]; ok {
		return &appInfo, nil
	}

	if err := Run(ctx, AppInfoPrint(appID)); err != nil {
		return nil, err
	}

	appInfo := appInfos[appID]
	return &appInfo, nil
}
