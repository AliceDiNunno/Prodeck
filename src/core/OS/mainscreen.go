package OS

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/OS/views/navigationview"
	"godeck/src/core/OS/views/statusbar"
	"godeck/src/core/connector"
)

type MainScreen struct {
	Screen *graphic.Screen

	Homescreen *HomeScreen
}

func (p *ProdeckOS) createMainScreen(connector connector.AppConnector) {
	main := &MainScreen{
		Screen: p.screenManager.screen,
	}

	statusBar := statusbar.CreateStatusBar(connector)

	navigationView := navigationview.CreateNavigationView(
		p.screenManager.screen.Width-statusBar.Screen.Width,
		p.screenManager.screen.Height, connector)

	// Create layout
	layout := &graphic.Layout{
		Parent: main.Screen,
	}

	layout.AddEntry(*statusBar.Screen)
	layout.AddEntry(*navigationView.Screen)

	main.Homescreen = CreateHomeScreen(navigationView.Screen.Width, navigationView.Screen.Height, connector)

	navigationView.SetRootView(main.Homescreen.Screen)

	main.Screen.SetLayout(layout)
}
