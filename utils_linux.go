// +build nacl linux solaris

package main

import (
	"bytes"
	"reflect"
	"syscall"
	"unsafe"
)

func nameFromDirent(de *syscall.Dirent) []byte {
	ml := int(uint64(de.Reclen) - uint64(unsafe.Offsetof(syscall.Dirent{}.Name)))

	var name []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&name))
	sh.Cap = ml
	sh.Len = ml
	sh.Data = uintptr(unsafe.Pointer(&de.Name[0]))

	if index := bytes.IndexByte(name, 0); index >= 0 {
		sh.Cap = index
		sh.Len = index
	}

	return name
}
