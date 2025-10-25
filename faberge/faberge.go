package faberge

import (
	"math/big"
)

// Faberge Eggs crush test
// https://www.codewars.com/kata/54cb771c9b30e8b5250011d4
func Height(n, m *big.Int) *big.Int {
	res := new(big.Int)

	prev := big.NewInt(1)
	cur := big.NewInt(0)
	coef := new(big.Int)
	for i := big.NewInt(1); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		coef.SetInt64(m.Int64() - i.Int64() + 1)
		cur.Mul(prev, coef)
		cur.Div(cur, i)

		res.Add(res, cur)
		prev.Set(cur)
	}

	return res
}
