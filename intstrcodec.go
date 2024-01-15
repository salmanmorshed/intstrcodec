package intstrcodec

import (
	"errors"
	"math"
	"strings"
)

type IntStrCodec struct {
	alphabet  string
	blockSize int
	minLength int
	mask      int
	mapping   []int
}

func CreateCodec(alphabet string, blockSize int, minLength int) (*IntStrCodec, error) {
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet must contain at least 2 characters")
	}

	if blockSize < 0 {
		return nil, errors.New("blockSize must be a positive integer")
	}

	cc := IntStrCodec{
		alphabet:  alphabet,
		blockSize: blockSize,
		minLength: minLength,
		mask:      (1 << blockSize) - 1,
		mapping:   make([]int, blockSize),
	}

	for i := range cc.mapping {
		cc.mapping[i] = blockSize - 1 - i
	}

	return &cc, nil
}

func (cc *IntStrCodec) IntToStr(n int) string {
	return cc.enbase(cc.encode(n))
}

func (cc *IntStrCodec) StrToInt(x string) int {
	return cc.decode(cc.debase(x))
}

func (cc *IntStrCodec) encode(n int) int {
	return (n & (^cc.mask)) | cc._encode(n&cc.mask)
}

func (cc *IntStrCodec) _encode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<i) != 0 {
			result |= 1 << b
		}
	}
	return result
}

func (cc *IntStrCodec) decode(n int) int {
	return (n & (^cc.mask)) | cc._decode(n&cc.mask)
}

func (cc *IntStrCodec) _decode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<b) != 0 {
			result |= 1 << i
		}
	}
	return result
}

func (cc *IntStrCodec) enbase(x int) string {
	result := cc._enbase(x)
	paddingLength := cc.minLength - len(result)
	if paddingLength <= 0 {
		return result
	}
	padding := strings.Repeat(string(cc.alphabet[0]), paddingLength)
	return padding + result
}

func (cc *IntStrCodec) _enbase(x int) string {
	n := len(cc.alphabet)
	if x < n {
		return string(cc.alphabet[x])
	}
	return cc._enbase(x/n) + string(cc.alphabet[x%n])
}

func (cc *IntStrCodec) debase(x string) int {
	n := len(cc.alphabet)
	result := 0
	for i := len(x) - 1; i >= 0; i-- {
		c := x[i]
		result += strings.IndexByte(cc.alphabet, c) * intPow(n, len(x)-1-i)
	}
	return result
}

func intPow(a, b int) int {
	// the margin of error introduced due to floating point calculations in this function becomes large enough to
	// cause the codec to break for input values beyond 2^55. todo: explore possible solutions
	return int(math.Pow(float64(a), float64(b)))
}
