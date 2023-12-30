package graphic

import Testing "testing"

var screen = Screen{Width: 8, Height: 4}

func TestButtonsCount(t *Testing.T) {
	if screen.NumberOfButtons() != 32 {
		t.Error("Expected 32, got ", screen.NumberOfButtons())
	}
}

func TestButtonPosition(t *Testing.T) {
	firstButtonX, firstButtonY := screen.ButtonPosition(0)
	if firstButtonX != 0 || firstButtonY != 0 {
		t.Error("Expected 0, 0, got ", firstButtonX, firstButtonY)
	}

	secondButtonX, secondButtonY := screen.ButtonPosition(8)
	if secondButtonX != 0 || secondButtonY != 1 {
		t.Error("Expected 0, 1, got ", secondButtonX, secondButtonY)
	}

	thirdButtonX, thirdButtonY := screen.ButtonPosition(21)

	if thirdButtonX != 5 || thirdButtonY != 2 {
		t.Error("Expected 5, 2, got ", thirdButtonX, thirdButtonY)
	}

	fourthButtonX, fourthButtonY := screen.ButtonPosition(31)
	if fourthButtonX != 7 || fourthButtonY != 3 {
		t.Error("Expected 7, 3, got ", fourthButtonX, fourthButtonY)
	}
}

func TestButtonNumber(t *Testing.T) {
	firstButton := screen.ButtonNumber(0, 0)
	if firstButton != 0 {
		t.Error("Expected 0, got ", firstButton)
	}

	secondButton := screen.ButtonNumber(0, 1)
	if secondButton != 8 {
		t.Error("Expected 8, got ", secondButton)
	}

	thirdButton := screen.ButtonNumber(5, 2)
	if thirdButton != 21 {
		t.Error("Expected 21, got ", thirdButton)
	}

	fourthButton := screen.ButtonNumber(7, 3)
	if fourthButton != 31 {
		t.Error("Expected 31, got ", fourthButton)
	}
}