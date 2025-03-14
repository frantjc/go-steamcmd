package steamcmd

var appInfos = map[int]AppInfo{}

// GetAppInfo returns the app info for the given app ID, assuming that
// steamcmd has previously ran AppInfoPrint for the given app ID.
func GetAppInfo(appID int) (*AppInfo, bool) {
	appInfo, ok := appInfos[appID]
	if ok {
		return &appInfo, true
	}
	return nil, false
}

type AppInfo struct {
	Common AppInfoCommon `vdf:"common"`
	Config AppInfoConfig `vdf:"config"`
	Depots AppInfoDepots `vdf:"depots"`
}

type AppInfoDepots struct {
	Branches map[string]AppInfoDepotsBranch `vdf:"branches"`
}

type AppInfoDepotsBranch struct {
	BuildID     int    `vdf:"buildid"`
	Description string `vdf:"description"`
	TimeUpdated int    `vdf:"timeupdated"`
	PwdRequired bool   `vdf:"pwdrequired"`
}

type AppInfoCommon struct {
	Name       string `vdf:"name"`
	Type       string `vdf:"type"`
	Parent     int    `vdf:"parent"`
	OSList     string `vdf:"oslist"`
	GameID     int    `vdf:"gameid"`
	OSArch     string `vdf:"osarch"`
	Icon       string `vdf:"icon"`
	Logo       string `vdf:"logo"`
	LogoSmall  string `vdf:"logo_small"`
	ClientIcon string `vdf:"clienticon"`
}

type AppInfoConfig struct {
	Launch      map[string]AppInfoConfigLaunch `vdf:"launch"`
	ContentType string                         `vdf:"contenttype"`
	InstallDir  string                         `vdf:"installdir"`
}

type AppInfoConfigLaunch struct {
	Executable string                     `vdf:"executable"`
	Arguments  string                     `vdf:"arguments"`
	Type       string                     `vdf:"type"`
	Config     *AppInfoConfigLaunchConfig `vdf:"config"`
}

type AppInfoConfigLaunchConfig struct {
	OSList string `vdf:"oslist"`
	OSArch string `vdf:"osarch"`
}
