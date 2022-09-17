package main

import (
	"github.com/thoriqadillah/terminal-path-finder/entity"
	"github.com/thoriqadillah/terminal-path-finder/service/cursor"
	"github.com/thoriqadillah/terminal-path-finder/service/graph"
)

func main() {
	const (
		// WIDTH  = 165
		// HEIGHT = 40
		WIDTH  = 60
		HEIGHT = 40
	)

	color := cursor.NewColor()
	cursor := cursor.NewCursor()
	canvas := entity.NewCanvas(WIDTH, HEIGHT).Draw()

	n := 0.40 * WIDTH * HEIGHT
	canvas.DrawBlock(int(n), "█")                         //hurdles
	start := canvas.DrawBlock(1, color.SetBlue("░"))      //start
	destination := canvas.DrawBlock(1, color.SetRed("░")) //end

	graph := graph.NewGraph(&start[0], &destination[0]).GetBlocks(&canvas)
	for {
		graph.BFS()
		cursor.Render(graph)

	}

}
