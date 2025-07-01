package bolt_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"text/tabwriter"

	"github.com/Konstantin8105/bolt"
)

func ExampleShearResistance() {
	b := bolt.New(bolt.D24, bolt.G5p8)
	sr := bolt.ShearResistance{B: b, Position: bolt.ThreadShear}
	fmt.Fprintf(os.Stdout, "%s\n", sr)

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
	b := bolt.New(bolt.D24, bolt.G5p8)
	t := bolt.TensionResistance{B: b, BT: bolt.UsuallyBolt}
	fmt.Fprintf(os.Stdout, "%s\n", t)

	// Output:
	// Calculation of tension resistance for HM24Cl5.8:
	// 	γM2 = 1.250
	// 	k2  = 0.900 - no-countersunk bolt
	// 	Fub = 500.0 MPa
	// 	As  = 352.8 mm²
	// 	In according to table 3.4 EN1993-1-8:
	// 	Tension resistance is 127.0 kN
}

func ExampleAddClass() {
	class := bolt.Class("S245")
	bolt.AddClass(
		class,
		230.0e6,
		360.0e6,
		0.5,
	)
	b := bolt.New(bolt.D24, class)
	fmt.Fprintf(os.Stdout, "Bolt : %s\n", b)
	fmt.Fprintf(os.Stdout, "%s\n", b.Do())
	fmt.Fprintf(os.Stdout, "Hole : %s\n", b.Do().Value())
	fmt.Fprintf(os.Stdout, "%s\n", b.Fyb())
	fmt.Fprintf(os.Stdout, "Fyb  : %s\n", b.Fyb().Value())
	fmt.Fprintf(os.Stdout, "%s\n", b.Fub())
	fmt.Fprintf(os.Stdout, "Fub  : %s\n", b.Fub().Value())
	fmt.Fprintf(os.Stdout, "%s\n", b.As())
	fmt.Fprintf(os.Stdout, "%s\n", b.A())
	// resistance
	sr := bolt.ShearResistance{B: b, Position: bolt.ThreadShear}
	fmt.Fprintf(os.Stdout, "%s\n", sr)
	t := bolt.TensionResistance{B: b, BT: bolt.UsuallyBolt}
	fmt.Fprintf(os.Stdout, "%s\n", t)

	// Output:
	// Bolt : HM24ClS245
	// For bolt HM24 hole is Ø26.0 mm
	// Hole : Ø26.0 mm
	// In according to table 3.1 EN1993-1-8 value Fyb is 230.0 MPa
	// Fyb  : 230.0 MPa
	// In according to table 3.1 EN1993-1-8 value Fub is 360.0 MPa
	// Fub  : 360.0 MPa
	// Tension stress area of the bolt HM24 is 352.8 mm²
	// The gross cross-section area of the bolt HM24 is 452.4 mm²
	// Calculation of shear resistance for HM24ClS245:
	// 	γM2 = 1.250
	// 	αν  = 0.500 - Shear plane passes through the threaded portion of the bolt
	// 	Fub = 360.0 MPa
	// 	As  = 352.8 mm²
	// 	In according to table 3.4 EN1993-1-8:
	// 	Shear resistance is 50.8 kN
	// Calculation of tension resistance for HM24ClS245:
	// 	γM2 = 1.250
	// 	k2  = 0.900 - no-countersunk bolt
	// 	Fub = 360.0 MPa
	// 	As  = 352.8 mm²
	// 	In according to table 3.4 EN1993-1-8:
	// 	Tension resistance is 91.4 kN
}

func Example() {
	var (
		D  = bolt.GetBoltDiameterList()
		Cl = bolt.GetBoltClassList()
	)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "| Diameter\t| Class\t| Tension, kN\t| Shear, kN\t|\n")
	for _, d := range D {
		for _, cl := range Cl {
			b := bolt.New(d, cl)
			nt := bolt.TensionResistance{B: b, BT: bolt.UsuallyBolt}
			ns := bolt.ShearResistance{B: b, Position: bolt.ThreadShear}
			fmt.Fprintf(w, "| %v\t| %v\t| %5.1f\t| %5.1f\t|\n", d, cl, nt.Value()/1000, ns.Value()/1000)
		}
		fmt.Fprintf(w, "|\t|\t|\t|\t|\n")
	}
	w.Flush()
	// Output:
	// | Diameter | Class  | Tension, kN | Shear, kN |
	// | HM12     | Cl4.6  |  24.3       |  16.2     |
	// | HM12     | Cl4.8  |  24.3       |  13.5     |
	// | HM12     | Cl5.6  |  30.4       |  20.2     |
	// | HM12     | Cl5.8  |  30.4       |  16.9     |
	// | HM12     | Cl6.8  |  36.4       |  20.2     |
	// | HM12     | Cl8.8  |  48.6       |  32.4     |
	// | HM12     | Cl10.9 |  60.7       |  33.7     |
	// |          |        |             |           |
	// | HM16     | Cl4.6  |  45.2       |  30.1     |
	// | HM16     | Cl4.8  |  45.2       |  25.1     |
	// | HM16     | Cl5.6  |  56.4       |  37.6     |
	// | HM16     | Cl5.8  |  56.4       |  31.4     |
	// | HM16     | Cl6.8  |  67.7       |  37.6     |
	// | HM16     | Cl8.8  |  90.3       |  60.2     |
	// | HM16     | Cl10.9 | 112.9       |  62.7     |
	// |          |        |             |           |
	// | HM20     | Cl4.6  |  70.6       |  47.0     |
	// | HM20     | Cl4.8  |  70.6       |  39.2     |
	// | HM20     | Cl5.6  |  88.2       |  58.8     |
	// | HM20     | Cl5.8  |  88.2       |  49.0     |
	// | HM20     | Cl6.8  | 105.8       |  58.8     |
	// | HM20     | Cl8.8  | 141.1       |  94.1     |
	// | HM20     | Cl10.9 | 176.4       |  98.0     |
	// |          |        |             |           |
	// | HM24     | Cl4.6  | 101.6       |  67.7     |
	// | HM24     | Cl4.8  | 101.6       |  56.4     |
	// | HM24     | Cl5.6  | 127.0       |  84.7     |
	// | HM24     | Cl5.8  | 127.0       |  70.6     |
	// | HM24     | Cl6.8  | 152.4       |  84.7     |
	// | HM24     | Cl8.8  | 203.2       | 135.5     |
	// | HM24     | Cl10.9 | 254.0       | 141.1     |
	// |          |        |             |           |
	// | HM30     | Cl4.6  | 161.6       | 107.7     |
	// | HM30     | Cl4.8  | 161.6       |  89.8     |
	// | HM30     | Cl5.6  | 202.0       | 134.6     |
	// | HM30     | Cl5.8  | 202.0       | 112.2     |
	// | HM30     | Cl6.8  | 242.4       | 134.6     |
	// | HM30     | Cl8.8  | 323.1       | 215.4     |
	// | HM30     | Cl10.9 | 403.9       | 224.4     |
	// |          |        |             |           |
	// | HM36     | Cl4.6  | 235.4       | 156.9     |
	// | HM36     | Cl4.8  | 235.4       | 130.8     |
	// | HM36     | Cl5.6  | 294.2       | 196.2     |
	// | HM36     | Cl5.8  | 294.2       | 163.5     |
	// | HM36     | Cl6.8  | 353.1       | 196.2     |
	// | HM36     | Cl8.8  | 470.8       | 313.9     |
	// | HM36     | Cl10.9 | 588.5       | 326.9     |
	// |          |        |             |           |
	// | HM42     | Cl4.6  | 323.1       | 215.4     |
	// | HM42     | Cl4.8  | 323.1       | 179.5     |
	// | HM42     | Cl5.6  | 403.8       | 269.2     |
	// | HM42     | Cl5.8  | 403.8       | 224.3     |
	// | HM42     | Cl6.8  | 484.6       | 269.2     |
	// | HM42     | Cl8.8  | 646.1       | 430.7     |
	// | HM42     | Cl10.9 | 807.6       | 448.7     |
	// |          |        |             |           |
	// | HM48     | Cl4.6  | 424.6       | 283.0     |
	// | HM48     | Cl4.8  | 424.6       | 235.9     |
	// | HM48     | Cl5.6  | 530.7       | 353.8     |
	// | HM48     | Cl5.8  | 530.7       | 294.8     |
	// | HM48     | Cl6.8  | 636.8       | 353.8     |
	// | HM48     | Cl8.8  | 849.1       | 566.1     |
	// | HM48     | Cl10.9 | 1061.4      | 589.7     |
	// |          |        |             |           |
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

func testCase(t *testing.T, filename, result string) {
	t.Run(filename, func(t *testing.T) {
		// name of test folder
		const folder string = "testdata"
		filename = filepath.Join(folder, filename)

		// for update test screens run in console:
		// UPDATE=true go test
		if os.Getenv("UPDATE") == "true" {
			if err := ioutil.WriteFile(filename, []byte(result), 0644); err != nil {
				t.Fatalf("Cannot write snapshot to file: %v", err)
			}
		}

		// compare datas
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatalf("Cannot read snapshot file: %v", err)
		}
		if !bytes.Equal([]byte(result), content) {
			t.Errorf("Snapshots is not same:\n%s\n%s", result, string(content))
		}
	})
}

func TestShearResistanceCases(t *testing.T) {
	for _, bd := range bolt.GetBoltDiameterList() {
		for _, bc := range bolt.GetBoltClassList() {
			b := bolt.New(bd, bc)

			filename := fmt.Sprintf("ShearResistance%s%s", bd, bc)

			result := boltShearResistance(b)
			testCase(t, filename, result)
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
	for _, bd := range bolt.GetBoltDiameterList() {
		for _, bc := range bolt.GetBoltClassList() {
			b := bolt.New(bd, bc)

			filename := fmt.Sprintf("TensionResistance%s%s", bd, bc)
			result := boltTensionResistance(b)
			testCase(t, filename, result)
		}
	}
}

func TestResistance(t *testing.T) {
	eps := 1e-8
	b := bolt.New(bolt.D24, bolt.G5p8)
	sr := bolt.Resistance{B: b}

	if f, _ := sr.Value(0.0, 0.0, bolt.NoView); float64(f) > eps {
		t.Errorf("Factor can not be not zero if load is zero")
	}
	if f, _ := sr.Value(1e10, 1e10, bolt.FullView); float64(f) < 1.0 {
		t.Errorf("Factor can not be less 1.0 if load is huge")
	}
}
