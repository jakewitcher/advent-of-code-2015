package light

type Light interface {
	Toggle() Light
	On() Light
	Off() Light
	Status() int
}

type Factory func(int, int) Light

type Position struct {
	X, Y int
}

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
