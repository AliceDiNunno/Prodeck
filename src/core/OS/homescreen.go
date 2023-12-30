package OS

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/OS/widgets"
	"godeck/src/core/connector"
	"image"
)

type HomeScreen struct {
	*graphic.Screen

	brightness   *widgets.BrightnessButton
	appConnector connector.AppConnector
}

func (n *HomeScreen) Tick(caller *graphic.Screen, x int, y int) {

}

func (n *HomeScreen) ScreenName() string {
	return "homescreen"
}

func (n *HomeScreen) ButtonPressed(caller *graphic.Screen, x int, y int) {
	println("HOME SCREEN BUTTON PRESSED")
	caller.Layout.ButtonPressed(x, y)
}

func (n *HomeScreen) GetButtonImage(caller *graphic.Screen, x int, y int) *image.RGBA {
	return caller.Layout.GetEntryImage(x, y)
}

func CreateHomeScreen(width int, height int, connector connector.AppConnector) *HomeScreen {
	homeScreen := &HomeScreen{
		Screen: &graphic.Screen{
			Width:        width,
			Height:       height,
			AppConnector: connector,
		},
		appConnector: connector,
	}

	homeScreen.Screen.Interactor = homeScreen

	layout := &graphic.Layout{
		Parent: homeScreen.Screen,
	}

	homeScreen.brightness = widgets.CreateBrightnessButton(homeScreen.appConnector)

	layout.AddEntry(*homeScreen.brightness.Screen)

	homeScreen.Screen.SetLayout(layout)

	return homeScreen
}
