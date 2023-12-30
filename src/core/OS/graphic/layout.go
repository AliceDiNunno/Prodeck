package graphic

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"godeck/src/core/imagebuilder"
	"image"
	"image/color"
)

type Layout struct {
	Parent  *Screen
	entries []*LayoutEntry
}

type ScreenUsecase interface {
	GetButtonColor(x int, y int) color.RGBA
}

type LayoutEntry struct {
	screen *Screen
	origin Point
}

func (l *Layout) ButtonPressed(x int, y int) {
	println("LAYOUT button pressed", x, y)
	for _, entry := range l.entries {
		entryHeight := entry.screen.Height
		entryWidth := entry.screen.Width
		entryOrigin := entry.origin

		if x >= entryOrigin.X && x < entryOrigin.X+entryWidth {
			if y >= entryOrigin.Y && y < entryOrigin.Y+entryHeight {
				println("REAL button pressed", entry.screen.ButtonNumber(x-entryOrigin.X, y-entryOrigin.Y))
				entry.screen.ButtonPressed(entry.screen.ButtonNumber(x-entryOrigin.X, y-entryOrigin.Y))
			}
		}
	}
}

func (l *Layout) isSpaceAvailable(origin Point) bool {
	for _, entry := range l.entries {
		entryHeight := entry.screen.Height
		entryWidth := entry.screen.Width
		entryOrigin := entry.origin

		if origin.X >= entryOrigin.X && origin.X < entryOrigin.X+entryWidth {
			if origin.Y >= entryOrigin.Y && origin.Y < entryOrigin.Y+entryHeight {
				return false
			}
		}
	}

	return true
}

func (l *Layout) canFitScreenInLayout(entry Screen) (error, Point) {
	canvasHeight := l.Parent.Height
	canvasWidth := l.Parent.Width

	entryHeight := entry.Height
	entryWidth := entry.Width

	if entryHeight > canvasHeight || entryWidth > canvasWidth {
		return errors.New("view is bigger than canvas"), Point{}
	}

	//For each point in the canvas we check if the screen can fit and is available
	log.Println("check for origins", canvasHeight, entryHeight, canvasWidth, entryWidth)

	for x := 0; x <= canvasWidth-entryWidth; x++ {
		for y := 0; y <= canvasHeight-entryHeight; y++ {

			isAvailable := true

			for entryX := x; entryX < x+entryWidth; entryX++ {
				for entryY := y; entryY < y+entryHeight; entryY++ {
					if !l.isSpaceAvailable(Point{entryX, entryY}) {
						isAvailable = false
					}
				}
			}

			if isAvailable {
				return nil, Point{x, y}
			}
		}
	}

	return errors.New("No available space is found"), Point{}
}

func (l *Layout) AddEntry(screenValue Screen) {
	e, p := l.canFitScreenInLayout(screenValue)

	if e != nil {
		log.Errorln(e)
		return
	}

	l.entries = append(l.entries, &LayoutEntry{
		screen: &screenValue,
		origin: p,
	})
}

func (l *Layout) GetEntryImage(x int, y int) *image.RGBA {
	for _, entry := range l.entries {
		entryHeight := entry.screen.Height
		entryWidth := entry.screen.Width
		entryOrigin := entry.origin

		if x >= entryOrigin.X && x < entryOrigin.X+entryWidth {
			if y >= entryOrigin.Y && y < entryOrigin.Y+entryHeight {
				if entry.screen.Interactor != nil {
					return entry.screen.GetEntryImage(x-entryOrigin.X, y-entryOrigin.Y)
				}
			}
		}
	}
	blank := imagebuilder.CreateImage(96, 96, color.RGBA{0, 0, 0, 0})
	return blank
}

func (l *Layout) ButtonsWaitingForRedraw() []Point {
	points := []Point{}

	subViews := l.entries
	for _, entry := range subViews {
		waiting := entry.screen.ButtonsWaitingForRedraw()
		for _, point := range waiting {
			points = append(points, Point{
				X: point.X + entry.origin.X,
				Y: point.Y + entry.origin.Y,
			})
		}
	}

	return points
}

func (l *Layout) Tick(x int, y int) {
	for _, entry := range l.entries {
		entryHeight := entry.screen.Height
		entryWidth := entry.screen.Width
		entryOrigin := entry.origin

		if x >= entryOrigin.X && x < entryOrigin.X+entryWidth {
			if y >= entryOrigin.Y && y < entryOrigin.Y+entryHeight {
				entry.screen.OSTick(x-entryOrigin.X, y-entryOrigin.Y)
			}
		}
	}
}
