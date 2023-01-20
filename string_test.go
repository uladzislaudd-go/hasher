package hasher

import "testing"

func benchmarkString(b *testing.B, h Hasher) {
	for i := 0; i < b.N; i++ {
		h.String("https://www.google.com")
	}
}

func Benchmark_s_String(b *testing.B) {
	benchmarkString(b, Safe())
}

func Benchmark_u_String(b *testing.B) {
	benchmarkString(b, Unsafe())
}
