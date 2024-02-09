package services

func SquareGrid(width int, height int) *Grid {
	grid := NewGrid(width, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			grid.Set(j, i, 1)
		}
	}
	return grid
}

func HorizontalLineGrid(width int, height int) *Grid {
	grid := NewGrid(width, height)
	for i := 0; i < height; i++ {
		grid.Set(i, height/2, 1)
	}
	return grid
}

func LShapeGrid(width int, height int) *Grid {
	grid := NewGrid(width, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				grid.Set(j, i, 1)
			}
		}
	}
	return grid
}

type Grid struct {
	width  int
	height int
	matrix [][]int
}

func NewGrid(width int, height int) *Grid {
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	return &Grid{width, height, matrix}
}

func (g *Grid) Set(x, y, value int) {
	g.matrix[y][x] = value
}
