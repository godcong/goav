// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavutil
	#include <libavutil/log.h>
	#include <stdlib.h>
*/
import "C"

type LogLevel int

const (
	LogLevelQuiet   LogLevel = C.AV_LOG_QUIET
	LogLevelPanic   LogLevel = C.AV_LOG_PANIC
	LogLevelFatal   LogLevel = C.AV_LOG_FATAL
	LogLevelError   LogLevel = C.AV_LOG_ERROR
	LogLevelWarning LogLevel = C.AV_LOG_WARNING
	LogLevelInfo    LogLevel = C.AV_LOG_INFO
	LogLevelVerbose LogLevel = C.AV_LOG_VERBOSE
	LogLevelDebug   LogLevel = C.AV_LOG_DEBUG
	LogLevelTrace   LogLevel = C.AV_LOG_TRACE
)

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

// AVLogSetLevel ...
func AVLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}

// AVLogGetLevel ...
func AVLogGetLevel() int {
	return int(C.av_log_get_level())
}
