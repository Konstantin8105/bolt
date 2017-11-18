package bolt_test

import (
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
		if db <= 0.0 {
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
