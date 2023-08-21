package intstrcodec

import (
	"fmt"
	"testing"
)

func TestCodecInputValueRanges(t *testing.T) {
	cc, err := CreateCodec("abcdefghijklmnopqrstuvwxyz", 20, 4)

	if err != nil {
		panic(err)
	}

	ranges := []struct {
		start, end int
	}{
		{start: 0, end: 1000},
		{start: 50000, end: 60000},
		{start: 3000000, end: 4000000},
		{start: 100000000, end: 100001000},
		{start: 1000000000, end: 1000001000},
		{start: 100000000000, end: 100000001000},
		{start: 10000000000000, end: 10000000001000},
		{start: 1000000000000000, end: 1000000000001000},
		{start: 1000000000000000000, end: 1000000000000001000},
	}

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			t.Run(fmt.Sprintf("Test IntToStr and StrToInt with input %d", i), func(t *testing.T) {
				encoded := cc.IntToStr(i)
				decoded := cc.StrToInt(encoded)

				if decoded != i {
					t.Errorf("Input: %d, Encoded: %s, Decoded: %d", i, encoded, decoded)
				}
			})
		}
	}
}
