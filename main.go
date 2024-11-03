package main

import (
	"http-tui/src/app"
	"http-tui/src/configs"
	"log"
)

func main() {
	confs := configs.NewAppConfigs()
	application, err := app.NewApp("HTTP TUI", confs)
	if err != nil {
		log.Fatal(err.Error())
	}

	Start(application)
}

func Start(application *app.App) {
	app.StartApp(application)
	app.LoopApp(application)
}
