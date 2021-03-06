// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavutil
	#include <libavutil/avutil.h>
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

//AVMalloc Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU).
func AVMalloc(s uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.av_malloc(C.size_t(s)))
}

// AvMallocArray ...
func AvMallocArray(n, s uintptr) unsafe.Pointer {
	return C.av_malloc_array(C.size_t(n), C.size_t(s))
}

//AVRealloc Allocate or reallocate a block of memory.
func AVRealloc(p *int, s uintptr) unsafe.Pointer {
	return C.av_realloc(unsafe.Pointer(&p), C.size_t(s))
}

//AVReallocF Allocate or reallocate a block of memory.
func AVReallocF(p int, n, e uintptr) unsafe.Pointer {
	return C.av_realloc_f(unsafe.Pointer(&p), C.size_t(n), C.size_t(e))
}

//AVReallocp Allocate or reallocate a block of memory.
func AVReallocp(p int, s uintptr) int {
	return int(C.av_reallocp(unsafe.Pointer(&p), C.size_t(s)))
}

// AVReallocArray ...
func AVReallocArray(p int, n, s uintptr) unsafe.Pointer {
	return C.av_realloc_array(unsafe.Pointer(&p), C.size_t(n), C.size_t(s))
}

// AVReallocpArray ...
func AVReallocpArray(p int, n, s uintptr) int {
	return int(C.av_reallocp_array(unsafe.Pointer(&p), C.size_t(n), C.size_t(s)))
}

//AVFree Free a memory block which has been allocated with av_malloc(z)() or av_realloc().
func AVFree(p unsafe.Pointer) {
	C.av_free(p)
}

//AVMallocz Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func AVMallocz(s uintptr) unsafe.Pointer {
	return C.av_mallocz(C.size_t(s))
}

//AVCalloc Allocate a block of nmemb * size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func AVCalloc(n, s uintptr) unsafe.Pointer {
	return C.av_calloc(C.size_t(n), C.size_t(s))
}

// AvMalloczArray ...
func AvMalloczArray(n, s uintptr) unsafe.Pointer {
	return C.av_mallocz_array(C.size_t(n), C.size_t(s))
}

//AVStrdup Duplicate the string s.
func AVStrdup(s string) string {
	return C.GoString(C.av_strdup(C.CString(s)))
}

//AVStrndup char * 	av_strndup (const char *s, size_t len) av_malloc_attrib
//Duplicate a substring of the string s.
func AVStrndup(s string, l uintptr) string {
	return C.GoString(C.av_strndup(C.CString(s), C.size_t(l)))
}

//AVMemdup Duplicate the buffer p.
func AVMemdup(p *int, s uintptr) unsafe.Pointer {
	return C.av_memdup(unsafe.Pointer(p), C.size_t(s))
}

//AVFreep Free a memory block which has been allocated with av_malloc(z)() or av_realloc() and set the pointer pointing to it to NULL.
func AVFreep(p unsafe.Pointer) {
	C.av_freep(p)
}

//AVDynArrayAdd Add an element to a dynamic array.
func AVDynArrayAdd(t unsafe.Pointer, n *int, e unsafe.Pointer) {
	C.av_dynarray_add(t, (*C.int)(unsafe.Pointer(n)), e)
}

//AVDynArrayAddNofree Add an element to a dynamic array.
func AVDynArrayAddNofree(p unsafe.Pointer, n *int, e unsafe.Pointer) int {
	return int(C.av_dynarray_add_nofree(p, (*C.int)(unsafe.Pointer(&n)), e))
}

//AVDynArray2Add Add an element of size elem_size to a dynamic array.
func AVDynArray2Add(t *unsafe.Pointer, n *int, e uintptr, d uint8) unsafe.Pointer {
	return C.av_dynarray2_add(t, (*C.int)(unsafe.Pointer(&n)), (C.size_t)(e), (*C.uint8_t)(unsafe.Pointer(&d)))
}

//AVSizeMult Multiply two size_t values checking for overflow.
func AVSizeMult(a, b uintptr, r *uintptr) int {
	return int(C.av_size_mult((C.size_t)(a), (C.size_t)(b), (*C.size_t)(unsafe.Pointer(r))))
}

//AVMaxAlloc Set the maximum size that may me allocated in one block.
func AVMaxAlloc(m uintptr) {
	C.av_max_alloc(C.size_t(m))
}

//AVMemcpyBackptr deliberately overlapping memcpy implementation
func AVMemcpyBackptr(d *uintptr, b, c int) {
	C.av_memcpy_backptr((*C.uint8_t)(unsafe.Pointer(d)), C.int(b), C.int(c))
}

//AVFastRealloc Reallocate the given block if it is not large enough, otherwise do nothing.
func AVFastRealloc(p unsafe.Pointer, s *uint, m uintptr) unsafe.Pointer {
	return C.av_fast_realloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(m))
}

//AVFastMalloc Allocate a buffer, reusing the given one if large enough.
func AVFastMalloc(p unsafe.Pointer, s *uint, m uintptr) {
	C.av_fast_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(m))
}
