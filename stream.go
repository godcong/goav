// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

//AVStreamGetRFrameRate av_stream_get_r_frame_rate (const Stream *s)
func (s *Stream) AVStreamGetRFrameRate() Rational {
	return newRational(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//AVStreamSetRFrameRate void av_stream_set_r_frame_rate (Stream *s, Rational r)
func (s *Stream) AVStreamSetRFrameRate(r Rational) {
	rat := C.struct_AVRational{
		num: C.int(r.Num()),
		den: C.int(r.Den()),
	}

	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), rat)
}

//AVStreamGetParser struct CodecParserContext * av_stream_get_parser (const Stream *s)
func (s *Stream) AVStreamGetParser() *CodecParserContext {
	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

//AVStreamGetEndPts int64_t av_stream_get_end_pts (const Stream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *Stream) AVStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}
