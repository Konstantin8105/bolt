# Calculation bolt property in according to EN1993-1-8.


[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/Eurocode3.Bolt/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/Eurocode3.Bolt?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/Eurocode3.Bolt)](https://goreportcard.com/report/github.com/Konstantin8105/Eurocode3.Bolt)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/Eurocode3.Bolt/blob/master/LICENSE)
[![Go](https://github.com/Konstantin8105/Eurocode3.Bolt/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/Konstantin8105/Eurocode3.Bolt/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Konstantin8105/Eurocode3.Bolt.svg)](https://pkg.go.dev/github.com/Konstantin8105/Eurocode3.Bolt)


Example of bolt property calculation:
```go
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
```


Example of shear resistance calculation for bolt HM24Cl5.8:
```go
func ExampleShearResistance() {
	b := bolt.New(bolt.D24, bolt.G5p8)
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
```
