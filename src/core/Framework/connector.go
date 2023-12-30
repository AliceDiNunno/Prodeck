package Framework

import "image"

func (p *ProdeckFramework) SetBrightness(value int) {
	p.setBrightness(value, true, true)

}

func (p *ProdeckFramework) SetButtonImage(button int, image image.Image) {
	if image == nil {
		return
	}

	rgba := imageToRGBA(image)

	p.setButtonImage(button, rgba, true)
}

func (p *ProdeckFramework) ForceSleep() {

}

func (p *ProdeckFramework) GetSerialNumber() string {
	return p.currentDevice.Serial
}

func (p *ProdeckFramework) GetWidth() int {
	return p.currentDevice.Width
}

func (p *ProdeckFramework) GetHeight() int {
	return p.currentDevice.Height
}
