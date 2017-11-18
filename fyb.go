package bolt

import "fmt"

type fycS struct {
	bc boltClass
}

func (f fycS) value() stress {
	var fyb = map[boltClass]stress{
		g4_6:  240.e6,
		g4_8:  320.e6,
		g5_6:  300.e6,
		g5_8:  400.e6,
		g6_8:  480.e6,
		g8_8:  640.e6,
		g10_9: 900.e6,
	}
	return fyb[f.bc]
}

func (f fycS) String() string {
	return fmt.Sprintf("In according to table 3.1 EN1993-1-8 value Fyb is %s", f.value())
}
