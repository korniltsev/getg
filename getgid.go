//go:build go1.19 && !go1.21 && darwin
// +build go1.19,!go1.21,darwin

package getg

import "unsafe"

func GetGID(g unsafe.Pointer) uint64 {
	return *(*uint64)(unsafe.Pointer(uintptr(g) + 152))
}
