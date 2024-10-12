package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	xerrors "github.com/frantjc/x/errors"
	xos "github.com/frantjc/x/os"
)

func main() {
	var (
		ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		err       = xerrors.Ignore(
			NewEntrypoint().ExecuteContext(ctx),
			context.Canceled,
		)
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	stop()
	xos.ExitFromError(err)
}
