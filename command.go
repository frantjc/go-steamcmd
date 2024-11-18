package steamcmd

import (
	"context"
	"io"
)

type cmd interface {
	Check(*promptFlags) error
	Args() ([]string, error)
	ReadOutput(context.Context, io.Reader) error
	Modify(*promptFlags) error
}
