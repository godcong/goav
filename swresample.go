// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package goav provides a high-level interface to the libswresample library audio resampling utilities
// The process of changing the sampling rate of a discrete signal to obtain a new discrete representation of the underlying continuous signal.
package goav

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"

// SwrContext ...
type (
	SwrContext C.struct_SwrContext
	//AVFrame C.struct_AVFrame
	//Class C.struct_AVClass
	//AvSampleFormat C.enum_AVSampleFormat
)

//SwrGetClass Get the Class for SwrContext.
func SwrGetClass() *Class {
	return (*Class)(C.swr_get_class())
}

//SwrAlloc SwrContext constructor functions.Allocate SwrContext.
func SwrAlloc() *SwrContext {
	return (*SwrContext)(C.swr_alloc())
}

//SwreSampleVersion Configuration accessors
func SwreSampleVersion() uint {
	return uint(C.swresample_version())
}

// SwreSampleConfiguration ...
func SwreSampleConfiguration() string {
	return C.GoString(C.swresample_configuration())
}

// SwreSampleLicense ...
func SwreSampleLicense() string {
	return C.GoString(C.swresample_license())
}
