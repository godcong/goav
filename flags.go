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

// CodecFlag ...
type CodecFlag int

// CodecFlagUnaligned ...
const CodecFlagUnaligned CodecFlag = C.AV_CODEC_FLAG_UNALIGNED

// CodecFlagQscale ...
const CodecFlagQscale CodecFlag = C.AV_CODEC_FLAG_QSCALE

// CodecFlag4mv ...
const CodecFlag4mv CodecFlag = C.AV_CODEC_FLAG_4MV

// CodecFlagOutputCorrupt ...
const CodecFlagOutputCorrupt CodecFlag = C.AV_CODEC_FLAG_OUTPUT_CORRUPT

// CodecFlagQpel ...
const CodecFlagQpel CodecFlag = C.AV_CODEC_FLAG_QPEL

// CodecFlagPass1 ...
const CodecFlagPass1 CodecFlag = C.AV_CODEC_FLAG_PASS1

// CodecFlagPass2 ...
const CodecFlagPass2 CodecFlag = C.AV_CODEC_FLAG_PASS2

// CodecFlagLoopFilter ...
const CodecFlagLoopFilter CodecFlag = C.AV_CODEC_FLAG_LOOP_FILTER

// CodecFlagGray ...
const CodecFlagGray CodecFlag = C.AV_CODEC_FLAG_GRAY

// CodecFlagPsnr ...
const CodecFlagPsnr CodecFlag = C.AV_CODEC_FLAG_PSNR

// CodecFlagTruncated ...
const CodecFlagTruncated CodecFlag = C.AV_CODEC_FLAG_TRUNCATED

// CodecFlagInterlacedDct ...
const CodecFlagInterlacedDct CodecFlag = C.AV_CODEC_FLAG_INTERLACED_DCT

// CodecFlagLowDelay ...
const CodecFlagLowDelay CodecFlag = C.AV_CODEC_FLAG_LOW_DELAY

// CodecFlagGlobalHeader ...
const CodecFlagGlobalHeader CodecFlag = C.AV_CODEC_FLAG_GLOBAL_HEADER

// CodecFlagBitexact ...
const CodecFlagBitexact CodecFlag = C.AV_CODEC_FLAG_BITEXACT

// CodecFlagAcPred ...
const CodecFlagAcPred CodecFlag = C.AV_CODEC_FLAG_AC_PRED

// CodecFlagInterlacedMe ...
const CodecFlagInterlacedMe CodecFlag = C.AV_CODEC_FLAG_INTERLACED_ME

// CodecFlagClosedGop ...
const CodecFlagClosedGop CodecFlag = C.AV_CODEC_FLAG_CLOSED_GOP

// AvioFlagRead ...
const (
	AVIOFlagRead      = int(C.AVIO_FLAG_READ)
	AVIOFlagWrite     = int(C.AVIO_FLAG_WRITE)
	AVIOFlagReadWrite = int(C.AVIO_FLAG_READ_WRITE)
)

// CodecFlag2 ...
type CodecFlag2 int

// CodecFlag2Fast ...
const CodecFlag2Fast CodecFlag2 = C.AV_CODEC_FLAG2_FAST

// CodecFlag2NoOutput ...
const CodecFlag2NoOutput CodecFlag2 = C.AV_CODEC_FLAG2_NO_OUTPUT

// CodecFlag2LocalHeader ...
const CodecFlag2LocalHeader CodecFlag2 = C.AV_CODEC_FLAG2_LOCAL_HEADER

// CodecFlag2DropFrameTimecode ...
const CodecFlag2DropFrameTimecode CodecFlag2 = C.AV_CODEC_FLAG2_DROP_FRAME_TIMECODE

// CodecFlag2IgnoreCrop ...
const CodecFlag2IgnoreCrop CodecFlag2 = C.AV_CODEC_FLAG2_IGNORE_CROP

// CodecFlag2Chunks ...
const CodecFlag2Chunks CodecFlag2 = C.AV_CODEC_FLAG2_CHUNKS

// CodecFlag2ShowAll ...
const CodecFlag2ShowAll CodecFlag2 = C.AV_CODEC_FLAG2_SHOW_ALL

// CodecFlag2ExportMvs ...
const CodecFlag2ExportMvs CodecFlag2 = C.AV_CODEC_FLAG2_EXPORT_MVS

// CodecFlag2SkipManual ...
const CodecFlag2SkipManual CodecFlag2 = C.AV_CODEC_FLAG2_SKIP_MANUAL

// CodecFlag2RoFlushNoop ...
const CodecFlag2RoFlushNoop CodecFlag = C.AV_CODEC_FLAG2_RO_FLUSH_NOOP
