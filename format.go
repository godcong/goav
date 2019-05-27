// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"

//Audio input devices iterator.
func (f *AVInputFormat) AvInputAudioDeviceNext() *AVInputFormat {
	return (*AVInputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(f)))
}

//Video input devices iterator.
func (f *AVInputFormat) AvInputVideoDeviceNext() *AVInputFormat {
	return (*AVInputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(f)))
}

//Audio output devices iterator.
func (f *AVOutputFormat) AvOutputAudioDeviceNext() *AVOutputFormat {
	return (*AVOutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(f)))
}

//Video output devices iterator.
func (f *AVOutputFormat) AvOutputVideoDeviceNext() *AVOutputFormat {
	return (*AVOutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(f)))
}
