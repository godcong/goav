// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// Buf ...
func (p *AVPacket) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}

// Duration ...
func (p *AVPacket) Duration() int {
	return int(p.duration)
}

// Flags ...
func (p *AVPacket) Flags() int {
	return int(p.flags)
}

// SetFlags ...
func (p *AVPacket) SetFlags(flags int) {
	p.flags = C.int(flags)
}

// SideDataElems ...
func (p *AVPacket) SideDataElems() int {
	return int(p.side_data_elems)
}

// Size ...
func (p *AVPacket) Size() int {
	return int(p.size)
}

// StreamIndex ...
func (p *AVPacket) StreamIndex() int {
	return int(p.stream_index)
}

// SetStreamIndex ...
func (p *AVPacket) SetStreamIndex(idx int) {
	p.stream_index = C.int(idx)
}

// ConvergenceDuration ...
func (p *AVPacket) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}

// Dts ...
func (p *AVPacket) Dts() int64 {
	return int64(p.dts)
}

// SetDts ...
func (p *AVPacket) SetDts(dts int64) {
	p.dts = C.int64_t(dts)
}

// Pos ...
func (p *AVPacket) Pos() int64 {
	return int64(p.pos)
}

// Pts ...
func (p *AVPacket) Pts() int64 {
	return int64(p.pts)
}

// SetPts ...
func (p *AVPacket) SetPts(pts int64) {
	p.dts = C.int64_t(pts)
}

// Data ...
func (p *AVPacket) Data() *uint8 {
	return (*uint8)(p.data)
}
