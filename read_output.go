package steamcmd

import (
	"bytes"
	"context"
	"io"
	"strings"

	vdf "github.com/frantjc/go-encoding-vdf"
)

var (
	promptBytes = []byte("Steam>")
	errBytes    = []byte("ERROR! ")
	failedBytes = []byte("FAILED ")
)

func readOutput(ctx context.Context, r io.Reader, appID int) error {
	errC := make(chan error, 1)

	go func() {
		defer close(errC)

		errC <- func() error {
			buf := new(bytes.Buffer)

			for {
				var b [512]byte

				n, err := r.Read(b[:])
				if err != nil {
					return err
				}

				if _, err := buf.Write(b[:n]); err != nil {
					return err
				}

				p := buf.Bytes()
				if _, msgB, found := bytes.Cut(p, errBytes); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					return &CommandError{
						Msg:    strings.ToLower(string(msgB)),
						Output: p,
					}
				} else if _, msgB, found := bytes.Cut(p, failedBytes); found {
					msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
					return &CommandError{
						Msg:    strings.ToLower(string(msgB)),
						Output: p,
					}
				} else if appID > 0 {
					if i := bytes.Index(p, []byte("{")); i >= 0 {
						appInfo := &AppInfo{}

						if err := vdf.NewDecoder(
							io.MultiReader(
								bytes.NewReader(p[i:buf.Len()]),
								r,
							),
						).Decode(appInfo); err != nil {
							return &CommandError{
								Msg:    "decoding vdf",
								Err:    err,
								Output: p,
							}
						}

						appInfos[appID] = *appInfo
						return nil
					}
				} else if bytes.Contains(p, promptBytes) {
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
