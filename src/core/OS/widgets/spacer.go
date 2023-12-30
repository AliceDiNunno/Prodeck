package widgets

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
	"image/color"
)

type Spacer struct {
	*graphic.Screen
}

func (s *Spacer) ScreenName() string {
	return "Spacer"
}

func (s *Spacer) Tick(caller *graphic.Screen, x int, y int) {
}

func (s *Spacer) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	return imagebuilder.CreateImage(96, 96, color.RGBA{0, 0, 0, 0})
}

func (s *Spacer) ButtonPressed(caller *graphic.Screen, x int, y int) {

}

func CreateSpacer(width int, height int, connector connector.AppConnector) *Spacer {
	spacer := &Spacer{
		Screen: &graphic.Screen{
			Width:        width,
			Height:       height,
			AppConnector: connector,
		},
	}
	spacer.Screen.Interactor = spacer

	return spacer
}
