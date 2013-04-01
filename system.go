package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import "unsafe"
import "runtime"
import "sync"

const (
	RESOURCES_PATH = C.ALLEGRO_RESOURCES_PATH
	TEMP_PATH = C.ALLEGRO_TEMP_PATH
	USER_HOME_PATH = C.ALLEGRO_USER_HOME_PATH
	USER_DOCUMENTS_PATH = C.ALLEGRO_USER_DOCUMENTS_PATH
	USER_DATA_PATH = C.ALLEGRO_USER_DATA_PATH
	USER_SETTINGS_PATH = C.ALLEGRO_USER_SETTINGS_PATH
	EXENAME_PATH = C.ALLEGRO_EXENAME_PATH
)

var allegroThread = make(chan func())
var threadRunning *sync.Once = new(sync.Once)

func startThread() {
	// Don't want to lock the only thread
	if runtime.GOMAXPROCS(0) < 2 {
		runtime.GOMAXPROCS(2)
	}
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		for f := range allegroThread {
			f()
		}
		threadRunning = new(sync.Once)
	}()
}

func stopThread() {
	close(allegroThread)
}

// Runs the function in the allegro thread
func RunInThread(f func()) {
	threadRunning.Do(startThread)

	done := make(chan bool, 1)
	wrapped := func() {
		f()
		done <- true
	}
	allegroThread <- wrapped
	
	// Wait till we're done in other thread
	<-done
}

func Init() {
	InstallSystem(int(C.ALLEGRO_VERSION_INT))
}

// Don't support callback version as no idea how to do that
func InstallSystem(version int) {
	RunInThread(func() {
		C.al_install_system(C.int(version), nil)
	})
}

func UninstallSystem() {
	RunInThread(func() {
		C.al_uninstall_system()
	})
}

func GetVersion() (int, int, int, int) {
	var version C.uint32_t
	RunInThread(func() {
		version = C.al_get_allegro_version()
	})
	maj := int(version >> 24)
	min := int((version >> 16) & 255)
	rev := int((version >> 8) & 255)
	rel := int(version & 255)
	return maj, min, rev, rel
}

func GetStandardPath(path int) Path {
	var pth *C.ALLEGRO_PATH
	RunInThread(func() {
		pth = C.al_get_standard_path(C.int(path))
	})
	return toPth(pth)
}

func SetExeName(name string) {
	ns := C.CString(name)
	defer C.free(unsafe.Pointer(ns))
	
	RunInThread(func() {
		C.al_set_exe_name(ns)
	})
}

func SetAppName(name string) {
	ns := C.CString(name)
	defer C.free(unsafe.Pointer(ns))
	
	RunInThread(func() {
		C.al_set_app_name(ns)
	})
}

func SetOrgName(name string) {
	ns := C.CString(name)
	defer C.free(unsafe.Pointer(ns))

	RunInThread(func() {
		C.al_set_org_name(ns)
	})
}

func GetAppName() string {
	var str string
	RunInThread(func() {
		str = C.GoString(C.al_get_app_name())
	})
	return str
}

func GetOrgName() string {
	var str string
	RunInThread(func() {
		str = C.GoString(C.al_get_org_name())
	})
	return str
}

func GetSystemConfig() *Config {
	var cnf *Config
	RunInThread(func() {
		cnf = (*Config)(C.al_get_system_config())
	})
	return cnf
}