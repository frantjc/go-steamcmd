package internal

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

var Cache = filepath.Join(xdg.CacheHome, "steamcmd")
