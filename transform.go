package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import "unsafe"
import "reflect"

type Transform C.ALLEGRO_TRANSFORM

func (t *Transform) GetMatrix() []float32 {
	length := 16
	data := (C.ALLEGRO_TRANSFORM)(*t).m

	var slice []float32
	// This make slice
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&slice)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(&data))
	return slice
}

func (t *Transform) Copy() *Transform {
	var cp Transform
	C.al_copy_transform((*C.ALLEGRO_TRANSFORM)(&cp), (*C.ALLEGRO_TRANSFORM)(t))
	return &cp
}

// Need to keep a reference to stop it being garbage collected (i think)
var current *Transform = nil

func (t *Transform) Use() {
	C.al_use_transform((*C.ALLEGRO_TRANSFORM)(t))
	current = t
}

func GetCurrentTransform() *Transform {
	return (*Transform)(C.al_get_current_transform())
}

func (t *Transform) Invert() {
	C.al_invert_transform((*C.ALLEGRO_TRANSFORM)(t))
}

func (t *Transform) HasInverse(tolerance float32) bool {
	return C.al_check_inverse((*C.ALLEGRO_TRANSFORM)(t), C.float(tolerance)) != C.int(0)
}

func (t *Transform) Identity() {
	C.al_identity_transform((*C.ALLEGRO_TRANSFORM)(t))
}

func (t *Transform) Build(x, y, sx, sy, theta float32) {
	C.al_build_transform((*C.ALLEGRO_TRANSFORM)(t),
		C.float(x), C.float(y), C.float(sx), C.float(sy), C.float(theta))
}

func (t *Transform) Translate(x, y float32) {
	C.al_translate_transform((*C.ALLEGRO_TRANSFORM)(t), C.float(x), C.float(y))
}

func (t *Transform) Compose(o *Transform) {
	C.al_compose_transform((*C.ALLEGRO_TRANSFORM)(t), (*C.ALLEGRO_TRANSFORM)(o))
}