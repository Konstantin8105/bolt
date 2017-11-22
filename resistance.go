package bolt

import "fmt"

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
	s += fmt.Sprintf("\tShear resistance is %s", sr.Value())
	return
}

/*
	private double EN1993_1_8_TABLE_3_4_FtRd(double Pub, double As, double gamma_M2)
	{
	    double k2 = 0.9;
	    return k2 * Pub * As / gamma_M2;
	}
*/
