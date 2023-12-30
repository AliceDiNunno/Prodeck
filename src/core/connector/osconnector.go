package connector

import (
	"image"
	"image/color"
)

type OSConnector interface {
	ButtonDown(button int)
	ButtonUp(button int)
	ButtonLongPressed(button int)
	SleepEntered()
	SleepExited()
}

type FrameworkConnector interface {
	SetBrightness(value int)
	SetButtonImage(button int, image image.Image)
	SetButtonColor(button int, color color.Color)
	ForceSleep()
	GetSerialNumber() string
	GetWidth() int
	GetHeight() int
}

type OSBuilder func() *OSConnector