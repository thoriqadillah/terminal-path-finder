package cursor

import (
	"fmt"
	"time"

	"github.com/thoriqadillah/terminal-path-finder/entity"
)

type iCursor interface {
	refresh()
	Render(canvas *entity.Canvas)
}

type cursor struct{}

func NewCursor() iCursor {
	return &cursor{}
}

func (c *cursor) refresh() {
	fmt.Print("\u001b[H")
}

func (c *cursor) Render(canvas *entity.Canvas) {
	time.Sleep(time.Second / 60) //this will make the renderer runs certain fps/s
	canvas.Display()
	c.refresh()
}
