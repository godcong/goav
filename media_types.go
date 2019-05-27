// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package goav

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

// AVMediaTypeUnknown ...
const AVMediaTypeUnknown = C.AVMEDIA_TYPE_UNKNOWN

// AVMediaTypeVideo ...
const AVMediaTypeVideo = C.AVMEDIA_TYPE_VIDEO

// AVMediaTypeAudio ...
const AVMediaTypeAudio = C.AVMEDIA_TYPE_AUDIO

// AVMediaTypeData ...
const AVMediaTypeData = C.AVMEDIA_TYPE_DATA

// AVMediaTypeSubtitle ...
const AVMediaTypeSubtitle = C.AVMEDIA_TYPE_SUBTITLE

// AVMediaTypeAttachment ...
const AVMediaTypeAttachment = C.AVMEDIA_TYPE_ATTACHMENT

// AVMediaTypeNb ...
const AVMediaTypeNb = C.AVMEDIA_TYPE_NB
