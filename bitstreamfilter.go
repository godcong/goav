// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//AVRegisterBitStreamFilter Register a bitstream filter.
func (f *BitStreamFilter) AVRegisterBitStreamFilter() {
	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(f))
}

//AVBitStreamFilterNext AVBitStreamFilterNext *av_bitstream_filter_next (const BitStreamFilter *f)
func (f *BitStreamFilter) AVBitStreamFilterNext() *BitStreamFilter {
	return (*BitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
}

//AVBitStreamFilterFilter Filter bitstream.
func (bfx *BitStreamFilterContext) AVBitStreamFilterFilter(ctx *CodecContext, a string, p **uint8, ps *int, b *uint8, bs, k int) int {
	return int(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bfx), (*C.struct_AVCodecContext)(ctx), C.CString(a), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

//AVBitStreamFilterClose Release bitstream filter context.
func (bfx *BitStreamFilterContext) AVBitStreamFilterClose() {
	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bfx))
}

//AVBitStreamFilterInit Create and initialize a bitstream filter context given a bitstream filter name.
func AVBitStreamFilterInit(n string) *BitStreamFilterContext {
	return (*BitStreamFilterContext)(C.av_bitstream_filter_init(C.CString(n)))
}
