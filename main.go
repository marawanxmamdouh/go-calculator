package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// Loads the user interface for the calculator.
func (c *calc) loadUI(app fyne.App) {
	c.initOutputLabel()

	equalsButton := c.createEqualsButton()

	c.setupMainWindow(app)
	c.setupContentGrid(equalsButton)
	c.configureCanvasEvents()
	c.finalizeWindowSetup()
}

// Creates a new calculator.
func newCalculator() *calc {
	return &calc{
		buttons: make(map[string]*widget.Button, 19),
	}
}

func main() {
	myApp := app.New()

	c := newCalculator()
	c.loadUI(myApp)
	myApp.Run()
}
