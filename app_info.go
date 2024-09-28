package steamcmd

type AppInfo struct {
	Common   *AppInfoCommon          `vdf:"common"`
	Extended *AppInfoExtended        `vdf:"extended"`
	Config   *AppInfoConfig          `vdf:"config"`
	Depots   map[string]AppInfoDepot `vdf:"depots"`
}

type AppInfoDepot struct {
	Config             *AppInfoConfig             `vdf:"config"`
	Manifests          map[string]AppInfoManifest `vdf:"manifests"`
	DepotFromApp       string                     `vdf:"depotfromapp"`
	EncryptedManifests map[string]AppInfoManifest `vdf:"encryptedmanifests"`
}

type AppInfoManifest struct {
	GID      string `vdf:"gid"`
	Size     string `vdf:"size"`
	Download string `vdf:"download"`
}

type AppInfoBranch struct {
	BuildID     string `vdf:"buildid"`
	Description string `vdf:"description"`
	TimeUpdated string `vdf:"timeupdated"`
	PwdRequired bool   `vdf:"pwdrequired"`
}

type AppInfoCommon struct {
	Name   string `vdf:"name"`
	Type   string `vdf:"type"`
	OSList string `vdf:"oslist"`
	GameID string `vdf:"gameid"`
}

type AppInfoExtended struct {
	Developer                 string `vdf:"developer"`
	GameDir                   string `vdf:"gamedir"`
	Homepage                  string `vdf:"homepage"`
	Icon                      string `vdf:"icon"`
	NoServers                 string `vdf:"noservers"`
	PrimaryCache              string `vdf:"primarycache"`
	SourceGame                string `vdf:"sourcegame"`
	State                     string `vdf:"state"`
	VisibleOnlyWhenInstalled  string `vdf:"visibleonlywheninstalled"`
	VisibleOnlyWhenSubscribed string `vdf:"visibleonlywhensubscribed"`
}

type AppInfoConfig struct {
	Launch      map[string]AppInfoConfigLaunch `vdf:"launch"`
	ContentType string                         `vdf:"contenttype"`
	InstallDir  string                         `vdf:"installdir"`
}

type AppInfoConfigLaunch struct {
	Executable string                     `vdf:"executable"`
	Arguments  string                     `vdf:"arguments"`
	Config     *AppInfoConfigLaunchConfig `vdf:"config"`
}

type AppInfoConfigLaunchConfig struct {
	OSList string `vdf:"oslist"`
	OSArch string `vdf:"osarch"`
}
