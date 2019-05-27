// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// Buf ...
func (p *Packet) Buf() *AVBufferRef {
	return (*AVBufferRef)(p.buf)
}

// Duration ...
func (p *Packet) Duration() int {
	return int(p.duration)
}

// Flags ...
func (p *Packet) Flags() int {
	return int(p.flags)
}

// SetFlags ...
func (p *Packet) SetFlags(flags int) {
	p.flags = C.int(flags)
}

// SideDataElems ...
func (p *Packet) SideDataElems() int {
	return int(p.side_data_elems)
}

// Size ...
func (p *Packet) Size() int {
	return int(p.size)
}

// StreamIndex ...
func (p *Packet) StreamIndex() int {
	return int(p.stream_index)
}

// SetStreamIndex ...
func (p *Packet) SetStreamIndex(idx int) {
	p.stream_index = C.int(idx)
}

// ConvergenceDuration ...
func (p *Packet) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}

// Dts ...
func (p *Packet) Dts() int64 {
	return int64(p.dts)
}

// SetDts ...
func (p *Packet) SetDts(dts int64) {
	p.dts = C.int64_t(dts)
}

// Pos ...
func (p *Packet) Pos() int64 {
	return int64(p.pos)
}

// Pts ...
func (p *Packet) Pts() int64 {
	return int64(p.pts)
}

// SetPts ...
func (p *Packet) SetPts(pts int64) {
	p.dts = C.int64_t(pts)
}

// Data ...
func (p *Packet) Data() *uint8 {
	return (*uint8)(p.data)
}
