package bolt_test

import (
	"fmt"

	"github.com/Konstantin8105/bolt"
)

func ExampleDistance() {
	b := bolt.New(bolt.D24, bolt.G5p8)
	fmt.Println(bolt.ShowAllDimensions(b, bolt.Dimension(40e-3)))

	// Output:
	// E1min  = 31.2 mm
	// E1max  = 200.0 mm
	// E2min  = 31.2 mm
	// E2max  = 200.0 mm
	// E3min  = 39.0 mm
	// E4min  = 39.0 mm
	// P1min  = 57.2 mm
	// P1max  = 200.0 mm
	// P1max  = 200.0 mm
	// P10max = 200.0 mm
	// P1imax = 400.0 mm
	// P2min  = 62.4 mm
	// P2max  = 200.0 mm
}
