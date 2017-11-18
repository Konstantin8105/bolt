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
		if fmt.Sprintf("%s", bc) == "" {
			t.Fatalf("Cannot convert to string: %v", bc)
		}
	}
}

func TestBoltDiameter(t *testing.T) {
	for pos, db := range bolt.GetBoltDiameterList() {
		if db <= 0.0 {
			t.Fatalf("Diameter of bolt cannot be : %v. See position in array: %d", float64(db), pos)
		}
		if fmt.Sprintf("%s", db) == "" {
			t.Fatalf("Cannot convert to string: %v", db)
		}
	}
}

func TestFyb(t *testing.T) {
	for pos, bc := range bolt.GetBoltClassList() {
		var fyb = bolt.Fyb{Bc: bc}
		if fyb.Value() <= 0.0 {
			t.Fatalf("Fyb cannot be : %v. See position in array: %d", fyb, pos)
		}
		if fmt.Sprintf("%s", fyb) == "" {
			t.Fatalf("Cannot convert to string: %v", fyb)
		}
	}
}
