package hasher

import (
	"fmt"
	"path/filepath"
	"testing"
)

func testPath(t *testing.T, h Hasher) {
	tests := []struct {
		d   uint64
		rv1 string
		rv2 string
		rv3 string
	}{
		{0x00000000000000ff, "00", "00", "0000000000ff"},
		{0x9a81c701c221b571, "9a", "81", "c701c221b571"},
	}

	roots := []string{"/home/user/"}

	for i, tt := range tests {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			want := filepath.Join(tt.rv1, tt.rv2, tt.rv3)
			if rv := h.Path(tt.d); rv != want {
				t.Errorf("h.Path() = %v, want %v", rv, want)
			}
		})
		t.Run(name, func(t *testing.T) {
			rv1, rv2, rv3 := h.Paths(tt.d)
			if rv1 != tt.rv1 {
				t.Errorf("h.Paths() got = %v, want %v", rv1, tt.rv1)
			}
			if rv2 != tt.rv2 {
				t.Errorf("h.Paths() got = %v, want %v", rv2, tt.rv2)
			}
			if rv3 != tt.rv3 {
				t.Errorf("h.Paths() got = %v, want %v", rv3, tt.rv3)
			}
		})
		for _, root := range roots {
			t.Run(name, func(t *testing.T) {
				want := filepath.Join(root, tt.rv1, tt.rv2, tt.rv3)
				if rv := h.PathRoot(root, tt.d); rv != want {
					t.Errorf("h.PathRoot() = %v, want %v", rv, want)
				}
			})
		}
	}
}

func Test_s_Path(t *testing.T) {
	testPath(t, Safe())
}

func Test_u_Path(t *testing.T) {
	testPath(t, Unsafe())
}

func benchmarkPath(b *testing.B, h Hasher) {
	for i := 0; i < b.N; i++ {
		h.Path(0x9a81c701c221b571)
	}
}

func Benchmark_s_Path(b *testing.B) {
	benchmarkPath(b, Safe())
}

func Benchmark_u_Path(b *testing.B) {
	benchmarkPath(b, Unsafe())
}

func benchmarkPaths(b *testing.B, h Hasher) {
	for i := 0; i < b.N; i++ {
		h.Paths(0x9a81c701c221b571)
	}
}

func Benchmark_s_Paths(b *testing.B) {
	benchmarkPaths(b, Safe())
}

func Benchmark_u_Paths(b *testing.B) {
	benchmarkPaths(b, Unsafe())
}

func benchmarkPathRoot(b *testing.B, h Hasher) {
	for i := 0; i < b.N; i++ {
		h.PathRoot("/home/user/", 0x9a81c701c221b571)
	}
}

func Benchmark_s_PathRoot(b *testing.B) {
	benchmarkPathRoot(b, Safe())
}

func Benchmark_u_PathRoot(b *testing.B) {
	benchmarkPathRoot(b, Unsafe())
}
