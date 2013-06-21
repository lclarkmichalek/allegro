package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Timer C.ALLEGRO_TIMER

func CreateTimer(speed_secs float64) *Timer {
	var t *Timer

	t = (*Timer)(C.al_create_timer(C.double(speed_secs)))

	return t
}

func (t *Timer) Start() {

	C.al_start_timer((*C.ALLEGRO_TIMER)(t))

}

func (t *Timer) Stop() {

	C.al_stop_timer((*C.ALLEGRO_TIMER)(t))

}

func (t *Timer) IsTimerStarted() bool {
	var b bool

	b = bool(C.al_get_timer_started((*C.ALLEGRO_TIMER)(t)))

	return b
}

func (t *Timer) Destroy() {

	C.al_destroy_timer((*C.ALLEGRO_TIMER)(t))

}

func (t *Timer) GetCount() int64 {
	var i int64

	i = int64(C.al_get_timer_count((*C.ALLEGRO_TIMER)(t)))

	return i
}

func (t *Timer) SetCount(count int64) {

	C.al_set_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(count))

}

func (t *Timer) AddCount(delta int64) {

	C.al_add_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(delta))

}

func (t *Timer) GetSpeed() float64 {
	var f float64

	f = float64(C.al_get_timer_speed((*C.ALLEGRO_TIMER)(t)))

	return f
}

func (t *Timer) SetSpeed(speed float64) {

	C.al_set_timer_speed((*C.ALLEGRO_TIMER)(t), C.double(speed))

}

func (t *Timer) GetEventSource() *EventSource {
	var es *C.ALLEGRO_EVENT_SOURCE

	es = C.al_get_timer_event_source((*C.ALLEGRO_TIMER)(t))

	return createEventSource(es)
}
