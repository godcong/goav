// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"
)

//AVFilterGraphAlloc Allocate a filter graph.
func AVFilterGraphAlloc() *FilterGraph {
	return (*FilterGraph)(C.avfilter_graph_alloc())
}

//AVFilterGraphAllocFilter Create a new filter instance in a filter graph.
func (g *FilterGraph) AVFilterGraphAllocFilter(f *Filter, n string) *FilterContext {
	return (*FilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), C.CString(n)))
}

//AVFilterGraphGetFilter Get a filter instance identified by instance name from graph.
func (g *FilterGraph) AVFilterGraphGetFilter(n string) *FilterContext {
	return (*FilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), C.CString(n)))
}

//AVFilterGraphSetAutoConvert Enable or disable automatic format conversion inside the graph.
func (g *FilterGraph) AVFilterGraphSetAutoConvert(f uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(f))
}

//AVFilterGraphConfig Check validity and configure all the links and formats in the graph.
func (g *FilterGraph) AVFilterGraphConfig(l int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(&l)))
}

//AVFilterGraphFree Free a graph, destroy its links, and set *graph to NULL.
func (g *FilterGraph) AVFilterGraphFree() {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(g)))
}

//AVFilterGraphParse Add a graph described by a string to a graph.
func (g *FilterGraph) AVFilterGraphParse(f string, i, o *FilterInOut, l int) int {
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), C.CString(f), (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l)))
}

//AVFilterGraphParsePtr Add a graph described by a string to a graph.
func (g *FilterGraph) AVFilterGraphParsePtr(f string, i, o **FilterInOut, l int) int {
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(&l)))
}

//AVFilterGraphParse2 Add a graph described by a string to a graph.
func (g *FilterGraph) AVFilterGraphParse2(f string, i, o **FilterInOut) int {
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o))))
}

//AVFilterGraphSendCommand Send a command to one or more filter instances.
func (g *FilterGraph) AVFilterGraphSendCommand(t, cmd, arg, res string, resl, f int) int {
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.CString(res), C.int(resl), C.int(f)))
}

//AVFilterGraphQueueCommand Queue a command for one or more filter instances.
func (g *FilterGraph) AVFilterGraphQueueCommand(t, cmd, arg string, f int, ts C.double) int {
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.int(f), ts))
}

//AVFilterGraphDump Dump a graph into a human-readable string representation.
func (g *FilterGraph) AVFilterGraphDump(o string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), C.CString(o)))
}

//AVFilterGraphRequestOldest Request a frame on the oldest sink
func (g *FilterGraph) AVFilterGraphRequestOldest() int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}

//AVFilterGraphCreateFilter Create and add a filter instance into an existing graph.
func AVFilterGraphCreateFilter(cx **FilterContext, f *Filter, n, a string, o int, g *FilterGraph) int {
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), C.CString(n), C.CString(a), unsafe.Pointer(&o), (*C.struct_AVFilterGraph)(g)))
}
