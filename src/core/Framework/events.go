package Framework

import (
	"github.com/AliceDiNunno/hid"
	log "github.com/sirupsen/logrus"
	eventDomain "godeck/src/domain/events"
)

func (p *ProdeckFramework) setupHIDEvents() {
	p.eventHub.Subscribe(eventDomain.DeviceConnectedEvent, func(topic eventDomain.Event, data eventDomain.EventData) {
		deviceInfo := data["device"].(hid.DeviceInfo)
		device := data["instance"].(*hid.Device)

		p.deviceConnected(deviceInfo, device)
	})

	p.eventHub.Subscribe(eventDomain.DeviceDisconnectedEvent, func(topic eventDomain.Event, data eventDomain.EventData) {
		log.Println("Device disconnected ", data["device"].(hid.DeviceInfo).Serial)

		if p.currentDevice != nil {
			p.currentDevice = nil
		}
	})
}

func (p *ProdeckFramework) setupEvents() {
	p.setupHIDEvents()
	p.setupSleepEvent()
	p.setupButtonEvent()
}