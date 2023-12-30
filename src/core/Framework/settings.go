package Framework

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (p *ProdeckFramework) setBrightness(value int, fade bool, save bool) {
	if save && value > 0 {
		p.currentBrightness = value
	}

	currentBrightness := p.currentDevice.Brightness

	if currentBrightness == value {
		return
	}
	log.Println("Setting brightness to ", value)

	if fade {
		go func() {
			if currentBrightness < value {
				for i := currentBrightness; i < value; i++ {
					p.currentDevice.SetBrightness(i)
					time.Sleep(10 * time.Millisecond)
				}
			} else {
				for i := currentBrightness; i > value; i-- {
					p.currentDevice.SetBrightness(i)
					time.Sleep(10 * time.Millisecond)
				}
			}
		}()
	} else {
		p.currentDevice.SetBrightness(value)
	}
}
