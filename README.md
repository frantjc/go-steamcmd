# go-steamcmd [![CI](https://github.com/frantjc/go-steamcmd/actions/workflows/ci.yml/badge.svg?branch=main&event=push)](https://github.com/frantjc/go-steamcmd/actions) [![godoc](https://pkg.go.dev/badge/github.com/frantjc/go-steamcmd.svg)](https://pkg.go.dev/github.com/frantjc/go-steamcmd) [![goreportcard](https://goreportcard.com/badge/github.com/frantjc/go-steamcmd)](https://goreportcard.com/report/github.com/frantjc/go-steamcmd)

[Go](https://go.dev) module download and interact with Valve's Steam's [`steamcmd`](https://developer.valvesoftware.com/wiki/SteamCMD).

## Install

```sh
go get github.com/frantjc/go-steamcmd
```

## Use

```go
package main

import (
  "context"
  "fmt"

  "github.com/frantjc/go-steamcmd"
)

func main() {
  if err := steamcmd.Run(context.Background(),
    steamcmd.ForceInstallDir("./"),
    steamcmd.Login{},
    steamcmd.AppUpdate{AppID: 896660},
  ); err != nil {
    panic(err)
  }

  fmt.Println("installed the Valheim Dedicated Server to current working directory")
}
```

See more examples from the [tests](steamcmd_test.go).
