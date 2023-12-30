package graphic

import (
	"fmt"
	"godeck/src/core/connector"
	"image"
)

// Even though the ScreenInteractor may be the same object as the Screen, it may not be the same object memory-wise so we need to pass the pointer of the Screen
// This may be sub-optimal
type ScreenInteractor interface {
	GetButtonImage(caller *Screen, x int, y int) *image.RGBA
	ButtonPressed(caller *Screen, x int, y int)
	Tick(caller *Screen, x int, y int)
	ScreenName() string
}

type Screen struct {
	Width  int
	Height int

	Interactor ScreenInteractor
	Layout     *Layout

	buttonsWaitingForRedraw []Point

	AppConnector connector.AppConnector
}

func (s *Screen) ButtonPressed(number int) {
	x, y := s.ButtonPosition(number)

	println("SCREEN BUTTON PRESSED")

	//Check if the button is in the layout
	if s.Layout != nil {
		if !s.Layout.isSpaceAvailable(Point{x, y}) {
			s.Layout.ButtonPressed(x, y)
			println("LAYOUT BUTTON PRESSED")
			return
		}
	}

	if s.Interactor != nil {
		s.Interactor.ButtonPressed(s, x, y)
	}
}

func (s *Screen) NumberOfButtons() int {
	return s.Width * s.Height
}

func (s *Screen) ButtonPosition(button int) (int, int) {
	return button % s.Width, button / s.Width
}

func (s *Screen) ButtonNumber(x int, y int) int {
	return y*s.Width + x
}

func (s *Screen) AskForRedraw(x int, y int) {
	if s.buttonsWaitingForRedraw == nil {
		s.buttonsWaitingForRedraw = []Point{}
	}
	//check if the button is already waiting for redraw
	for _, point := range s.buttonsWaitingForRedraw {
		if point.X == x && point.Y == y {
			return
		}
	}
	s.buttonsWaitingForRedraw = append(s.buttonsWaitingForRedraw, Point{X: x, Y: y})
	s.AppConnector.Redraw()
}

func (s *Screen) ButtonsWaitingForRedraw() []Point {
	selfButtons := []Point{}

	if s.buttonsWaitingForRedraw != nil {
		selfButtons = s.buttonsWaitingForRedraw
	}

	if s.Layout != nil {
		layoutButtons := s.Layout.ButtonsWaitingForRedraw()

		selfButtons = append(selfButtons, layoutButtons...)
	}

	s.ClearButtonsWaitingForRedraw()

	return selfButtons
}

func (s *Screen) SetLayout(layout *Layout) {
	s.Layout = layout
}

func (s *Screen) GetEntryImage(x int, y int) *image.RGBA {
	if s.Layout != nil {
		return s.Layout.GetEntryImage(x, y)
	}

	//remove from the list of buttons waiting for redraw
	for i, point := range s.buttonsWaitingForRedraw {
		if point.X == x && point.Y == y {
			println("Remove from redraw list" + fmt.Sprintf("%d, %d", x, y))
			s.buttonsWaitingForRedraw = append(s.buttonsWaitingForRedraw[:i], s.buttonsWaitingForRedraw[i+1:]...)
			break
		}
	}
	return s.Interactor.GetButtonImage(s, x, y)
}

func (s *Screen) OSTick(x int, y int) {
	if s.Layout != nil {
		s.Layout.Tick(x, y)
	}

	if s.Interactor != nil {
		s.Interactor.Tick(s, x, y)
	}
}

func (s *Screen) ClearButtonsWaitingForRedraw() {
	s.buttonsWaitingForRedraw = []Point{}
}
