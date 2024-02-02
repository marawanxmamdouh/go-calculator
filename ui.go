package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// Initializes the output label for the calculator.
func (c *calc) initOutputLabel() {
	c.output = widget.NewLabel("")
	c.output.Alignment = fyne.TextAlignTrailing
	c.output.TextStyle.Monospace = true
}

// Creates the equals button for the calculator.
func (c *calc) createEqualsButton() *widget.Button {
	equals := c.actionButton("=", c.evaluate)
	equals.Importance = widget.HighImportance
	return equals
}

// Sets up the main window for the calculator.
func (c *calc) setupMainWindow(app fyne.App) {
	c.window = app.NewWindow("Calc")
}

// Sets up the content grid for the calculator.
func (c *calc) setupContentGrid(equalsButton *widget.Button) {
	c.window.SetContent(container.NewGridWithColumns(1,
		c.output,
		c.createButtonsRow('C', '(', ')', '/'),
		c.createButtonsRow(7, 8, 9, '*'),
		c.createButtonsRow(4, 5, 6, '-'),
		c.createButtonsRow(1, 2, 3, '+'),
		c.createButtonsRow(0, '.', equalsButton),
	))
}

// Refactored to abstract button creation logic into a separate method.
func (c *calc) createButtonsRow(items ...interface{}) fyne.CanvasObject {
	buttons := make([]fyne.CanvasObject, len(items))
	for i, item := range items {
		buttons[i] = c.createButton(item)
	}
	return container.NewGridWithColumns(len(items), buttons...)
}

// Creates various types of buttons for a calculator application.
func (c *calc) createButton(item interface{}) fyne.CanvasObject {
	switch v := item.(type) {
	case rune:
		if v == 'C' {
			return c.actionButton(string(v), c.clear)
		} else {
			return c.charButton(v)
		}
	case int:
		return c.digitButton(v)
	case *widget.Button:
		return v
	default:
		// Optionally handle unexpected item types, potentially returning a default button or nil.
		return nil
	}
}

// Finalizes the setup for the calculator's window.
func (c *calc) finalizeWindowSetup() {
	c.window.Resize(fyne.NewSize(400, 600))
	c.window.Show()
}

// -------------------------------------------------- Button Creation --------------------------------------------------
// Adds a button with the provided text and action to the calculator.
func (c *calc) actionButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button
	return button
}

// Creates a button for a digit.
func (c *calc) digitButton(digit int) *widget.Button {
	return c.actionButton(strconv.Itoa(digit), func() {
		c.appendToEquation(strconv.Itoa(digit))
	})
}

// Creates a button for a character.
func (c *calc) charButton(char rune) *widget.Button {
	return c.actionButton(string(char), func() {
		c.appendToEquation(string(char))
	})
}
