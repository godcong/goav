// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

//SwsAllocVec Allocate and return an uninitialized vector with length coefficients.
func SwsAllocVec(l int) *SwsVector {
	return (*SwsVector)(C.sws_allocVec(C.int(l)))
}

//SwsGetGaussianVec Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func SwsGetGaussianVec(v, q float64) *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(v), C.double(q))))
}

//SwsGetConstVec Allocate and return a vector with length coefficients, all with the same value c.
func SwsGetConstVec(c float64, l int) *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getConstVec(C.double(c), C.int(l))))
}

//SwsGetIdentityVec Allocate and return a vector with just one coefficient, with value 1.0.
func SwsGetIdentityVec() *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getIdentityVec()))
}

//SwsScaleVec Scale all the coefficients of a by the scalar value.
func (a *SwsVector) SwsScaleVec(s float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(a)), C.double(s))
}

//SwsNormalizeVec Scale all the coefficients of a so that their sum equals height.
func (a *SwsVector) SwsNormalizeVec(h float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), C.double(h))
}

// SwsConvVec ...
func (a *SwsVector) SwsConvVec(b *SwsVector) {
	C.sws_convVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// SwsAddVec ...
func (a *SwsVector) SwsAddVec(b *SwsVector) {
	C.sws_addVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// Deprecated:SwsSubVec ...
func (a *SwsVector) SwsSubVec(b *SwsVector) {
	C.sws_subVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// Deprecated:SwsShiftVec ...
func (a *SwsVector) SwsShiftVec(s int) {
	C.sws_shiftVec((*C.struct_SwsVector)(a), C.int(s))
}

//SwsCloneVec Allocate and return a clone of the vector a, that is a vector with the same coefficients as a.
func (a *SwsVector) SwsCloneVec() *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_cloneVec((*C.struct_SwsVector)(a))))
}

//Deprecated:SwsPrintVec2 Print with av_log() a textual representation of the vector a if log_level <= av_log_level.
func (a *SwsVector) SwsPrintVec2(lctx *Class, l int) {
	C.sws_printVec2((*C.struct_SwsVector)(a), (*C.struct_AVClass)(lctx), C.int(l))
}

// SwsFreeVec ...
func (a *SwsVector) SwsFreeVec() {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}
