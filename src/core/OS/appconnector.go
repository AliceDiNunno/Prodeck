package OS

func (p *ProdeckOS) SetBrightness(value int) {
	println("CALL SET BRIGHTNESS WITH VALUE: ", value)

	p.deviceConfig.Brightness = value
	p.connector.SetBrightness(value)
	p.config.UpdateDeviceConfig(p.deviceConfig)
}

func (p *ProdeckOS) GetBrightness() int {
	println("CALL GET BRIGHTNESS WITH VALUE: ", p.deviceConfig.Brightness)
	return p.deviceConfig.Brightness
}

func (p *ProdeckOS) Redraw() {
	p.eventHub.Publish("redraw", nil)
}
