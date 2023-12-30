package navigationview

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
	"image/color"
)

type NavigationHistory struct {
	previousView *NavigationHistory
	currentView  *graphic.Screen
	nextView     *NavigationHistory
}

type NavigationView struct {
	*graphic.Screen
	history *NavigationHistory
}

func (n *NavigationView) ScreenName() string {
	return "NavigationView"
}

func (n *NavigationView) ButtonPressed(caller *graphic.Screen, x int, y int) {
	caller.AskForRedraw(x, y)
	log.Warn("redraw ", x, y)

	n.history.currentView.Interactor.ButtonPressed(n.history.currentView, x, y)
}

func (n *NavigationView) Tick(caller *graphic.Screen, x int, y int) {

}

func (n *NavigationView) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	currentView := n.history.currentView
	if currentView == nil || currentView.Interactor == nil {
		return imagebuilder.CreateImage(96, 96, color.RGBA{0, 0, 0, 0})
	}

	return currentView.Interactor.GetButtonImage(currentView, x, y)
}

func (n *NavigationView) SetRootView(view *graphic.Screen) {
	n.history = &NavigationHistory{
		previousView: nil,
		currentView:  view,
		nextView:     nil,
	}
}

func CreateNavigationView(width int, height int, connector connector.AppConnector) *NavigationView {
	navigationView := &NavigationView{
		Screen: &graphic.Screen{
			Width:        width,
			Height:       height,
			AppConnector: connector,
		},
	}
	navigationView.Screen.Interactor = navigationView

	return navigationView
}
