package wires

type Identifier string

type Signal uint16

type Shift int

type Wire struct {
	Identifier
	Signal
}

func NewWire(identifier Identifier, signal Signal) Wire {
	return Wire{
		Identifier: identifier,
		Signal:     signal,
	}
}
