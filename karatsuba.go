package karatsuba

import (
	//"fmt"
	"math/big"
)

const (
	THRESHOLD = 1000
)

func Multiply(x, y *big.Int) *big.Int {

	x1 := big.NewInt(x.Int64())
	y1 := big.NewInt(x.Int64())

	m := pivot(x1, y1)

	if m < THRESHOLD {
		z := big.NewInt(0)
		return z.Mul(x1, y1)
	}

	return big.NewInt(0)
}

func split(x *big.Int, m uint) []*big.Int {
	x1 := big.NewInt(x.Int64())
	left := bigint(int64(x.Uint64())).Rsh(x1, m)

	t1 := bigint(left.Int64())
	t1 = t1.Lsh(t1, m)

	t3 := bigint(x.Int64())
	right := t3.Sub(t3, t1)

	return []*big.Int{left, right}
}

func bigint(n int64) *big.Int {
	return big.NewInt(n)
}

func pivot(a, b *big.Int) int {
	return min(a.BitLen(), b.BitLen()) / 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
