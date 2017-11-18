package bolt

import "fmt"

// Bolt - base property of bolt
type Bolt struct {
	bc  Class
	dia Diameter
}

// Class is class of bolt
type Class string

// Typical names of bolt classes
const (
	G4_6  Class = "4.6"
	G4_8  Class = "4.8"
	G5_6  Class = "5.6"
	G5_8  Class = "5.8"
	G6_8  Class = "6.8"
	G8_8  Class = "8.8"
	G10_9 Class = "10.9"
)

// GetBoltClassList - list of all allowable bolt classes
func GetBoltClassList() []Class {
	return []Class{G4_6, G4_8, G5_6, G5_8, G6_8, G8_8, G10_9}
}

func (bc Class) String() string {
	return fmt.Sprintf("Bolt class : Cl %s", string(bc))
}

// Diameter is diameter of bolt
type Diameter float64

// Typical bolt diameters
const (
	D12 Diameter = 12.e-3
	D16 Diameter = 16.e-3
	D20 Diameter = 20.e-3
	D24 Diameter = 24.e-3
	D30 Diameter = 30.e-3
	D36 Diameter = 36.e-3
	D42 Diameter = 42.e-3
	D48 Diameter = 48.e-3
)

// GetBoltDiameterList - list of all allowable bolt classes
func GetBoltDiameterList() []Diameter {
	return []Diameter{D12, D16, D20, D24, D30, D36, D42, D48}
}

func (bd Diameter) String() string {
	return fmt.Sprintf("%.1 mm", float64(bd)*1e3)
}

// Table of Fyb
var fyb = map[Class]Stress{
	G4_6:  240.e6,
	G4_8:  320.e6,
	G5_6:  300.e6,
	G5_8:  400.e6,
	G6_8:  480.e6,
	G8_8:  640.e6,
	G10_9: 900.e6,
}

// Fyb - stress of bolt in according to table 3.1. EN1993-1-8
type Fyb struct {
	Stress
	BoltClass Class
}

// Value - return value of Fyb
func (f Fyb) Value() Stress {
	return fyb[f.BoltClass]
}

func (f Fyb) String() string {
	return fmt.Sprintf("In according to table 3.1 EN1993-1-8 value Fyb is %s", f.Value())
}

// Stress - struct of float64 for Stress values
type Stress float64

func (s Stress) String() string {
	return fmt.Sprintf("%.1f MPa", float64(s)*1.e-6)
}
