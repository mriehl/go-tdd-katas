package rover

const (
	NOTHING  = iota
	OBSTACLE = iota
	ROVER    = iota
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

func (grid *Grid) OverflowPosition(startingPoint Coordinates, dX, dY int) (Coordinates, bool) {
	newCoords := startingPoint.Move(dX, dY)
	if newCoords.X >= grid.Width {
		newCoords.X -= grid.Width
	}
	if newCoords.Y >= grid.Height {
		newCoords.Y -= grid.Height
	}

	if grid.At(newCoords) != NOTHING {
		return startingPoint, false
	}

	return newCoords, true
}
