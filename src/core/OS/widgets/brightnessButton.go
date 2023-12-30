package widgets

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
	"math/rand"
)

type BrightnessButton struct {
	*graphic.Screen
	appConnector connector.AppConnector
}

func (s *BrightnessButton) ScreenName() string {
	return "Brightness"
}

func (s *BrightnessButton) Tick(caller *graphic.Screen, x int, y int) {

}

func (s *BrightnessButton) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	println("ASK TO UPDATE")
	text := fmt.Sprintf("%d%%", s.appConnector.GetBrightness())
	return imagebuilder.IconWithText("sun.max.fill", text)
}

func (s *BrightnessButton) ButtonPressed(caller *graphic.Screen, x int, y int) {
	s.appConnector.SetBrightness(rand.Intn(100))

	spew.Dump("ask for redraw", x+1, y)

	caller.AskForRedraw(x+1, y)
	caller.AskForRedraw(x, y)
}

func CreateBrightnessButton(appConnector connector.AppConnector) *BrightnessButton {
	timeWidget := &BrightnessButton{
		Screen: &graphic.Screen{
			Width:        1,
			Height:       1,
			AppConnector: appConnector,
		},
		appConnector: appConnector,
	}
	timeWidget.Screen.Interactor = timeWidget

	return timeWidget
}
