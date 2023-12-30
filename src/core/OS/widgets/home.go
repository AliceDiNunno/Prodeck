package widgets

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
)

type Home struct {
	*graphic.Screen
}

func (s *Home) ScreenName() string {
	return "Home"
}

func (s *Home) Tick(caller *graphic.Screen, x int, y int) {

}

func (s *Home) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	return imagebuilder.Icon("icon")
}

func (s *Home) ButtonPressed(caller *graphic.Screen, x int, y int) {

}

func CreateHome(connector connector.AppConnector) *Home {
	timeWidget := &Home{
		Screen: &graphic.Screen{
			Width:        1,
			Height:       1,
			AppConnector: connector,
		},
	}
	timeWidget.Screen.Interactor = timeWidget

	return timeWidget
}
