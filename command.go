package steamcmd

import (
	"context"
	"io"
)

type cmd interface {
	check(*promptFlags) error
	args() ([]string, error)
	readOutput(context.Context, io.Reader) error
	modify(*promptFlags) error
}
