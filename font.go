package allegro

/*
#cgo pkg-config: allegro_font-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_font.h>
*/
import "C"
import "unsafe"
import "reflect"

type Font C.ALLEGRO_FONT

const (
	ALIGN_LEFT    = C.ALLEGRO_ALIGN_LEFT
	ALIGN_RIGHT   = C.ALLEGRO_ALIGN_RIGHT
	ALIGN_CENTER  = C.ALLEGRO_ALIGN_CENTER
	ALIGN_INTEGER = C.ALLEGRO_ALIGN_INTEGER
)

func InitFont() {

	C.al_init_font_addon()

}

func ShutdownFont() {

	C.al_shutdown_font_addon()

}

func LoadFont(fname string, size, flags int) *Font {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))
	var f *Font

	f = (*Font)(C.al_load_font(cfname, C.int(size), C.int(flags)))

	return f
}

func (f *Font) Destroy() {

	C.al_destroy_font((*C.ALLEGRO_FONT)(f))

}

func (f *Font) GetLineHeight() int {
	var i int

	i = int(C.al_get_font_line_height((*C.ALLEGRO_FONT)(f)))

	return i
}

func (f *Font) GetAscent() int {
	var i int

	i = int(C.al_get_font_ascent((*C.ALLEGRO_FONT)(f)))

	return i
}

func (f *Font) GetDescent() int {
	var d int

	d = int(C.al_get_font_descent((*C.ALLEGRO_FONT)(f)))

	return d
}

func (f *Font) GetTextWidth(text string) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	var i int

	i = int(C.al_get_text_width((*C.ALLEGRO_FONT)(f), ctext))

	return i
}

func (f *Font) Draw(color Color, x, y float32, flags int, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.al_draw_text((*C.ALLEGRO_FONT)(f), (C.ALLEGRO_COLOR)(color),
		C.float(x), C.float(y), C.int(flags), ctext)

}

func (f *Font) DrawJustified(color Color, x1, x2, y, diff float32, flags int, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.al_draw_justified_text((*C.ALLEGRO_FONT)(f), C.ALLEGRO_COLOR(color),
		C.float(x1), C.float(x2), C.float(y), C.float(diff), C.int(flags), ctext)

}

func (f *Font) GetTextDimensions(text string) (int, int, int, int) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	var x, y, w, h C.int

	C.al_get_text_dimensions((*C.ALLEGRO_FONT)(f), ctext,
		&x, &y, &w, &h)

	return int(x), int(y), int(w), int(h)
}

func GrabFontFromBitmap(bmp *Bitmap, ranges []int) *Font {
	n := len(ranges) / 2 * 2
	data := (*reflect.SliceHeader)(unsafe.Pointer(&ranges)).Data

	var f *Font

	f = (*Font)(C.al_grab_font_from_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.int(n), (*C.int)(unsafe.Pointer(data))))

	return f
}

func LoadBitmapFont(fname string) *Font {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))

	var f *Font

	f = (*Font)(C.al_load_bitmap_font(cfname))

	return f
}

func CreateBuiltinFont() *Font {
	var f *Font

	f = (*Font)(C.al_create_builtin_font())

	return f
}
