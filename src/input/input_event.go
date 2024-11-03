package input

type InputEvent struct {
	KeyCode   int8
	KeyString string
	handled   bool
}

func (event *InputEvent) StopProgation() {
	event.handled = true
}

type InputListener interface {
	OnInputEvent(event *InputEvent) bool
}

func (inputManager *InputManager) AddListener(listener InputListener) bool {
	for _, knownListener := range inputManager.listeners {
		if knownListener == listener {
			return false
		}
	}
	inputManager.listeners = append(inputManager.listeners, listener)
	return true
}

func (inputManager *InputManager) RemovListener(listener InputListener) bool {
	found := false
	filteredListeners := []InputListener{}
	for _, each := range inputManager.listeners {
		if each != listener {
			filteredListeners = append(filteredListeners, each)
		} else {
			found = true
		}
	}
	if !found {
		return false
	}
	inputManager.listeners = filteredListeners
	return true
}
