package bolt

import "fmt"

// Force - type of force
type Force float64

func (f Force) String() string {
	return fmt.Sprintf("%.1f kN", float64(f)*1e-3)
}

// αν - factor
// TODO : add more details
var αν = map[Class]float64{
	G4p6:  0.6,
	G4p8:  0.5,
	G5p6:  0.6,
	G5p8:  0.5,
	G6p8:  0.5,
	G8p8:  0.6,
	G10p9: 0.5,
}

// γM2 - factor
var γM2 = 1.25

type ShearResistance struct {
	bolt Bolt
}

func (sr ShearResistance) Value() Force {
	return Force(αν[sr.bolt.bc] * float64(sr.bolt.Fub().Value()) * float64(sr.bolt.As().Value()) / γM2)
}

func (sr ShearResistance) String() string {
	// TODO : Add more calculation formula
	return fmt.Sprintf("Shear resistance is %s", sr.Value())
}

/*
	private double EN1993_1_8_TABLE_3_4_FtRd(double Pub, double As, double gamma_M2)
	{
	    double k2 = 0.9;
	    return k2 * Pub * As / gamma_M2;
	}
*/
