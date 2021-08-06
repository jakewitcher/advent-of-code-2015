package domain

const (
	on  Status = true
	off Status = false
)

const (
	turnOn  string = "turn on"
	turnOff string = "turn off"
	toggle  string = "toggle"
)

type Position struct {
	X, Y int
}

type Status bool

type Light struct {
	Position
	Status Status
}

func (l Light) Toggle() Light {
	return newLight(l.X, l.Y, !l.Status)
}

func (l Light) On() Light {
	return newLight(l.X, l.Y, on)
}

func (l Light) Off() Light {
	return newLight(l.X, l.Y, off)
}

type Grid [][]Light

func (g Grid) ApplyDirection(d Direction) {
	for i := d.Start.X; i <= d.End.X; i++ {
		for j := d.Start.Y; j <= d.End.Y; j++ {
			switch d.Action {
			case turnOn:
				light := g[i][j]
				g[i][j] = light.On()
			case turnOff:
				light := g[i][j]
				g[i][j] = light.Off()
			case toggle:
				light := g[i][j]
				g[i][j] = light.Toggle()
			}
		}
	}
}

func (g Grid) CountLightsOn() int {
	var lightsOn int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			light := g[i][j]
			if light.Status {
				lightsOn++
			}
		}
	}

	return lightsOn
}

type Direction struct {
	Start  Position
	End    Position
	Action string
}

func NewLight(x, y int) Light {
	return Light{
		Position: Position{x, y},
		Status:   off,
	}
}

func newLight(x int, y int, status Status) Light {
	return Light{
		Position: Position{x, y},
		Status:   status,
	}
}

func NewGrid(width, height int) Grid {
	grid := make(Grid, height)

	for i := 0; i < height; i++ {
		row := make([]Light, width)

		for j := 0; j < width; j++ {
			row[j] = NewLight(i, j)
		}

		grid[i] = row
	}

	return grid
}
