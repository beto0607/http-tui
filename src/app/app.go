package app

import (
	"errors"
	"http-tui/src/configs"
	"http-tui/src/input"
	"http-tui/src/logger"
	"http-tui/src/ui"
	"log"
	"time"
)

type App struct {
	Name          string
	Configs       *configs.AppConfigs
	UIManager     *ui.UIManager
	InputManager  *input.InputManager
	ticker        time.Ticker
	tickerStopper chan bool
	Logger        *log.Logger
}

func (app App) OnInputEvent(event *input.InputEvent) bool {
	switch event.KeyCode {
	case 27, 81, 113: // esc, Q, q
		StopApp(&app)
	default:
		app.Logger.Println(event.KeyCode)
		app.Logger.Println(event.KeyString)
	}

	return false
}

func NewApp(name string, appConfigs *configs.AppConfigs) (*App, error) {
	newLogger := logger.NewLogger()

	uiManager, err := ui.NewUIManager(appConfigs, newLogger)
	if err != nil {
		return nil, errors.New("Couldn't not initialize UIManager: " + err.Error())

	}
	inputManager := input.NewInputManager(appConfigs, newLogger)

	ticker := time.NewTicker(1 * time.Second)
	tickerStopper := make(chan bool)

	app := App{
		Name:          name,
		Configs:       appConfigs,
		UIManager:     uiManager,
		InputManager:  inputManager,
		ticker:        *ticker,
		tickerStopper: tickerStopper,
		Logger:        newLogger,
	}

	ok := inputManager.AddListener(&app)
	if !ok {
		return nil, errors.New("Couldn't add listener")
	}

	return &app, nil
}

func StartApp(app *App) {
	app.Logger.Printf("Start app")
	ui.StartUI(app.UIManager)
	input.StartInput(app.InputManager)
}

func StopApp(app *App) {
	app.Logger.Printf("Stop app")
	app.tickerStopper <- true
	app.InputManager.RemovListener(app)
	input.StopInput(app.InputManager)
	ui.StopUI(app.UIManager)
	logger.StopLogger()
}

func LoopApp(app *App) {
	for {
		select {
		case <-app.tickerStopper:
			app.Logger.Printf("Close app")
			return
		case <-app.ticker.C:
			break
		}
	}
}
