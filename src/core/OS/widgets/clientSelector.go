package widgets

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
)

type ClientSelector struct {
	*graphic.Screen
}

func (s *ClientSelector) ScreenName() string {
	return "ClientSelector"
}

func (s *ClientSelector) Tick(caller *graphic.Screen, x int, y int) {

}

func (s *ClientSelector) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	return imagebuilder.Icon("display")
}

func (s *ClientSelector) ButtonPressed(caller *graphic.Screen, x int, y int) {

}

func CreateClientSelector(connector connector.AppConnector) *ClientSelector {
	timeWidget := &ClientSelector{
		Screen: &graphic.Screen{
			Width:        1,
			Height:       1,
			AppConnector: connector,
		},
	}
	timeWidget.Screen.Interactor = timeWidget

	return timeWidget
}
