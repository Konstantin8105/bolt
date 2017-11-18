# Calculation of bolt


[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/Eurocode3.Bolt/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/Eurocode3.Bolt?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/Eurocode3.Bolt.svg?branch=master)](https://travis-ci.org/Konstantin8105/Eurocode3.Bolt)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/Eurocode3.Bolt)](https://goreportcard.com/report/github.com/Konstantin8105/Eurocode3.Bolt)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/Eurocode3.Bolt/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/Eurocode3.Bolt?status.svg)](https://godoc.org/github.com/Konstantin8105/Eurocode3.Bolt)

How to do that with good view and good result.

Generally we have formules, like:
```
c := a * b
```
We want to convert to another view:
```
Calculation of `c` by formula `a*b` in according to formula ....
Calculation: 
	a := 12 item
	b := 4 pieces
	c := a * b = 12 * 4 = 48 pieces
```
Convertion must be automatically.

Let's continue:
```
d := c * a
```
Convertion:
```
Caluclation `d` by formula `c * a` in according to formula ....
Calculation:
	a := 12 item
	c := 48 pieces
	d := c * a = 12 * 48 = 576 pieces
```
In that case, we see value `c` not in detail.

