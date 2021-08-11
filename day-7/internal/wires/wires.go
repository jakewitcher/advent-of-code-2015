package wires

type Identifier string

type Shift int

type SignalProducer interface {
	ProduceSignal() Signal
}

type Signal uint16

func (s Signal) ProduceSignal() Signal {
	return s
}

type Wire struct {
	Identifier
	Signal
}

func (w Wire) ProduceSignal() Signal {
	return w.Signal
}

func NewWire(identifier Identifier, signal Signal) Wire {
	return Wire{
		Identifier: identifier,
		Signal:     signal,
	}
}
