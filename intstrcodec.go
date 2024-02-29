package intstrcodec

import (
	"errors"
	"log"
	"math"
	"strings"
)

const defaultMinLength = 5

type IntPowFn func(int, int) int

type Codec struct {
	alphabet  string
	blockSize int
	minLength int
	mask      int
	mapping   []int
	intPowFn  func(int, int) int
}

type OptFn func(*Codec)

func WithMinLength(ml int) OptFn {
	if ml < 1 {
		log.Fatalln("minimum length value should be greater than zero")
	}
	return func(cc *Codec) {
		cc.minLength = ml
	}
}

func WithIntPowerFn(pf IntPowFn) OptFn {
	return func(cc *Codec) {
		cc.intPowFn = pf
	}
}

func New(alphabet string, blockSize int, optFns ...OptFn) (*Codec, error) {
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet must contain at least 2 characters")
	}

	if blockSize < 0 {
		return nil, errors.New("blockSize must be a positive integer")
	}

	cc := Codec{
		alphabet:  alphabet,
		blockSize: blockSize,
		minLength: defaultMinLength,
		mask:      (1 << blockSize) - 1,
		mapping:   make([]int, blockSize),
		intPowFn:  NativeIntPower,
	}

	for i := range cc.mapping {
		cc.mapping[i] = blockSize - 1 - i
	}

	for _, optFunc := range optFns {
		optFunc(&cc)
	}

	return &cc, nil
}

func (cc *Codec) IntToStr(n int) string {
	return cc.enbase(cc.encode(n))
}

func (cc *Codec) StrToInt(x string) int {
	return cc.decode(cc.debase(x))
}

func (cc *Codec) encode(n int) int {
	return (n & (^cc.mask)) | cc._encode(n&cc.mask)
}

func (cc *Codec) _encode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<i) != 0 {
			result |= 1 << b
		}
	}
	return result
}

func (cc *Codec) decode(n int) int {
	return (n & (^cc.mask)) | cc._decode(n&cc.mask)
}

func (cc *Codec) _decode(n int) int {
	result := 0
	for i, b := range cc.mapping {
		if n&(1<<b) != 0 {
			result |= 1 << i
		}
	}
	return result
}

func (cc *Codec) enbase(x int) string {
	result := cc._enbase(x)
	paddingLength := cc.minLength - len(result)
	if paddingLength <= 0 {
		return result
	}
	padding := strings.Repeat(string(cc.alphabet[0]), paddingLength)
	return padding + result
}

func (cc *Codec) _enbase(x int) string {
	n := len(cc.alphabet)
	if x < n {
		return string(cc.alphabet[x])
	}
	return cc._enbase(x/n) + string(cc.alphabet[x%n])
}

func (cc *Codec) debase(x string) int {
	n := len(cc.alphabet)
	result := 0
	for i := len(x) - 1; i >= 0; i-- {
		c := x[i]
		result += strings.IndexByte(cc.alphabet, c) * cc.intPowFn(n, len(x)-1-i)
	}
	return result
}

func NativeIntPower(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

func CustomIntPower(base, exponent int) int {
	result := 1
	if exponent < 0 {
		return NativeIntPower(base, exponent)
	}
	for exponent > 0 {
		result *= base
		exponent--
	}
	return result
}