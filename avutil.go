// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package goav is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package goav

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavutil/pixdesc.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// Options ...
type Options C.struct_AVOptions

// Tree ...
type Tree C.struct_AVTree

// Rational ...
type Rational C.struct_AVRational

// PixelFormat ...
type PixelFormat C.enum_AVPixelFormat

// ChromaLocation ...
type ChromaLocation C.enum_AVChromaLocation

// ChromaLocationUnspecified ...
const ChromaLocationUnspecified ChromaLocation = C.AVCHROMA_LOC_UNSPECIFIED

// ChromaLocationLeft ...
const ChromaLocationLeft ChromaLocation = C.AVCHROMA_LOC_LEFT

// ChromaLocationCenter ...
const ChromaLocationCenter ChromaLocation = C.AVCHROMA_LOC_CENTER

// ChromaLocationTopLeft ...
const ChromaLocationTopLeft ChromaLocation = C.AVCHROMA_LOC_TOPLEFT

// ChromaLocationTop ...
const ChromaLocationTop ChromaLocation = C.AVCHROMA_LOC_TOP

// ChromaLocationBottomLeft ...
const ChromaLocationBottomLeft ChromaLocation = C.AVCHROMA_LOC_BOTTOMLEFT

// ChromaLocationBottom ...
const ChromaLocationBottom ChromaLocation = C.AVCHROMA_LOC_BOTTOM

// ErrorCode ...
type ErrorCode int

// OptionSearchFlags ...
type OptionSearchFlags int

// OptionSearchChildren ...
const (
	OptionSearchChildren OptionSearchFlags = C.AV_OPT_SEARCH_CHILDREN
	OptionSearchFakeObj  OptionSearchFlags = C.AV_OPT_SEARCH_FAKE_OBJ
)

// LossFlags ...
type LossFlags int

// LossFlagNone ...
const LossFlagNone LossFlags = 0

// LossFlagResolution ...
const LossFlagResolution LossFlags = C.FF_LOSS_RESOLUTION

// LossFlagDepth ...
const LossFlagDepth LossFlags = C.FF_LOSS_DEPTH

// LossFlagColorspace ...
const LossFlagColorspace LossFlags = C.FF_LOSS_COLORSPACE

// LossFlagAlpha ...
const LossFlagAlpha LossFlags = C.FF_LOSS_ALPHA

// LossFlagColorquant ...
const LossFlagColorquant LossFlags = C.FF_LOSS_COLORQUANT

// LossFlagChroma ...
const LossFlagChroma LossFlags = C.FF_LOSS_CHROMA

// LossFlagAll ...
const LossFlagAll LossFlags = -1

// File ...
type File C.FILE

func init() {
	AVLogSetLevel(LogLevelQuiet)
}

//AVUtilVersion Return the LIBAvUTIL_VERSION_INT constant.
func AVUtilVersion() uint {
	return uint(C.avutil_version())
}

//AVUtilConfiguration Return the libavutil build-time configuration.
func AVUtilConfiguration() string {
	return C.GoString(C.avutil_configuration())
}

//AVUtilLicense Return the libavutil license.
func AVUtilLicense() string {
	return C.GoString(C.avutil_license())
}

//AVGetMediaTypeString Return a string describing the media_type enum, NULL if media_type is unknown.
func AVGetMediaTypeString(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

//AVGetPictureTypeChar Return a single letter to describe the given picture type pict_type.
func AVGetPictureTypeChar(pt PictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

//AVXIfNull Return x default pointer in case p is NULL.
func AVXIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

//AVIntListLengthForSize Compute the length of an integer list.
func AVIntListLengthForSize(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

//AVFopenUtf8 Open a file using a UTF-8 filename.
func AVFopenUtf8(p, m string) *File {
	f := C.av_fopen_utf8(C.CString(p), C.CString(m))
	return (*File)(f)
}

//AVGetTimeBaseQ Return the fractional representation of the internal time base.
func AVGetTimeBaseQ() Rational {
	return (Rational)(C.av_get_time_base_q())
}
