// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"
)

// CodecParameters ...
func (avs *Stream) CodecParameters() *AvCodecParameters {
	return (*AvCodecParameters)(unsafe.Pointer(avs.codecpar))
}

// Codec ...
func (avs *Stream) Codec() *CodecContext {
	return (*CodecContext)(unsafe.Pointer(avs.codec))
}

// Metadata ...
func (avs *Stream) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(avs.metadata))
}

// IndexEntries ...
func (avs *Stream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
}

// AttachedPic ...
func (avs *Stream) AttachedPic() Packet {
	return *fromCPacket(&avs.attached_pic)
}

// SideData ...
func (avs *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

// ProbeData ...
func (avs *Stream) ProbeData() AvProbeData {
	return AvProbeData(avs.probe_data)
}

// AvgFrameRate ...
func (avs *Stream) AvgFrameRate() Rational {
	return newRational(avs.avg_frame_rate)
}

// RFrameRate ...
func (avs *Stream) RFrameRate() Rational {
	return newRational(avs.r_frame_rate)
}

// SampleAspectRatio ...
func (avs *Stream) SampleAspectRatio() Rational {
	return newRational(avs.sample_aspect_ratio)
}

// TimeBase ...
func (avs *Stream) TimeBase() Rational {
	return newRational(avs.time_base)
}

// Discard ...
func (avs *Stream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

// NeedParsing ...
func (avs *Stream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(avs.need_parsing)
}

// CodecInfoNbFrames ...
func (avs *Stream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

// Disposition ...
func (avs *Stream) Disposition() int {
	return int(avs.disposition)
}

// EventFlags ...
func (avs *Stream) EventFlags() int {
	return int(avs.event_flags)
}

// Id ...
func (avs *Stream) Id() int {
	return int(avs.id)
}

// Index ...
func (avs *Stream) Index() int {
	return int(avs.index)
}

// InjectGlobalSideData ...
func (avs *Stream) InjectGlobalSideData() int {
	return int(avs.inject_global_side_data)
}

// LastIpDuration ...
func (avs *Stream) LastIpDuration() int {
	return int(avs.last_IP_duration)
}

// NbDecodedFrames ...
func (avs *Stream) NbDecodedFrames() int {
	return int(avs.nb_decoded_frames)
}

// NbIndexEntries ...
func (avs *Stream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

// NbSideData ...
func (avs *Stream) NbSideData() int {
	return int(avs.nb_side_data)
}

// ProbePackets ...
func (avs *Stream) ProbePackets() int {
	return int(avs.probe_packets)
}

// PtsWrapBehavior ...
func (avs *Stream) PtsWrapBehavior() int {
	return int(avs.pts_wrap_behavior)
}

// RequestProbe ...
func (avs *Stream) RequestProbe() int {
	return int(avs.request_probe)
}

// SkipSamples ...
func (avs *Stream) SkipSamples() int {
	return int(avs.skip_samples)
}

// SkipToKeyframe ...
func (avs *Stream) SkipToKeyframe() int {
	return int(avs.skip_to_keyframe)
}

// StreamIdentifier ...
func (avs *Stream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

// UpdateInitialDurationsDone ...
func (avs *Stream) UpdateInitialDurationsDone() int {
	return int(avs.update_initial_durations_done)
}

// CurDts ...
func (avs *Stream) CurDts() int64 {
	return int64(avs.cur_dts)
}

// Duration ...
func (avs *Stream) Duration() int64 {
	return int64(avs.duration)
}

// FirstDts ...
func (avs *Stream) FirstDts() int64 {
	return int64(avs.first_dts)
}

// InterleaverChunkDuration ...
func (avs *Stream) InterleaverChunkDuration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

// InterleaverChunkSize ...
func (avs *Stream) InterleaverChunkSize() int64 {
	return int64(avs.interleaver_chunk_size)
}

// LastDtsForOrderCheck ...
func (avs *Stream) LastDtsForOrderCheck() int64 {
	return int64(avs.last_dts_for_order_check)
}

// LastIPPts ...
func (avs *Stream) LastIPPts() int64 {
	return int64(avs.last_IP_pts)
}

// MuxTsOffset ...
func (avs *Stream) MuxTsOffset() int64 {
	return int64(avs.mux_ts_offset)
}

// NbFrames ...
func (avs *Stream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

// PtsBuffer ...
func (avs *Stream) PtsBuffer() int64 {
	return int64(avs.pts_buffer[0])
}

// PtsReorderError ...
func (avs *Stream) PtsReorderError() int64 {
	return int64(avs.pts_reorder_error[0])
}

// PtsWrapReference ...
func (avs *Stream) PtsWrapReference() int64 {
	return int64(avs.pts_wrap_reference)
}

// StartTime ...
func (avs *Stream) StartTime() int64 {
	return int64(avs.start_time)
}

// Parser ...
func (avs *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(avs.parser))
}

// LastInPacketBuffer ...
func (avs *Stream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

// DtsMisordered ...
func (avs *Stream) DtsMisordered() uint8 {
	return uint8(avs.dts_misordered)
}

// DtsOrdered ...
func (avs *Stream) DtsOrdered() uint8 {
	return uint8(avs.dts_ordered)
}

// PtsReorderErrorCount ...
func (avs *Stream) PtsReorderErrorCount() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

// IndexEntriesAllocatedSize ...
func (avs *Stream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}

// Free ...
func (avs *Stream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}
