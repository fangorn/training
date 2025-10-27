package decode_bits

import "strings"

func DecodeBits(bits string) (res string) {
	res = strings.Trim(bits, "0")
	var prev string
	var l int
	var sl []int

	for i, bit := range res {
		if i == 0 {
			prev = string(bit)
			l = 1
			continue
		}

		if string(bit) == prev {
			l++
			continue
		}

		sl = append(sl, l)
		l = 1
		prev = string(bit)
	}

	if len(sl) == 0 {
		if l > 0 {
			sl = append(sl, l)
		} else {
			return ""
		}
	}

	minl := sl[0]
	for _, l := range sl {
		if l < minl {
			minl = l
		}
	}

	var dot, dash, space, div string
	for i := 0; i < minl; i++ {
		dot += "1"
		dash += "111"
		div += "0"
		space += "000"
	}
	res = strings.ReplaceAll(res, dash, "-")
	res = strings.ReplaceAll(res, space, "  ")
	res = strings.ReplaceAll(res, dot, ".")
	res = strings.ReplaceAll(res, div, "")

	return res
}
