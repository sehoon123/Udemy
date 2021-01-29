package word

import (
	"testing"

	"github.com/sehoon123/Hands-on/13-2/quote"
)

func BenchmarkUseCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
func BenchmarkCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
