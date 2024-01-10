package steamcmd

// TODO: Needs vdf equivalent of e.g. stdlib's encoding/json.

// import (
// 	"bytes"
// 	"context"
//  "exec"
// 	"fmt"
// 	"io"

// 	vdf "github.com/frantjc/go-encoding-vdf"
// )

// type AppInfoPrintCommand string

// func (c AppInfoPrintCommand) String() string {
// 	return string(c)
// }

// var (
// 	appInfos map[string]AppInfo
// )

// func (c AppInfoPrintCommand) Check(flags *promptFlags) error {
// 	return new(baseCommand).Check(flags)
// }

// func (c AppInfoPrintCommand) Args() ([]string, error) {
// 	if c == "" {
// 		return nil, fmt.Errorf("app_info_print requires app ID")
// 	} else if _, ok := appInfos[c.String()]; ok {
// 		return new(baseCommand).Args()
// 	}

// 	return []string{"app_info_print", c.String()}, nil
// }

// func (c AppInfoPrintCommand) ReadOutput(ctx context.Context, r io.Reader) error {
// 	if _, ok := appInfos[c.String()]; ok {
// 		return new(baseCommand).ReadOutput(ctx, r)
// 	}

// 	var (
// 		pr, pw    = io.Pipe()
// 		buf       = new(bytes.Buffer)
// 		mw        = io.MultiWriter(pw, buf)
// 		errB      = []byte("ERROR! ")
// 		notFoundB = []byte("No app info for AppID")
// 	)

// 	go func() {
// 		for {
// 			var b [4096]byte

// 			n, err := r.Read(b[:])
// 			if err != nil {
// 				_ = pw.CloseWithError(err)
// 				return
// 			}

// 			if _, err := mw.Write(b[:n]); err != nil {
// 				_ = pw.CloseWithError(err)
// 				return
// 			}

// 			p := buf.Bytes()
// 			if _, msgB, found := bytes.Cut(p, errB); found {
// 				msgB, _, _ = bytes.Cut(msgB, []byte("\n"))
// 				_ = pw.CloseWithError(fmt.Errorf(string(msgB)))
// 				return
// 			} else if bytes.Contains(p, notFoundB) {
// 				_ = pw.CloseWithError(fmt.Errorf("app info for app ID %s not found", c.String()))
// 				return
// 			}
// 		}
// 	}()

// 	appInfo := &AppInfo{}
// 	if err := vdf.NewDecoder(pr).Decode(appInfo); err != nil {
// 		return err
// 	}

// 	if appInfos == nil {
// 		appInfos = make(map[string]AppInfo)
// 	}

// 	appInfos[c.String()] = *appInfo
// 	return nil
// }

// func (c AppInfoPrintCommand) Modify(flags *promptFlags) error {
// 	return new(baseCommand).Modify(flags)
// }

// func (c AppInfoPrintCommand) AppInfo() *AppInfo {
// 	if appInfo, ok := appInfos[c.String()]; ok {
// 		return &appInfo
// 	}

// 	return nil
// }

// type AppInfoPrintCommandLine AppInfoPrintCommand

// func (s *Steamcmd) AppInfoPrint(ctx context.Context, l AppInfoPrintCommandLine) (*AppInfo, error) {
// 	c := AppInfoPrintCommand(l)

// 	if appInfo := c.AppInfo(); appInfo != nil {
// 		return appInfo, nil
// 	}

// 	if err := s.init(); err != nil {
// 		return nil, err
// 	}

// 	args, err := c.Args()
// 	if err != nil {
// 		return nil, err
// 	} else if len(args) > 0 {
// 		args[0] = "+" + args[0]
// 	}

// 	args = append(args, "+quit")

// 	var (
// 		rErrC  = make(chan error, 1)
// 		oErrC  = make(chan error, 1)
// 		cmd    = exec.CommandContext(ctx, s.Path, args...)
// 		pr, pw = io.Pipe()
// 	)
// 	cmd.Stdout = pw

// 	defer pw.Close()
// 	defer pr.Close()

// 	go func() {
// 		rErrC <- cmd.Run()
// 	}()
// 	go func() {
// 		oErrC <- c.ReadOutput(ctx, pr)
// 	}()

// 	select {
// 	case err := <-rErrC:
// 		return nil, err
// 	case err := <-oErrC:
// 		if err != nil {
// 			return nil, err
// 		}

// 		return c.AppInfo(), nil
// 	}
// }
