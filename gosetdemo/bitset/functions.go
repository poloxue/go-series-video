package bitset

const (
	shift = 6    // 2^n = 64  中的 n, 即 6
	mask  = 0x3f // 2^n - 1，即 63，即 0x3f
)

// 所在的索引
func index(n int) int {
	return n >> shift
}

func posVal(n int) uint64 {
	return 1 << uint64(n&mask)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
