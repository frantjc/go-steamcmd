package steamcmd

import "context"

type anyCommand struct {
	checkFn      func(flags *promptFlags) error
	argsFn       func() ([]string, error)
	readOutputFn func(context.Context, *Prompt) error
	modifyFn     func(flags *promptFlags) error
}

var _ cmd = &anyCommand{}

func (c *anyCommand) check(flags *promptFlags) error {
	return c.checkFn(flags)
}

func (c *anyCommand) args() ([]string, error) {
	return c.argsFn()
}

func (c *anyCommand) readOutput(ctx context.Context, p *Prompt) error {
	return c.readOutputFn(ctx, p)
}

func (c *anyCommand) modify(flags *promptFlags) error {
	return c.modifyFn(flags)
}
