// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"

//Get a filter definition matching the given name.
func AVFilterGetByName(n string) *Filter {
	return (*Filter)(C.avfilter_get_by_name(C.CString(n)))
}

//Register a filter.
func (f *Filter) AVFilterRegister() int {
	return int(C.avfilter_register((*C.struct_AVFilter)(f)))
}

//Iterate over all registered filters.
func (f *Filter) AVFilterNext() *Filter {
	return (*Filter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}
