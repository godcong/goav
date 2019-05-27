// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"

//AVInputAudioDeviceNext Audio input devices iterator.
func (f *InputFormat) AVInputAudioDeviceNext() *InputFormat {
	return (*InputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(f)))
}

//AVInputVideoDeviceNext Video input devices iterator.
func (f *InputFormat) AVInputVideoDeviceNext() *InputFormat {
	return (*InputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(f)))
}

//AVOutputAudioDeviceNext Audio output devices iterator.
func (f *OutputFormat) AVOutputAudioDeviceNext() *OutputFormat {
	return (*OutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(f)))
}

//AVOutputVideoDeviceNext Video output devices iterator.
func (f *OutputFormat) AVOutputVideoDeviceNext() *OutputFormat {
	return (*OutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(f)))
}
