package connector

type AppConnector interface {
	SetBrightness(value int)
	GetBrightness() int
	Redraw()
}
