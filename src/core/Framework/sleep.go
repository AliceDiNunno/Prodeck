package Framework

import (
	log "github.com/sirupsen/logrus"
	eventDomain "godeck/src/domain/events"
	"time"
)

func (p *ProdeckFramework) setupSleepEvent() {
	p.eventHub.Subscribe(eventDomain.DeviceWillSleepEvent, func(topic eventDomain.Event, data eventDomain.EventData) {
		log.Println("Device will sleep")

		p.isSleeping = true
		go p.setBrightness(0, true, false)
		(*p.currentOS).SleepEntered()

		p.eventHub.CancelPublishLater(eventDomain.DeviceWillSleepEvent)
	})
}

func (p *ProdeckFramework) prepareForSleep() {
	log.Println("Device will sleep in an hour")
	p.eventHub.PublishLater(eventDomain.DeviceWillSleepEvent, nil, time.Hour)
}

func (p *ProdeckFramework) wakeUpFromSleep() bool {
	p.eventHub.CancelPublishLater(eventDomain.DeviceWillSleepEvent)
	if p.isSleeping {
		p.isSleeping = false
		go p.setBrightness(p.currentBrightness, true, false)
		(*p.currentOS).SleepExited()
		return true
	}
	return false
}
