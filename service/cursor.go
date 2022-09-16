package service

import "fmt"

type iColor interface {
	reset() string
	SetRed(str string) string
	SetGreen(str string) string
	SetBlue(str string) string
	SetYellow(str string) string
}

type iCursor interface {
	refresh()
}

type cursor struct{}

func NewColor() iColor {
	return &cursor{}
}

func NewCursor() iCursor {
	return &cursor{}
}

func (c *cursor) refresh() {
	fmt.Print("\u001b[H")
}

func (c *cursor) reset() string {
	return "\u001b[0m"
}

func (c *cursor) SetRed(str string) string {
	return "\u001b[31;1m" + str + c.reset()
}

func (c *cursor) SetGreen(str string) string {
	return "\u001b[32;1m" + str + c.reset()
}

func (c *cursor) SetYellow(str string) string {
	return "\u001b[33;1m" + str + c.reset()
}

func (c *cursor) SetBlue(str string) string {
	return "\u001b[34;1m" + str + c.reset()
}
