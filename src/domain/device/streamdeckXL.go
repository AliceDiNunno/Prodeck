package device

const StreamDeckXLProductId = 0x006c

type StreamDeckXL struct {
}

func (s StreamDeckXL) Rows() int {
	return 4
}

func (s StreamDeckXL) Columns() int {
	return 8
}

func (s StreamDeckXL) ClearCommand() []byte {
	return []byte{0x03, 0x02}
}

func (s StreamDeckXL) BrightnessCommand(brightness int) []byte {
	if brightness <= 0 {
		brightness = 0
	}

	if brightness >= 100 {
		brightness = 100
	}

	return []byte{0x03, 0x08, byte(brightness)}
}

func (s StreamDeckXL) ImageReportLength() int {
	return 1024
}

func (s StreamDeckXL) ImageHeaderLength() int {
	return 8
}

func (s StreamDeckXL) ImageFormat() string {
	return "jpeg"
}

func NewStreamDeckXL() DeckDevice {
	return StreamDeckXL{}
}
