package graph

import (
	"strconv"

	"github.com/thoriqadillah/terminal-path-finder/entity"
	"github.com/thoriqadillah/terminal-path-finder/lib"
)

type Graph struct {
	start       entity.Block
	destination entity.Block
	queue       lib.IQueue[entity.Block]
	Canvas      *entity.Canvas
	path        map[string]entity.Block
}

func NewGraph(start entity.Block, destination entity.Block) *Graph {
	graph := Graph{
		start:       start,
		destination: destination,
		queue:       lib.NewQueue[entity.Block](),
		path:        make(map[string]entity.Block),
	}
	graph.queue.Enqueue(start)
	key := strconv.Itoa(start.X) + "|" + strconv.Itoa(start.Y)
	graph.path[key] = entity.Block{}
	return &graph
}

func (g *Graph) GetBlocks(canvas *entity.Canvas) *Graph {
	g.Canvas = canvas
	return g
}

func (g *Graph) BFS() (map[string]entity.Block, bool) {
	neighbours := []int{-1, 1, 0, 0}

	if g.queue.IsEmpty() {
		return g.path, false
	}

	current := g.queue.Dequeue()
	for i := range neighbours {
		x := current.X + neighbours[i]
		y := current.Y + neighbours[(i+2)%4]

		if x < 0 || y < 0 || x >= len(g.Canvas.Cells[0]) || y >= len(g.Canvas.Cells) || // if index out of bound or
			g.Canvas.Cells[y][x] == "░" || g.Canvas.Cells[y][x] == "█" || g.Canvas.Cells[y][x] == g.start.Char { //if visited or blocked
			continue
		}

		g.queue.Enqueue(entity.NewBlock(x, y, ""))
		key := strconv.Itoa(x) + "|" + strconv.Itoa(y)
		g.path[key] = current //make current node as parrent

		if x == g.destination.X && y == g.destination.Y {
			return g.path, true
		}
		g.Canvas.Cells[y][x] = "░"
	}

	return g.path, false
}

func (g *Graph) ReconstructPath(paths map[string]entity.Block, char string) {
	current := paths[strconv.Itoa(g.destination.X)+"|"+strconv.Itoa(g.destination.Y)]
	var empty entity.Block

	for current != empty {
		if (current.X != g.destination.X && current.Y != g.destination.Y) ||
			(current.X != g.start.X && current.Y != g.start.Y) {
			g.Canvas.Cells[current.Y][current.X] = char
		}
		current = paths[strconv.Itoa(current.X)+"|"+strconv.Itoa(current.Y)]
	}
}
