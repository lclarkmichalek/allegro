package color

/*
#cgo pkg-config: allegro_color-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_color.h>
*/
import "C"
import "unsafe"

func CreateColorCMYK(c, m, y, k float32) Color {
	return Color(C.al_color_cmyk(C.float(c), C.float(m), C.float(y), C.float(k)))
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

func GetVersion() int32 {
	return int32(C.al_get_allegro_color_version())
}