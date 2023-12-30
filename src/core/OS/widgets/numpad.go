package widgets

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
	"image/color"
)

type Numpad struct {
	*graphic.Screen
	NumberPressed func(value int)
}

func (s *Numpad) ScreenName() string {
	return "Numpad"
}

func (s *Numpad) Tick(caller *graphic.Screen, x int, y int) {
}

const (
	classicNumpadForm = iota
	unsupportedNumpadForm
)

func (s *Numpad) getForm() int {
	if s.Width == 3 && s.Height == 4 {
		return classicNumpadForm
	}
	return unsupportedNumpadForm
}

func (s *Numpad) numpadButtonForForm(x int, y int) int {
	switch s.getForm() {
	case classicNumpadForm:
		button := s.Screen.ButtonNumber(x, y) + 1
		if button == 10 || button == 12 {
			return -1
		} else if button == 11 {
			return 0
		}
		return button
	}
	return s.Screen.ButtonNumber(x, y)
}

func (s *Numpad) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	if s.numpadButtonForForm(x, y) == -1 {
		return imagebuilder.CreateImage(96, 96, color.RGBA{0, 0, 0, 0})
	}
	return imagebuilder.Icon(fmt.Sprintf("%d", s.numpadButtonForForm(x, y)))
}

func (s *Numpad) ButtonPressed(caller *graphic.Screen, x int, y int) {
	numpadButton := s.numpadButtonForForm(x, y)

	if numpadButton == -1 {
		return
	}

	println("Numpad button pressed: ", numpadButton)
	if s.NumberPressed != nil {
		s.NumberPressed(numpadButton)
		println("Redraw numpad")
	}
}

func CreateNumPad(width int, height int, connector connector.AppConnector) *Numpad {
	//must have at least 10 buttons
	if width*height < 10 {
		log.Errorln("Numpad must have at least 10 buttons")
		return nil
	}

	statusBar := &Numpad{
		Screen: &graphic.Screen{
			Width:        width,
			Height:       height,
			AppConnector: connector,
		},
	}
	statusBar.Screen.Interactor = statusBar

	return statusBar
}
