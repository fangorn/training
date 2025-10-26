package decode

import (
	"strings"
)

func DecodeMorse(morseCode string) string {
	morse := strings.Trim(morseCode, " ")
	morseWords := strings.Split(morse, "   ")

	res := ""
	for _, mw := range morseWords {
		mletters := strings.Split(mw, " ")
		for _, ml := range mletters {
			res = res + MORSE_CODE[string(ml)]
		}
		res = res + " "
	}
	return strings.Trim(res, " ")
}
