// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//Utility function to access log2_chroma_w log2_chroma_h from the pixel format AvPixFmtDescriptor.
func (pf AVPixelFormat) AVCodecGetChromaSubSample(h, v *int) {
	C.avcodec_get_chroma_sub_sample((C.enum_AVPixelFormat)(pf), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(v)))
}

//Return a value representing the fourCC code associated to the pixel format pix_fmt, or 0 if no associated fourCC code can be found.
func (pf AVPixelFormat) AVCodecPixFmtToCodecTag() uint {
	return uint(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(pf)))
}

// AVCodecGetPixFmtLoss ...
func (pf AVPixelFormat) AVCodecGetPixFmtLoss(f AVPixelFormat, a int) int {
	return int(C.avcodec_get_pix_fmt_loss((C.enum_AVPixelFormat)(pf), (C.enum_AVPixelFormat)(f), C.int(a)))
}

//Find the best pixel format to convert to given a certain source pixel format.
func (pf *AVPixelFormat) AVCodecFindBestPixFmtOfList(s AVPixelFormat, a int, l *int) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_find_best_pix_fmt_of_list((*C.enum_AVPixelFormat)(pf), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(l))))
}

// AVCodecFindBestPixFmtOf2 ...
func (pf AVPixelFormat) AVCodecFindBestPixFmtOf2(f2, s AVPixelFormat, a int, l *int) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_find_best_pix_fmt_of_2((C.enum_AVPixelFormat)(pf), (C.enum_AVPixelFormat)(f2), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(l))))
}
