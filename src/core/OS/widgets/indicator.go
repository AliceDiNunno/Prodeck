package widgets

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
	"image/color"
)

type Indicator struct {
	*graphic.Screen
	color color.RGBA
}

func (s *Indicator) ScreenName() string {
	return "Indicator"
}

func (s *Indicator) Tick(caller *graphic.Screen, x int, y int) {

}

func (s *Indicator) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	return imagebuilder.CreateImage(96, 96, s.color)
}

func (s *Indicator) SetColor(c color.RGBA) {
	s.color = c
	s.AskForRedraw(0, 0)
}

func (s *Indicator) ButtonPressed(caller *graphic.Screen, x int, y int) {

}

func CreateIndicator(connector connector.AppConnector) *Indicator {
	indicator := &Indicator{
		Screen: &graphic.Screen{
			Width:        1,
			Height:       1,
			AppConnector: connector,
		},
	}
	indicator.Screen.Interactor = indicator

	return indicator
}
