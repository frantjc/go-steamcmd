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

type Prompt interface {
	Run(context.Context) error
	ForceInstallDir(context.Context, string) error
	Login(context.Context, *Login) error
	ForcePlatformType(context.Context, PlatformType) error
	AppUpdate(context.Context, *AppUpdate) error
	Quit() error
}

func (p *prompt) Run(ctx context.Context) error {
	return errors.Join(p.err, p.run(ctx, base))
}

func (p *prompt) ForceInstallDir(ctx context.Context, dir string) error {
	return errors.Join(p.err, p.run(ctx, forceInstallDir(dir)))
}

func (p *prompt) ForcePlatformType(ctx context.Context, platformType PlatformType) error {
	return errors.Join(p.err, p.run(ctx, forcePlatformType(platformType)))
}

func (p *prompt) Login(ctx context.Context, cmd *Login) error {
	return errors.Join(p.err, p.run(ctx, cmd))
}

func (p *prompt) AppUpdate(ctx context.Context, cmd *AppUpdate) error {
	return errors.Join(p.err, p.run(ctx, cmd))
}

func (p *prompt) Quit() error {
	return errors.Join(p.err, p.run(context.TODO(), quit))
}

func (p *prompt) run(ctx context.Context, cmd command) error {
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
