package hasher

import (
	"reflect"
	"unsafe"
)

func string_to_byte_unsafe(str string) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&str)).Data)), len(str))
}

func string_to_byte(str string) []byte {
	return []byte(str)
}

func (u *u) string(str string, fn func(string) []byte) uint64 {
	u.h.Reset()
	u.h.Write(fn(str))
	return u.h.Sum64()
}

func (u *u) String(str string) uint64 {
	return u.string(str, string_to_byte_unsafe)
}

func (s *s) String(str string) uint64 {
	s.m.Lock()
	defer s.m.Unlock()
	return s.u.string(str, string_to_byte)
}
