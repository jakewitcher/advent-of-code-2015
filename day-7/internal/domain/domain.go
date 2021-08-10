package domain

type Identifier string

type Signal int

type Gate string

const (
	And    Gate = "AND"
	Or     Gate = "OR"
	LShift Gate = "LSHIFT"
	RShift Gate = "RSHIFT"
	Not    Gate = "NOT"
)

type Wire struct {
	Identifier
	Signal
}

type SignalCreator interface {
	Create() Signal
}

type SignalProvider interface {
	Provide(Wire, Signal) Wire
}

func NewWire(identifier Identifier, signal Signal) Wire {
	return Wire{
		Identifier: identifier,
		Signal:     signal,
	}
}
