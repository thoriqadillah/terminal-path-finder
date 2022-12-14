package graph

import (
	"strconv"

	"github.com/thoriqadillah/terminal-path-finder/entity"
	"github.com/thoriqadillah/terminal-path-finder/lib/list"
)

type graph struct {
	start       *entity.Block
	destination *entity.Block
	queue       list.IQueue[entity.Block]
	stack       list.IStack[entity.Block]
	Canvas      *entity.Canvas
	path        map[string]entity.Block
}

func NewGraph(start *entity.Block, destination *entity.Block) *graph {
	graph := graph{
		start:       start,
		destination: destination,
		queue:       list.NewQueue[entity.Block](),
		stack:       list.NewStack[entity.Block](),
		path:        make(map[string]entity.Block),
	}
	graph.queue.Enqueue(*start)
	graph.stack.Push(*start)
	key := strconv.Itoa(start.X) + "|" + strconv.Itoa(start.Y)
	graph.path[key] = entity.Block{}
	return &graph
}

func (g *graph) GetBlocks(canvas *entity.Canvas) *graph {
	g.Canvas = canvas
	return g
}

func (g *graph) BFS() bool {
	if g.queue.IsEmpty() {
		return false
	}

	neighbours := []int{-1, 1, 0, 0}
	current := g.queue.Dequeue()

	for i := range neighbours {
		x := current.X + neighbours[i]
		y := current.Y + neighbours[(i+2)%4]

		if x < 0 || y < 0 || x >= len(g.Canvas.Cells[0]) || y >= len(g.Canvas.Cells) || // if index out of bound or
			g.Canvas.Cells[y][x] == "░░" || g.Canvas.Cells[y][x] == "██" || g.Canvas.Cells[y][x] == g.start.Char { //if visited or blocked
			continue
		}

		g.queue.Enqueue(entity.NewBlock(x, y, ""))
		key := strconv.Itoa(x) + "|" + strconv.Itoa(y)
		g.path[key] = current //make current node as parrent

		if x == g.destination.X && y == g.destination.Y {
			return true
		}
		g.Canvas.Cells[y][x] = "░░"
	}

	return false
}

func (g *graph) DFS() bool {
	if g.queue.IsEmpty() {
		return false
	}

	neighbours := []int{-1, 1, 0, 0}
	current := g.stack.Pop()

	for i := range neighbours {
		x := current.X + neighbours[i]
		y := current.Y + neighbours[(i+2)%4]

		if x < 0 || y < 0 || x >= len(g.Canvas.Cells[0]) || y >= len(g.Canvas.Cells) || // if index out of bound or
			g.Canvas.Cells[y][x] == "░░" || g.Canvas.Cells[y][x] == "██" || g.Canvas.Cells[y][x] == g.start.Char { //if visited or blocked
			continue
		}

		g.stack.Push(entity.NewBlock(x, y, ""))
		key := strconv.Itoa(x) + "|" + strconv.Itoa(y)
		g.path[key] = current //make current node as parrent

		if x == g.destination.X && y == g.destination.Y {
			return true
		}
		g.Canvas.Cells[y][x] = "░░"
	}

	return false
}

func (g *graph) ReconstructPath(char string) {
	current := g.path[strconv.Itoa(g.destination.X)+"|"+strconv.Itoa(g.destination.Y)]

	for current != *g.start {
		if current != *g.destination || current != *g.start {
			g.Canvas.Cells[current.Y][current.X] = char
		}
		current = g.path[strconv.Itoa(current.X)+"|"+strconv.Itoa(current.Y)]
	}
}
