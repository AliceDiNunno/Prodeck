package OS

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/adapters/events/hub"
	"godeck/src/core/OS/config"
	"godeck/src/core/connector"
)

type ProdeckOS struct {
	connector connector.FrameworkConnector

	serial        string
	screenManager *ScreenManager
	deviceConfig  *config.DeviceConfig
	config        *config.Config

	buttonsState map[int]bool

	wasLocked bool
	eventHub  *hub.Hub
}

func (p *ProdeckOS) ButtonLongPressed(button int) {
	p.buttonsState[button] = false
	log.Println("Button long pressed")
}

func (p *ProdeckOS) SleepEntered() {
	println("lock screen")
	p.screenManager.lockScreen.isLocked = true
}

func (p *ProdeckOS) SleepExited() {

}

func (p *ProdeckOS) ButtonDown(button int) {
	p.buttonsState[button] = true
	p.screenManager.ButtonPressed(button)
}

func (p *ProdeckOS) ButtonUp(button int) {
	p.buttonsState[button] = false
	/*
		if button == 8 {
			p.screenManager.lockScreen.isLocked = true
		}*/
}

func initOS(connector connector.FrameworkConnector, eventhub *hub.Hub) *ProdeckOS {
	serial := connector.GetSerialNumber()
	log.WithFields(log.Fields{
		"serial": serial,
	}).Println("Init OS")

	cfg := config.NewConfig()
	deviceConfig := cfg.GetDeviceConfig(serial)

	proos := ProdeckOS{
		connector:    connector,
		serial:       serial,
		buttonsState: map[int]bool{},
		deviceConfig: deviceConfig,
		eventHub:     eventhub,
		config:       cfg,
	}

	for i := 0; i < connector.GetHeight()*connector.GetWidth(); i++ {
		proos.buttonsState[i] = false
	}

	proos.createScreenManager(connector.GetHeight(), connector.GetWidth(), &proos)

	connector.SetBrightness(deviceConfig.Brightness)

	return &proos
}

func StartOS(connector connector.FrameworkConnector, eventhub *hub.Hub) *ProdeckOS {
	proos := initOS(connector, eventhub)

	return proos
}
