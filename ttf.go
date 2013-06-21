package allegro

/*
#cgo pkg-config: allegro_ttf-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_ttf.h>
*/
import "C"
import "unsafe"

const (
	TTF_NO_KERNING  = C.ALLEGRO_TTF_NO_KERNING
	TTF_MONOCHROME  = C.ALLEGRO_TTF_MONOCHROME
	TTF_NO_AUTOHINT = C.ALLEGRO_TTF_NO_AUTOHINT
)

func InitTTF() {

	C.al_init_ttf_addon()

}

func ShutdownTTF() {

	C.al_shutdown_ttf_addon()

}

func LoadTTFFont(fname string, size, flags int) *Font {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))

	var f *Font

	f = (*Font)(C.al_load_ttf_font(cfname, C.int(size), C.int(flags)))

	return f
}

func LoadTTFFontStretch(fname string, w, h, flags int) *Font {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))

	var f *Font

	f = (*Font)(C.al_load_ttf_font_stretch(cfname, C.int(w), C.int(h), C.int(flags)))

	return f
}
