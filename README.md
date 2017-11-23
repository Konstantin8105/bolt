# Calculation of bolt


[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/Eurocode3.Bolt/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/Eurocode3.Bolt?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/Eurocode3.Bolt.svg?branch=master)](https://travis-ci.org/Konstantin8105/Eurocode3.Bolt)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/Eurocode3.Bolt)](https://goreportcard.com/report/github.com/Konstantin8105/Eurocode3.Bolt)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/Eurocode3.Bolt/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/Eurocode3.Bolt?status.svg)](https://godoc.org/github.com/Konstantin8105/Eurocode3.Bolt)

Calculation bolt property in according to Eurocode 3.

Example of shear resistance calculation for bolt HM24Cl5.8:
```go
func ExampleShearResistance() {
	b := bolt.New(bolt.G5p8, bolt.D24)
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
