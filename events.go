package main

import (
	"fyne.io/fyne/v2"
	"strconv"
)

// Handles the event of a rune being typed.
func (c *calc) onTypedRune(r rune) {
	if action, ok := c.buttons[string(r)]; ok {
		action.OnTapped()
	}
}

// Handles the event of a key being typed.
func (c *calc) onTypedKey(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		c.evaluate()
	case fyne.KeyBackspace:
		c.backspace()
	}
}

// Handles the event of the paste shortcut being used.
func (c *calc) onPasteShortcut(shortcut fyne.Shortcut) {
	content := shortcut.(*fyne.ShortcutPaste).Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		c.appendToEquation(content)
	}
}

// Handles the event of the copy shortcut being used.
func (c *calc) onCopyShortcut(shortcut fyne.Shortcut) {
	shortcut.(*fyne.ShortcutCopy).Clipboard.SetContent(c.equation)
}

// Configures the events for the calculator's canvas.
func (c *calc) configureCanvasEvents() {
	canvas := c.window.Canvas()
	canvas.SetOnTypedRune(c.onTypedRune)
	canvas.SetOnTypedKey(c.onTypedKey)
	canvas.AddShortcut(&fyne.ShortcutCopy{}, c.onCopyShortcut)
	canvas.AddShortcut(&fyne.ShortcutPaste{}, c.onPasteShortcut)
}
