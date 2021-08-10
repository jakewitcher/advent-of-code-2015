package domain

type Identifier string

type Signal uint16

type Gate string

type Shift int

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

func NewWire(identifier Identifier, signal Signal) Wire {
	return Wire{
		Identifier: identifier,
		Signal:     signal,
	}
}
