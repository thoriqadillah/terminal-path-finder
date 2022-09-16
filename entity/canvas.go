package entity

import (
	"fmt"

	"github.com/thoriqadillah/terminal-path-finder/lib"
)

type canvas struct {
	width  int
	height int
	cells  cell
	blocks []block
}

func NewCanvas(width int, height int) *canvas {
	return &canvas{
		width:  width,
		height: height,
		cells:  make(cell, height),
	}
}

func (c *canvas) DrawCanvas() {
	for i := range c.cells {
		c.cells[i] = make([]string, c.width)
	}

	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			c.cells[i][j] = " "
		}
	}
}

func (c *canvas) DrawBlock(total int) {
	blocks := make([]block, total)

	for i := range blocks {
		x := lib.Random(c.width)
		y := lib.Random(c.height)

		blocks[i] = newBlock(x, y)
		c.cells[y][x] = blocks[i].char
	}

	c.blocks = append(c.blocks, blocks...)
}

func (c *canvas) Render() {
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			fmt.Printf("%s", c.cells[i][j])
		}
		println()
	}
}
