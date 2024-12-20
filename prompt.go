package steamcmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	xslice "github.com/frantjc/x/slice"
)

type promptFlags struct {
	loggedIn bool
}

type Prompt struct {
	flags  *promptFlags
	stdin  io.WriteCloser
	stdout io.ReadCloser
	err    error
	mu     sync.Mutex
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

func (p *Prompt) ForceInstallDir(ctx context.Context, dir string) error {
	return p.run(ctx, forceInstallDir(dir))
}

func (p *Prompt) ForcePlatformType(ctx context.Context, platformType PlatformType) error {
	return p.run(ctx, forcePlatformType(platformType))
}

func (p *Prompt) Login(ctx context.Context, opts ...LoginOpt) error {
	cmd := &login{}

	for _, opt := range opts {
		opt(cmd)
	}

	return p.run(ctx, cmd)
}

func (p *Prompt) AppUpdate(ctx context.Context, appID int, opts ...AppUpdateOpt) error {
	cmd := &appUpdate{
		AppID: appID,
	}

	for _, opt := range opts {
		opt(cmd)
	}

	return p.run(ctx, cmd)
}

func (p *Prompt) AppInfoPrint(ctx context.Context, appID int) (*AppInfo, error) {
	if appInfo, ok := appInfos[appID]; ok {
		return &appInfo, nil
	}

	if err := p.run(ctx, appInfoRequest(appID)); err != nil {
		return nil, err
	}

	time.Sleep(time.Second)

	if err := p.run(ctx, appInfoPrint(appID)); err != nil {
		return nil, err
	}

	appInfo := appInfos[appID]

	return &appInfo, nil
}

func (p *Prompt) AppInfoRequest(ctx context.Context, appID int) error {
	if _, ok := appInfos[appID]; ok {
		return nil
	}

	return p.run(ctx, appInfoRequest(appID))
}

func (p *Prompt) WorkshopDownloadItem(ctx context.Context, appID, publishedFileID int) error {
	return p.run(ctx, &workshopDownloadItem{
		AppID:           appID,
		PublishedFileID: publishedFileID,
	})
}

func (p *Prompt) Close(ctx context.Context) error {
	return errors.Join(p.run(ctx, quit), p.stdin.Close(), p.stdout.Close())
}

func (p *Prompt) run(ctx context.Context, cmd cmd) error {
	if p.err != nil {
		return p.err
	}

	if err := cmd.check(p.flags); err != nil {
		return err
	}

	args, err := cmd.args()
	if err != nil {
		return err
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, err := fmt.Fprintln(p.stdin, xslice.Map(args, func(arg string, _ int) any {
		return arg
	})...); err != nil {
		return err
	}

	if err := cmd.readOutput(ctx, p); err != nil {
		return err
	}

	return cmd.modify(p.flags)
}
