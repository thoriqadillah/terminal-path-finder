package entity

type Block struct {
	X    int
	Y    int
	Char string
}

func NewBlock(x int, y int, char string) Block {
	return Block{x, y, char}
}
