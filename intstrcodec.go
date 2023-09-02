package intstrcodec

import (
	"errors"
	"math"
	"strings"
)

type CodecConfig struct {
	alphabet  string
	blockSize int
	minLength int
	mask      int
	mapping   []int
}

func CreateCodec(alphabet string, blockSize int, minLength ...int) (*CodecConfig, error) {
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet must contain at least 2 characters")
	}

	if blockSize < 1 {
		return nil, errors.New("blockSize should be a positive integer")
	}

	cc := &CodecConfig{
		alphabet:  alphabet,
		blockSize: blockSize,
		minLength: 5,
		mask:      (1 << blockSize) - 1,
	}

	if len(minLength) > 0 {
		cc.minLength = minLength[0]
	}

	cc.mapping = make([]int, blockSize)
	for i := 0; i < blockSize; i++ {
		cc.mapping[i] = blockSize - 1 - i
	}
	return cc, nil
}

func (cc *CodecConfig) IntToStr(n int) string {
	return cc.enbase(cc.encode(n))
}

func (cc *CodecConfig) StrToInt(x string) int {
	return cc.decode(cc.debase(x))
}

func (cc *CodecConfig) encode(n int) int {
	return (n & (^cc.mask)) | cc._encode(n&cc.mask)
}

func (cc *CodecConfig) _encode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<i) != 0 {
			result |= (1 << b)
		}
	}
	return result
}

func (cc *CodecConfig) decode(n int) int {
	return (n & (^cc.mask)) | cc._decode(n&cc.mask)
}

func (cc *CodecConfig) _decode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<b) != 0 {
			result |= (1 << i)
		}
	}
	return result
}

func (cc *CodecConfig) enbase(x int) string {
	result := cc._enbase(x)
	paddingLength := cc.minLength - len(result)
	if paddingLength <= 0 {
		return result
	}
	padding := strings.Repeat(string(cc.alphabet[0]), paddingLength)
	return padding + result
}

func (cc *CodecConfig) _enbase(x int) string {
	n := len(cc.alphabet)
	if x < n {
		return string(cc.alphabet[x])
	}
	return cc._enbase(x/n) + string(cc.alphabet[x%n])
}

func (cc *CodecConfig) debase(x string) int {
	n := len(cc.alphabet)
	result := 0
	for i := len(x) - 1; i >= 0; i-- {
		c := x[i]
		result += strings.IndexByte(cc.alphabet, c) * int(math.Pow(float64(n), float64(len(x)-1-i)))
	}
	return result
}
