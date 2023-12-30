package Framework

import (
	"github.com/AliceDiNunno/hid"
	"godeck/src/adapters/streamdeck"
	eventDomain "godeck/src/domain/events"
)

func (p *ProdeckFramework) startDiscovery() {
	streamdeck.NewDiscovery(p.eventHub).Discover()
}

func (p *ProdeckFramework) initDevice() {
	p.initializeButtons()
	go p.pollButtonState(p.currentDevice.HID)
	p.setBrightness(25, false, true)
}

func (p *ProdeckFramework) pollButtonState(dev *hid.Device) {
	//32 button state
	buttonStates := make([]bool, 32)

	for i := 0; i < 32; i++ {
		buttonStates[i] = false
	}

	for {
		var arr = make([]byte, 32+4)

		read, err := dev.Read(arr)
		if err != nil {
			println("error", err)
			return
		}

		if read > 0 {
			for i := 0; i < 32; i++ {
				currentButtonState := arr[i+4] == 1

				if currentButtonState != buttonStates[i] {
					buttonStates[i] = currentButtonState

					p.eventHub.Publish(eventDomain.ButtonStateChangedEvent, map[string]interface{}{
						"button": i,
						"state":  currentButtonState,
					})
				}
			}
		}
	}
}