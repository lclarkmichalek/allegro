package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Joystick C.ALLEGRO_JOYSTICK

type JoystickState struct {
	Axes [][]float32
	Buttons []int
}

func InstallJoystick() {
	C.al_install_joystick()
}

func UninstallJoystick() {
	C.al_uninstall_joystick()
}

func IsJoystickInstalled() bool {
	return bool(C.al_is_joystick_installed())
}

func ReconfigureJoysticks() bool {
	return bool(C.al_reconfigure_joysticks())
}

func GetJoysticks() []*Joystick {
	n := int(C.al_get_num_joysticks())
	sticks := make([]*Joystick, n)
	for i := 0; i < n; i++ {
		stick := (*Joystick)(C.al_get_joystick(C.int(i)))
		sticks[i] = stick
	}
	return sticks
}

func (j *Joystick) IsActive() bool {
	return bool(C.al_get_joystick_active((*C.ALLEGRO_JOYSTICK)(j)))
}

func (j *Joystick) GetName() string {
	cs := C.al_get_joystick_name((*C.ALLEGRO_JOYSTICK)(j))
	return C.GoString(cs)
}