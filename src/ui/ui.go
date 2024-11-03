package ui

import (
	"fmt"
	"http-tui/src/configs"
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
	fmt.Print(ANSI_MAP["HOME"] + ANSI_MAP["CLEAR"]) // Clear and move cursor to home
	fmt.Print(drawRectangle(uiManager.WindowSize.Width, uiManager.WindowSize.Height))

	fmt.Print(MoveCursorTo(0, 2))
	fmt.Print(" HTTP TUI ")

	fmt.Print(ANSI_MAP["HIDE_CURSOR"])
	// fmt.Print("\033[H") // Clear and move cursor to home
}

// UI Loop
func StartUI(uiManager *UIManager) {
	uiManager.logger.Printf("Start uiManager")
	go func() {
		for {
			select {
			case <-uiManager.tickerStopper:
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
