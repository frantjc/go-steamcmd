package steamcmd

import (
	"context"
	"io"
)

var Quit = &anyCommand{
	checkFn: func(_ *promptFlags) error {
		return nil
	},
	argsFn: func() ([]string, error) {
		return []string{"quit"}, nil
	},
	readOutputFn: func(_ context.Context, _ io.Reader) error {
		return nil
	},
	modifyFn: func(_ *promptFlags) error {
		return nil
	},
}
