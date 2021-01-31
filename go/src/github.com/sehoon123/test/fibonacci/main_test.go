package fibonacci

import "testing"

func BenchmarkFibo1(b *testing.B) {
	b.ResetTimer()
	cache := make(map[int64]int64)
	for i := 0; i < b.N; i++ {

		fib(140, cache)
	}

}
func BenchmarkFibo2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		fiboarray(140)
	}

}
