package steamcmd

import (
	"context"
	"io"
)

var quit = &anyCommand{
	checkFn: func(_ *promptFlags) error {
		return nil
	},
	argsFn: func() ([]string, error) {
		return []string{"quit"}, nil
	},
	readOutputFn: func(ctx context.Context, r io.Reader) error {
		return nil
	},
	modifyFn: func(_ *promptFlags) error {
		return nil
	},
}
