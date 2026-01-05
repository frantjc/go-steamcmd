package steamcmd

import "fmt"

type CommandError struct {
	Err    error
	Msg    string
	Output []byte
}

func (e *CommandError) Error() string {
	if e == nil {
		return ""
	}
	msg := e.Msg
	if e.Err != nil {
		msg = fmt.Sprintf("steamcmd: %s: %s", msg, e.Err.Error())
	}
	return msg
}

func (e *CommandError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}
