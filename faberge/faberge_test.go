package faberge

import (
	"math/big"
	"testing"
)

func TestHeight(t *testing.T) {
	tests := []struct {
		n, m, expected int64
	}{
		{0, 14, 0},
		{2, 0, 0},
		{2, 14, 105},
		{7, 20, 137979},
	}

	for _, test := range tests {
		result := Height(big.NewInt(test.n), big.NewInt(test.m))
		expected := big.NewInt(test.expected)

		if result.Cmp(expected) != 0 {
			t.Errorf("Height(%d, %d) = %s, expected %s",
				test.n, test.m, result, expected)
		}
	}
}
