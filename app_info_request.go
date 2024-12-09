package steamcmd

import (
	"context"
	"fmt"
)

type appInfoRequest int

var _ cmd = appInfoRequest(0)

func (c appInfoRequest) String() string {
	return fmt.Sprintf("%d", c)
}

func (c appInfoRequest) check(flags *promptFlags) error {
	if !flags.loggedIn {
		return fmt.Errorf("cannot app_info_request before login")
	}

	return nil
}

func (c appInfoRequest) args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_request requires app ID")
	}

	return []string{"app_info_request", c.String()}, nil
}

func (c appInfoRequest) readOutput(ctx context.Context, p *Prompt) error {
	return readOutput(ctx, p, 0)
}

func (c appInfoRequest) modify(_ *promptFlags) error {
	return nil
}
