package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Joystick C.ALLEGRO_JOYSTICK

type JoystickState struct {
	Axes    [][]float32
	Buttons []int
}

func InstallJoystick() {
	RunInThread(func() {
		C.al_install_joystick()
	})
}

func UninstallJoystick() {
	RunInThread(func() {
		C.al_uninstall_joystick()
	})
}

func IsJoystickInstalled() bool {
	var b bool
	RunInThread(func() {
		b = bool(C.al_is_joystick_installed())
	})
	return b
}

func ReconfigureJoysticks() bool {
	var b bool
	RunInThread(func() {
		b = bool(C.al_reconfigure_joysticks())
	})
	return b
}

func GetJoysticks() []*Joystick {
	var n int
	RunInThread(func() {
		n = int(C.al_get_num_joysticks())
	})
	sticks := make([]*Joystick, n)
	for i := 0; i < n; i++ {
		var stick *Joystick
		RunInThread(func() {
			stick = (*Joystick)(C.al_get_joystick(C.int(i)))
		})
		sticks[i] = stick
	}
	return sticks
}

func (j *Joystick) IsActive() bool {
	var b bool
	RunInThread(func() {
		b = bool(C.al_get_joystick_active((*C.ALLEGRO_JOYSTICK)(j)))
	})
	return b
}

func (j *Joystick) GetName() string {
	var cs *C.char
	RunInThread(func() {
		cs = C.al_get_joystick_name((*C.ALLEGRO_JOYSTICK)(j))
	})
	return C.GoString(cs)
}
