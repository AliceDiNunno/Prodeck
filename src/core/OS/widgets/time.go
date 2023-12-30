package widgets

import (
	"fmt"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/connector"
	"godeck/src/core/imagebuilder"
	"image"
	"time"
)

type Time struct {
	*graphic.Screen
	lastMinute int
}

func (s *Time) ScreenName() string {
	return "Time"
}

func (s *Time) Tick(caller *graphic.Screen, x int, y int) {
	if s.shouldUpdate() {
		println("Ask for redraw")
		caller.AskForRedraw(x, y)
	}
}

func (s *Time) shouldUpdate() bool {
	//init the loc
	loc, _ := time.LoadLocation("Europe/Paris")

	//set timezone,
	now := time.Now().In(loc)

	minute := now.Minute()

	return s.lastMinute != minute
}

func (s *Time) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	//init the loc
	loc, _ := time.LoadLocation("Europe/Paris")

	//set timezone,
	now := time.Now().In(loc)

	hour := now.Hour()
	minute := now.Minute()

	println(fmt.Sprintf("%d vs %d", s.lastMinute, minute))

	s.lastMinute = minute

	s.AskForRedraw(x, y)

	text := fmt.Sprintf("%02d:%02d", hour, minute)

	println("Current time: " + text)

	return imagebuilder.IconWithText("clock", text)
}

func (s *Time) ButtonPressed(caller *graphic.Screen, x int, y int) {

}

func CreateTime(connector connector.AppConnector) *Time {
	timeWidget := &Time{
		Screen: &graphic.Screen{
			Width:        1,
			Height:       1,
			AppConnector: connector,
		},
	}
	timeWidget.Screen.Interactor = timeWidget

	return timeWidget
}
