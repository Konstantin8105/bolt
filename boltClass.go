package bolt

import "fmt"

// boltClass is class of bolt
type boltClass string

const (
	g4_6  boltClass = "4.6"
	g4_8  boltClass = "4.8"
	g5_6  boltClass = "5.6"
	g5_8  boltClass = "5.8"
	g6_8  boltClass = "6.8"
	g8_8  boltClass = "8.8"
	g10_9 boltClass = "10.9"
)

func (bc boltClass) String() string {
	return fmt.Sprintf("Bolt class : Cl %s", string(bc))
}

func (bc boltClass) description() string {
	return bc.String()
}
