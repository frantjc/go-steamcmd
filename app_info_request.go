package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type appInfoRequest int

var _ cmd = appInfoRequest(0)

func (c appInfoRequest) String() string {
	return fmt.Sprintf("%d", c)
}

func (c appInfoRequest) check(_ *promptFlags) error {
	return nil
}

func (c appInfoRequest) args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_request requires app ID")
	}

	return []string{"app_info_request", c.String()}, nil
}

func (c appInfoRequest) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c appInfoRequest) modify(_ *promptFlags) error {
	return nil
}
