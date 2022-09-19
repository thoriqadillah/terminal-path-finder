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
		WIDTH  = 70
		HEIGHT = 20
	)

	color := cursor.NewColor()
	cursor := cursor.NewCursor()
	canvas := entity.NewCanvas(WIDTH, HEIGHT).Draw()

	n := 0.4 * WIDTH * HEIGHT
	canvas.DrawBlock(int(n), "█")                         //hurdles
	start := canvas.DrawBlock(1, color.SetBlue("█"))      //start
	destination := canvas.DrawBlock(1, color.SetRed("█")) //end

	graph := graph.NewGraph(start[0], destination[0]).GetBlocks(&canvas)
	var path map[string]entity.Block
	found := false

	for {
		if !found {
			path, found = graph.DFS()
		} else {
			graph.ReconstructPath(path, color.SetGreen("█"))
		}

		cursor.Render(&canvas)
	}
}
