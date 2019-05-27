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

// DeviceRect ...
type DeviceRect C.struct_AVDeviceRect

// DeviceCapabilitiesQuery ...
type DeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery

// DeviceInfo ...
type DeviceInfo C.struct_AVDeviceInfo

// DeviceInfoList ...
type DeviceInfoList C.struct_AVDeviceInfoList

// AppToDevMessageType ...
type AppToDevMessageType C.enum_AVAppToDevMessageType

// DevToAppMessageType ...
type DevToAppMessageType C.enum_AVDevToAppMessageType

//AVDeviceVersion unsigned 	avdevice_version (void)
func AVDeviceVersion() uint {
	return uint(C.avdevice_version())
}

//AVDeviceConfiguration Return the libavdevice build-time configuration.
func AVDeviceConfiguration() string {
	return C.GoString(C.avdevice_configuration())
}

//AVDeviceLicense Return the libavdevice license.
func AVDeviceLicense() string {
	return C.GoString(C.avdevice_license())
}

//AVDeviceRegisterAll Initialize libavdevice and register all the input and output devices.
func AVDeviceRegisterAll() {
	C.avdevice_register_all()
}

//AVDeviceAppToDevControlMessage Send control message from application to device.
func AVDeviceAppToDevControlMessage(ctx *FormatContext, m AppToDevMessageType, da int, d uintptr) int {
	return int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(ctx), (C.enum_AVAppToDevMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//AVDeviceDevToAppControlMessage Send control message from device to application.
func AVDeviceDevToAppControlMessage(ctx *FormatContext, m DevToAppMessageType, da int, d uintptr) int {
	return int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(ctx), (C.enum_AVDevToAppMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//AVDeviceCapabilitiesCreate Initialize capabilities probing API based on AvOption API.
func AVDeviceCapabilitiesCreate(c **DeviceCapabilitiesQuery, ctx *FormatContext, d **AVDictionary) int {
	return int(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AVDeviceCapabilitiesFree Free resources created by avdevice_capabilities_create()
func AVDeviceCapabilitiesFree(c **DeviceCapabilitiesQuery, ctx *FormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(ctx))
}

//AVDeviceListDevices List devices.
func AVDeviceListDevices(ctx *FormatContext, d **DeviceInfoList) int {
	return int(C.avdevice_list_devices((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(d))))
}

//AVDeviceFreeListDevices Convenient function to free result of avdeviceListDevices().
func AVDeviceFreeListDevices(d **DeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(d)))
}

// //int 	avdevice_list_input_sources (struct InputFormat *device, const char *device_name, AVDictionary *device_options, DeviceInfoList **device_list)
// //List devices.
// func AVDeviceListInputSources(d *InputFormat, dv string, do *AVDictionary, dl **DeviceInfoList) int {
// 	return int(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(d), C.CString(dv), (*C.struct_AVDictionary)(do), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }

// //int 	avdevice_list_output_sinks (struct OutputFormat *device, const char *device_name, AVDictionary *device_options, DeviceInfoList **device_list)
// func AVDeviceListOutputSinks(d *OutputFormat, dn string, di *AVDictionary, dl **DeviceInfoList) int {
// 	return int(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(d), C.CString(dn), (*C.struct_AVDictionary)(di), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }
