package bolt

import (
	"fmt"
	"math"
)

// Force - type of force
type Force float64

func (f Force) String() string {
	return fmt.Sprintf("%.1f kN", float64(f)*1e-3)
}

// Factor - type of factors
type Factor float64

func (f Factor) String() string {
	return fmt.Sprintf("%.3f", float64(f))
}

// ανThreadShear - factor if shear by thread of bolt
var ανThreadShear = map[Class]Factor{
	G4p6:  0.6,
	G4p8:  0.5,
	G5p6:  0.6,
	G5p8:  0.5,
	G6p8:  0.5,
	G8p8:  0.6,
	G10p9: 0.5,
}

// ανThreadShear - factor if shear not by thread of bolt
var ανUnthreadShear Factor = 0.6

// PositionShear - position of shear on thread or not
type PositionShear bool

// Constants
const (
	ThreadShear   PositionShear = false
	UnthreadShear               = true
)

func (pos PositionShear) String() string {
	if pos == ThreadShear {
		return "Shear plane passes through the threaded portion of the bolt"
	}
	return "Shear plane passes through the unthreaded portion of the bolt"
}

// FactorγM2 - factor
var FactorγM2 Factor = 1.25

// ShearResistance - force of resistance on shear
type ShearResistance struct {
	B        Bolt
	Position PositionShear
}

func (sr ShearResistance) αν() Factor {
	switch sr.Position {
	case UnthreadShear:
		return ανUnthreadShear
	}
	//case ThreadShear:
	return ανThreadShear[sr.B.bc]
}

// Value - return Force of shear resistance
func (sr ShearResistance) Value() Force {
	return Force(float64(sr.αν()) * float64(sr.B.Fub().Value()) * float64(sr.B.As().Value()) / float64(FactorγM2))
}

func (sr ShearResistance) String() (s string) {
	s += fmt.Sprintf("Calculation of shear resistance for %s%s:\n", sr.B.bd, sr.B.bc)
	s += fmt.Sprintf("\tγM2 = %s\n", FactorγM2)
	s += fmt.Sprintf("\tαν  = %s - %s\n", sr.αν(), sr.Position)
	s += fmt.Sprintf("\tFub = %s\n", sr.B.Fub().Value())
	s += fmt.Sprintf("\tAs  = %s\n", sr.B.As().Value())
	s += fmt.Sprintf("\tIn according to table 3.4 EN1993-1-8:\n")
	s += fmt.Sprintf("\tShear resistance is %s", sr.Value())
	return
}

// Type - configuration of bolt
type Type bool

// Constants
const (
	UsuallyBolt     Type = false
	CountersunkBolt      = true
)

func (bt Type) String() string {
	if bt { // == CountersunkBolt
		return "countersunk bolt"
	}
	return "no-countersunk bolt"
}

// TensionResistance - force of resistance on tension
type TensionResistance struct {
	B  Bolt
	BT Type
}

// K2 - Factor
func (t TensionResistance) K2() Factor {
	if t.BT { // == CountersunkBolt
		return Factor(0.63)
	}
	return 0.9
}

// Value - return Force of tension resistance
func (t TensionResistance) Value() Force {
	return Force(float64(t.K2()) * float64(t.B.Fub().Value()) * float64(t.B.As().Value()) / float64(FactorγM2))
}

func (t TensionResistance) String() (s string) {
	s += fmt.Sprintf("Calculation of tension resistance for %s%s:\n", t.B.bd, t.B.bc)
	s += fmt.Sprintf("\tγM2 = %s\n", FactorγM2)
	s += fmt.Sprintf("\tk2  = %s - %s\n", t.K2(), t.BT)
	s += fmt.Sprintf("\tFub = %s\n", t.B.Fub().Value())
	s += fmt.Sprintf("\tAs  = %s\n", t.B.As().Value())
	s += fmt.Sprintf("\tIn according to table 3.4 EN1993-1-8:\n")
	s += fmt.Sprintf("\tTension resistance is %s", t.Value())
	return
}

// Resistance - combined resistance shear and tension
type Resistance struct {
	B        Bolt
	BT       Type
	Position PositionShear
}

// ViewResult - type of result view
type ViewResult int

// ViewResult constants
const (
	NoView ViewResult = iota
	FullView
)

// Value - return result of combined resistance calculation
func (r Resistance) Value(FvEd, FtEd Force, view ViewResult) (_ Factor, s string) {
	max := 0.0

	FvRd := ShearResistance{B: r.B, Position: r.Position}
	f1 := float64(FvEd) / float64(FvRd.Value())
	if view == FullView {
		s += fmt.Sprintf("%s\n", FvRd)
		s += fmt.Sprintf("Factor %s\n", Factor(f1))
	}
	max = math.Max(max, f1)

	FtRd := TensionResistance{B: r.B, BT: r.BT}
	f2 := float64(FtEd) / float64(FtRd.Value())
	if view == FullView {
		s += fmt.Sprintf("%s\n", FtRd)
		s += fmt.Sprintf("Factor %s\n", Factor(f2))
	}
	max = math.Max(max, f2)

	max = math.Max(max, float64(FvEd)/float64(FvRd.Value())+float64(FtEd)/(1.4*float64(FtRd.Value())))
	if view == FullView {
		s += fmt.Sprintf("Summary factor of combined loads is %s\n", Factor(max))
	}

	return Factor(max), s
}
