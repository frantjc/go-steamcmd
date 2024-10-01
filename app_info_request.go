package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type AppInfoRequest int

var _ Cmd = AppInfoRequest(0)

func (c AppInfoRequest) String() string {
	return fmt.Sprintf("%d", c)
}

func (c AppInfoRequest) check(_ *promptFlags) error {
	return nil
}

func (c AppInfoRequest) args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_request requires app ID")
	}

	return []string{"app_info_request", c.String()}, nil
}

func (c AppInfoRequest) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c AppInfoRequest) modify(_ *promptFlags) error {
	return nil
}
