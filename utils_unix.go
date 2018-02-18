// +build darwin dragonfly freebsd netbsd openbsd

package main

import (
	"reflect"
	"syscall"
	"unsafe"
)

func nameFromDirent(de *syscall.Dirent) []byte {
	ml := int(de.Namlen)

	var name []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&name))
	sh.Cap = ml
	sh.Len = ml
	sh.Data = uintptr(unsafe.Pointer(&de.Name[0]))

	return name
}
