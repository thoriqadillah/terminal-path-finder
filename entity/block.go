package entity

type Block struct {
	X    int
	Y    int
	char string
}

func newBlock(x int, y int, char string) Block {
	return Block{x, y, char}
}
