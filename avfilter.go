// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package goav contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package goav

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"
)

// Filter ...
type Filter C.struct_AVFilter

// FilterContext ...
type FilterContext C.struct_AVFilterContext

// FilterLink ...
type FilterLink C.struct_AVFilterLink

// FilterGraph ...
type FilterGraph C.struct_AVFilterGraph

// FilterInOut ...
type FilterInOut C.struct_AVFilterInOut

// FilterPad ...
type FilterPad C.struct_AVFilterPad

//AVFilterVersion AVFilterVersion Return the LIBAvFILTER_VERSION_INT constant.
func AVFilterVersion() uint {
	return uint(C.avfilter_version())
}

//AVFilterConfiguration Return the libavfilter build-time configuration.
func AVFilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

//AVFilterLicense Return the libavfilter license.
func AVFilterLicense() string {
	return C.GoString(C.avfilter_license())
}

//AVFilterPadCount Get the number of elements in a NULL-terminated array of Pads (e.g.
func AVFilterPadCount(p *FilterPad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

//AVFilterPadGetName Get the name of an FilterPad.
func AVFilterPadGetName(p *FilterPad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//AVFilterPadGetType Get the type of an FilterPad.
func AVFilterPadGetType(p *FilterPad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//AVFilterLink FilterLink two filters together.
func AVFilterLink(s *FilterContext, sp uint, d *FilterContext, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

//AVFilterLinkFree Free the link in *link, and set its pointer to NULL.
func AVFilterLinkFree(l **FilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

//AVFilterLinkGetChannels Get the number of channels of a link.
func AVFilterLinkGetChannels(l *FilterLink) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

//AVFilterLinkSetClosed Set the closed field of a link.
func AVFilterLinkSetClosed(l *FilterLink, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

//AVFilterConfigLinks Negotiate the media format, dimensions, etc of all inputs to a filter.
func AVFilterConfigLinks(f *FilterContext) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

//AVFilterProcessCommand Make the filter instance process a command.
func AVFilterProcessCommand(f *FilterContext, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

//AVFilterRegisterAll Initialize the filter system.
func AVFilterRegisterAll() {
	C.avfilter_register_all()
}

//AVFilterInitStr Initialize a filter with the supplied parameters.
func (c *FilterContext) AVFilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(c), C.CString(args)))
}

//AVFilterInitDict Initialize a filter with the supplied dictionary of options.
func (c *FilterContext) AVFilterInitDict(o **AVDictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(c), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//AVFilterFree Free a filter context.
func (c *FilterContext) AVFilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(c))
}

//AVFilterInsertFilter Insert a filter in the middle of an existing link.
func AVFilterInsertFilter(l *FilterLink, f *FilterContext, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

// AVFilterGetClass ...
func AVFilterGetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

//AVFilterInOutAlloc Allocate a single FilterInOut entry.
func AVFilterInOutAlloc() *FilterInOut {
	return (*FilterInOut)(C.avfilter_inout_alloc())
}

//AVFilterInOutFree Free the supplied list of FilterInOut and set *inout to NULL.
func AVFilterInOutFree(i *FilterInOut) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}
