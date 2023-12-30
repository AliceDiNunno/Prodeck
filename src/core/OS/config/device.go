package config

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DeviceConfig struct {
	gorm.Model

	Serial     string `gorm:"uniqueIndex"`
	Brightness int
}

func (c *Config) initialDeviceConfig(serial string) DeviceConfig {
	return DeviceConfig{
		Serial:     serial,
		Brightness: 50,
	}
}

func (c *Config) GetDeviceConfig(serial string) *DeviceConfig {
	//Check if device exists
	var device DeviceConfig
	err := c.db.First(&device, "serial = ?", serial).Error

	if device.Serial == "" || err != nil {
		log.Println("Creating new device config")
		device = c.initialDeviceConfig(serial)
		c.db.Create(&device)
	}

	return &device
}

func (c *Config) UpdateDeviceConfig(device *DeviceConfig) {
	c.db.Save(&device)
}
