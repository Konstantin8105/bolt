package bolt

import "fmt"

// Bolt - base property of bolt
type Bolt struct {
	bc  BoltClass
	dia BoltDiameter
}

// BoltClass is class of bolt
type BoltClass string

// Typical names of bolt classes
const (
	G4_6  BoltClass = "4.6"
	G4_8  BoltClass = "4.8"
	G5_6  BoltClass = "5.6"
	G5_8  BoltClass = "5.8"
	G6_8  BoltClass = "6.8"
	G8_8  BoltClass = "8.8"
	G10_9 BoltClass = "10.9"
)

// GetBoltClassList - list of all allowable bolt classes
func GetBoltClassList() []BoltClass {
	return []BoltClass{G4_6, G4_8, G5_6, G5_8, G6_8, G8_8, G10_9}
}

func (bc BoltClass) String() string {
	return fmt.Sprintf("Bolt class : Cl %s", string(bc))
}

// BoltDiameter is diameter of bolt
type BoltDiameter float64

// Typical bolt diameters
const (
	D12 BoltDiameter = 12.e-3
	D16 BoltDiameter = 16.e-3
	D20 BoltDiameter = 20.e-3
	D24 BoltDiameter = 24.e-3
	D30 BoltDiameter = 30.e-3
	D36 BoltDiameter = 36.e-3
	D42 BoltDiameter = 42.e-3
	D48 BoltDiameter = 48.e-3
)

// GetBoltDiameterList - list of all allowable bolt classes
func GetBoltDiameterList() []BoltDiameter {
	return []BoltDiameter{D12, D16, D20, D24, D30, D36, D42, D48}
}

func (bd BoltDiameter) String() string {
	return fmt.Sprintf("%.1 mm", float64(bd)*1e3)
}

// Table of Fyb
var fyb = map[BoltClass]stress{
	G4_6:  240.e6,
	G4_8:  320.e6,
	G5_6:  300.e6,
	G5_8:  400.e6,
	G6_8:  480.e6,
	G8_8:  640.e6,
	G10_9: 900.e6,
}

type Fyb struct {
	stress
	Class BoltClass
}

// Value - return value of Fyb
func (f Fyb) Value() stress {
	return fyb[f.Class]
}

func (f Fyb) String() string {
	return fmt.Sprintf("In according to table 3.1 EN1993-1-8 value Fyb is %s", f.Value())
}

type stress float64

func (s stress) String() string {
	return fmt.Sprintf("%.1f MPa", float64(s)*1.e-6)
}
