package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import "unsafe"
import "reflect"

const (
	// Bitmap flags
	VIDEO_BITMAP = C.ALLEGRO_VIDEO_BITMAP
	MEMORY_BITMAP = C.ALLEGRO_MEMORY_BITMAP
	KEEP_BITMAP_FORMAT = C.ALLEGRO_KEEP_BITMAP_FORMAT
	FORCE_LOCKING = C.ALLEGRO_FORCE_LOCKING
	NO_PRESERVE_TEXTURE = C.ALLEGRO_NO_PRESERVE_TEXTURE
	ALPHA_TEST = C.ALLEGRO_ALPHA_TEST
	MIN_LINEAR = C.ALLEGRO_MIN_LINEAR
	MAG_LINEAR = C.ALLEGRO_MAG_LINEAR
	MIPMAP = C.ALLEGRO_MIPMAP
	NO_PREMULTIPLIED_ALPHA = C.ALLEGRO_NO_PREMULTIPLIED_ALPHA

	// Flip Stuff
	FLIP_HORIZONTAL = C.ALLEGRO_FLIP_HORIZONTAL
	FLIP_VERTICAL = C.ALLEGRO_FLIP_VERTICAL

	// Blender Operations
	ADD = C.ALLEGRO_ADD
	DEST_MINUS_SRC = C.ALLEGRO_DEST_MINUS_SRC
	SRC_MINUS_DEST = C.ALLEGRO_SRC_MINUS_DEST
	ZERO = C.ALLEGRO_ZERO
	ONE = C.ALLEGRO_ONE
	ALPHA = C.ALLEGRO_ALPHA
	INVERSE_ALPHA = C.ALLEGRO_INVERSE_ALPHA

	// Pixel Formats
	FORMAT_ANY = C.ALLEGRO_PIXEL_FORMAT_ANY
	FORMAT_ANY_NO_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_NO_ALPHA
	FORMAT_ANY_WITH_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_WITH_ALPHA
	FORMAT_ANY_15_NO_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_15_NO_ALPHA
	FORMAT_ANY_16_NO_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_16_NO_ALPHA
	FORMAT_ANY_16_WITH_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_16_WITH_ALPHA
	FORMAT_ANY_24_NO_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_24_NO_ALPHA
	FORMAT_ANY_32_NO_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_32_NO_ALPHA
	FORMAT_ANY_32_WITH_ALPHA = C.ALLEGRO_PIXEL_FORMAT_ANY_32_WITH_ALPHA
	FORMAT_ARGB_8888 = C.ALLEGRO_PIXEL_FORMAT_ARGB_8888
	FORMAT_RGBA_8888 = C.ALLEGRO_PIXEL_FORMAT_RGBA_8888
	FORMAT_ARGB_4444 = C.ALLEGRO_PIXEL_FORMAT_ARGB_4444
	FORMAT_RGB_888 = C.ALLEGRO_PIXEL_FORMAT_RGB_888
	FORMAT_RGB_565 = C.ALLEGRO_PIXEL_FORMAT_RGB_565
	FORMAT_RGB_555 = C.ALLEGRO_PIXEL_FORMAT_RGB_555
	FORMAT_RGBA_5551 = C.ALLEGRO_PIXEL_FORMAT_RGBA_5551
	FORMAT_ARGB_1555 = C.ALLEGRO_PIXEL_FORMAT_ARGB_1555
	FORMAT_ABGR_8888 = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888
	FORMAT_XBGR_8888 = C.ALLEGRO_PIXEL_FORMAT_XBGR_8888
	FORMAT_BGR_888 = C.ALLEGRO_PIXEL_FORMAT_BGR_888
	FORMAT_BGR_565 = C.ALLEGRO_PIXEL_FORMAT_BGR_565
	FORMAT_BGR_555 = C.ALLEGRO_PIXEL_FORMAT_BGR_555
	FORMAT_RGBX_8888 = C.ALLEGRO_PIXEL_FORMAT_RGBX_8888
	FORMAT_XRGB_8888 = C.ALLEGRO_PIXEL_FORMAT_XRGB_8888
	FORMAT_ABGR_F32 = C.ALLEGRO_PIXEL_FORMAT_ABGR_F32
	FORMAT_ABGR_8888_LE = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888_LE
	FORMAT_RGBA_4444 = C.ALLEGRO_PIXEL_FORMAT_RGBA_4444

	// Bitmap lock flags
	LOCK_READONLY = C.ALLEGRO_LOCK_READONLY
	LOCK_WRITEONLY = C.ALLEGRO_LOCK_WRITEONLY
	LOCK_READWRITE = C.ALLEGRO_LOCK_READWRITE
)

type Color C.ALLEGRO_COLOR

func CreateColor(r, g, b, a byte) Color {
	return (Color)(C.al_map_rgba(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a)))
}

func (c Color) GetRGBA() (byte, byte, byte, byte) {
	var r, g, b, a C.uchar
	C.al_unmap_rgba((C.ALLEGRO_COLOR)(c), &r, &g, &b, &a)
	return byte(r), byte(g), byte(b), byte(a)
}

func (c Color) Clear() {
	C.al_clear_to_color((C.ALLEGRO_COLOR)(c))
}

func (c Color) DrawPixel(x, y float32) {
	C.al_draw_pixel(C.float(x), C.float(y), (C.ALLEGRO_COLOR)(c))
}

func (c Color) PutPixel(x, y int) {
	C.al_put_pixel(C.int(x), C.int(y), (C.ALLEGRO_COLOR)(c))
}

func (c Color) PutBlendedPixel(x, y int) {
	C.al_put_blended_pixel(C.int(x), C.int(y), (C.ALLEGRO_COLOR)(c))
}

type LockedRegion struct {
	data *C.ALLEGRO_LOCKED_REGION
	Width, Height int
}

func CreateLockedRegion(p *C.ALLEGRO_LOCKED_REGION, width, height int) LockedRegion {
	return LockedRegion{p, width, height}
}

func (l *LockedRegion) GetRaw() *C.ALLEGRO_LOCKED_REGION {
	return l.data
}

func (l *LockedRegion) GetData() []byte {
	ptr := l.GetRaw()
	length := l.Width * l.Height * int(ptr.pixel_size)
	var slice []byte
	// This make slice
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&slice)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(ptr.data))
	return slice
}

func (l *LockedRegion) GetFormat() int {
	return int(l.GetRaw().format)
}

func (l *LockedRegion) GetPitch() int {
	return int(l.GetRaw().pitch)
}

func (l *LockedRegion) GetPixelSize() int {
	return int(l.GetRaw().pixel_size)
}

func (b *Bitmap) Lock(format, flags int) LockedRegion {
	w, h := b.GetDimensions()
	ptr := C.al_lock_bitmap((*C.ALLEGRO_BITMAP)(b), C.int(format), C.int(flags))
	return CreateLockedRegion(ptr, w, h)
}

func (b *Bitmap) LockRegion(x, y, w, h, format, flags int) LockedRegion {
	ptr := C.al_lock_bitmap_region((*C.ALLEGRO_BITMAP)(b), C.int(x), C.int(y),
		C.int(w), C.int(h), C.int(format), C.int(flags))
	return CreateLockedRegion(ptr, w, h)
}

func (b *Bitmap) Unlock() {
	C.al_unlock_bitmap((*C.ALLEGRO_BITMAP)(b))
}

func GetPixelSize(format int) int {
	return int(C.al_get_pixel_size(C.int(format)))
}

func GetNewBitmapFlags() int {
	return int(C.al_get_new_bitmap_flags())
}

func GetNewBitmapFormat() int {
	return int(C.al_get_new_bitmap_format())
}

func SetNewBitmapFlags(flags int) {
	C.al_set_new_bitmap_flags(C.int(flags))
}

func SetNewBitmapFormat(format int) {
	C.al_set_new_bitmap_format(C.int(format))
}

type Bitmap C.ALLEGRO_BITMAP

func NewBitmap(w, h int) *Bitmap {
	return (*Bitmap)(C.al_create_bitmap(C.int(w), C.int(h)))
}

func GetTargetBitmap() *Bitmap {
	return (*Bitmap)(C.al_get_target_bitmap())
}

func (b *Bitmap) Destroy() {
	C.al_destroy_bitmap((*C.ALLEGRO_BITMAP)(b))
}

func (b *Bitmap) CreateSubBitmap(x, y, w, h int) *Bitmap {
	sub := C.al_create_sub_bitmap((*C.ALLEGRO_BITMAP)(b), C.int(x), C.int(y), C.int(w), C.int(h))
	return (*Bitmap)(sub)
}

func (b *Bitmap) Clone() *Bitmap {
	n := C.al_clone_bitmap((*C.ALLEGRO_BITMAP)(b))
	return (*Bitmap)(n)
}

func (b *Bitmap) GetDimensions() (int, int) {
	ptr := (*C.ALLEGRO_BITMAP)(b)
	w := C.al_get_bitmap_width(ptr)
	h := C.al_get_bitmap_height(ptr)
	return int(w), int(h)
}

func (b *Bitmap) GetPixel(x, y int) Color {
	return (Color)(C.al_get_pixel((*C.ALLEGRO_BITMAP)(b), C.int(x), C.int(y)))
}

func (b *Bitmap) IsCompatible() bool {
	return bool(C.al_is_compatible_bitmap((*C.ALLEGRO_BITMAP)(b)))
}

func (b *Bitmap) IsSubBitmap() bool {
	return bool(C.al_is_sub_bitmap((*C.ALLEGRO_BITMAP)(b)))
}

func (b *Bitmap) GetParentBitmap() *Bitmap {
	return (*Bitmap)(C.al_get_parent_bitmap((*C.ALLEGRO_BITMAP)(b)))
}

func (b *Bitmap) Draw(dx, dy float32, flags int) {
	C.al_draw_bitmap((*C.ALLEGRO_BITMAP)(b), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawTinted(dx, dy float32, flags int, color Color) {
	C.al_draw_tinted_bitmap((*C.ALLEGRO_BITMAP)(b),
		(C.ALLEGRO_COLOR)(color), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawRegion(sx, sy, sw, sh, dx, dy float32, flags int) {
	C.al_draw_bitmap_region((*C.ALLEGRO_BITMAP)(b), C.float(sx), C.float(sy),
		C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawTintedRegion(sx, sy, sw, sh, dx, dy float32, flags int, color Color) {
	C.al_draw_tinted_bitmap_region((*C.ALLEGRO_BITMAP)(b), (C.ALLEGRO_COLOR)(color), C.float(sx), C.float(sy),
		C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawRotated(cx, cy, dx, dy, angle float32, flags int) {
	C.al_draw_rotated_bitmap((*C.ALLEGRO_BITMAP)(b), C.float(cx), C.float(cy), C.float(dx), C.float(dy),
		C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawTintedRotated(cx, cy, dx, dy, angle float32, flags int, color Color) {
	C.al_draw_tinted_rotated_bitmap((*C.ALLEGRO_BITMAP)(b), (C.ALLEGRO_COLOR)(color), C.float(cx), C.float(cy),
		C.float(dx), C.float(dy), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawScaledRotated(cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
	C.al_draw_scaled_rotated_bitmap((*C.ALLEGRO_BITMAP)(b),
		C.float(cx), C.float(cy), C.float(dx), C.float(dy),
		C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawTintedScaledRotated(cx, cy, dx, dy, xscale, yscale, angle float32,
	flags int, color Color) {
	C.al_draw_tinted_scaled_rotated_bitmap((*C.ALLEGRO_BITMAP)(b),
		(C.ALLEGRO_COLOR)(color), C.float(cx), C.float(cy),
		C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle),
		C.int(flags))
}

func (b *Bitmap) DrawTintedScaledRotatedRegion(sx, sy, sw, sh, cx, cy, dx, dy, xscale, yscale,
	angle float32, flags int, color Color) {
	C.al_draw_tinted_scaled_rotated_bitmap_region((*C.ALLEGRO_BITMAP)(b), C.float(sx), C.float(sy),
		C.float(sw), C.float(sh), (C.ALLEGRO_COLOR)(color), C.float(cx), C.float(cy),
		C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle),
		C.int(flags))
}
func (b *Bitmap) DrawScaled(sx, sy, sw, sh, cx, cy, cw, ch float32, flags int) {
	C.al_draw_scaled_bitmap((*C.ALLEGRO_BITMAP)(b), C.float(sx), C.float(sy), C.float(sw), C.float(sh),
		C.float(cx), C.float(cy), C.float(cw), C.float(ch), C.int(flags))
}

func (b *Bitmap) DrawTintedScaled(sx, sy, sw, sh, cx, cy, cw, ch float32, flags int, color Color) {
	C.al_draw_tinted_scaled_bitmap((*C.ALLEGRO_BITMAP)(b), (C.ALLEGRO_COLOR)(color),
		C.float(sx), C.float(sy), C.float(sw), C.float(sh),
		C.float(cx), C.float(cy), C.float(cw), C.float(ch), C.int(flags))
}

func (b *Bitmap) SetTargetBitmap() {
	C.al_set_target_bitmap((*C.ALLEGRO_BITMAP)(b))
}

func (d *Display) SetTargetBackbuffer() {
	C.al_set_target_backbuffer((*C.ALLEGRO_DISPLAY)(d))
}

func GetCurrentDisplay() *Display {
	return (*Display)(C.al_get_current_display())
}

func GetBlender() (int, int, int) {
	var a, b, c C.int
	C.al_get_blender(&a, &b, &c)
	return int(a), int(b), int(c)
}

func SetBlender(op, src, dst int) {
	C.al_set_blender(C.int(op), C.int(src), C.int(dst))
}

func GetClippingRectangle() (int, int, int, int) {
	var x, y, w, h C.int
	C.al_get_clipping_rectangle(&x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func SetClippingRectangle(x, y, w, h int) {
	C.al_set_clipping_rectangle(C.int(x), C.int(y), C.int(w), C.int(h))
}

func ResetClippingRectangle() {
	C.al_reset_clipping_rectangle()
}

func (b *Bitmap) ConvertMaskToAlpha(color Color) {
	C.al_convert_mask_to_alpha((*C.ALLEGRO_BITMAP)(b), (C.ALLEGRO_COLOR)(color))
}

func HoldBitmapDrawing(hold bool) {
	C.al_hold_bitmap_drawing(C.bool(hold))
}

func IsBitmapDrawingHeld() bool {
	return bool(C.al_is_bitmap_drawing_held())
}

/*
// Can't see a way to do this (various problems over closures and jazz)
type BitmapLoader interface {
	LoadBitmap(string) *Bitmap
	GetExtensions() []string
}

type BitmapSaver interface {
	SaveBitmap(string, *Bitmap) bool
	GetExtensions() []string
}

func RegisterBitmapLoader(BitmapLoader loader) {
	
}

//export goLoaderCallback
func goLoaderCallback(p unsafe.Pointer, ) {
	(*(*func(string) C.ALLEGRO_BITMAP))
}
*/