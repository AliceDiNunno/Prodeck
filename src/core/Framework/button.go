package Framework

import (
	log "github.com/sirupsen/logrus"
	eventDomain "godeck/src/domain/events"
	"time"
)

func (p *ProdeckFramework) setupButtonEvent() {
	p.eventHub.Subscribe(eventDomain.ButtonStateChangedEvent, func(topic eventDomain.Event, data eventDomain.EventData) {
		button := data["button"].(int)
		pressed := data["state"].(bool)

		if p.isSleeping && button == 31 && pressed && p.wakeUpFromSleep() {
			p.lastInteractionTime = time.Now().Unix()
			p.prepareForSleep()
			return
		}

		log.WithFields(log.Fields(data)).Info(topic)

		//setting last interaction time only if the button is released

		if !pressed {
			p.lastInteractionTime = time.Now().Unix()

			p.buttonPressedUp(button)

			p.prepareForSleep()
		} else {
			p.buttonPressedDown(button)
		}
	})
}

func (p *ProdeckFramework) buttonPressedDown(button int) {
	p.outlineCurrentButton(button)

	p.pressedTime.buttons[button] = time.Now().Unix()

	if p.currentOS != nil {
		(*p.currentOS).ButtonDown(button)
	}
}

func (p *ProdeckFramework) buttonPressedUp(button int) {
	image, ok := p.imageCache.images[button]
	if ok {
		p.setButtonImage(button, image, false)
	}

	if p.currentOS != nil {
		if p.pressedTime.buttons[button] != 0 && time.Now().Unix()-p.pressedTime.buttons[button] > 2 {
			(*p.currentOS).ButtonLongPressed(button)
		} else {
			(*p.currentOS).ButtonUp(button)
		}
	}
}