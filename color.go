package allegro

/*
#cgo pkg-config: allegro_color-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_color.h>
*/
import "C"
import "unsafe"

func CreateColorCMYK(c, m, y, k float32) Color {
	col := C.al_color_cmyk(C.float(c), C.float(m), C.float(y), C.float(k))
	return Color{col.r, col.g, col.b, col.a}
}

func CreateColorHSL(h, s, l float32) Color {
	return Color(C.al_color_hsl(C.float(h), C.float(s), C.float(l)))
}

func CreateColorHSV(h, s, v float32) Color {
	return Color(C.al_color_hsv(C.float(h), C.float(s), C.float(v)))
}

func CreateColorHTML(html string) Color {
	cs := C.CString(html)
	defer C.free(unsafe.Pointer(cs))

	return Color(C.al_color_html(cs))
}

func CreateColorName(name string) Color {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return Color(C.al_color_name(cs))
}

func CreateColorYUV(y, u, v float32) Color {
	return Color(C.al_color_yuv(C.float(y), C.float(u), C.float(v)))
}

func (col Color) GetCMYK() (float32, float32, float32, float32) {
	var c, m, y, k C.float
	cc := C.ALLEGRO_COLOR(col)
	C.al_color_rgb_to_cmyk(cc.r, cc.g, cc.b, &c, &m, &y, &k)
	return float32(c), float32(m), float32(y), float32(k)
}

func (col Color) GetHSL() (float32, float32, float32) {
	var h, s, l C.float
	cc := C.ALLEGRO_COLOR(col)
	C.al_color_rgb_to_hsl(cc.r, cc.g, cc.b, &h, &s, &l)
	return float32(h), float32(s), float32(l)
}

func (col Color) GetHSV() (float32, float32, float32) {
	var h, s, l C.float
	cc := C.ALLEGRO_COLOR(col)
	C.al_color_rgb_to_hsv(cc.r, cc.g, cc.b, &h, &s, &l)
	return float32(h), float32(s), float32(l)
}

func (col Color) GetYUV() (float32, float32, float32) {
	var h, s, l C.float
	cc := C.ALLEGRO_COLOR(col)
	C.al_color_rgb_to_yuv(cc.r, cc.g, cc.b, &h, &s, &l)
	return float32(h), float32(s), float32(l)
}