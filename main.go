//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "qbittorrent-portable"
	Papp.Name = "qBittorrent"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))
	Papp.Process = PathJoin(Papp.AppPath, "qbittorrent.exe")
	Papp.Args = nil
	Papp.WorkingDir = Papp.AppPath

	profilePath := CreateFolder(PathJoin(Papp.DataPath, "profile"))
	OverrideEnv("QBT_PROFILE", profilePath)

	Launch(os.Args[1:])
}
