package intstrcodec

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	alphabet  = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	blockSize = 20
	minLength = 5
)

func TestEncodeDecodeRanges(t *testing.T) {
	codec, err := CreateCodec(alphabet, blockSize, minLength)
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
		{start: 9223372036854774806, end: 9223372036854775806},
	}

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			t.Run(fmt.Sprintf("Test IntToStr and StrToInt with input %d", i), func(t *testing.T) {
				encoded := codec.IntToStr(i)
				decoded := codec.StrToInt(encoded)
				if decoded != i {
					t.Errorf("Decode failed: input=%d, encoded=%s, decoded=%d", i, encoded, decoded)
				}
			})
		}
	}
}

func TestMinLength(t *testing.T) {
	for ml := 4; ml <= 32; ml++ {
		codec, err := CreateCodec(alphabet, blockSize, ml)
		if err != nil {
			t.Fatalf("Failed to create codec: %v", err)
		}

		for i := 0; i < 10; i++ {
			input := rand.Intn(100)
			encoded := codec.IntToStr(input)
			if len(encoded) < ml {
				t.Errorf("Failed: Expected padded length %d, got %d", ml, len(encoded))
			}
		}
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	cc, err := CreateCodec(alphabet, blockSize, minLength)
	if err != nil {
		b.Fatalf("Failed to create codec: %v", err)
	}

	for i := 0; i < b.N; i++ {
		input := rand.Intn(1000000000)
		encoded := cc.IntToStr(input)
		cc.StrToInt(encoded)
	}
}
