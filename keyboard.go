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
	return bool(C.al_is_keyboard_installed())
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
	return bool(C.al_key_down((*C.ALLEGRO_KEYBOARD_STATE)(k), C.int(keycode)))
}

func KeycodeToName(keycode int) string {
	cs := C.al_keycode_to_name(C.int(keycode))
	return C.GoString(cs)
}

func SetKeyboardLEDs(leds int) {
	C.al_set_keyboard_leds(C.int(leds))
}

func GetKeyboardEventSource() *EventSource {
	return (*EventSource)(C.al_get_keyboard_event_source())
}