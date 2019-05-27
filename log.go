// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavutil
	#include <libavutil/log.h>
	#include <stdlib.h>
*/
import "C"

// AVLogQuiet ...
const (
	AVLogQuiet   = -8
	AVLogPanic   = 0
	AVLogFatal   = 8
	AVLogError   = 16
	AVLogWarning = 24
	AVLogInfo    = 32
	AVLogVerbose = 40
	AVLogDebug   = 48
	AVLogTrace   = 56
)

// AvLogSetLevel ...
func AvLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}

// AvLogGetLevel ...
func AvLogGetLevel() int {
	return int(C.av_log_get_level())
}
