package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type KeyboardState C.ALLEGRO_KEYBOARD_STATE

func InstallKeyboard() {

	C.al_install_keyboard()

}

func IsKeyboardInstalled() bool {
	var b bool

	b = bool(C.al_is_keyboard_installed())

	return b
}

func UninstallKeyboard() {

	C.al_uninstall_keyboard()

}

func GetKeyboadState() *KeyboardState {
	var state C.ALLEGRO_KEYBOARD_STATE

	C.al_get_keyboard_state(&state)

	return (*KeyboardState)(&state)
}

func (k *KeyboardState) KeyDown(keycode int) bool {
	var b bool

	b = bool(C.al_key_down((*C.ALLEGRO_KEYBOARD_STATE)(k), C.int(keycode)))

	return b
}

func KeycodeToName(keycode int) string {
	var cs *C.char

	cs = C.al_keycode_to_name(C.int(keycode))

	return C.GoString(cs)
}

func SetKeyboardLEDs(leds int) {

	C.al_set_keyboard_leds(C.int(leds))

}

func GetKeyboardEventSource() *EventSource {
	var es *C.ALLEGRO_EVENT_SOURCE

	es = C.al_get_keyboard_event_source()

	return createEventSource(es)
}
