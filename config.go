package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import "unsafe"

type Config C.ALLEGRO_CONFIG

func CreateConfig() *Config {
	return (*Config)(C.al_create_config())
}

func LoadConfig(filename string) *Config {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	return (*Config)(C.al_load_config_file(cs))
}

func (c *Config) Destroy() {
	C.al_destroy_config((*C.ALLEGRO_CONFIG)(c))
}

func (c *Config) AddSection(section string) {
	cs := C.CString(section)
	defer C.free(unsafe.Pointer(cs))

	C.al_add_config_section((*C.ALLEGRO_CONFIG)(c), cs)
}

func (c *Config) AddComment(section string, comment string) {
	ss := C.CString(section)
	defer C.free(unsafe.Pointer(ss))
	cs := C.CString(comment)
	defer C.free(unsafe.Pointer(cs))

	C.al_add_config_comment((*C.ALLEGRO_CONFIG)(c), ss, cs)
}

func (c *Config) Get(section string, key string) (string, bool) {
	ss := C.CString(section)
	defer C.free(unsafe.Pointer(ss))
	ks := C.CString(key)
	defer C.free(unsafe.Pointer(ks))

	cstr := C.al_get_config_value((*C.ALLEGRO_CONFIG)(c), ss, ks)
	if (cstr == nil) {
		return "", false
	}
	return C.GoString(cstr), true
}

func (c *Config) Set(section string, key string, val string) {
	ss := C.CString(section)
	defer C.free(unsafe.Pointer(ss))
	ks := C.CString(key)
	defer C.free(unsafe.Pointer(ks))
	vs := C.CString(val)
	defer C.free(unsafe.Pointer(vs))

	C.al_set_config_value((*C.ALLEGRO_CONFIG)(c), ss, ks, vs)
}

func (c *Config) IterSections() chan string {
	channel := make(chan string)
	go func() {
		var iter **C.ALLEGRO_CONFIG_SECTION

		name_ptr := C.al_get_first_config_section((*C.ALLEGRO_CONFIG)(c), iter)
		for name_ptr != nil {
			name := C.GoString(name_ptr)
			channel <- name
			name_ptr = C.al_get_next_config_section(iter)
		}
		close(channel)
	}()
	return channel
}

func (c *Config) IterKeys(sec string) chan string {
	channel := make(chan string)
	go func() {
		var iter **C.ALLEGRO_CONFIG_ENTRY
		ss := C.CString(sec)
		defer C.free(unsafe.Pointer(ss))

		key_ptr := C.al_get_first_config_entry((*C.ALLEGRO_CONFIG)(c), ss, iter)
		for key_ptr != nil {
			key := C.GoString(key_ptr)
			channel <- key
			key_ptr = C.al_get_next_config_entry(iter)
		}
		close(channel)
	}()
	return channel
}

func (m *Config) Merge(child *Config) {
	m_ptr := (*C.ALLEGRO_CONFIG)(m)
	c_ptr := (*C.ALLEGRO_CONFIG)(child)

	C.al_merge_config_into(m_ptr, c_ptr)
}