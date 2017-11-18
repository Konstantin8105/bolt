package bolt

import "fmt"

type boltDiameter float64

const (
	D12 boltDiameter = 12.e-3
	D16 boltDiameter = 16.e-3
	D20 boltDiameter = 20.e-3
	D24 boltDiameter = 24.e-3
	D30 boltDiameter = 30.e-3
	D36 boltDiameter = 36.e-3
	D42 boltDiameter = 42.e-3
	D48 boltDiameter = 48.e-3
)

func (bd boldDiameter) String() string {
	return fmt.Sprintf("%.1 mm", float64(bd)*1e3)
}
