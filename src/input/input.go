package input

import (
	"bufio"
	"http-tui/src/configs"
	"log"
	"os"
	"os/exec"
	"time"
)

type InputManager struct {
	reader           bufio.Reader
	inputRefreshRate int
	ticker           time.Ticker
	tickerStopper    chan bool
	listeners        []InputListener
	logger           *log.Logger
}

func NewInputManager(appConfigs *configs.AppConfigs, logger *log.Logger) *InputManager {
	msPerTick := time.Millisecond * time.Duration(1000/appConfigs.InputRefreshRate)
	ticker := time.NewTicker(time.Duration(msPerTick))
	tickerStopper := make(chan bool)

	bufreader := bufio.NewReader(os.Stdin)

	InputManager := InputManager{
		reader:           *bufreader,
		inputRefreshRate: appConfigs.InputRefreshRate,
		ticker:           *ticker,
		tickerStopper:    tickerStopper,
		listeners:        make([]InputListener, 0),
		logger:           logger,
	}

	return &InputManager
}

func StartInput(inputManager *InputManager) {
	inputManager.logger.Println("Start InputManager")

	configTerm()

	go func() {
		for {
			select {
			case <-inputManager.tickerStopper:
				return
			case t := <-inputManager.ticker.C:
				var _ time.Time = t
				tick(inputManager)
			}
		}
	}()
}

func tick(inputManager *InputManager) {
	buffer := make([]byte, 1)

	n, err := inputManager.reader.Read(buffer)
	if err != nil {
		log.Panic(err)
	}
	if n == 0 {
		return
	}

	event := InputEvent{
		handled:   false,
		KeyCode:   int8(buffer[0]),
		KeyString: string(buffer),
	}

	for i := len(inputManager.listeners) - 1; i >= 0; i-- {
		listener := inputManager.listeners[i]
		handledByListener := listener.OnInputEvent(&event)
		if event.handled || handledByListener {
			break
		}
	}
}

func configTerm() {
	// disable input buffering
	err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	if err != nil {
		log.Fatal("Couldn't disable input buffering")
	}
	// do not display entered characters on the screen
	err = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	if err != nil {
		log.Fatal("Couldn't avoid rendering input on screen")
	}
}

func StopInput(inputManager *InputManager) {
	inputManager.logger.Println("Stop InputManager")
	inputManager.tickerStopper <- true
	err := exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	if err != nil {
		log.Fatal("Couldn't revert rendering input on screen")
	}
}
