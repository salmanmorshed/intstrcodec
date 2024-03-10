package intstrcodec

import (
	"errors"
	"slices"
	"strings"
)

type Codec struct {
	alphabet  string
	blockSize int
	blockMask int
}

func New(alphabet string, blockSize int) (*Codec, error) {
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet must contain at least 2 characters")
	}

	for _, r := range alphabet {
		if r >= 128 {
			return nil, errors.New("alphabet must contain only ASCII characters")
		}
	}

	if blockSize < 0 {
		return nil, errors.New("blockSize must be a positive integer")
	}

	cc := Codec{
		alphabet:  alphabet,
		blockSize: blockSize,
		blockMask: (1 << blockSize) - 1,
	}

	return &cc, nil
}

func (cc *Codec) Encode(inputInt int) string {
	return cc.convertToBaseN(cc.reverseLowerBlock(inputInt))
}

func (cc *Codec) Decode(encodedStr string) int {
	return cc.reverseLowerBlock(cc.convertToBase10(encodedStr))
}

func (cc *Codec) reverseLowerBlock(n int) int {
	return (n & (^cc.blockMask)) | cc.reverseZeroPaddedBits(n&cc.blockMask)
}

func (cc *Codec) reverseZeroPaddedBits(num int) int {
	var reversed int
	for range cc.blockSize {
		lastBit := num & 1
		reversed = reversed << 1
		reversed |= lastBit
		num >>= 1
	}
	return reversed
}

func (cc *Codec) convertToBaseN(value10 int) string {
	var result strings.Builder
	var bytes []uint8
	baseN := len(cc.alphabet)
	for value10 > 0 {
		rem := value10 % baseN
		bytes = append(bytes, cc.alphabet[rem])
		value10 /= baseN
	}
	slices.Reverse(bytes)
	result.Write(bytes)
	return result.String()
}

func (cc *Codec) convertToBase10(valueN string) int {
	var result int
	baseN := len(cc.alphabet)
	for _, char := range []byte(valueN) {
		result *= baseN
		result += strings.IndexByte(cc.alphabet, char)
	}
	return result
}