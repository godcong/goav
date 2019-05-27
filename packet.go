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

//AVInitPacket Initialize optional fields of a packet with default values.
func (p *Packet) AVInitPacket() {
	C.av_init_packet((*C.struct_AVPacket)(p))
	p.size = 0
	p.data = nil
}

//AVNewPacket Allocate the payload of a packet and initialize its fields with default values.
func (p *Packet) AVNewPacket(s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//AVShrinkPacket Reduce packet size, correctly zeroing padding.
func (p *Packet) AVShrinkPacket(s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

//AVGrowPacket Increase packet size, correctly zeroing padding.
func (p *Packet) AVGrowPacket(s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//AVPacketFromData Initialize a reference-counted packet from av_malloc()ed data.
func (p *Packet) AVPacketFromData(d *uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(d), C.int(s)))

}

// AVDupPacket ...
func (p *Packet) AVDupPacket() int {
	return int(C.av_dup_packet((*C.struct_AVPacket)(p)))

}

//AVCopyPacket Copy packet, including contents.
func (p *Packet) AVCopyPacket(r *Packet) int {
	return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//AVCopyPacketSideData Copy packet side data.
func (p *Packet) AVCopyPacketSideData(r *Packet) int {
	return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//AVFreePacket Free a packet.
func (p *Packet) AVFreePacket() {
	C.av_free_packet((*C.struct_AVPacket)(p))

}

//AVPacketNewSideData Allocate new information of a packet.
func (p *Packet) AVPacketNewSideData(t PacketSideDataType, s int) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//AVPacketShrinkSideData Shrink the already allocated side data buffer.
func (p *Packet) AVPacketShrinkSideData(t PacketSideDataType, s int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//AVPacketGetSideData Get side information from packet.
func (p *Packet) AVPacketGetSideData(t PacketSideDataType, s *int) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(s))))
}

//AVPacketMergeSideData int 	av_packet_merge_side_data (Packet *pkt)
func (p *Packet) AVPacketMergeSideData() int {
	return int(C.av_packet_merge_side_data((*C.struct_AVPacket)(p)))
}

//AVPacketSplitSideData int 	av_packet_split_side_data (Packet *pkt)
func (p *Packet) AVPacketSplitSideData() int {
	return int(C.av_packet_split_side_data((*C.struct_AVPacket)(p)))
}

//AVPacketFreeSideData Convenience function to free all the side data stored.
func (p *Packet) AVPacketFreeSideData() {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

//AVPacketRef Setup a new reference to the data described by a given packet.
func (p *Packet) AVPacketRef(s *Packet) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//AVPacketUnref Wipe the packet.
func (p *Packet) AVPacketUnref() {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

//AVPacketMoveRef Move every field in src to dst and reset src.
func (p *Packet) AVPacketMoveRef(s *Packet) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s))
}

//AVPacketCopyProps only "properties" fields from src to dst.
func (p *Packet) AVPacketCopyProps(s *Packet) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//AVPacketRescaleTs Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
func (p *Packet) AVPacketRescaleTs(r, r2 Rational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), (C.struct_AVRational)(r), (C.struct_AVRational)(r2))
}

func toCPacket(pkt *Packet) *C.struct_AVPacket {
	return (*C.struct_AVPacket)(unsafe.Pointer(pkt))
}

func fromCPacket(pkt *C.struct_AVPacket) *Packet {
	return (*Packet)(unsafe.Pointer(pkt))
}
