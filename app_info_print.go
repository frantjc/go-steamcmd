package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type appInfoPrint int

var _ cmd = appInfoPrint(0)

func (c appInfoPrint) String() string {
	return fmt.Sprintf("%d", c)
}

var appInfos = map[int]AppInfo{}

func (c appInfoPrint) Check(_ *promptFlags) error {
	return nil
}

func (c appInfoPrint) Args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_print requires app ID")
	}

	return []string{"app_info_print", c.String()}, nil
}

func (c appInfoPrint) ReadOutput(ctx context.Context, r io.Reader) error {
	return readOutput(ctx, r, int(c))
}

func (c appInfoPrint) Modify(_ *promptFlags) error {
	return nil
}
