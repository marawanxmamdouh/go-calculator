package main

import (
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

// Struct represents the calculator.
type calc struct {
	equation string
	output   *widget.Label
	buttons  map[string]*widget.Button
	window   fyne.Window
}

// Display updates the calculator's display with the provided text.
func (c *calc) display(newText string) {
	if newText == "" {
		c.equation = ""
	} else if strings.Contains(c.equation, "error") {
		c.equation = "error"
	} else {
		c.equation = newText
	}
	c.output.SetText(c.equation)
}

// Appends the provided text to the current equation.
func (c *calc) appendToEquation(text string) {
	c.display(c.equation + text)
}

// Clears the current equation.
func (c *calc) clear() {
	c.display("")
}

// Removes the last character from the current equation.
func (c *calc) backspace() {
	if len(c.equation) > 0 {
		c.display(c.equation[:len(c.equation)-1])
	}
}

// Evaluates the current equation and updates the display with the result.
func (c *calc) evaluate() {
	if strings.Contains(c.equation, "error") {
		c.display("error")
		return
	}
	expression, err := govaluate.NewEvaluableExpression(c.equation)
	if err != nil {
		log.Println("Error in calculation:", err)
		c.display("error")
		return
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Println("Error evaluating expression:", err)
		c.display("error")
		return
	}

	c.display(formatResult(result))
}

// Formats the result of an evaluation into a string.
func formatResult(result interface{}) string {
	switch value := result.(type) {
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	default:
		return "error"
	}
}
