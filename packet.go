// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavcodec libavformat
//#include <libavcodec/avcodec.h>
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"
)

// AvPktFlagKey ...
const (
	AVPktFlagKey     = int(C.AV_PKT_FLAG_KEY)
	AVPktFlagCorrupt = int(C.AV_PKT_FLAG_CORRUPT)
	AVPktFlagDiscard = int(C.AV_PKT_FLAG_DISCARD)
)

//Initialize optional fields of a packet with default values.
func (p *AVPacket) AVInitPacket() {
	C.av_init_packet((*C.struct_AVPacket)(p))
	p.size = 0
	p.data = nil
}

//Allocate the payload of a packet and initialize its fields with default values.
func (p *AVPacket) AvNewPacket(s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Reduce packet size, correctly zeroing padding.
func (p *AVPacket) AvShrinkPacket(s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

//Increase packet size, correctly zeroing padding.
func (p *AVPacket) AvGrowPacket(s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Initialize a reference-counted packet from av_malloc()ed data.
func (p *AVPacket) AVPacketFromData(d *uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(d), C.int(s)))

}

// AvDupPacket ...
func (p *AVPacket) AvDupPacket() int {
	return int(C.av_dup_packet((*C.struct_AVPacket)(p)))

}

//Copy packet, including contents.
func (p *AVPacket) AvCopyPacket(r *AVPacket) int {
	return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Copy packet side data.
func (p *AVPacket) AvCopyPacketSideData(r *AVPacket) int {
	return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Free a packet.
func (p *AVPacket) AvFreePacket() {
	C.av_free_packet((*C.struct_AVPacket)(p))

}

//AVPacketNewSideData Allocate new information of a packet.
func (p *AVPacket) AVPacketNewSideData(t AVPacketSideDataType, s int) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//AVPacketShrinkSideData Shrink the already allocated side data buffer.
func (p *AVPacket) AVPacketShrinkSideData(t AVPacketSideDataType, s int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//AVPacketGetSideData Get side information from packet.
func (p *AVPacket) AVPacketGetSideData(t AVPacketSideDataType, s *int) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(s))))
}

//AVPacketMergeSideData int 	av_packet_merge_side_data (AVPacket *pkt)
func (p *AVPacket) AVPacketMergeSideData() int {
	return int(C.av_packet_merge_side_data((*C.struct_AVPacket)(p)))
}

//AVPacketSplitSideData int 	av_packet_split_side_data (AVPacket *pkt)
func (p *AVPacket) AVPacketSplitSideData() int {
	return int(C.av_packet_split_side_data((*C.struct_AVPacket)(p)))
}

//AVPacketFreeSideData Convenience function to free all the side data stored.
func (p *AVPacket) AVPacketFreeSideData() {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

//AVPacketRef Setup a new reference to the data described by a given packet.
func (p *AVPacket) AVPacketRef(s *AVPacket) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//AVPacketUnref Wipe the packet.
func (p *AVPacket) AVPacketUnref() {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

//AVPacketMoveRef Move every field in src to dst and reset src.
func (p *AVPacket) AVPacketMoveRef(s *AVPacket) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s))
}

//AVPacketCopyProps only "properties" fields from src to dst.
func (p *AVPacket) AVPacketCopyProps(s *AVPacket) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//AVPacketRescaleTs Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
func (p *AVPacket) AVPacketRescaleTs(r, r2 AVRational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), (C.struct_AVRational)(r), (C.struct_AVRational)(r2))
}

func toCPacket(pkt *AVPacket) *C.struct_AVPacket {
	return (*C.struct_AVPacket)(unsafe.Pointer(pkt))
}

func fromCPacket(pkt *C.struct_AVPacket) *AVPacket {
	return (*AVPacket)(unsafe.Pointer(pkt))
}
