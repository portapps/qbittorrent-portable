//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
package main

import (
	"os"
	"path/filepath"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
)

var (
	app *portapps.App
)

func init() {
	var err error

	// Init app
	if app, err = portapps.New("qbittorrent-portable", "qBittorrent"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	if err := os.MkdirAll(app.DataPath, 0o755); err != nil {
		log.Fatal().Err(err).Msg("Cannot create data path")
	}
	app.Process = filepath.Join(app.AppPath, "qbittorrent.exe")

	profilePath := filepath.Join(app.DataPath, "profile")
	if err := os.MkdirAll(profilePath, 0o755); err != nil {
		log.Fatal().Err(err).Msg("Cannot create profile path")
	}
	os.Setenv("QBT_PROFILE", profilePath)

	defer app.Close()
	app.Launch(os.Args[1:])
}
