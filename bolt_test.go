package bolt_test

import (
	"fmt"
	"testing"

	bolt "github.com/Konstantin8105/Eurocode3.Bolt"
)

func TestBoltClass(t *testing.T) {
	for pos, bc := range bolt.GetBoltClassList() {
		if string(bc) == "" {
			t.Fatalf("Bolt class  haven`t name. Position in array = %d", pos)
		}
		if bc.String() == "" {
			t.Fatalf("Cannot convert to string: %v", bc)
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
			b := bolt.New(bc, bd)
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
	b := bolt.New(bolt.G8p8, bolt.D24)
	fmt.Printf("Bolt : %s%s\n", b.D(), b.Cl())
	fmt.Printf("%s\n", b.Do())
	fmt.Printf("Hole : %s\n", b.Do().Value())
	fmt.Printf("%s\n", b.Fyb())
	fmt.Printf("Fyb  : %s\n", b.Fyb().Value())
	// Output:
	// Bolt : HM24Cl8.8
	// For bolt HM24 hole is Ø26.0 mm
	// Hole : Ø26.0 mm
	// In according to table 3.1 EN1993-1-8 value Fyb is 640.0 MPa
	// Fyb  : 640.0 MPa
}
