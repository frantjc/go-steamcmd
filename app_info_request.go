package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type AppInfoRequest string

func (c AppInfoRequest) String() string {
	return string(c)
}

func (c AppInfoRequest) check(_ *promptFlags) error {
	return nil
}

func (c AppInfoRequest) args() ([]string, error) {
	if c == "" {
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
