package intstrcodec

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	alphabet  = "bkus2y8ng9dch5xam3t6r7pqe4zfwjv"
	blockSize = 10
)

func TestEncodeDecodeRanges(t *testing.T) {
	codec, err := New(alphabet, blockSize)
	if err != nil {
		t.Fatal("failed to initialize codec")
	}

	ranges := []struct {
		start, end int
	}{
		{start: 0, end: 1000},
		{start: 10000, end: 11000},
		{start: 1000000, end: 1001000},
		{start: 100000000, end: 100001000},
		{start: 10000000000, end: 10000001000},
		{start: 1000000000000, end: 1000000001000},
		{start: 100000000000000, end: 100000000001000},
		{start: 10000000000000000, end: 10000000000001000},
		{start: 9223372036854774807, end: 9223372036854775807},
	}

	for _, r := range ranges {
		for input := r.start; input < r.end; input++ {
			t.Run(fmt.Sprintf("test encode and decode with input %d", input), func(t *testing.T) {
				encoded := codec.Encode(input)
				decoded := codec.Decode(encoded)
				if decoded != input {
					t.Errorf("decode failed: input=%d, encoded=%s, decoded=%d", input, encoded, decoded)
				}
			})
		}
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	cc, err := New(alphabet, blockSize)
	if err != nil {
		b.Fatalf("failed to initialize codec: %v", err)
	}

	for i := 0; i < b.N; i++ {
		input := rand.Intn(1000000000)
		encoded := cc.Encode(input)
		cc.Decode(encoded)
	}
}