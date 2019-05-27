// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

//AVRational av_stream_get_r_frame_rate (const AVStream *s)
func (s *AVStream) AvStreamGetRFrameRate() AVRational {
	return newRational(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (AVStream *s, AVRational r)
func (s *AVStream) AvStreamSetRFrameRate(r AVRational) {
	rat := C.struct_AVRational{
		num: C.int(r.Num()),
		den: C.int(r.Den()),
	}

	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), rat)
}

//struct AVCodecParserContext * av_stream_get_parser (const AVStream *s)
func (s *AVStream) AvStreamGetParser() *AVCodecParserContext {
	return (*AVCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

//int64_t av_stream_get_end_pts (const AVStream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *AVStream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}
