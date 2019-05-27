// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//Register a bitstream filter.
func (f *AVBitStreamFilter) AVRegisterBitStreamFilter() {
	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(f))
}

//AVBitStreamFilterNext *av_bitstream_filter_next (const AVBitStreamFilter *f)
func (f *AVBitStreamFilter) AVBitStreamFilterNext() *AVBitStreamFilter {
	return (*AVBitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
}

//AVBitStreamFilter Filter bitstream.
func (bfx *AVBitStreamFilterContext) AVBitStreamFilterFilter(ctx *AVCodecContext, a string, p **uint8, ps *int, b *uint8, bs, k int) int {
	return int(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bfx), (*C.struct_AVCodecContext)(ctx), C.CString(a), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

//AVBitStreamFilterClose Release bitstream filter context.
func (bfx *AVBitStreamFilterContext) AVBitStreamFilterClose() {
	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bfx))
}

//AVBitStreamFilterInit Create and initialize a bitstream filter context given a bitstream filter name.
func AVBitStreamFilterInit(n string) *AVBitStreamFilterContext {
	return (*AVBitStreamFilterContext)(C.av_bitstream_filter_init(C.CString(n)))
}
