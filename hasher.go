package hasher

import (
	"hash"
	"sync"

	"github.com/cespare/xxhash"
)

type (
	Hasher interface {
		String(string) uint64

		Paths(uint64) (string, string, string)
		Path(uint64) string
		PathRoot(string, uint64) string
	}

	u struct {
		h hash.Hash64
	}

	s struct {
		u *u
		m sync.Mutex
	}
)

func new() *u {
	rv := &u{
		h: xxhash.New(),
	}

	return rv
}

func Unsafe() Hasher {
	return new()
}

func Safe() Hasher {
	rv := &s{
		u: new(),
	}
	return rv
}
