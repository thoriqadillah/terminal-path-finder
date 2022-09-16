package entity

import (
	"fmt"

	"github.com/thoriqadillah/terminal-path-finder/lib"
)

type canvas struct {
	width  int
	height int
	cells  cell
}

func NewCanvas(width int, height int) *canvas {
	cells := make(cell, height)
	for i := range cells {
		cells[i] = make([]string, width)
	}

	return &canvas{width, height, cells}
}

func (c *canvas) Draw() *canvas {
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			c.cells[i][j] = " "
		}
	}

	return c
}

func (c *canvas) DrawBlock(total int, char string) []block {
	blocks := make([]block, total)

	for i := 0; i < total; i++ {
		x := lib.Random(c.width)
		y := lib.Random(c.height)

		blocks[i] = newBlock(x, y, char)
		c.cells[y][x] = char
	}

	return blocks
}

func (c *canvas) Render() {
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			fmt.Printf("%+v", c.cells[i][j])
		}
		println()
	}
}
