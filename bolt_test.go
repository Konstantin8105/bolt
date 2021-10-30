package bolt_test

import (
	"fmt"
	"os"
	"testing"

	bolt "github.com/Konstantin8105/Eurocode3.Bolt"
	"github.com/Konstantin8105/description"
)

func TestBoltClass(t *testing.T) {
	for pos, bc := range bolt.GetBoltClassList() {
		if string(bc) == "" {
			t.Fatalf("Bolt class  haven`t name. Position in array = %d", pos)
		}
		if bc.String() == "" {
			t.Fatalf("Cannot convert to string: %v", bc)
		}
		if pos > 0 {
			present := bolt.Fub{BoltClass: bc}
			last := bolt.Fub{BoltClass: bolt.GetBoltClassList()[pos-1]}
			if float64(present.Value()) < float64(last.Value()) {
				t.Fatalf("Next class of bolt is not strong")
			}
		}
	}
}

func TestBoltDiameter(t *testing.T) {
	for pos, db := range bolt.GetBoltDiameterList() {
		if db <= 0.0 || db > 0.1 {
			t.Fatalf("Diameter of bolt cannot be : %v. See position in array: %d", float64(db), pos)
		}
		if db.String() == "" {
			t.Fatalf("Cannot convert to string: %v", db)
		}
		if pos > 0 {
			last := float64(bolt.GetBoltDiameterList()[pos-1])
			present := float64(db)
			if last > present {
				t.Fatalf("Next diameter is not more")
			}
		}
	}
}

func TestFyb(t *testing.T) {
	for pos, bc := range bolt.GetBoltClassList() {
		var fyb = bolt.Fyb{BoltClass: bc}
		if fyb.Value() <= 0.0 {
			t.Fatalf("Fyb cannot be : %v. See position in array: %d", fyb, pos)
		}
		if fyb.String() == "" {
			t.Fatalf("Cannot convert to string: %v", fyb)
		}
	}
}

func TestBolt(t *testing.T) {
	for _, bd := range bolt.GetBoltDiameterList() {
		for _, bc := range bolt.GetBoltClassList() {
			b := bolt.New(bd, bc)
			if float64(b.Fyb().Value()) <= 0.0 {
				t.Fatal("Cannot be Fyb <= 0.0")
			}
			if float64(b.D()) <= 0.0 {
				t.Fatal("Cannot be D <= 0.0")
			}
			if string(b.Cl()) == "" {
				t.Fatal("Cannot Cl == \"\"")
			}
		}
	}
}

func ExampleBolt() {
	b := bolt.New(bolt.D24, bolt.G8p8)
	fmt.Printf("Bolt : %s%s\n", b.D(), b.Cl())
	fmt.Printf("%s\n", b.Do())
	fmt.Printf("Hole : %s\n", b.Do().Value())
	fmt.Printf("%s\n", b.Fyb())
	fmt.Printf("Fyb  : %s\n", b.Fyb().Value())
	fmt.Printf("%s\n", b.Fub())
	fmt.Printf("Fub  : %s\n", b.Fub().Value())
	fmt.Printf("%s\n", b.As())
	fmt.Printf("%s\n", b.A())

	// Output:
	// Bolt : HM24Cl8.8
	// For bolt HM24 hole is Ø26.0 mm
	// Hole : Ø26.0 mm
	// In according to table 3.1 EN1993-1-8 value Fyb is 640.0 MPa
	// Fyb  : 640.0 MPa
	// In according to table 3.1 EN1993-1-8 value Fub is 800.0 MPa
	// Fub  : 800.0 MPa
	// Tension stress area of the bolt HM24 is 352.8 mm²
	// The gross cross-section area of the bolt HM24 is 452.4 mm²
}

func boltProperty(b bolt.Bolt) (s string) {
	s += fmt.Sprintf("Bolt : %s%s\n", b.D(), b.Cl())
	s += fmt.Sprintf("%s\n", b.Do())
	s += fmt.Sprintf("Hole : %s\n", b.Do().Value())
	s += fmt.Sprintf("%s\n", b.Fyb())
	s += fmt.Sprintf("Fyb  : %s\n", b.Fyb().Value())
	s += fmt.Sprintf("%s\n", b.Fub())
	s += fmt.Sprintf("Fub  : %s\n", b.Fub().Value())
	s += fmt.Sprintf("%s\n", b.As())
	s += fmt.Sprintf("%s\n", b.A())
	return
}

func TestCases(t *testing.T) {
	for _, bd := range bolt.GetBoltDiameterList() {
		for _, bc := range bolt.GetBoltClassList() {
			b := bolt.New(bd, bc)

			// create filename of test
			filename := fmt.Sprintf("%s%s", bd, bc)
			result := boltProperty(b)
			testCase(t, filename, result)
		}
	}
}

func TestDescription(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	descr, err := description.New(dir)
	if err != nil {
		t.Fatal(err)
	}
	rep, err := descr.Report()
	if err != nil {
		t.Logf("%s", dir)
		t.Fatal(err)
	}
	t.Logf("%s", rep)
}
