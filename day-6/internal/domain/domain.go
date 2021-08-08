package domain

const (
	turnOn  string = "turn on"
	turnOff string = "turn off"
	toggle  string = "toggle"
)

type Position struct {
	X, Y int
}

type Light interface {
	Toggle() Light
	On() Light
	Off() Light
	Status() int
}

type LightFactory func(int, int) Light

type BinaryLight struct {
	Position
	IsOn bool
}

func (l BinaryLight) Toggle() Light {
	return newBinaryLight(l.X, l.Y, !l.IsOn)
}

func (l BinaryLight) On() Light {
	return newBinaryLight(l.X, l.Y, true)
}

func (l BinaryLight) Off() Light {
	return newBinaryLight(l.X, l.Y, false)
}

func (l BinaryLight) Status() int {
	if l.IsOn {
		return 1
	}

	return 0
}

type BrightnessLight struct {
	Position
	Brightness int
}

func (l BrightnessLight) Toggle() Light {
	return newBrightnessLight(l.X, l.Y, l.Brightness + 2)
}

func (l BrightnessLight) On() Light {
	return newBrightnessLight(l.X, l.Y, l.Brightness + 1)
}

func (l BrightnessLight) Off() Light {
	var brightness int
	if l.Brightness - 1 > 0 {
		brightness = l.Brightness - 1
	}
	return newBrightnessLight(l.X, l.Y, brightness)
}

func (l BrightnessLight) Status() int {
	return l.Brightness
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
			lightsOn += light.Status()
		}
	}

	return lightsOn
}

type Direction struct {
	Start  Position
	End    Position
	Action string
}

func NewBinaryLight(x, y int) Light {
	return newBinaryLight(x, y, false)
}

func newBinaryLight(x, y int, isOn bool) Light {
	return BinaryLight{
		Position: Position{x, y},
		IsOn: isOn,
	}
}

func NewBrightnessLight(x, y int) Light {
	return newBrightnessLight(x, y, 0)
}

func newBrightnessLight(x, y, brightness int) Light {
	return BrightnessLight{
		Position: Position{x, y},
		Brightness: brightness,
	}
}

func NewGrid(width, height int, lightFactory LightFactory) Grid {
	grid := make(Grid, height)

	for i := 0; i < height; i++ {
		row := make([]Light, width)

		for j := 0; j < width; j++ {
			row[j] = lightFactory(i, j)
		}

		grid[i] = row
	}

	return grid
}
