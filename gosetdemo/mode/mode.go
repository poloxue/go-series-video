package mode

type Mode uint8

const (
	F1 = 1 << iota // 0b00000001
	F2             // 0b00000010
	F3             // 0b00000100
	F4             // 0b00001000
	F5             // 0b00010000
	F6             // 0b00100000
	F7             // 0b01000000
	F8             // 0b10000000
)

func Set(m, f Mode) Mode {
	return m | f
}

func Has(m, f Mode) bool {
	return m&f != 0
}

func Clear(m, f Mode) Mode {
	return m &^ f
}

func Toggle(m, f Mode) Mode {
	return m ^ f
}
