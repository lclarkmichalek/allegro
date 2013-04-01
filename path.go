package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import "unsafe"

type Path struct { ptr unsafe.Pointer }

func fromPth(p Path) *[0]byte {
	return (*[0]byte)(p.ptr)
}

func toPth(p *[0]byte) Path {
	return Path{ptr: unsafe.Pointer(p)}
}