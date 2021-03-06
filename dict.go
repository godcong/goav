// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package goav is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package goav

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/dict.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// AVDictionary ...
type AVDictionary C.struct_AVDictionary

// AVDictionaryEntry ...
type AVDictionaryEntry C.struct_AVDictionaryEntry

// AV_DICT_MATCH_CASE ...
const (
	AVDictMatchCase     = int(C.AV_DICT_MATCH_CASE)
	AVDictIgnoreSuffix  = int(C.AV_DICT_IGNORE_SUFFIX)
	AVDictDontStrdupKey = int(C.AV_DICT_DONT_STRDUP_KEY)
	AVDictDontStrdupVal = int(C.AV_DICT_DONT_STRDUP_VAL)
	AVDictDontOverwrite = int(C.AV_DICT_DONT_OVERWRITE)
	AVDictAppend        = int(C.AV_DICT_APPEND)
	AVDictMultikey      = int(C.AV_DICT_MULTIKEY)
)

// AvDictGet ...
func (d *AVDictionary) AvDictGet(key string, prev *AVDictionaryEntry, flags int) *AVDictionaryEntry {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	return (*AVDictionaryEntry)(C.av_dict_get(
		(*C.struct_AVDictionary)(d),
		Ckey,
		(*C.struct_AVDictionaryEntry)(prev),
		C.int(flags),
	))
}

// AvDictCount ...
func (d *AVDictionary) AvDictCount() int {
	return int(C.av_dict_count((*C.struct_AVDictionary)(d)))
}

// AvDictSet ...
func (d *AVDictionary) AvDictSet(key, value string, flags int) int {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	Cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(Cvalue))

	return int(C.av_dict_set(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		Ckey,
		Cvalue,
		C.int(flags),
	))
}

// AvDictSetInt ...
func (d *AVDictionary) AvDictSetInt(key string, value int64, flags int) int {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	return int(C.av_dict_set_int(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		Ckey,
		C.int64_t(value),
		C.int(flags),
	))
}

// AVDictParseString ...
func (d *AVDictionary) AVDictParseString(str, keyValSep, pairsSep string, flags int) int {
	Cstr := C.CString(str)
	defer C.free(unsafe.Pointer(Cstr))

	CkeyValSep := C.CString(keyValSep)
	defer C.free(unsafe.Pointer(CkeyValSep))

	CpairsSep := C.CString(pairsSep)
	defer C.free(unsafe.Pointer(CpairsSep))

	return int(C.av_dict_parse_string(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		Cstr,
		CkeyValSep,
		CpairsSep,
		C.int(flags),
	))
}

// AvDictCopy ...
func (d *AVDictionary) AvDictCopy(src *AVDictionary, flags int) int {
	return int(C.av_dict_copy(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		(*C.struct_AVDictionary)(unsafe.Pointer(src)),
		C.int(flags),
	))
}

// AvDictFree ...
func (d *AVDictionary) AvDictFree() {
	C.av_dict_free((**C.struct_AVDictionary)(unsafe.Pointer(&d)))
}

// AvDictGetString ...
func (d *AVDictionary) AvDictGetString(keyValSep, pairsSep byte) (int, string) {
	var Cbuf *C.char

	ret := int(C.av_dict_get_string(
		(*C.struct_AVDictionary)(d),
		(**C.char)(&Cbuf),
		C.char(keyValSep),
		C.char(pairsSep),
	))

	var buf string
	if ret == 0 {
		buf = C.GoString(Cbuf)
		C.free(unsafe.Pointer(Cbuf))
	}

	return ret, buf
}

// Key ...
func (d *AVDictionaryEntry) Key() string {
	return C.GoString(d.key)
}

// Value ...
func (d *AVDictionaryEntry) Value() string {
	return C.GoString(d.value)
}
