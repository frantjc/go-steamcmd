package steamcmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	vdf "github.com/frantjc/go-encoding-vdf"
)

type appInfoPrint int

var _ cmd = appInfoPrint(0)

func (c appInfoPrint) String() string {
	return fmt.Sprintf("%d", c)
}

var appInfos = map[int]AppInfo{}

func (c appInfoPrint) check(_ *promptFlags) error {
	return nil
}

func (c appInfoPrint) args() ([]string, error) {
	if c == 0 {
		return nil, fmt.Errorf("app_info_print requires app ID")
	}

	return []string{"app_info_print", c.String()}, nil
}

func (c appInfoPrint) readOutput(ctx context.Context, r io.Reader) error {
	var (
		errC    = make(chan error, 1)
		buf     = new(bytes.Buffer)
		errB    = []byte("ERROR! ")
		failedB = []byte("FAILED ")
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
				errC <- fmt.Errorf("%s", strings.ToLower(string(msgB)))
				return
			} else if _, msgB, found := bytes.Cut(p, failedB); found {
				msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
				errC <- fmt.Errorf("%s", strings.ToLower(string(msgB)))
				return
			} else if i := bytes.Index(p, []byte("{")); i >= 0 {
				appInfo := &AppInfo{}

				if err := vdf.NewDecoder(
					io.MultiReader(
						bytes.NewReader(p[i:buf.Len()]),
						r,
					),
				).Decode(appInfo); err != nil {
					errC <- err
					return
				}

				appInfos[int(c)] = *appInfo
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
}

func (c appInfoPrint) modify(_ *promptFlags) error {
	return nil
}
