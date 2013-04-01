Bindings for the Allegro library for Go
---------------------------------------

Documentation? Ha! Read the allegro docs. Almost everything mirrors the allegro
C library (major exceptions are anything to do with iteration). Things have been
"object orientified" somewhat: `al_destroy_display` -> `display.Destroy` and the
like.

If you want to do your own stuff using the C api, you can convert most of the
types simply by casting the types (i.e. (*C.ALLEGRO_DISPLAY)(disp) and vice
versa). There are a couple of exceptions (well, just LockedRegion), but you can
use the `.GetRaw()` method to get the raw pointer/data.

Allegro requires everything to be called from one hardware thread; to run a
function in that thread, use `RunInThread`.

Things which aren't supported and should be:

 - `allegro_image`
 - `allegro_primitives`
 - `allegro_color`
 - `allegro_font`
 - `allegro_ttf`
 - `allegro_audio`
 - `allegro_acodec`
 - OpenGl compat/integration

Things which aren't supported and won't be:

 - `allegro_memfs`
 - `allegro_physfs`
 - `allegro_native_dialogue`
 - file io
 - fixed point math
 - memory management
 - file paths
 - "state" (whatever that is wrt allegro)
 - threading
 - UTF-8 (unless it becomes needed)
 - iphone methods
 - Direct3D compat/integration

Things which aren't supported and I don't know how:

 - Lack of memory leaks? AHAHAHA
 - Support for writing bitmap loaders/savers in Go
 - Emitting/catching user events
