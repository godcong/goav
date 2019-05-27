// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavcodec/avcodec.h>
//#include <libavformat/avformat.h>
//#include <libavformat/avio.h>
import "C"

// AVCodecFlagUnaligned ...
const (
	AVCodecFlagUnaligned          = int(C.AV_CODEC_FLAG_UNALIGNED)
	AVCodecFlagQscale             = int(C.AV_CODEC_FLAG_QSCALE)
	AVCodecFlag4mv                = int(C.AV_CODEC_FLAG_4MV)
	AVCodecFlagOutputCorrupt      = int(C.AV_CODEC_FLAG_OUTPUT_CORRUPT)
	AVCodecFlagQpel               = int(C.AV_CODEC_FLAG_QPEL)
	AVCodecFlagPass1              = int(C.AV_CODEC_FLAG_PASS1)
	AVCodecFlagPass2              = int(C.AV_CODEC_FLAG_PASS2)
	AVCodecFlagLoopFilter         = int(C.AV_CODEC_FLAG_LOOP_FILTER)
	AVCodecFlagGray               = int(C.AV_CODEC_FLAG_GRAY)
	AVCodecFlagPsnr               = int(C.AV_CODEC_FLAG_PSNR)
	AVCodecFlagTruncated          = int(C.AV_CODEC_FLAG_TRUNCATED)
	AVCodecFlagInterlacedDct      = int(C.AV_CODEC_FLAG_INTERLACED_DCT)
	AVCodecFlagLowDelay           = int(C.AV_CODEC_FLAG_LOW_DELAY)
	AVCodecFlagGlobalHeader       = int(C.AV_CODEC_FLAG_GLOBAL_HEADER)
	AVCodecFlagBitexact           = int(C.AV_CODEC_FLAG_BITEXACT)
	AVCodecFlagAcPred             = int(C.AV_CODEC_FLAG_AC_PRED)
	AVCodecFlagInterlacedMe       = int(C.AV_CODEC_FLAG_INTERLACED_ME)
	AVCodecFlagClosedGop          = int(C.AV_CODEC_FLAG_CLOSED_GOP)
	AVCodecFlag2Fast              = int(C.AV_CODEC_FLAG2_FAST)
	AVCodecFlag2NoOutput          = int(C.AV_CODEC_FLAG2_NO_OUTPUT)
	AVCodecFlag2LocalHeader       = int(C.AV_CODEC_FLAG2_LOCAL_HEADER)
	AVCodecFlag2DropFrameTimecode = int(C.AV_CODEC_FLAG2_DROP_FRAME_TIMECODE)
	AVCodecFlag2Chunks            = int(C.AV_CODEC_FLAG2_CHUNKS)
	AVCodecFlag2IgnoreCrop        = int(C.AV_CODEC_FLAG2_IGNORE_CROP)
	AVCodecFlag2ShowAll           = int(C.AV_CODEC_FLAG2_SHOW_ALL)
	AVCodecFlag2ExportMvs         = int(C.AV_CODEC_FLAG2_EXPORT_MVS)
	AVCodecFlag2SkipManual        = int(C.AV_CODEC_FLAG2_SKIP_MANUAL)
	AVCodecFlag2RoFlushNoop       = int(C.AV_CODEC_FLAG2_RO_FLUSH_NOOP)
)

// AvioFlagRead ...
const (
	AvioFlagRead      = int(C.AVIO_FLAG_READ)
	AvioFlagWrite     = int(C.AVIO_FLAG_WRITE)
	AvioFlagReadWrite = int(C.AVIO_FLAG_READ_WRITE)
)
