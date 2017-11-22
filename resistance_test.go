package bolt_test

import (
	"fmt"
	"testing"

	bolt "github.com/Konstantin8105/Eurocode3.Bolt"
	"github.com/bradleyjkemp/cupaloy"
)

func ExampleShearResistance() {
	b := bolt.New(bolt.G5p8, bolt.D24)
	sr := bolt.ShearResistance{B: b, Position: bolt.ThreadShear}
	fmt.Printf("%s\n", sr)

	// Output:
	// Calculation of shear resistance for HM24Cl5.8:
	// 	γM2 = 1.250
	// 	αν  = 0.500 - Shear plane passes through the threaded portion of the bolt
	// 	Fub = 500.0 MPa
	// 	As  = 352.8 mm²
	//	In according to table 3.4 EN1993-1-8:
	// 	Shear resistance is 70.6 kN
}

func boltResistance(b bolt.Bolt) (s string) {
	{
		sr := bolt.ShearResistance{B: b, Position: bolt.ThreadShear}
		s += fmt.Sprintf("%s\n", sr)
	}
	{
		sr := bolt.ShearResistance{B: b, Position: bolt.UnthreadShear}
		s += fmt.Sprintf("%s\n", sr)
	}
	return
}

func TestResistanceCases(t *testing.T) {
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))
	for _, bd := range bolt.GetBoltDiameterList() {
		for _, bc := range bolt.GetBoltClassList() {
			b := bolt.New(bc, bd)

			testName := fmt.Sprintf("ShearResistance%s%s", bd, bc)
			t.Run(testName, func(t *testing.T) {
				result := boltResistance(b)
				err := snapshotter.SnapshotMulti(testName, result)
				if err != nil {
					t.Fatalf("error: %s", err)
				}
			})
		}
	}
}
