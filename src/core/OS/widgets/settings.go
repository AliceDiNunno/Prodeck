package widgets

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
	color2 "image/color"
)

type Settings struct {
	*graphic.Screen

	active bool
}

func (s *Settings) ScreenName() string {
	return "Settings"
}

func (s *Settings) Tick(caller *graphic.Screen, x int, y int) {

}

func (s *Settings) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	color := color2.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}
	if s.active {
		color = color2.RGBA{
			R: 255,
			G: 98,
			B: 38,
			A: 0,
		}
	}
	/*return imagebuilder.IconWithBackground("gear", color.RGBA{
		R: 255,
		G: 98,
		B: 38,
		A: 0,
	})*/
	return imagebuilder.IconWithBackground("gear", color)
}

func (s *Settings) ButtonPressed(caller *graphic.Screen, x int, y int) {
	s.active = !s.active
	println("is active: ", s.active)
	caller.AskForRedraw(x, y)
}

func CreateSettings(connector connector.AppConnector) *Settings {
	timeWidget := &Settings{
		Screen: &graphic.Screen{
			Width:        1,
			Height:       1,
			AppConnector: connector,
		},
		active: false,
	}
	timeWidget.Screen.Interactor = timeWidget

	return timeWidget
}
