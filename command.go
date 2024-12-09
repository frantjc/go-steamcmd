package steamcmd

import "context"

type cmd interface {
	check(*promptFlags) error
	args() ([]string, error)
	readOutput(context.Context, *Prompt) error
	modify(*promptFlags) error
}
