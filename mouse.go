package allegro 

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

const (
	CURSOR_DEFAULT = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_DEFAULT
	CURSOR_ARROW = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ARROW
	CURSOR_BUSY = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_BUSY
	CURSOR_QUESTION = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_QUESTION
	CURSOR_EDIT = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_EDIT
	CURSOR_MOVE = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_MOVE
	CURSOR_RESIZE_N = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_N
	CURSOR_RESIZE_W = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_W
	CURSOR_RESIZE_S = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_S
	CURSOR_RESIZE_E = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_E
	CURSOR_RESIZE_NW = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NW
	CURSOR_RESIZE_SW = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SW
	CURSOR_RESIZE_SE = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SE
	CURSOR_RESIZE_NE = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NE
	CURSOR_PROGRESS = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PROGRESS
	CURSOR_PRECISION = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PRECISION
	CURSOR_LINK = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_LINK
	CURSOR_ALT_SELECT = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ALT_SELECT
	CURSOR_UNAVAILABLE = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_UNAVAILABLE
)

type MouseState C.ALLEGRO_MOUSE_STATE

func GetMouseState() MouseState {
	var ms MouseState
	C.al_get_mouse_state((*C.ALLEGRO_MOUSE_STATE)(&ms))
	return ms
}

func (m MouseState) GetPos() (int, int, int, int) {
	ms := (C.ALLEGRO_MOUSE_STATE)(m)
	return int(ms.x), int(ms.y), int(ms.w), int(ms.z)
}

func InstallMouse() {
	C.al_install_mouse()
}

func IsMouseInstalled() bool {
	return bool(C.al_is_mouse_installed())
}

func UninstallMouse() {
	C.al_uninstall_mouse()
}

func (m MouseState) GetAxes() []int {
	n := int(C.al_get_mouse_num_axes())
	ms := (C.ALLEGRO_MOUSE_STATE)(m)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		axis := int(C.al_get_mouse_state_axis(&ms, C.int(i)))
		slice[i] = axis
	}
	return slice
}

func (m MouseState) GetButtons() []bool {
	n := int(C.al_get_mouse_num_buttons())
	ms := (C.ALLEGRO_MOUSE_STATE)(m)
	slice := make([]bool, n)
	for i := uint(0); i < uint(n); i++ {
		button := (ms.buttons & (1 << i)) != 0
		slice[i] = button
	}
	return slice
}

func (d *Display) SetMouse(x, y int) bool {
	return bool(C.al_set_mouse_xy((*C.ALLEGRO_DISPLAY)(d), C.int(x), C.int(y)))
}

func SetMouseWZ(w, z int) bool {
	return bool(C.al_set_mouse_w(C.int(w))) && bool(C.al_set_mouse_z(C.int(z)))
}

func SetMouseAxis(which, value int) bool {
	return bool(C.al_set_mouse_axis(C.int(which), C.int(value)))
}

func GetMouseEventSource() *EventSource {
	return (*EventSource)(C.al_get_mouse_event_source())
}

type MouseCursor C.ALLEGRO_MOUSE_CURSOR

func CreateMouseCursor(bmp *Bitmap, x, y int) *MouseCursor {
	return (*MouseCursor)(C.al_create_mouse_cursor(
		(*C.ALLEGRO_MOUSE_CURSOR)(bmp), C.int(x), C.int(y)))
}

func (mc *MouseCursor) Destroy() {
	C.al_destroy_mouse_cursor((*C.ALLEGRO_MOUSE_CURSOR)(mc))
}

func (d *Display) SetMouseCursor(mc *MouseCursor) {
	C.al_set_mouse_cursor((*C.ALLEGRO_DISPLAY)(d), (*C.ALLEGRO_MOUSE_CURSOR)(mc))
}

func (d *Display) SetSystemCursor(cursor C.ALLEGRO_SYSTEM_MOUSE_CURSOR) bool {
	return bool(C.al_set_system_mouse_cursor((*C.ALLEGRO_DISPLAY)(d), cursor))
}

func (d *Display) ShowCursor(show bool) {
	if show {
		C.al_show_mouse_cursor((*C.ALLEGRO_DISPLAY)(d))
	} else {
		C.al_hide_mouse_cursor((*C.ALLEGRO_DISPLAY)(d))
	}
}

func (d *Display) GrabCursor(grab bool) {
	if grab {
		C.al_grab_mouse((*C.ALLEGRO_DISPLAY)(d))
	} else {
		C.al_ungrab_mouse()
	}
}