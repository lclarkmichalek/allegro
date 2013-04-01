package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Timer C.ALLEGRO_TIMER

func CreateTimer(speed_secs float64) *Timer {
	return (*Timer)(C.al_create_timer(C.double(speed_secs)))
}

func (t *Timer) Start() {
	C.al_start_timer((*C.ALLEGRO_TIMER)(t))
}

func (t *Timer) Stop() {
	C.al_stop_timer((*C.ALLEGRO_TIMER)(t))
}

func (t *Timer) IsTimerStarted() bool {
	return bool(C.al_get_timer_started((*C.ALLEGRO_TIMER)(t)))
}

func (t *Timer) Destroy() {
	C.al_destroy_timer((*C.ALLEGRO_TIMER)(t))
}

func (t *Timer) GetCount() int64 {
	return int64(C.al_get_timer_count((*C.ALLEGRO_TIMER)(t)))
}

func (t *Timer) SetCount(count int64) {
	C.al_set_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(count))
}

func (t *Timer) AddCount(delta int64) {
	C.al_add_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(delta))
}

func (t *Timer) GetSpeed() float64 {
	return float64(C.al_get_timer_speed((*C.ALLEGRO_TIMER)(t)))
}

func (t *Timer) SetSpeed(speed float64) {
	C.al_set_timer_speed((*C.ALLEGRO_TIMER)(t), C.double(speed))
}

func (t *Timer) GetEventSource() *EventSource {
	return (*EventSource)(C.al_get_timer_event_source((*C.ALLEGRO_TIMER)(t)))
}