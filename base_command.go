package steamcmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

var base = &anyCommand{
	checkFn: func(_ *promptFlags) error {
		return nil
	},
	argsFn: func() ([]string, error) {
		return make([]string, 0), nil
	},
	readOutputFn: func(ctx context.Context, r io.Reader) error {
		var (
			errC     = make(chan error, 1)
			buf      = new(bytes.Buffer)
			promptB  = []byte("Steam>")
			successB = []byte("Success!")
			errB     = []byte("ERROR! ")
			failedB  = []byte("FAILED ")
		)

		go func() {
			defer close(errC)

			for {
				var b [4096]byte

				if _, err := r.Read(b[:]); err != nil {
					errC <- err
					return
				}

				if _, err := buf.Write(b[:]); err != nil {
					errC <- err
					return
				}

				p := buf.Bytes()
				if _, msgB, found := bytes.Cut(p, errB); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					errC <- fmt.Errorf(string(msgB))
					return
				} else if _, msgB, found := bytes.Cut(p, failedB); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					errC <- fmt.Errorf(string(msgB))
					return
				} else if bytes.Contains(p, successB) || bytes.Contains(p, promptB) {
					return
				}
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errC:
			return err
		}
	},
	modifyFn: func(_ *promptFlags) error {
		return nil
	},
}
