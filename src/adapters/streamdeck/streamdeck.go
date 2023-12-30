package streamdeck

import (
	"bytes"
	"github.com/AliceDiNunno/hid"
	log "github.com/sirupsen/logrus"
	"godeck/src/domain/device"
	"image"
	"image/jpeg"
)

type StreamDeck struct {
	Device     device.DeckDevice
	Serial     string
	Brightness int
	HID        *hid.Device
	Height     int
	Width      int
}

func NewStreamDeck(device device.DeckDevice, serial string, instance *hid.Device) *StreamDeck {
	return &StreamDeck{
		Device:     device,
		Serial:     serial,
		Brightness: -1,
		HID:        instance,

		Height: 4,
		Width:  8,
	}
}

func padding(size int) []byte {
	padding := make([]byte, size)
	for i := 0; i < size; i++ {
		padding[i] = 0x00
	}
	return padding
}

func splitIntToBytes(num int) []byte {
	return []byte{byte(num >> 8), byte(num)}
}

func (deck *StreamDeck) SetBrightness(brightness int) {
	feature := []byte{0x03, 0x08, byte(brightness)}

	_, err := deck.HID.SendFeatureReport(feature)

	if err != nil {
		println("Error sending feature report", err)
		return
	}
	deck.Brightness = brightness
}

func (deck *StreamDeck) SetButtonImage(key int, img *image.RGBA) {
	// Encode as PNG.
	if img == nil {
		return
	}
	var buf bytes.Buffer

	jpeg.Encode(&buf, img, nil)

	if !(key >= 0 && key < 32) {
		panic("Invalid key index.")
	}

	image := buf.Bytes()

	pageNumber := 0
	bytesRemaining := len(image)

	//While there are still bytes to send
	ImageReportLength := 1024

	for bytesRemaining > 0 {
		thisLength := bytesRemaining
		ImageReportHeaderLength := 8
		ImageReportPayloadLength := ImageReportLength - ImageReportHeaderLength

		if thisLength > ImageReportPayloadLength {
			thisLength = ImageReportPayloadLength
		}

		bytesSent := pageNumber * ImageReportPayloadLength

		isLastPage := byte(0)
		if thisLength == bytesRemaining {
			isLastPage = 1
		}

		lengthBytes := splitIntToBytes(thisLength)
		pageNumberBytes := splitIntToBytes(pageNumber)

		header := []byte{
			0x02,
			0x07,
			byte(key),
			isLastPage,
			lengthBytes[1],
			lengthBytes[0],
			pageNumberBytes[1],
			pageNumberBytes[0],
		}

		payload := append(header, image[bytesSent:bytesSent+thisLength]...)
		pdng := padding(ImageReportLength - len(payload))
		_, err := deck.HID.Write(append(payload, pdng...))
		if err != nil {
			log.Warn("Error writing to device", err)
			return
		}

		bytesRemaining = bytesRemaining - thisLength
		pageNumber = pageNumber + 1
	}
}
