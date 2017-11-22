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

func ExampleTensionResistance() {
	b := bolt.New(bolt.G5p8, bolt.D24)
	t := bolt.TensionResistance{B: b, BT: bolt.UsuallyBolt}
	fmt.Printf("%s\n", t)

	// Output:
	// Calculation of tension resistance for HM24Cl5.8:
	// 	γM2 = 1.250
	// 	k2  = 0.900 - no-countersunk bolt
	// 	Fub = 500.0 MPa
	// 	As  = 352.8 mm²
	// 	In according to table 3.4 EN1993-1-8:
	// 	Tension resistance is 127.0 kN
}

func boltShearResistance(b bolt.Bolt) (s string) {
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

func TestShearResistanceCases(t *testing.T) {
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))
	for _, bd := range bolt.GetBoltDiameterList() {
		for _, bc := range bolt.GetBoltClassList() {
			b := bolt.New(bc, bd)

			testName := fmt.Sprintf("ShearResistance%s%s", bd, bc)
			t.Run(testName, func(t *testing.T) {
				result := boltShearResistance(b)
				err := snapshotter.SnapshotMulti(testName, result)
				if err != nil {
					t.Fatalf("error: %s", err)
				}
			})
		}
	}
}

func boltTensionResistance(b bolt.Bolt) (s string) {
	{
		sr := bolt.TensionResistance{B: b, BT: bolt.UsuallyBolt}
		s += fmt.Sprintf("%s\n", sr)
	}
	{
		sr := bolt.TensionResistance{B: b, BT: bolt.CountersunkBolt}
		s += fmt.Sprintf("%s\n", sr)
	}
	return
}

func TestTensionResistanceCases(t *testing.T) {
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))
	for _, bd := range bolt.GetBoltDiameterList() {
		for _, bc := range bolt.GetBoltClassList() {
			b := bolt.New(bc, bd)

			testName := fmt.Sprintf("TensionResistance%s%s", bd, bc)
			t.Run(testName, func(t *testing.T) {
				result := boltTensionResistance(b)
				err := snapshotter.SnapshotMulti(testName, result)
				if err != nil {
					t.Fatalf("error: %s", err)
				}
			})
		}
	}
}
