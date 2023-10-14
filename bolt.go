// Package bolt for calculate bolt property in according to EN1993-1-8
package bolt

import (
	"fmt"
	"math"
)

// Bolt - base property of bolt
type Bolt struct {
	bc Class
	bd Diameter
}

// New - create a new bolt
func New(bd Diameter, bc Class) Bolt {
	return Bolt{bc: bc, bd: bd}
}

// Fyb - return Fyb stress.
// unit: Pa
func (b Bolt) Fyb() Fyb {
	return Fyb{BoltClass: b.bc}
}

// Fub - return Fub stress.
// unit: Pa
func (b Bolt) Fub() Fub {
	return Fub{BoltClass: b.bc}
}

// D - diameter of bolt.
// unit: meter
func (b Bolt) D() Diameter {
	return b.bd
}

// Do - diameter of bolt hole.
// unit: meter
func (b Bolt) Do() HoleDiameter {
	return HoleDiameter{Dia: b.bd}
}

// Cl - class of bolt
func (b Bolt) Cl() Class {
	return b.bc
}

// As - area of As
func (b Bolt) As() AreaAs {
	return AreaAs{Dia: b.bd}
}

// A - area of A
func (b Bolt) A() AreaA {
	return AreaA{Dia: b.bd}
}

// Class is class of bolt
type Class string

// Typical names of bolt classes
const (
	G4p6  Class = "4.6"
	G4p8  Class = "4.8"
	G5p6  Class = "5.6"
	G5p8  Class = "5.8"
	G6p8  Class = "6.8"
	G8p8  Class = "8.8"
	G10p9 Class = "10.9"
)

// GetBoltClassList - list of all allowable bolt classes
func GetBoltClassList() []Class {
	return []Class{G4p6, G4p8, G5p6, G5p8, G6p8, G8p8, G10p9}
}

func (bc Class) String() string {
	return fmt.Sprintf("Cl%s", string(bc))
}

// Diameter is diameter of bolt
// unit: meter
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
	return fmt.Sprintf("HM%.0f", float64(bd)*1e3)
}

// HoleDiameter - struct of bolt hole
type HoleDiameter struct {
	Dia Diameter
}

var holeDiameter = map[Diameter]DiameterDimension{
	D12: 13e-3,
	D16: 18e-3,
	D20: 22e-3,
	D24: 26e-3,
	D30: 33e-3,
	D36: 39e-3,
	D42: 45e-3,
	D48: 51e-3,
}

// Value - return value diameter of hole for bolt
func (hd HoleDiameter) Value() DiameterDimension {
	return holeDiameter[hd.Dia]
}
func (hd HoleDiameter) String() string {
	return fmt.Sprintf("For bolt %s hole is %s", hd.Dia, hd.Value())
}

// DiameterDimension - dimension of diameter
type DiameterDimension float64

func (dia DiameterDimension) String() string {
	return fmt.Sprintf("Ø%s", Dimension(float64(dia)))
}

// Table of Fyb.
// unit: Pa
var fyb = map[Class]Stress{
	G4p6:  240.e6,
	G4p8:  320.e6,
	G5p6:  300.e6,
	G5p8:  400.e6,
	G6p8:  480.e6,
	G8p8:  640.e6,
	G10p9: 900.e6,
}

// Table of Fub.
// unit: Pa
var fub = map[Class]Stress{
	G4p6:  400.e6,
	G4p8:  400.e6,
	G5p6:  500.e6,
	G5p8:  500.e6,
	G6p8:  600.e6,
	G8p8:  800.e6,
	G10p9: 1000.e6,
}

// Fyb - stress of bolt in according to table 3.1. EN1993-1-8.
// unit: Pa
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

// Fub - stress of bolt in according to table 3.1. EN1993-1-8.
// unit: Pa
type Fub struct {
	Stress
	BoltClass Class
}

// Value - return value of Fub
func (f Fub) Value() Stress {
	return fub[f.BoltClass]
}

func (f Fub) String() string {
	return fmt.Sprintf("In according to table 3.1 EN1993-1-8 value Fub is %s", f.Value())
}

// Stress - struct of float64 for Stress values.
// unit: Pa
type Stress float64

func (s Stress) String() string {
	return fmt.Sprintf("%.1f MPa", float64(s)*1.e-6)
}

// Pinch - struct of bolt pinch
type Pinch struct {
	Dia Diameter
}

var boltPinch = map[Diameter]Dimension{
	D12: 1.75e-3,
	D16: 2.00e-3,
	D20: 2.50e-3,
	D24: 3.00e-3,
	D30: 3.50e-3,
	D36: 4.00e-3,
	D42: 4.50e-3,
	D48: 5.00e-3,
}

// Value - return value of bolt pinch
func (bp Pinch) Value() Dimension {
	return boltPinch[bp.Dia]
}

// Dimension - type for linear dimension sizes (height, thk, width)
type Dimension float64

func (d Dimension) String() string {
	return fmt.Sprintf("%.1f mm", float64(d)*1.e3)
}

// Area - type of area.
// unit - sq.meter
type Area float64

func (a Area) String() string {
	return fmt.Sprintf("%.1f mm\u00B2", float64(a)*1.e6)
}

// AreaAs tension stress area of the bolt
type AreaAs struct {
	Dia Diameter
}

// Value - return value of area As (tension stress area of the bolt)
func (as AreaAs) Value() Area {
	var pd = Pinch(as) // Use a type conversion
	p := float64(pd.Value())
	dia := float64(as.Dia)
	return Area(math.Pi / 4. * math.Pow(dia-0.935229*p, 2.0))
}

func (as AreaAs) String() string {
	return fmt.Sprintf("Tension stress area of the bolt %s is %s", as.Dia, as.Value())
}

// AreaA - the gross cross-section area of bolt
type AreaA struct {
	Dia Diameter
}

// Value - return value of area A (the gross cross-section area of bolt)
func (aa AreaA) Value() Area {
	dia := float64(aa.Dia)
	return Area(math.Pi / 4. * math.Pow(dia, 2.0))
}

func (aa AreaA) String() string {
	return fmt.Sprintf("The gross cross-section area of the bolt %s is %s", aa.Dia, aa.Value())
}
