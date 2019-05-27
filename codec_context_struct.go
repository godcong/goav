// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package goav provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package goav

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"reflect"
	"unsafe"
)

// Type ...
func (ctx *AVCodecContext) Type() AVMediaType {
	return AVMediaType(ctx.codec_type)
}

// SetBitRate ...
func (ctx *AVCodecContext) SetBitRate(br int64) {
	ctx.bit_rate = C.int64_t(br)
}

// GetCodecID ...
func (ctx *AVCodecContext) GetCodecID() AVCodecID {
	return AVCodecID(ctx.codec_id)
}

// SetCodecID ...
func (ctx *AVCodecContext) SetCodecID(id AVCodecID) {
	ctx.codec_id = C.enum_AVCodecID(id)
}

// GetCodecType ...
func (ctx *AVCodecContext) GetCodecType() AVMediaType {
	return AVMediaType(ctx.codec_type)
}

// SetCodecType ...
func (ctx *AVCodecContext) SetCodecType(ctype AVMediaType) {
	ctx.codec_type = C.enum_AVMediaType(ctype)
}

// GetTimeBase ...
func (ctx *AVCodecContext) GetTimeBase() AVRational {
	return newRational(ctx.time_base)
}

// SetTimeBase ...
func (ctx *AVCodecContext) SetTimeBase(timeBase AVRational) {
	ctx.time_base.num = C.int(timeBase.Num())
	ctx.time_base.den = C.int(timeBase.Den())
}

// GetWidth ...
func (ctx *AVCodecContext) GetWidth() int {
	return int(ctx.width)
}

// SetWidth ...
func (ctx *AVCodecContext) SetWidth(w int) {
	ctx.width = C.int(w)
}

// GetHeight ...
func (ctx *AVCodecContext) GetHeight() int {
	return int(ctx.height)
}

// SetHeight ...
func (ctx *AVCodecContext) SetHeight(h int) {
	ctx.height = C.int(h)
}

// GetPixelFormat ...
func (ctx *AVCodecContext) GetPixelFormat() AVPixelFormat {
	return AVPixelFormat(C.int(ctx.pix_fmt))
}

// SetPixelFormat ...
func (ctx *AVCodecContext) SetPixelFormat(fmt AVPixelFormat) {
	ctx.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

// GetFlags ...
func (ctx *AVCodecContext) GetFlags() int {
	return int(ctx.flags)
}

// SetFlags ...
func (ctx *AVCodecContext) SetFlags(flags int) {
	ctx.flags = C.int(flags)
}

// GetMeRange ...
func (ctx *AVCodecContext) GetMeRange() int {
	return int(ctx.me_range)
}

// SetMeRange ...
func (ctx *AVCodecContext) SetMeRange(r int) {
	ctx.me_range = C.int(r)
}

// GetMaxQDiff ...
func (ctx *AVCodecContext) GetMaxQDiff() int {
	return int(ctx.max_qdiff)
}

// SetMaxQDiff ...
func (ctx *AVCodecContext) SetMaxQDiff(v int) {
	ctx.max_qdiff = C.int(v)
}

// GetQMin ...
func (ctx *AVCodecContext) GetQMin() int {
	return int(ctx.qmin)
}

// SetQMin ...
func (ctx *AVCodecContext) SetQMin(v int) {
	ctx.qmin = C.int(v)
}

// GetQMax ...
func (ctx *AVCodecContext) GetQMax() int {
	return int(ctx.qmax)
}

// SetQMax ...
func (ctx *AVCodecContext) SetQMax(v int) {
	ctx.qmax = C.int(v)
}

// GetQCompress ...
func (ctx *AVCodecContext) GetQCompress() float32 {
	return float32(ctx.qcompress)
}

// SetQCompress ...
func (ctx *AVCodecContext) SetQCompress(v float32) {
	ctx.qcompress = C.float(v)
}

// GetExtraData ...
func (ctx *AVCodecContext) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctx.extradata)),
		Len:  int(ctx.extradata_size),
		Cap:  int(ctx.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

// SetExtraData ...
func (ctx *AVCodecContext) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	ctx.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	ctx.extradata_size = C.int(header.Len)
}

// Release ...
func (ctx *AVCodecContext) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(ctx)))
	C.av_freep(unsafe.Pointer(ctx))
}
