package OS

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	eventDomain "godeck/src/domain/events"
	"time"
)

type ScreenManager struct {
	lockScreen *LockScreen
	screen     *graphic.Screen
}

func (m ScreenManager) ButtonPressed(button int) {
	if m.lockScreen.isLocked {
		m.lockScreen.Screen.ButtonPressed(button)
	} else {
		m.screen.ButtonPressed(button)
	}
}

func (p *ProdeckOS) createScreenManager(height int, width int, connector connector.AppConnector) {
	p.screenManager = &ScreenManager{
		screen: &graphic.Screen{
			Width:        width,
			Height:       height,
			AppConnector: connector,
		},
		lockScreen: CreateLockScreen(width, height, connector),
	}

	p.createMainScreen(connector)

	p.startRefreshLoop()
}

func (p *ProdeckOS) currentScreen() *graphic.Screen {
	if p.screenManager.lockScreen.isLocked {
		return p.screenManager.lockScreen.Screen
	} else {
		return p.screenManager.screen
	}
}

func (p *ProdeckOS) refreshAllScreen() {
	screen := p.currentScreen()

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			img := screen.GetEntryImage(x, y)
			p.connector.SetButtonImage(screen.ButtonNumber(x, y), img)
		}
	}
}

func (p *ProdeckOS) startRefreshLoop() {
	p.refreshAllScreen()

	p.eventHub.Subscribe("refresh", func(topic eventDomain.Event, data eventDomain.EventData) {
		go func() {
			screen := p.currentScreen()

			waitingForRedrew := screen.ButtonsWaitingForRedraw()

			if p.wasLocked && !p.screenManager.lockScreen.isLocked ||
				!p.wasLocked && p.screenManager.lockScreen.isLocked {
				p.wasLocked = p.screenManager.lockScreen.isLocked

				p.refreshAllScreen()
			}

			for x := 0; x < screen.Width; x++ {
				for y := 0; y < screen.Height; y++ {
					screen.OSTick(x, y)
				}
			}

			for _, button := range waitingForRedrew {
				//If button is pressed, don't redraw it now
				log.Info(button, " is waiting to be drawn")

				if p.buttonsState[screen.ButtonNumber(button.X, button.Y)] {
					continue
				}
				entryImage := screen.GetEntryImage(button.X, button.Y)
				if entryImage == nil {
					continue
				}
				p.connector.SetButtonImage(screen.ButtonNumber(button.X, button.Y), entryImage)

			}

			screen.ClearButtonsWaitingForRedraw()
		}()
	})

	p.wasLocked = p.screenManager.lockScreen.isLocked

	go func() {
		for _ = range time.Tick(time.Second / 2) {
			p.eventHub.Publish("refresh", nil)
		}
	}()
}
