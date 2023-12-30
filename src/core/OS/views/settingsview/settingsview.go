package settingsview

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"image"
)

type SettingsView struct {
	*graphic.Screen
}

func (n *SettingsView) ScreenName() string {
	return "SettingsView"
}

func (n *SettingsView) ButtonPressed(caller *graphic.Screen, x int, y int) {
	caller.AskForRedraw(x, y)
	log.Warn("redraw ", x, y)
}

func (n *SettingsView) Tick(caller *graphic.Screen, x int, y int) {

}

func (n *SettingsView) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	return nil
}

func CreateNavigationView(width int, height int, connector connector.AppConnector) *SettingsView {
	settingsView := &SettingsView{
		Screen: &graphic.Screen{
			Width:        width,
			Height:       height,
			AppConnector: connector,
		},
	}
	settingsView.Screen.Interactor = settingsView

	return settingsView
}
