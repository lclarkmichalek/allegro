package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import "unsafe"

const (
	// Window flags
	WINDOWED                  = C.ALLEGRO_WINDOWED
	FULLSCREEN                = C.ALLEGRO_FULLSCREEN
	FULLSCREEN_WINDOW         = C.ALLEGRO_FULLSCREEN_WINDOW
	RESIZABLE                 = C.ALLEGRO_RESIZABLE
	OPENGL                    = C.ALLEGRO_OPENGL
	OPENGL_3_0                = C.ALLEGRO_OPENGL_3_0
	OPENGL_FORWARD_COMPATIBLE = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
	FRAMELESS                 = C.ALLEGRO_FRAMELESS
	GENERATE_EXPOSE_EVENTS    = C.ALLEGRO_GENERATE_EXPOSE_EVENTS

	// Priorities
	REQUIRE  = C.ALLEGRO_REQUIRE
	SUGGEST  = C.ALLEGRO_SUGGEST
	DONTCARE = C.ALLEGRO_DONTCARE

	// Configurable stuff
	COLOR_SIZE             = C.ALLEGRO_COLOR_SIZE
	RED_SIZE               = C.ALLEGRO_RED_SIZE
	GREEN_SIZE             = C.ALLEGRO_GREEN_SIZE
	BLUE_SIZE              = C.ALLEGRO_BLUE_SIZE
	ALPHA_SIZE             = C.ALLEGRO_ALPHA_SIZE
	RED_SHIFT              = C.ALLEGRO_RED_SHIFT
	GREEN_SHIFT            = C.ALLEGRO_GREEN_SHIFT
	BLUE_SHIFT             = C.ALLEGRO_BLUE_SHIFT
	ALPHA_SHIFT            = C.ALLEGRO_ALPHA_SHIFT
	ACC_RED_SIZE           = C.ALLEGRO_ACC_RED_SIZE
	ACC_GREEN_SIZE         = C.ALLEGRO_ACC_GREEN_SIZE
	ACC_BLUE_SIZE          = C.ALLEGRO_ACC_BLUE_SIZE
	ACC_ALPHA_SIZE         = C.ALLEGRO_ACC_ALPHA_SIZE
	STEREO                 = C.ALLEGRO_STEREO
	AUX_BUFFERS            = C.ALLEGRO_AUX_BUFFERS
	DEPTH_SIZE             = C.ALLEGRO_DEPTH_SIZE
	STENCIL_SIZE           = C.ALLEGRO_STENCIL_SIZE
	SAMPLE_BUFFERS         = C.ALLEGRO_SAMPLE_BUFFERS
	SAMPLES                = C.ALLEGRO_SAMPLES
	RENDER_METHOD          = C.ALLEGRO_RENDER_METHOD
	FLOAT_COLOR            = C.ALLEGRO_FLOAT_COLOR
	FLOAT_DEPTH            = C.ALLEGRO_FLOAT_DEPTH
	SINGLE_BUFFER          = C.ALLEGRO_SINGLE_BUFFER
	SWAP_METHOD            = C.ALLEGRO_SWAP_METHOD
	COMPATIBLE_DISPLAY     = C.ALLEGRO_COMPATIBLE_DISPLAY
	UPDATE_DISPLAY_REGION  = C.ALLEGRO_UPDATE_DISPLAY_REGION
	VSYNC                  = C.ALLEGRO_VSYNC
	MAX_BITMAP_SIZE        = C.ALLEGRO_MAX_BITMAP_SIZE
	SUPPORT_NPOT_BITMAP    = C.ALLEGRO_SUPPORT_NPOT_BITMAP
	CAN_DRAW_INTO_BITMAP   = C.ALLEGRO_CAN_DRAW_INTO_BITMAP
	SUPPORT_SEPARATE_ALPHA = C.ALLEGRO_SUPPORT_SEPARATE_ALPHA
)

type Display C.ALLEGRO_DISPLAY

func CreateDisplay(width, height, flags int) *Display {
	var d *Display
	RunInThread(func() {
		C.al_set_new_display_flags((C.int)(flags))
		ptr := C.al_create_display((C.int)(width), (C.int)(height))
		d = (*Display)(ptr)
	})
	return d
}

func (d *Display) Destroy() {
	RunInThread(func() {
		C.al_destroy_display((*C.ALLEGRO_DISPLAY)(d))
	})
}

func SetNewDisplayOption(option, value, importance int) {
	RunInThread(func() {
		C.al_set_new_display_option((C.int)(option), (C.int)(value), (C.int)(importance))
	})
}

func ResetNewDisplayOptions() {
	RunInThread(func() {
		C.al_reset_new_display_options()
	})
}

func GetNewWindowPosition() (int, int) {
	var w, h C.int
	RunInThread(func() {
		C.al_get_new_window_position(&w, &h)
	})
	return (int)(w), (int)(h)
}

func SetNewWindowPosition(w, h int) {
	RunInThread(func() {
		C.al_set_new_window_position((C.int)(w), (C.int)(h))
	})
}

func GetNewDisplayRefreshRate() int {
	var v int
	RunInThread(func() {
		v = (int)(C.al_get_new_display_refresh_rate())
	})
	return v
}

func SetNewDisplayRefreshRate(ref int) {
	RunInThread(func() {
		C.al_set_new_display_refresh_rate((C.int)(ref))
	})
}

func Flip() {
	RunInThread(func() {
		C.al_flip_display()
	})
}

func FlipRegion(x, y, w, h int) {
	RunInThread(func() {
		C.al_update_display_region(C.int(x), C.int(y), C.int(w), C.int(h))
	})
}

func WaitVSync() bool {
	var v bool
	RunInThread(func() {
		v = bool(C.al_wait_for_vsync())
	})
	return v
}

func (d *Display) GetEventSource() *EventSource {
	var es *C.ALLEGRO_EVENT_SOURCE
	RunInThread(func() {
		src_ptr := C.al_get_display_event_source((*C.ALLEGRO_DISPLAY)(d))
		es = src_ptr
	})
	return createEventSource(es)
}

func (d *Display) GetBackbuffer() *Bitmap {
	var b *Bitmap
	RunInThread(func() {
		b = (*Bitmap)(C.al_get_backbuffer((*C.ALLEGRO_DISPLAY)(d)))
	})
	return b
}

func (d *Display) GetDimensions() (int, int) {
	ptr := (*C.ALLEGRO_DISPLAY)(d)
	var w, h int
	RunInThread(func() {
		w = int(C.al_get_display_width(ptr))
		h = int(C.al_get_display_height(ptr))
	})
	return w, h
}

func (d *Display) Resize(w, h int) bool {
	var b bool
	RunInThread(func() {
		b = bool(C.al_resize_display((*C.ALLEGRO_DISPLAY)(d), C.int(w), C.int(h)))
	})
	return b
}

func (d *Display) AcknowledgeResize() bool {
	var b bool
	RunInThread(func() {
		b = bool(C.al_acknowledge_resize((*C.ALLEGRO_DISPLAY)(d)))
	})
	return b
}

func (d *Display) GetPosition() (int, int) {
	var x, y C.int
	RunInThread(func() {
		C.al_get_window_position((*C.ALLEGRO_DISPLAY)(d), &x, &y)
	})
	return int(x), int(y)
}

func (d *Display) SetPosition(x, y int) {
	C.al_set_window_position((*C.ALLEGRO_DISPLAY)(d), C.int(x), C.int(y))
}

func (d *Display) GetFlags() int {
	return int(C.al_get_display_flags((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) SetFlag(flag int, on bool) bool {
	return bool(C.al_set_display_flag((*C.ALLEGRO_DISPLAY)(d), C.int(flag), C.bool(on)))
}

func (d *Display) GetOption(option int) int {
	return int(C.al_get_display_option((*C.ALLEGRO_DISPLAY)(d), C.int(option)))
}

func (d *Display) GetFormat() int {
	return int(C.al_get_display_format((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) GetRefreshRate() int {
	return int(C.al_get_display_refresh_rate((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) SetTitle(title string) {
	ss := C.CString(title)
	defer C.free(unsafe.Pointer(ss))
	C.al_set_window_title((*C.ALLEGRO_DISPLAY)(d), ss)
}

func (d *Display) SetIcon(bmp *Bitmap) {
	C.al_set_display_icon((*C.ALLEGRO_DISPLAY)(d), (*C.ALLEGRO_BITMAP)(bmp))
}

func InhibitScreensaver(inhibit bool) {
	C.al_inhibit_screensaver(C.bool(inhibit))
}
