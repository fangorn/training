package decode_bits

import (
	"testing"
	"training/morse/decode"
)

func TestDecodeBits(t *testing.T) {
	expectedMessage := "HEY JUDE"

	decodedBits := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	decodedMessage := decode.DecodeMorse(DecodeBits(decodedBits))

	if decodedMessage != expectedMessage {
		t.Errorf("Expected %s, got %s", expectedMessage, decodedMessage)
	}
}
