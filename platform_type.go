package steamcmd

type PlatformType string

func (t PlatformType) String() string {
	return string(t)
}

var (
	PlatformTypeWindows PlatformType = "windows"
	PlatformTypeLinux   PlatformType = "linux"
	PlatformTypeMacOS   PlatformType = "macos"
)
