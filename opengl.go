package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_opengl.h>
*/
import "C"

// We don't return a gl.GLuint as I don't want to add the full gl package
// as a dependency. If you need this, cast it. Otherwise, ignore
func (b *Bitmap) GetGLTexture() uint32 {
	var tex uint32
	RunInThread(func() {
		tex = uint32(C.al_get_opengl_texture((*C.ALLEGRO_BITMAP)(b)))
	})
	return tex
}

func (b *Bitmap) GetGLTextureSize() (int, int) {
	var w, h C.int
	RunInThread(func() {
		C.al_get_opengl_texture_size((*C.ALLEGRO_BITMAP)(b), &w, &h)
	})
	return int(w), int(h)
}

func (b *Bitmap) GetGLTexturePosition() (int, int) {
	var x, y C.int
	RunInThread(func() {
		C.al_get_opengl_texture_position((*C.ALLEGRO_BITMAP)(b), &x, &y)
	})
	return int(x), int(y)
}

func (b *Bitmap) GetGLFBO() uint32 {
	var fbo uint32
	RunInThread(func() {
		fbo = uint32(C.al_get_opengl_fbo((*C.ALLEGRO_BITMAP)(b)))
	})
	return fbo
}