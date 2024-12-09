package steamcmd

import "context"

var quit = &anyCommand{
	checkFn: func(_ *promptFlags) error {
		return nil
	},
	argsFn: func() ([]string, error) {
		return []string{"quit"}, nil
	},
	readOutputFn: func(_ context.Context, _ *Prompt) error {
		return nil
	},
	modifyFn: func(_ *promptFlags) error {
		return nil
	},
}
