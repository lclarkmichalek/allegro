package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Timeout C.ALLEGRO_TIMEOUT

func GetTime() float64 {
	return float64(C.al_get_time())
}

func InitTimeout(secs float64) Timeout {
	var timeout C.ALLEGRO_TIMEOUT
	C.al_init_timeout(&timeout, C.double(secs))
	return Timeout(timeout)
}

func Rest(secs float64) {
	C.al_rest(C.double(secs))
}
