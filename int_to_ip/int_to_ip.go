package int_to_ip

import (
	"strconv"
	"strings"
)

func Int32ToIp(n uint32) string {
	var mask, part uint32
	var res string

	mask = 255 << 24
	for i := 0; i < 4; i++ {
		part = (n & mask) >> (24 - 8*i)
		mask >>= 8

		res += strconv.Itoa(int(part)) + "."
	}

	return strings.Trim(res, ".")
}
