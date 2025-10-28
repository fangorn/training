package duration_format

import (
	"math"
	"strconv"
	"strings"
)

type timeUnit struct {
	singular string
	plural   string
	period   int64
}

var timeUnits []timeUnit = []timeUnit{
	timeUnit{
		singular: "second",
		plural:   "seconds",
		period:   60,
	},
	timeUnit{
		singular: "minute",
		plural:   "minutes",
		period:   60,
	},
	timeUnit{
		singular: "hour",
		plural:   "hours",
		period:   24,
	},
	timeUnit{
		singular: "day",
		plural:   "days",
		period:   365,
	},
	timeUnit{
		singular: "year",
		plural:   "years",
		period:   math.MaxInt64,
	},
}

// https://www.codewars.com/kata/52742f58faf5485cae000b9a
func FormatDuration(seconds int64) string {
	parts := make([]string, 5, 5)
	var value, t int64

	t = seconds
	for i, tu := range timeUnits {
		value = t % tu.period

		if value == 1 {
			parts[4-i] = "1 " + tu.singular
		} else if value > 1 {
			parts[4-i] = strconv.FormatInt(value, 10) + " " + tu.plural
		}

		t /= tu.period
	}

	for i := 0; i < len(parts); {
		if parts[i] == "" {
			parts = append(parts[:i], parts[i+1:]...)
		} else {
			i++
		}
	}

	var res string
	if len(parts) >= 2 {
		res = strings.Join(parts[:len(parts)-1], ", ") + " and "
	}
	res = res + parts[len(parts)-1]

	return res
}
