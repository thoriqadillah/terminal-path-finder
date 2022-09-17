package graph

import (
	"github.com/thoriqadillah/terminal-path-finder/entity"
	"github.com/thoriqadillah/terminal-path-finder/lib"
)

type Graph struct {
	start       *entity.Block
	destination *entity.Block
	queue       lib.IQueue[entity.Block]
	Canvas      *entity.Canvas
}

func NewGraph(start *entity.Block, destination *entity.Block) *Graph {
	queue := lib.NewQueue[entity.Block]()
	queue.Enqueue(*start)

	return &Graph{
		start:       start,
		destination: destination,
		queue:       queue,
	}
}

func (g *Graph) GetBlocks(canvas *entity.Canvas) *Graph {
	g.Canvas = canvas
	return g
}

func (g *Graph) BFS() {
	neighbours := []int{-1, 1, 0, 0}
	height := len(g.Canvas.Cells)
	width := len(g.Canvas.Cells[0])

	current := g.queue.Dequeue()
	for i := range neighbours {
		x := current.X + neighbours[i]
		y := current.Y + neighbours[(i+2)%4]

		if x <= 0 || y <= 0 || x >= width || y >= height || // if index out of bound or
			g.Canvas.Cells[y][x] == "░" || g.Canvas.Cells[y][x] == "█" { //if visited or blocked
			continue
		}
		if x == g.destination.X || y == g.destination.Y {
			return
		}
		g.Canvas.Cells[y][x] = "░"
		g.queue.Enqueue(entity.NewBlock(x, y, "░"))
	}
}
