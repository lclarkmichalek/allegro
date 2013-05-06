package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

const (
	ORIENTATION_0_DEGREES   = C.ALLEGRO_DISPLAY_ORIENTATION_0_DEGREES
	ORIENTATION_90_DEGREES  = C.ALLEGRO_DISPLAY_ORIENTATION_90_DEGREES
	ORIENTATION_180_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_180_DEGREES
	ORIENTATION_270_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_270_DEGREES
	ORIENTATION_FACE_UP     = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_UP
	ORIENTATION_FACE_DOWN   = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_DOWN
)

type JoystickAxisEvent struct {
	Source    *EventSource
	Timestamp float64
	Joystick  *Joystick
	JStick    int
	JAxis     int
	JPos      float32
}

type JoystickButtonDownEvent struct {
	Source    *EventSource
	Timestamp float64
	Joystick  *Joystick
	JButton   int
}

type JoystickButtonUpEvent struct {
	Source    *EventSource
	Timestamp float64
	Joystick  *Joystick
	JButton   int
}

type JoystickConfigurationEvent struct {
	Source    *EventSource
	Timestamp float64
}

type KeyDownEvent struct {
	Source    *EventSource
	Timestamp float64
	Keycode   int
	Display   *Display
}

type KeyUpEvent struct {
	Source    *EventSource
	Timestamp float64
	Keycode   int
	Display   *Display
}

type KeyCharEvent struct {
	Source    *EventSource
	Timestamp float64
	Keycode   int
	Unichar   int
	Modifiers int
	Repeat    bool
	Display   *Display
}

type MouseAxesEvent struct {
	Source         *EventSource
	Timestamp      float64
	X, Y, Z, W     int
	DX, DY, DZ, DW int
	Display        *Display
}

type MouseButtonDown struct {
	Source     *EventSource
	Timestamp  float64
	X, Y, Z, W int
	Button     int
	Display    *Display
}

type MouseButtonUp struct {
	Source     *EventSource
	Timestamp  float64
	X, Y, Z, W int
	Button     int
	Display    *Display
}

type MouseWarpEvent struct {
	Source         *EventSource
	Timestamp      float64
	X, Y, Z, W     int
	DX, DY, DZ, DW int
	Display        *Display
}

type MouseEnterDisplayEvent struct {
	Source     *EventSource
	Timestamp  float64
	X, Y, Z, W int
	Display    *Display
}

type MouseLeaveDisplayEvent struct {
	Source     *EventSource
	Timestamp  float64
	X, Y, Z, W int
	Display    *Display
}

type TimerEvent struct {
	Source    *EventSource
	Timestamp float64
	Timer     *Timer
	Count     int64
}

type DisplayExposeEvent struct {
	Source     *EventSource
	Timestamp  float64
	Display    *Display
	X, Y, W, H int
}

type DisplayResizeEvent struct {
	Source     *EventSource
	Timestamp  float64
	Display    *Display
	X, Y, W, H int
}

type DisplayCloseEvent struct {
	Source    *EventSource
	Timestamp float64
	Display   *Display
}

type DisplaySwitchInEvent struct {
	Source    *EventSource
	Timestamp float64
	Display   *Display
}

type DisplaySwitchOutEvent struct {
	Source    *EventSource
	Timestamp float64
	Display   *Display
}

type DisplayOrientationEvent struct {
	Source      *EventSource
	Timestamp   float64
	Display     *Display
	Orientation int
}

/*
Dont know if this can be used to support interface{} types
type UserEvent struct {
	Source *EventSource
	Timestamp float64

}
*/
