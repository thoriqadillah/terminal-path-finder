package entity

type block struct {
	x    int
	y    int
	char string
}

func newBlock(x int, y int) block {
	return block{
		x:    x,
		y:    y,
		char: "░",
	}
}
