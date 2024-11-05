package ui

import (
	"fmt"
	"http-tui/src/configs"
	"http-tui/src/utils"
	"log"
	"time"
)

type Size struct {
	Width  int
	Height int
}
type UIManager struct {
	WindowSize    *Size
	FPS           int
	ticker        time.Ticker
	tickerStopper chan bool
	logger        *log.Logger
}

func NewUIManager(appConfigs *configs.AppConfigs, logger *log.Logger) (*UIManager, error) {
	msPerTick := time.Millisecond * time.Duration(1000/appConfigs.FPS)
	ticker := time.NewTicker(time.Duration(msPerTick))
	tickerStopper := make(chan bool)

	windowSize, err := getTerminalSize()
	if err != nil {
		return nil, err
	}

	uiManager := UIManager{
		FPS:           appConfigs.FPS,
		ticker:        *ticker,
		tickerStopper: tickerStopper,
		logger:        logger,
		WindowSize:    windowSize,
	}
	return &uiManager, nil
}

func tick(uiManager *UIManager, _ time.Time) {
	// fmt.Print(utils.ANSI_MAP["CLEAR_TO_END"])
	// fmt.Print(utils.ANSI_MAP["CLEAR_TO_BEGINNING"])
	fmt.Print(utils.ANSI_MAP["HOME"])
	// fmt.Print(utils.ANSI_MAP["CLEAR"]) // Clear and move cursor to home
	fmt.Print(drawRectangle(uiManager.WindowSize.Width, uiManager.WindowSize.Height))

	fmt.Print(utils.MoveCursorTo(0, 2))
	fmt.Print("\033[32m")
	fmt.Print(" HTTP TUI ")
	fmt.Print("\033[0m")
}

func setup() {
	fmt.Print(utils.ANSI_MAP["SCREEN_SAVE"])
	fmt.Print(utils.ANSI_MAP["SCREEN_MODE_SET"])
	fmt.Print(utils.ANSI_MAP["HIDE_CURSOR"])
}

func revertSetup() {
	fmt.Print(utils.ANSI_MAP["SHOW_CURSOR"])
	fmt.Print(utils.ANSI_MAP["CLEAR"]) // Clear and move cursor to home
	fmt.Print(utils.ANSI_MAP["SCREEN_MODE_UNSET"])
	fmt.Print(utils.ANSI_MAP["SCREEN_RESTORE"])
}

// UI Loop
func StartUI(uiManager *UIManager) {
	setup()
	uiManager.logger.Printf("Start uiManager")
	go func() {
		for {
			select {
			case <-uiManager.tickerStopper:
				revertSetup()
				return
			case t := <-uiManager.ticker.C:
				tick(uiManager, t)
			}
		}
	}()
}

func StopUI(uiManager *UIManager) {
	uiManager.tickerStopper <- true
}
