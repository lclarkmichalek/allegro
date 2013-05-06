package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
#include "./event_getter.h"
*/
import "C"

type EventSource C.ALLEGRO_EVENT_SOURCE

func GetEvents(sources []*EventSource) chan interface{} {
	ch := make(chan interface{}, 10)

	go func() {
		var queue *C.ALLEGRO_EVENT_QUEUE
		queue = C.al_create_event_queue()
		defer C.al_destroy_event_queue(queue)

		for _, src := range sources {
			ptr := (*C.ALLEGRO_EVENT_SOURCE)(src)
			C.al_register_event_source(queue, ptr)
		}

		for {
			var al_event C.ALLEGRO_EVENT
			C.al_wait_for_event(queue, &al_event)

			ev := toEv(al_event)
			if ev != nil {
				ch <- ev
			}
		}
	}()
	return ch
}

func toEv(ev C.ALLEGRO_EVENT) interface{} {
	any := C.get_any(ev)
	src := (*EventSource)(any.source)
	ts := float64(any.timestamp)
	joystick := C.get_joystick(ev)
	mouse := C.get_mouse(ev)
	display := C.get_display(ev)
	keyboard := C.get_keyboard(ev)
	timer := C.get_timer(ev)
	switch C.get_type(ev) {
	case C.ALLEGRO_EVENT_JOYSTICK_AXIS:
		return JoystickAxisEvent{
			src, ts,
			(*Joystick)(joystick.id),
			int(joystick.stick),
			int(joystick.axis),
			float32(joystick.pos)}
	case C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN:
		return JoystickButtonDownEvent{
			src, ts,
			(*Joystick)(joystick.id),
			int(joystick.button)}
	case C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP:
		return JoystickButtonUpEvent{
			src, ts,
			(*Joystick)(joystick.id),
			int(joystick.button)}
	case C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION:
		return JoystickConfigurationEvent{src, ts}
	case C.ALLEGRO_EVENT_KEY_DOWN:
		return KeyDownEvent{
			src, ts,
			int(keyboard.keycode),
			(*Display)(keyboard.display)}
	case C.ALLEGRO_EVENT_KEY_UP:
		return KeyUpEvent{
			src, ts,
			int(keyboard.keycode),
			(*Display)(keyboard.display)}
	case C.ALLEGRO_EVENT_KEY_CHAR:
		return KeyCharEvent{
			src, ts,
			int(keyboard.keycode),
			int(keyboard.unichar),
			int(keyboard.modifiers),
			bool(keyboard.repeat),
			(*Display)(keyboard.display)}
	case C.ALLEGRO_EVENT_MOUSE_AXES:
		return MouseAxesEvent{
			src, ts,
			int(mouse.x),
			int(mouse.y),
			int(mouse.z),
			int(mouse.w),
			int(mouse.dx),
			int(mouse.dy),
			int(mouse.dz),
			int(mouse.dw),
			(*Display)(mouse.display)}
	case C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN:
		return MouseButtonDown{
			src, ts,
			int(mouse.x),
			int(mouse.y),
			int(mouse.z),
			int(mouse.w),
			int(mouse.button),
			(*Display)(mouse.display)}
	case C.ALLEGRO_EVENT_MOUSE_BUTTON_UP:
		return MouseButtonUp{
			src, ts,
			int(mouse.x),
			int(mouse.y),
			int(mouse.z),
			int(mouse.w),
			int(mouse.button),
			(*Display)(mouse.display)}
	case C.ALLEGRO_EVENT_MOUSE_WARPED:
		return MouseWarpEvent{
			src, ts,
			int(mouse.x),
			int(mouse.y),
			int(mouse.z),
			int(mouse.w),
			int(mouse.dx),
			int(mouse.dy),
			int(mouse.dz),
			int(mouse.dw),
			(*Display)(mouse.display)}
	case C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY:
		return MouseEnterDisplayEvent{
			src, ts,
			int(mouse.x), int(mouse.y),
			int(mouse.z), int(mouse.w),
			(*Display)(mouse.display)}
	case C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY:
		return MouseLeaveDisplayEvent{
			src, ts,
			int(mouse.x), int(mouse.y),
			int(mouse.z), int(mouse.w),
			(*Display)(mouse.display)}
	case C.ALLEGRO_EVENT_TIMER:
		return TimerEvent{
			src, ts,
			(*Timer)(timer.source),
			int64(timer.count)}
	case C.ALLEGRO_EVENT_DISPLAY_EXPOSE:
		return DisplayExposeEvent{
			src, ts,
			(*Display)(display.source),
			int(display.x), int(display.y),
			int(display.width), int(display.height)}
	case C.ALLEGRO_EVENT_DISPLAY_CLOSE:
		return DisplayCloseEvent{
			src, ts,
			(*Display)(display.source)}
	case C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT:
		return DisplaySwitchOutEvent{
			src, ts,
			(*Display)(display.source)}
	case C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN:
		return DisplaySwitchInEvent{
			src, ts,
			(*Display)(display.source)}
	case C.ALLEGRO_EVENT_DISPLAY_ORIENTATION:
		return DisplayOrientationEvent{
			src, ts,
			(*Display)(display.source),
			int(display.orientation)}
	}
	return nil
}
