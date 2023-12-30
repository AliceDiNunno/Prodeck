package Framework

import (
	"github.com/AliceDiNunno/hid"
	log "github.com/sirupsen/logrus"
	"godeck/src/adapters/events/hub"
	"godeck/src/adapters/streamdeck"
	"godeck/src/core/connector"
	deviceDomain "godeck/src/domain/device"
	"image"
	"image/color"
)

type ButtonImageCache struct {
	images map[int]*image.RGBA
}

type ButtonPressedTime struct {
	buttons map[int]int64
}

type ProdeckFramework struct {
	eventHub *hub.Hub

	currentDevice *streamdeck.StreamDeck
	currentOS     *connector.OSConnector

	imageCache  ButtonImageCache
	pressedTime ButtonPressedTime

	lastInteractionTime int64
	isSleeping          bool
	currentBrightness   int

	osBuilder func() connector.OSConnector
}

func (p *ProdeckFramework) initializeButtons() {
	for i := 0; i < 32; i++ {
		p.SetButtonColor(i, color.Black)
	}
}

func (p *ProdeckFramework) deviceConnected(deviceInfo hid.DeviceInfo, instance *hid.Device) {
	log.WithFields(log.Fields{
		"sn": deviceInfo.Serial,
	}).Println("Device connected")

	p.currentDevice = streamdeck.NewStreamDeck(deviceDomain.NewStreamDeckXL(), deviceInfo.Serial, instance)

	p.initDevice()
	p.prepareForSleep()

	os := p.osBuilder()

	p.currentOS = &os
}

func (p *ProdeckFramework) Start() {
	log.Println("Starting Framework")

	p.setupEvents()

	go p.startDiscovery()
}

// TODO: there is only a basic support for one StreamDeck XL for now
func NewProdeckFramework(eventHub *hub.Hub, builder func() connector.OSConnector) *ProdeckFramework {
	return &ProdeckFramework{
		eventHub: eventHub,
		imageCache: ButtonImageCache{
			images: make(map[int]*image.RGBA),
		},
		pressedTime: ButtonPressedTime{
			buttons: make(map[int]int64),
		},
		currentOS: nil,
		osBuilder: builder,
	}
}