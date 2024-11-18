package steamcmd

import (
	"context"
	"io"
)

type anyCommand struct {
	checkFn      func(flags *promptFlags) error
	argsFn       func() ([]string, error)
	readOutputFn func(context.Context, io.Reader) error
	modifyFn     func(flags *promptFlags) error
}

var _ cmd = &anyCommand{}

func (c *anyCommand) Check(flags *promptFlags) error {
	return c.checkFn(flags)
}

func (c *anyCommand) Args() ([]string, error) {
	return c.argsFn()
}

func (c *anyCommand) ReadOutput(ctx context.Context, r io.Reader) error {
	return c.readOutputFn(ctx, r)
}

func (c *anyCommand) Modify(flags *promptFlags) error {
	return c.modifyFn(flags)
}
