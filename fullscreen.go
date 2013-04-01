package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
#include "./event_getter.h"
*/
import "C"

type DisplayMode struct {
	Width, Height int
	Format int
	RefreshRate int
}

func DisplayModes() []DisplayMode {
	var n int
	RunInThread(func() {
		n = int(C.al_get_num_display_modes())
	})
	modes := make([]DisplayMode, n)
	for i := 0; i < n; i++ {
		var mode C.ALLEGRO_DISPLAY_MODE
		RunInThread(func() {
			C.al_get_display_mode(C.int(i), &mode)
		})
		modes[i] = DisplayMode{
			int(mode.width), int(mode.height),
			int(mode.format), int(mode.refresh_rate)}
	}
	return modes
}