#include "./event_getter.h"

int get_type(ALLEGRO_EVENT ev) {
    return ev.type;
}

ALLEGRO_ANY_EVENT get_any(ALLEGRO_EVENT ev) {
    return ev.any;
}
ALLEGRO_DISPLAY_EVENT get_display(ALLEGRO_EVENT ev) {
    return ev.display;
}
ALLEGRO_JOYSTICK_EVENT get_joystick(ALLEGRO_EVENT ev) {
    return ev.joystick;
}
ALLEGRO_KEYBOARD_EVENT get_keyboard(ALLEGRO_EVENT ev) {
    return ev.keyboard;
}
ALLEGRO_MOUSE_EVENT get_mouse(ALLEGRO_EVENT ev) {
    return ev.mouse;
}
ALLEGRO_TIMER_EVENT get_timer(ALLEGRO_EVENT ev) {
    return ev.timer;
}
