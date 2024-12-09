package steamcmd

import (
	"context"
	"fmt"
)

type appInfoPrint int

var _ cmd = appInfoPrint(0)

func (c appInfoPrint) String() string {
	return fmt.Sprintf("%d", c)
}

var appInfos = map[int]AppInfo{}

func (c appInfoPrint) check(flags *promptFlags) error {
	if !flags.loggedIn {
		return fmt.Errorf("cannot app_info_print before login")
	}

	return nil
}

func (c appInfoPrint) args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_print requires app ID")
	}

	return []string{"app_info_print", c.String()}, nil
}

func (c appInfoPrint) readOutput(ctx context.Context, p *Prompt) error {
	return readOutput(ctx, p, int(c))
}

func (c appInfoPrint) modify(_ *promptFlags) error {
	return nil
}
