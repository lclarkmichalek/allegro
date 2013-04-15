package allegro

/*
#cgo pkg-config: allegro_image-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_image.h>
*/
import "C"

func InitImage() {
	RunInThread(func() {
		C.al_init_image_addon()
	})
}

func ShutdownImage() {
	RunInThread(func() {
		C.al_shutdown_image_addon()
	})
}
