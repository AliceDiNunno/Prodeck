package streamdeck

import (
	"github.com/AliceDiNunno/hid"
	log "github.com/sirupsen/logrus"
	"godeck/src/adapters/events/hub"
	device2 "godeck/src/domain/device"
	eventDomain "godeck/src/domain/events"
)

type Discovery struct {
	hub     *hub.Hub
	devices []StreamDeck
}

func (discovery *Discovery) Discover() {
	for {
		devices := hid.Enumerate(device2.VendorId, device2.StreamDeckXLProductId)

		knownDevices := make([]StreamDeck, len(discovery.devices))
		copy(knownDevices, discovery.devices)

		for i, device := range devices {
			if !discovery.isDeviceKnown(device) {
				if device.Serial == "" {
					log.Errorln("Device has no serial, skipping")
					continue
				}

				hidInstance, err := device.Open()

				if err == nil {
					deck := NewStreamDeck(device2.NewStreamDeckXL(), device.Serial, hidInstance)
					discovery.devices = append(discovery.devices, *deck)

					discovery.hub.Publish(eventDomain.DeviceConnectedEvent, eventDomain.EventData{
						"device":   device,
						"instance": hidInstance,
					})
				}
			} else {
				knownDevices = append(knownDevices[:i], knownDevices[i+1:]...)
			}
		}

		for _, device := range knownDevices {
			discovery.hub.Publish(eventDomain.DeviceDisconnectedEvent, eventDomain.EventData{
				"device": device,
			})

			for j, knownDevice := range discovery.devices {
				if knownDevice.Serial == device.Serial {
					discovery.devices = append(discovery.devices[:j], discovery.devices[j+1:]...)
				}
			}
		}
	}
}

func (discovery *Discovery) isDeviceKnown(scannedDevice hid.DeviceInfo) bool {
	for _, device := range discovery.devices {
		serial := device.Serial
		if serial == scannedDevice.Serial {
			return true
		}
	}
	return false
}

func NewDiscovery(hub *hub.Hub) *Discovery {
	return &Discovery{
		hub: hub,
	}
}