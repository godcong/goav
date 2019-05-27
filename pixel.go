// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package goav

//#cgo pkg-config: libavformat libavcodec libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libswscale/swscale.h>
import "C"

// AvPixFmtYuv ...
const (
	AvPixFmtYuv        = 0
	AvPixFmtYuv420p9   = C.AV_PIX_FMT_YUV420P9
	AvPixFmtYuv422p9   = C.AV_PIX_FMT_YUV422P9
	AvPixFmtYuv444p9   = C.AV_PIX_FMT_YUV444P9
	AvPixFmtYuv420p10  = C.AV_PIX_FMT_YUV420P10
	AvPixFmtYuv422p10  = C.AV_PIX_FMT_YUV422P10
	AvPixFmtYuv440p10  = C.AV_PIX_FMT_YUV440P10
	AvPixFmtYuv444p10  = C.AV_PIX_FMT_YUV444P10
	AvPixFmtYuv420p12  = C.AV_PIX_FMT_YUV420P12
	AvPixFmtYuv422p12  = C.AV_PIX_FMT_YUV422P12
	AvPixFmtYuv440p12  = C.AV_PIX_FMT_YUV440P12
	AvPixFmtYuv444p12  = C.AV_PIX_FMT_YUV444P12
	AvPixFmtYuv420p14  = C.AV_PIX_FMT_YUV420P14
	AvPixFmtYuv422p14  = C.AV_PIX_FMT_YUV422P14
	AvPixFmtYuv444p14  = C.AV_PIX_FMT_YUV444P14
	AvPixFmtYuv420p16  = C.AV_PIX_FMT_YUV420P16
	AvPixFmtYuv422p16  = C.AV_PIX_FMT_YUV422P16
	AvPixFmtYuv444p16  = C.AV_PIX_FMT_YUV444P16
	AvPixFmtYuva420p9  = C.AV_PIX_FMT_YUVA420P9
	AvPixFmtYuva422p9  = C.AV_PIX_FMT_YUVA422P9
	AvPixFmtYuva444p9  = C.AV_PIX_FMT_YUVA444P9
	AvPixFmtYuva420p10 = C.AV_PIX_FMT_YUVA420P10
	AvPixFmtYuva422p10 = C.AV_PIX_FMT_YUVA422P10
	AvPixFmtYuva444p10 = C.AV_PIX_FMT_YUVA444P10
	AvPixFmtYuva420p16 = C.AV_PIX_FMT_YUVA420P16
	AvPixFmtYuva422p16 = C.AV_PIX_FMT_YUVA422P16
	AvPixFmtYuva444p16 = C.AV_PIX_FMT_YUVA444P16
	AvPixFmtRgb24      = C.AV_PIX_FMT_RGB24
	AvPixFmtRgba       = C.AV_PIX_FMT_RGBA

	SwsFastBilinear     = C.SWS_FAST_BILINEAR
	SwsBilinear         = C.SWS_BILINEAR
	SwsBicubic          = C.SWS_BICUBIC
	SwsX                = C.SWS_X
	SwsPoint            = C.SWS_POINT
	SwsArea             = C.SWS_AREA
	SwsBicublin         = C.SWS_BICUBLIN
	SwsGauss            = C.SWS_GAUSS
	SwsSinc             = C.SWS_SINC
	SwsLanczos          = C.SWS_LANCZOS
	SwsSpline           = C.SWS_SPLINE
	SwsSrcVChrDropMask  = C.SWS_SRC_V_CHR_DROP_MASK
	SwsSrcVChrDropShift = C.SWS_SRC_V_CHR_DROP_SHIFT
	SwsParamDefault     = C.SWS_PARAM_DEFAULT
	SwsPrintInfo        = C.SWS_PRINT_INFO
	SwsFullChrHInt      = C.SWS_FULL_CHR_H_INT
	SwsFullChrHInp      = C.SWS_FULL_CHR_H_INP
	SwsDirectBgr        = C.SWS_DIRECT_BGR
	SwsAccurateRnd      = C.SWS_ACCURATE_RND
	SwsBitexact         = C.SWS_BITEXACT
	SwsErrorDiffusion   = C.SWS_ERROR_DIFFUSION
	SwsMaxReduceCutoff  = C.SWS_MAX_REDUCE_CUTOFF
	SwsCsItu709         = C.SWS_CS_ITU709
	SwsCsFcc            = C.SWS_CS_FCC
	SwsCsItu601         = C.SWS_CS_ITU601
	SwsCsItu624         = C.SWS_CS_ITU624
	SwsCsSmpte170m      = C.SWS_CS_SMPTE170M
	SwsCsSmpte240m      = C.SWS_CS_SMPTE240M
	SwsCsDefault        = C.SWS_CS_DEFAULT
	SwsCsBt2020         = C.SWS_CS_BT2020
)

// String ...
func (pf PixelFormat) String() string {
	switch int(pf) {
	case AvPixFmtYuv420p9:
		return "YUV420P9"

	case AvPixFmtYuv422p9:
		return "YUV422P9"

	case AvPixFmtYuv444p9:
		return "YUV444P9"

	case AvPixFmtYuv420p10:
		return "YUV420P10"

	case AvPixFmtYuv422p10:
		return "YUV422P10"

	case AvPixFmtYuv440p10:
		return "YUV440P10"

	case AvPixFmtYuv444p10:
		return "YUV444P10"

	case AvPixFmtYuv420p12:
		return "YUV420P12"

	case AvPixFmtYuv422p12:
		return "YUV422P12"

	case AvPixFmtYuv440p12:
		return "YUV440P12"

	case AvPixFmtYuv444p12:
		return "YUV444P12"

	case AvPixFmtYuv420p14:
		return "YUV420P14"

	case AvPixFmtYuv422p14:
		return "YUV422P14"

	case AvPixFmtYuv444p14:
		return "YUV444P14"

	case AvPixFmtYuv420p16:
		return "YUV420P16"

	case AvPixFmtYuv422p16:
		return "YUV422P16"

	case AvPixFmtYuv444p16:
		return "YUV444P16"

	case AvPixFmtYuva420p9:
		return "YUVA420P9"

	case AvPixFmtYuva422p9:
		return "YUVA422P9"

	case AvPixFmtYuva444p9:
		return "YUVA444P9"

	case AvPixFmtYuva420p10:
		return "YUVA420P10"

	case AvPixFmtYuva422p10:
		return "YUVA422P10"

	case AvPixFmtYuva444p10:
		return "YUVA444P10"

	case AvPixFmtYuva420p16:
		return "YUVA420P16"

	case AvPixFmtYuva422p16:
		return "YUVA422P16"

	case AvPixFmtYuva444p16:
		return "YUVA444P16"

	case AvPixFmtRgb24:
		return "RGB24"

	case AvPixFmtRgba:
		return "RGBA"
	}

	return "{UNKNOWN}"
}
