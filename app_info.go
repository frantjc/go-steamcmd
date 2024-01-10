package steamcmd

// TODO: Needs vdf equivalent of e.g. stdlib's encoding/json.

// type AppInfo struct {
// 	Common   *AppInfoCommon   `vdf:"common,omitempty"`
// 	Extended *AppInfoExtended `vdf:"extended,omitempty"`
// 	Config   *AppInfoConfig   `vdf:"config,omitempty"`
// }

// type AppInfoCommon struct {
// 	Name   string `vdf:"name,omitempty"`
// 	Type   string `vdf:"type,omitempty"`
// 	OSList string `vdf:"oslist,omitempty"`
// 	GameID string `vdf:"gameid,omitempty"`
// }

// type AppInfoExtended struct {
// 	Developer                 string `vdf:"developer,omitempty"`
// 	GameDir                   string `vdf:"gamedir,omitempty"`
// 	Homepage                  string `vdf:"homepage,omitempty"`
// 	Icon                      string `vdf:"icon,omitempty"`
// 	NoServers                 string `vdf:"noservers,omitempty"`
// 	PrimaryCache              string `vdf:"primarycache,omitempty"`
// 	SourceGame                string `vdf:"sourcegame,omitempty"`
// 	State                     string `vdf:"state,omitempty"`
// 	VisibleOnlyWhenInstalled  string `vdf:"visibleonlywheninstalled,omitempty"`
// 	VisibleOnlyWhenSubscribed string `vdf:"visibleonlywhensubscribed,omitempty"`
// }

// type AppInfoConfig struct {
// 	Launch      map[string]AppInfoConfigLaunch `vdf:"launch,omitempty"`
// 	ContentType string                         `vdf:"contenttype,omitempty"`
// 	InstallDir  string                         `vdf:"installdir,omitempty"`
// }

// type AppInfoConfigLaunch struct {
// 	Executable string                     `vdf:"executable,omitempty"`
// 	Arguments  string                     `vdf:"arguments,omitempty"`
// 	Config     *AppInfoConfigLaunchConfig `vdf:"config,omitempty"`
// }

// type AppInfoConfigLaunchConfig struct {
// 	OSList string `vdf:"oslist"`
// 	OSArch string `vdf:"osarch"`
// }
