package bolt

import (
	"fmt"
	"math"
)

// Distance - distances in according to table 3.3 EN1993-1-8.
// Unit - meter
type Distance struct {
	b   Bolt
	thk Dimension
}

// GetDistances - create a struct with all dimensions
// Unit - meter
func GetDistances(b Bolt, thk Dimension) Distance {
	return Distance{b: b, thk: thk}
}

// E1min - dimension e1min in according to table 3.3 EN1993-1-8
func (d Distance) E1min() Dimension {
	return Dimension(float64(d.b.Do().Value()) * 1.2)
}

// E1max - dimension e1max in according to table 3.3 EN1993-1-8
func (d Distance) E1max() Dimension {
	return Dimension(4.0*float64(d.thk) + 40.0e-3)
}

// E2min - dimension e2min in according to table 3.3 EN1993-1-8
func (d Distance) E2min() Dimension {
	return Dimension(float64(d.b.Do().Value()) * 1.2)
}

// E2max - dimension e2max in according to table 3.3 EN1993-1-8
func (d Distance) E2max() Dimension {
	return Dimension(4.0*float64(d.thk) + 40.0e-3)
}

// E3min - dimension e3min in according to table 3.3 EN1993-1-8
func (d Distance) E3min() Dimension {
	return Dimension(float64(d.b.Do().Value()) * 1.5)
}

// E4min - dimension e4min in according to table 3.3 EN1993-1-8
func (d Distance) E4min() Dimension {
	return Dimension(float64(d.b.Do().Value()) * 1.5)
}

// P1min - dimension p1min in according to table 3.3 EN1993-1-8
func (d Distance) P1min() Dimension {
	return Dimension(float64(d.b.Do().Value()) * 2.2)
}

// P1max - dimension p1max in according to table 3.3 EN1993-1-8
func (d Distance) P1max() Dimension {
	return Dimension(math.Min(14.0*float64(d.thk), 200.0e-3))
}

// P10max - dimension p10max in according to table 3.3 EN1993-1-8
func (d Distance) P10max() Dimension {
	return Dimension(math.Min(14.0*float64(d.thk), 200.0e-3))
}

// P1imax - dimension p1imax in according to table 3.3 EN1993-1-8
func (d Distance) P1imax() Dimension {
	return Dimension(math.Min(28.0*float64(d.thk), 400.0e-3))
}

// P2min - dimension p2min in according to table 3.3 EN1993-1-8
func (d Distance) P2min() Dimension {
	return Dimension(float64(d.b.Do().Value()) * 2.4)
}

// P2max - dimension p2max in according to table 3.3 EN1993-1-8
func (d Distance) P2max() Dimension {
	return Dimension(math.Min(14.0*float64(d.thk), 200.0e-3))
}

// ShowAllDimensions - print all dimensions in according to table 3.3 EN1993-1-8
// Unit - meter
func ShowAllDimensions(b Bolt, thk Dimension) (s string) {
	d := GetDistances(b, thk)
	s += fmt.Sprintf("E1min  = %s\n", d.E1min())
	s += fmt.Sprintf("E1max  = %s\n", d.E1max())

	s += fmt.Sprintf("E2min  = %s\n", d.E2min())
	s += fmt.Sprintf("E2max  = %s\n", d.E2max())

	s += fmt.Sprintf("E3min  = %s\n", d.E3min())

	s += fmt.Sprintf("E4min  = %s\n", d.E4min())

	s += fmt.Sprintf("P1min  = %s\n", d.P1min())
	s += fmt.Sprintf("P1max  = %s\n", d.P1max())

	s += fmt.Sprintf("P1max  = %s\n", d.P1max())

	s += fmt.Sprintf("P10max = %s\n", d.P10max())

	s += fmt.Sprintf("P1imax = %s\n", d.P1imax())

	s += fmt.Sprintf("P2min  = %s\n", d.P2min())
	s += fmt.Sprintf("P2max  = %s\n", d.P2max())

	return
}
