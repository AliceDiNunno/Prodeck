package statusbar

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/OS/widgets"
	"godeck/src/core/connector"
	"image"
)

type StatusBar struct {
	Screen *graphic.Screen

	// Widgets
	home           *widgets.Home
	time           *widgets.Time
	clientSelector *widgets.ClientSelector
	gear           *widgets.Settings
}

func (n *StatusBar) ScreenName() string {
	return "StatusBar"
}

func (n *StatusBar) Tick(caller *graphic.Screen, x int, y int) {
	n.home.Tick(caller, 0, 0)
	n.time.Tick(caller, 0, 1)
}

func (n *StatusBar) ButtonPressed(caller *graphic.Screen, x int, y int) {
	caller.AskForRedraw(x, y)
	log.Warn("redrawstusbar ", x, y)
}

func (n *StatusBar) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	return nil
}

func CreateStatusBar(connector connector.AppConnector) *StatusBar {
	statusbar := &StatusBar{
		Screen: &graphic.Screen{
			Width:        1,
			Height:       4,
			AppConnector: connector,
		},
	}

	statusbar.Screen.Interactor = statusbar

	home := widgets.CreateHome(connector)
	time := widgets.CreateTime(connector)
	clientSelector := widgets.CreateClientSelector(connector)
	gear := widgets.CreateSettings(connector)

	statusbar.home = home
	statusbar.time = time
	statusbar.clientSelector = clientSelector
	statusbar.gear = gear

	// Create layout
	layout := &graphic.Layout{
		Parent: statusbar.Screen,
	}

	layout.AddEntry(*home.Screen)
	layout.AddEntry(*time.Screen)
	layout.AddEntry(*clientSelector.Screen)
	layout.AddEntry(*gear.Screen)

	statusbar.Screen.SetLayout(layout)

	return statusbar
}
