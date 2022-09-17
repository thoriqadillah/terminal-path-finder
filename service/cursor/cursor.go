package cursor

import (
	"fmt"
	"time"

	"github.com/thoriqadillah/terminal-path-finder/service/graph"
)

type iCursor interface {
	refresh()
	Render(graph *graph.Graph)
}

type cursor struct{}

func NewCursor() iCursor {
	return &cursor{}
}

func (c *cursor) refresh() {
	fmt.Print("\u001b[H")
}

func (c *cursor) Render(graph *graph.Graph) {
	time.Sleep(1000 * time.Millisecond / 60) //this will make the renderer runs certain fps/s
	graph.Canvas.Display()
	c.refresh()
}
