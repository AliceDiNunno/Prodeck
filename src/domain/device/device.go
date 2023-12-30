package device

const VendorId = 0x0fd9

type DeckDevice interface {
	Rows() int
	Columns() int

	ClearCommand() []byte
	BrightnessCommand(brightness int) []byte

	ImageReportLength() int
	ImageHeaderLength() int
	ImageFormat() string
}
