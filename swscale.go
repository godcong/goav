// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package goav performs highly optimized image scaling and colorspace and pixel format conversion operations.
//Rescaling: is the process of changing the video size. Several rescaling options and algorithms are available.
//Pixel format conversion: is the process of converting the image format and colorspace of the image.
package goav

//#cgo pkg-config: libswscale libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <string.h>
//#include <stdint.h>
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

// SwsContext ...
type SwsContext C.struct_SwsContext

// SwsFilter ...
type SwsFilter C.struct_SwsFilter

// SwsVector ...
type SwsVector C.struct_SwsVector

// Class ...
type Class C.struct_AVClass

//SwscaleVersion Return the LIBSWSCALE_VERSION_INT constant.
func SwscaleVersion() uint {
	return uint(C.swscale_version())
}

//SwscaleConfiguration Return the libswscale build-time configuration.
func SwscaleConfiguration() string {
	return C.GoString(C.swscale_configuration())
}

//SwscaleLicense Return the libswscale license.
func SwscaleLicense() string {
	return C.GoString(C.swscale_license())
}

//SwsGetCoefficients Return a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
func SwsGetCoefficients(c int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(c))))
}

//SwsIsSupportedInput Return a positive value if pix_fmt is a supported input format, 0 otherwise.
func SwsIsSupportedInput(p PixelFormat) int {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(p)))
}

//SwsIsSupportedOutput Return a positive value if pix_fmt is a supported output format, 0 otherwise.
func SwsIsSupportedOutput(p PixelFormat) int {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(p)))
}

// SwsIssupportedendiannessconversion ...
func SwsIssupportedendiannessconversion(p PixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(p)))
}

//SwsScale Scale the image slice in srcSlice and put the resulting scaled slice in the image in dst.
func SwsScale(ctxt *SwsContext, src *uint8, str int, y, h int, d *uint8, ds int) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (*C.uint8_t)(unsafe.Pointer(src))
	cstr := (*C.int)(unsafe.Pointer(&str))
	cd := (*C.uint8_t)(unsafe.Pointer(d))
	cds := (*C.int)(unsafe.Pointer(&ds))
	return int(C.sws_scale(cctxt, &csrc, cstr, C.int(y), C.int(h), &cd, cds))
}

// SwsScale2 ...
func SwsScale2(ctxt *SwsContext, srcData [8]*uint8, srcStride [8]int32, y, h int, dstData [8]*uint8, dstStride [8]int32) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (**C.uint8_t)(unsafe.Pointer(&srcData[0]))
	cstr := (*C.int)(unsafe.Pointer(&srcStride[0]))
	cd := (**C.uint8_t)(unsafe.Pointer(&dstData[0]))
	cds := (*C.int)(unsafe.Pointer(&dstStride))
	return int(C.sws_scale(cctxt, csrc, cstr, C.int(y), C.int(h), cd, cds))
}

// SwsSetColorspaceDetails ...
func SwsSetColorspaceDetails(ctxt *SwsContext, it *int, sr int, t *int, dr, b, c, s int) int {
	cit := (*C.int)(unsafe.Pointer(it))
	ct := (*C.int)(unsafe.Pointer(t))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, C.int(sr), ct, C.int(dr), C.int(b), C.int(c), C.int(s)))
}

// SwsGetColorspaceDetails ...
func SwsGetColorspaceDetails(ctxt *SwsContext, it, sr, t, dr, b, c, s *int) int {
	cit := (**C.int)(unsafe.Pointer(it))
	csr := (*C.int)(unsafe.Pointer(sr))
	ct := (**C.int)(unsafe.Pointer(t))
	cdr := (*C.int)(unsafe.Pointer(dr))
	cb := (*C.int)(unsafe.Pointer(b))
	cc := (*C.int)(unsafe.Pointer(c))
	cs := (*C.int)(unsafe.Pointer(s))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, csr, ct, cdr, cb, cc, cs))
}

// SwsGetDefaultFilter ...
func SwsGetDefaultFilter(lb, cb, ls, cs, chs, cvs float32, v int) *SwsFilter {
	return (*SwsFilter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lb), C.float(cb), C.float(ls), C.float(cs), C.float(chs), C.float(cvs), C.int(v))))
}

// SwsFreeFilter ...
func SwsFreeFilter(f *SwsFilter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(f))
}

//SwsConvertPalette8ToPacked32 Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
func SwsConvertPalette8ToPacked32(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//SwsConvertPalette8ToPacked24 Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
func SwsConvertPalette8ToPacked24(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//SwsGetClass Get the Class for swsContext.
func SwsGetClass() *Class {
	return (*Class)(C.sws_get_class())
}
