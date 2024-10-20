package steamcmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	xslice "github.com/frantjc/x/slice"
)

type promptFlags struct {
	loggedIn bool
}

type prompt struct {
	flags  *promptFlags
	stdin  io.Writer
	stdout io.Reader
	mu     sync.Mutex
	err    error
}

type AppUpdateOpt func(*appUpdate)

func WithValidate(au *appUpdate) {
	au.Validate = true
}

func WithBeta(beta, betaPassword string) AppUpdateOpt {
	return func(au *appUpdate) {
		au.Beta = beta
		au.BetaPassword = betaPassword
	}
}

type LoginOpt func(*login)

func WithAccount(username, password string) LoginOpt {
	return func(l *login) {
		l.Username = username
		l.Password = password
	}
}

func WithSteamGuardCode(steamGuardCode string) LoginOpt {
	return func(l *login) {
		l.SteamGuardCode = steamGuardCode
	}
}

type Prompt interface {
	ForceInstallDir(context.Context, string) error
	Login(context.Context, ...LoginOpt) error
	ForcePlatformType(context.Context, PlatformType) error
	AppUpdate(context.Context, int, ...AppUpdateOpt) error
	AppInfoPrint(context.Context, int) (*AppInfo, error)
	AppInfoRequest(context.Context, int) error
	WorkshopDownloadItem(context.Context, int, int) error
	Quit(context.Context) error
}

func (p *prompt) ForceInstallDir(ctx context.Context, dir string) error {
	return errors.Join(p.err, p.run(ctx, forceInstallDir(dir)))
}

func (p *prompt) ForcePlatformType(ctx context.Context, platformType PlatformType) error {
	return errors.Join(p.err, p.run(ctx, forcePlatformType(platformType)))
}

func (p *prompt) Login(ctx context.Context, opts ...LoginOpt) error {
	cmd := &login{}

	for _, opt := range opts {
		opt(cmd)
	}

	return errors.Join(p.err, p.run(ctx, cmd))
}

func (p *prompt) AppUpdate(ctx context.Context, appID int, opts ...AppUpdateOpt) error {
	cmd := &appUpdate{
		AppID: appID,
	}

	for _, opt := range opts {
		opt(cmd)
	}

	return errors.Join(p.err, p.run(ctx, cmd))
}

func (p *prompt) AppInfoPrint(ctx context.Context, appID int) (*AppInfo, error) {
	if appInfo, ok := appInfos[appID]; ok {
		return &appInfo, nil
	}

	if err := errors.Join(p.err, p.run(ctx, appInfoPrint(appID))); err != nil {
		return nil, err
	}

	appInfo := appInfos[appID]

	return &appInfo, nil
}

func (p *prompt) AppInfoRequest(ctx context.Context, appID int) error {
	if _, ok := appInfos[appID]; ok {
		return nil
	}

	return errors.Join(p.err, p.run(ctx, appInfoRequest(appID)))
}

func (p *prompt) WorkshopDownloadItem(ctx context.Context, appID, publishedFileID int) error {
	return errors.Join(p.err, p.run(ctx, &workshopDownloadItem{
		AppID:           appID,
		PublishedFileID: publishedFileID,
	}))
}

func (p *prompt) Quit(ctx context.Context) error {
	return errors.Join(p.err, p.run(ctx, quit))
}

func (p *prompt) run(ctx context.Context, cmd cmd) error {
	if err := cmd.check(p.flags); err != nil {
		return err
	}

	args, err := cmd.args()
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(p.stdin, xslice.Map(args, func(arg string, _ int) any {
		return arg
	})...); err != nil {
		return err
	}

	if err := cmd.readOutput(ctx, p.stdout); err != nil {
		return err
	}

	return cmd.modify(p.flags)
}
