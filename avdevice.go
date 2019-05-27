// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package goav deals with the input and output devices provided by the libavdevice library
// The libavdevice library provides the same interface as libavformat.
// Namely, an input device is considered like a demuxer, and an output device like a muxer,
// and the interface and generic device options are the same provided by libavformat
package goav

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"
import (
	"unsafe"
)

// AvDeviceRect ...
type (
	AvDeviceRect              C.struct_AVDeviceRect
	AvDeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery
	AvDeviceInfo              C.struct_AVDeviceInfo
	AvDeviceInfoList          C.struct_AVDeviceInfoList
	//AVInputFormat               C.struct_AVInputFormat
	//AVOutputFormat              C.struct_AVOutputFormat
	//AVFormatContext C.struct_AVFormatContext
	//Dictionary                C.struct_AVDictionary
	AvAppToDevMessageType C.enum_AVAppToDevMessageType
	AvDevToAppMessageType C.enum_AVDevToAppMessageType
)

//unsigned 	avdevice_version (void)
func AVDeviceVersion() uint {
	return uint(C.avdevice_version())
}

//Return the libavdevice build-time configuration.
func AVDeviceConfiguration() string {
	return C.GoString(C.avdevice_configuration())
}

//Return the libavdevice license.
func AVDeviceLicense() string {
	return C.GoString(C.avdevice_license())
}

//Initialize libavdevice and register all the input and output devices.
func AVDeviceRegisterAll() {
	C.avdevice_register_all()
}

//AVDeviceAppToDevControlMessage Send control message from application to device.
func AVDeviceAppToDevControlMessage(ctx *AVFormatContext, m AvAppToDevMessageType, da int, d uintptr) int {
	return int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(ctx), (C.enum_AVAppToDevMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//AVDeviceDevToAppControlMessage Send control message from device to application.
func AVDeviceDevToAppControlMessage(ctx *AVFormatContext, m AvDevToAppMessageType, da int, d uintptr) int {
	return int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(ctx), (C.enum_AVDevToAppMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//AVDeviceCapabilitiesCreate Initialize capabilities probing API based on AvOption API.
func AVDeviceCapabilitiesCreate(c **AvDeviceCapabilitiesQuery, ctx *AVFormatContext, d **Dictionary) int {
	return int(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AVDeviceCapabilitiesFree Free resources created by avdevice_capabilities_create()
func AVDeviceCapabilitiesFree(c **AvDeviceCapabilitiesQuery, ctx *AVFormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(ctx))
}

//AVDeviceListDevices List devices.
func AVDeviceListDevices(ctx *AVFormatContext, d **AvDeviceInfoList) int {
	return int(C.avdevice_list_devices((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(d))))
}

//AVDeviceFreeListDevices Convenient function to free result of avdeviceListDevices().
func AVDeviceFreeListDevices(d **AvDeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(d)))
}

// //int 	avdevice_list_input_sources (struct AVInputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// //List devices.
// func AVDeviceListInputSources(d *AVInputFormat, dv string, do *Dictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(d), C.CString(dv), (*C.struct_AVDictionary)(do), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }

// //int 	avdevice_list_output_sinks (struct AVOutputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// func AVDeviceListOutputSinks(d *AVOutputFormat, dn string, di *Dictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(d), C.CString(dn), (*C.struct_AVDictionary)(di), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }
