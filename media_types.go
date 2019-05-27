// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package goav

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

// AvmediaTypeUnknown ...
const AVmediaTypeUnknown = C.AVMEDIA_TYPE_UNKNOWN

// AVMediaTypeVideo ...
const AVMediaTypeVideo = C.AVMEDIA_TYPE_VIDEO

// AVmediaTypeAudio ...
const AVmediaTypeAudio = C.AVMEDIA_TYPE_AUDIO

// AVmediaTypeData ...
const AVmediaTypeData = C.AVMEDIA_TYPE_DATA

// AVmediaTypeSubtitle ...
const AVmediaTypeSubtitle = C.AVMEDIA_TYPE_SUBTITLE

// AVmediaTypeAttachment ...
const AVmediaTypeAttachment = C.AVMEDIA_TYPE_ATTACHMENT

// AVmediaTypeNb ...
const AVmediaTypeNb = C.AVMEDIA_TYPE_NB
