package slvch

import (
	"testing"
)

func BenchmarkChannel(b *testing.B) {
	ch := make(chan string)

	go func() {
		for i := range ch {
			_ = i
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- "Hello, World!"
	}
	close(ch)
}

func BenchmarkSlice(b *testing.B) {
	sl := make([]string, 0)

	for i := 0; i < b.N; i++ {
		sl = append(sl, "Hello, World!")
	}

	_ = sl
}
