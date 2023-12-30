package OS

import (
	"fmt"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/OS/widgets"
	"godeck/src/core/connector"
)

type LockScreen struct {
	isLocked bool

	*graphic.Screen

	enteredPin string
}

func CreateLockScreen(width int, height int, connector connector.AppConnector) *LockScreen {
	lockScreen := &LockScreen{
		Screen: &graphic.Screen{
			Width:        width,
			Height:       height,
			AppConnector: connector,
		},
		isLocked: false,
	}

	numpad := widgets.CreateNumPad(3, 4, connector)

	layout := &graphic.Layout{
		Parent: lockScreen.Screen,
	}

	spacer := widgets.CreateSpacer(1, 4, connector)

	layout.AddEntry(*spacer.Screen)
	layout.AddEntry(*numpad.Screen)

	numpad.NumberPressed = func(number int) {
		lockScreen.enteredPin = fmt.Sprintf("%s%d", lockScreen.enteredPin, number)
		println("current pin: " + lockScreen.enteredPin)

		if len(lockScreen.enteredPin) == 4 {
			lockScreen.isLocked = false
			lockScreen.enteredPin = ""
		}
	}

	lockScreen.Screen.SetLayout(layout)

	return lockScreen
}
