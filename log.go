// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavutil
	#include <libavutil/log.h>
	#include <stdlib.h>
*/
import "C"

// LogLevel ...
type LogLevel int

// LogLevelQuiet ...
const LogLevelQuiet LogLevel = C.AV_LOG_QUIET

// LogLevelPanic ...
const LogLevelPanic LogLevel = C.AV_LOG_PANIC

// LogLevelFatal ...
const LogLevelFatal LogLevel = C.AV_LOG_FATAL

// LogLevelError ...
const LogLevelError LogLevel = C.AV_LOG_ERROR

// LogLevelWarning ...
const LogLevelWarning LogLevel = C.AV_LOG_WARNING

// LogLevelInfo ...
const LogLevelInfo LogLevel = C.AV_LOG_INFO

// LogLevelVerbose ...
const LogLevelVerbose LogLevel = C.AV_LOG_VERBOSE

// LogLevelDebug ...
const LogLevelDebug LogLevel = C.AV_LOG_DEBUG

// LogLevelTrace ...
const LogLevelTrace LogLevel = C.AV_LOG_TRACE

// AVLogSetLevel ...
func AVLogSetLevel(level LogLevel) {
	C.av_log_set_level(C.int(level))
}

// AVLogGetLevel ...
func AVLogGetLevel() LogLevel {
	return (LogLevel)(C.av_log_get_level())
}
