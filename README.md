# Calculation of bolt

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

