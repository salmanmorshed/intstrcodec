package intstrcodec

import (
	"errors"
	"math"
	"strings"
)

type CodecConfig struct {
	alphabet  string
	blockSize uint
	minLen    int
	mask      uint
	mapping   []int
}

func CreateCodec(alphabet string, blockSize uint, minLen ...int) (*CodecConfig, error) {
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet must contain at least 2 characters")
	}

	var mLen int
	if len(minLen) > 0 {
		mLen = minLen[0]
	} else {
		mLen = 5
	}
	cc := &CodecConfig{
		alphabet:  alphabet,
		blockSize: blockSize,
		minLen:    mLen,
		mask:      (1 << blockSize) - 1,
	}
	cc.mapping = make([]int, blockSize)
	for i := uint(0); i < blockSize; i++ {
		cc.mapping[i] = int(blockSize - 1 - i)
	}
	return cc, nil
}

func (cc *CodecConfig) encodeInt(n int, minLen int) string {
	return cc.enbase(cc.encode(n), minLen)
}

func (cc *CodecConfig) decodeStr(n string) int {
	return cc.decode(cc.debase(n))
}

func (cc *CodecConfig) encode(n int) int {
	return (n & (^int(cc.mask))) | cc._encode(n&int(cc.mask))
}

func (cc *CodecConfig) _encode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<uint(i)) != 0 {
			result |= (1 << uint(b))
		}
	}
	return result
}

func (cc *CodecConfig) decode(n int) int {
	return (n & (^int(cc.mask))) | cc._decode(n&int(cc.mask))
}

func (cc *CodecConfig) _decode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<uint(b)) != 0 {
			result |= (1 << uint(i))
		}
	}
	return result
}

func (cc *CodecConfig) enbase(x int, minLen int) string {
	result := cc._enbase(x)
	paddingLength := minLen - len(result)
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

func (cc *CodecConfig) IntToStr(x int) string {
	return cc.encodeInt(x, cc.minLen)
}

func (cc *CodecConfig) StrToInt(x string) int {
	return cc.decodeStr(x)
}
