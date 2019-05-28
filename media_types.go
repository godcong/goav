// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package goav

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

// MediaType ...
type MediaType C.enum_AVMediaType

// AVMediaTypeUnknown ...
const AVMediaTypeUnknown MediaType = C.AVMEDIA_TYPE_UNKNOWN

// AVMediaTypeVideo ...
const AVMediaTypeVideo MediaType = C.AVMEDIA_TYPE_VIDEO

// AVMediaTypeAudio ...
const AVMediaTypeAudio MediaType = C.AVMEDIA_TYPE_AUDIO

// AVMediaTypeData ...
const AVMediaTypeData MediaType = C.AVMEDIA_TYPE_DATA

// AVMediaTypeSubtitle ...
const AVMediaTypeSubtitle MediaType = C.AVMEDIA_TYPE_SUBTITLE

// AVMediaTypeAttachment ...
const AVMediaTypeAttachment MediaType = C.AVMEDIA_TYPE_ATTACHMENT

// AVMediaTypeNb ...
const AVMediaTypeNb MediaType = C.AVMEDIA_TYPE_NB
