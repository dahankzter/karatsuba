package karatsuba

import (
	"math/big"
)

const (
	THRESHOLD = 1000
)

func Multiply(x, y *big.Int) *big.Int {

	m := min(x.BitLen(), y.BitLen()) / 2

	if m < THRESHOLD {
		z := big.NewInt(0)
		return z.Mul(x, y)
	}

	return big.NewInt(0)
}

func split(x *big.Int, m uint) []*big.Int {
	left := bigint(int64(x.Uint64())).Rsh(x, m)
	right := bigint(int64(x.Uint64())).Sub(x, bigint(int64(left.Uint64())))
	return []*big.Int{left, right}
}

func bigint(n int64) *big.Int {
	return big.NewInt(int64(n))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
