package cursor

type iColor interface {
	reset() string
	SetRed(str string) string
	SetGreen(str string) string
	SetBlue(str string) string
	SetYellow(str string) string
}

func NewColor() iColor {
	return &cursor{}
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
