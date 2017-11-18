package bolt

type stress float64

type (s stress) String() string{
	return fmt.Sprintf("%.1f MPa",float64(s)*1.e-6)
}
