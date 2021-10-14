package fastabs_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/dangermike/fastabs"
)

func slowabs(v int) int {
	if v >= 0 {
		return v
	}
	return v * -1
}

func TestFastAbs(t *testing.T) {
	for i := 0; i < 10000; i++ {
		value := rand.Int()
		if rand.Intn(2) == 0 {
			value *= -1
		}
		t.Run(strconv.Itoa(value), func(t *testing.T) {
			if slowabs(value) != fastabs.Abs(value) {
				t.Errorf("%d: expected %d, received %d", value, slowabs(value), fastabs.Abs(value))
			}
		})
	}
}

func BenchmarkFastAbs(b *testing.B) {
	source := make([]int, 100000)
	for i := 0; i < len(source); i++ {
		source[i] = rand.Int()
		if rand.Intn(2) == 0 {
			source[i] *= -1
		}
	}
	b.ResetTimer()
	for n := 0; n < 4; n++ {
		b.Run("round_"+strconv.Itoa(n), func(b *testing.B) {
			b.Run("slowabs", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					slowabs(i - 500000000)
				}
			})
			b.Run("fastabs", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fastabs.Abs(i - 500000000)
				}
			})
		})
	}
}
