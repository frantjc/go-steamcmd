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

var _ command = &anyCommand{}

func (c *anyCommand) check(flags *promptFlags) error {
	return c.checkFn(flags)
}

func (c *anyCommand) args() ([]string, error) {
	return c.argsFn()
}

func (c *anyCommand) readOutput(ctx context.Context, r io.Reader) error {
	return c.readOutputFn(ctx, r)
}

func (c *anyCommand) modify(flags *promptFlags) error {
	return c.modifyFn(flags)
}
