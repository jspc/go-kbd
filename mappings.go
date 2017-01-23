package main

import (
	"fmt"
)

// Mapping is a representation of a code to value for a given keyboard
//
// Note: both shiftcase and uppercase exist. This is because while caps
// will logically return the same value as shift on text, it behaves
// differently with character, numbers.
type Mapping struct {
	Lowercase, Uppercase, Shiftcase, Altcase string
}

// Empty determines whether a Mapping is empty for the purpose of returning
// key values
func (m Mapping) Empty() bool {
	return m == Mapping{}
}

// Mapper contains logic for mapping codes to keys
type Mapper struct {
	Caps     bool
	Shifted  bool
	Alted    bool
	Mappings []Mapping

	LeftShift, RightShift, Alt, AltGr, CapsLock uint16
}

// MapManager contains a direct translation of Scancodes to keys
type MapManager interface {
	Mappings() []Mapping
	CapsLock() uint16
	LeftShift() uint16
	RightShift() uint16
	Alt() uint16
	AltGr() uint16
}

// NewMapper returns a Mapper for mapping presses to values
func NewMapper(mm MapManager) (m Mapper) {
	// TODO: How would one ensure these *are* false at start up?
	m.Caps = false
	m.Shifted = false
	m.Alted = false

	m.Mappings = mm.Mappings()
	m.LeftShift = mm.LeftShift()
	m.RightShift = mm.RightShift()
	m.Alt = mm.Alt()
	m.AltGr = mm.AltGr()
	m.CapsLock = mm.CapsLock()

	return
}

func (m *Mapper) ShiftOn() {
	m.Shifted = true
}

func (m *Mapper) ShiftOff() {
	m.Shifted = false
}

func (m *Mapper) AltOn() {
	m.Alted = true
}

func (m *Mapper) AltOff() {
	m.Alted = false
}

func (m *Mapper) CapsLockFlip() {
	if m.Caps {
		m.Caps = false
	} else {
		m.Caps = true
	}
}

func (m Mapper) Print(sc uint16) {
	var output string
	mappedCode := m.Mappings[int(sc)]

	if mappedCode.Empty() {
		output = fmt.Sprintf("<%d>", sc)
	} else {
		switch {
		case m.Shifted:
			output = mappedCode.Shiftcase
		case m.Caps:
			output = mappedCode.Uppercase
		case m.Alted:
			output = mappedCode.Altcase
		default:
			output = mappedCode.Lowercase
		}
	}

	fmt.Print(output)
}
