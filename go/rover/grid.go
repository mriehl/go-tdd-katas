package rover

const (
	NOTHING = 0
	ROVER   = 42
)

type Grid struct {
	Width  int
	Height int
	Field  [][]int
}

func NewGrid(width, height int) *Grid {
	grid := new(Grid)
	grid.Width = width
	grid.Height = height
	grid.Field = make([][]int, width)
	for index := range grid.Field {
		grid.Field[index] = make([]int, height)
	}
	return grid
}

func (grid *Grid) Snap(rover *Rover) *Grid {
	grid.Insert(rover.Coords, ROVER)
	return grid
}

func (grid *Grid) Insert(coords Coordinates, value int) {
	grid.Field[coords.X][coords.Y] = value
}

func (grid *Grid) At(coords Coordinates) int {
	return grid.Field[coords.X][coords.Y]
}

func (grid *Grid) OverflowPosition(startingPoint Coordinates, dX, dY int) Coordinates {
	newCoords := Coordinates{startingPoint.X + dX, startingPoint.Y + dY}
	if newCoords.X >= grid.Width {
		newCoords.X -= grid.Width
	}
	if newCoords.Y >= grid.Height {
		newCoords.Y -= grid.Height
	}
	return newCoords
}
