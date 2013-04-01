package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Timer C.ALLEGRO_TIMER

func CreateTimer(speed_secs float64) *Timer {
	var t *Timer
	RunInThread(func() {
		t = (*Timer)(C.al_create_timer(C.double(speed_secs)))
	})
	return t
}

func (t *Timer) Start() {
	RunInThread(func() {
		C.al_start_timer((*C.ALLEGRO_TIMER)(t))
	})
}

func (t *Timer) Stop() {
	RunInThread(func() {
		C.al_stop_timer((*C.ALLEGRO_TIMER)(t))
	})
}

func (t *Timer) IsTimerStarted() bool {
	var b bool
	RunInThread(func() {
		b = bool(C.al_get_timer_started((*C.ALLEGRO_TIMER)(t)))
	})
	return b
}

func (t *Timer) Destroy() {
	RunInThread(func() {
		C.al_destroy_timer((*C.ALLEGRO_TIMER)(t))
	})
}

func (t *Timer) GetCount() int64 {
	var i int64
	RunInThread(func() {
		i = int64(C.al_get_timer_count((*C.ALLEGRO_TIMER)(t)))
	})
	return i
}

func (t *Timer) SetCount(count int64) {
	RunInThread(func() {
		C.al_set_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(count))
	})
}

func (t *Timer) AddCount(delta int64) {
	RunInThread(func() {
		C.al_add_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(delta))
	})
}

func (t *Timer) GetSpeed() float64 {
	var f float64
	RunInThread(func() {
		f = float64(C.al_get_timer_speed((*C.ALLEGRO_TIMER)(t)))
	})
	return f
}

func (t *Timer) SetSpeed(speed float64) {
	RunInThread(func() {
		C.al_set_timer_speed((*C.ALLEGRO_TIMER)(t), C.double(speed))
	})
}

func (t *Timer) GetEventSource() *EventSource {
	var es *EventSource
	RunInThread(func() {
		es = (*EventSource)(C.al_get_timer_event_source((*C.ALLEGRO_TIMER)(t)))
	})
	return es
}