package steamcmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"

	vdf "github.com/frantjc/go-encoding-vdf"
)

var (
	promptBytes = []byte("Steam>")
	errBytes    = []byte("ERROR! ")
	failedBytes = []byte("FAILED ")
)

func readOutput(ctx context.Context, p *Prompt, appID int) error {
	errC := make(chan error, 1)

	go func() {
		defer close(errC)

		errC <- func() error {
			buf := new(bytes.Buffer)

			for {
				var b [512]byte

				n, err := p.stdout.Read(b[:])
				if err != nil {
					return err
				}

				if _, err := buf.Write(b[:n]); err != nil {
					return err
				}

				q := buf.Bytes()
				if _, msgB, found := bytes.Cut(q, errBytes); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					return &CommandError{
						Msg:    strings.ToLower(string(msgB)),
						Output: q,
					}
				} else if _, msgB, found := bytes.Cut(q, failedBytes); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					return &CommandError{
						Msg:    strings.ToLower(string(msgB)),
						Output: q,
					}
				} else if appID > 0 {
					if i := bytes.Index(q, []byte("{")); i >= 0 {
						appInfo := &AppInfo{}

						// On Linux, steamcmd needs a kick to make it
						// finish printing the output of app_info_print.
						if runtime.GOOS == "linux" {
							cctx, cancel := context.WithTimeout(ctx, time.Second*9)
							defer cancel()

							go func() {
								for {
									select {
									case <-cctx.Done():
										return
									case <-time.NewTimer(time.Second).C:
										_, _ = fmt.Fprintln(p.stdin)
									}
								}
							}()
						}

						if err := vdf.NewDecoder(
							io.MultiReader(
								bytes.NewReader(q[i:buf.Len()]),
								p.stdout,
							),
						).Decode(appInfo); err != nil {
							return &CommandError{
								Msg:    "decoding vdf",
								Err:    err,
								Output: q,
							}
						}

						appInfos[appID] = *appInfo
						return nil
					}
				} else if bytes.Contains(q, promptBytes) {
					return nil
				}
			}
		}()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errC:
		return err
	}
}
