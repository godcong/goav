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
func (s *Stream) CodecParameters() *AvCodecParameters {
	return (*AvCodecParameters)(unsafe.Pointer(s.codecpar))
}

// Codec ...
func (s *Stream) Codec() *CodecContext {
	return (*CodecContext)(unsafe.Pointer(s.codec))
}

// Metadata ...
func (s *Stream) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(s.metadata))
}

// IndexEntries ...
func (s *Stream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(s.index_entries))
}

// AttachedPic ...
func (s *Stream) AttachedPic() AVPacket {
	return *fromCPacket(&s.attached_pic)
}

// SideData ...
func (s *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(s.side_data))
}

// ProbeData ...
func (s *Stream) ProbeData() AvProbeData {
	return AvProbeData(s.probe_data)
}

// AvgFrameRate ...
func (s *Stream) AvgFrameRate() Rational {
	return newRational(s.avg_frame_rate)
}

// RFrameRate ...
func (s *Stream) RFrameRate() Rational {
	return newRational(s.r_frame_rate)
}

// SampleAspectRatio ...
func (s *Stream) SampleAspectRatio() Rational {
	return newRational(s.sample_aspect_ratio)
}

// TimeBase ...
func (s *Stream) TimeBase() Rational {
	return newRational(s.time_base)
}

// Discard ...
func (s *Stream) Discard() AVDiscard {
	return AVDiscard(s.discard)
}

// NeedParsing ...
func (s *Stream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(s.need_parsing)
}

// CodecInfoNbFrames ...
func (s *Stream) CodecInfoNbFrames() int {
	return int(s.codec_info_nb_frames)
}

// Disposition ...
func (s *Stream) Disposition() int {
	return int(s.disposition)
}

// EventFlags ...
func (s *Stream) EventFlags() int {
	return int(s.event_flags)
}

// Id ...
func (s *Stream) Id() int {
	return int(s.id)
}

// Index ...
func (s *Stream) Index() int {
	return int(s.index)
}

// InjectGlobalSideData ...
func (s *Stream) InjectGlobalSideData() int {
	return int(s.inject_global_side_data)
}

// LastIpDuration ...
func (s *Stream) LastIpDuration() int {
	return int(s.last_IP_duration)
}

// NbDecodedFrames ...
func (s *Stream) NbDecodedFrames() int {
	return int(s.nb_decoded_frames)
}

// NbIndexEntries ...
func (s *Stream) NbIndexEntries() int {
	return int(s.nb_index_entries)
}

// NbSideData ...
func (s *Stream) NbSideData() int {
	return int(s.nb_side_data)
}

// ProbePackets ...
func (s *Stream) ProbePackets() int {
	return int(s.probe_packets)
}

// PtsWrapBehavior ...
func (s *Stream) PtsWrapBehavior() int {
	return int(s.pts_wrap_behavior)
}

// RequestProbe ...
func (s *Stream) RequestProbe() int {
	return int(s.request_probe)
}

// SkipSamples ...
func (s *Stream) SkipSamples() int {
	return int(s.skip_samples)
}

// SkipToKeyframe ...
func (s *Stream) SkipToKeyframe() int {
	return int(s.skip_to_keyframe)
}

// StreamIdentifier ...
func (s *Stream) StreamIdentifier() int {
	return int(s.stream_identifier)
}

// UpdateInitialDurationsDone ...
func (s *Stream) UpdateInitialDurationsDone() int {
	return int(s.update_initial_durations_done)
}

// CurDts ...
func (s *Stream) CurDts() int64 {
	return int64(s.cur_dts)
}

// Duration ...
func (s *Stream) Duration() int64 {
	return int64(s.duration)
}

// FirstDts ...
func (s *Stream) FirstDts() int64 {
	return int64(s.first_dts)
}

// InterleaverChunkDuration ...
func (s *Stream) InterleaverChunkDuration() int64 {
	return int64(s.interleaver_chunk_duration)
}

// InterleaverChunkSize ...
func (s *Stream) InterleaverChunkSize() int64 {
	return int64(s.interleaver_chunk_size)
}

// LastDtsForOrderCheck ...
func (s *Stream) LastDtsForOrderCheck() int64 {
	return int64(s.last_dts_for_order_check)
}

// LastIPPts ...
func (s *Stream) LastIPPts() int64 {
	return int64(s.last_IP_pts)
}

// MuxTsOffset ...
func (s *Stream) MuxTsOffset() int64 {
	return int64(s.mux_ts_offset)
}

// NbFrames ...
func (s *Stream) NbFrames() int64 {
	return int64(s.nb_frames)
}

// PtsBuffer ...
func (s *Stream) PtsBuffer() int64 {
	return int64(s.pts_buffer[0])
}

// PtsReorderError ...
func (s *Stream) PtsReorderError() int64 {
	return int64(s.pts_reorder_error[0])
}

// PtsWrapReference ...
func (s *Stream) PtsWrapReference() int64 {
	return int64(s.pts_wrap_reference)
}

// StartTime ...
func (s *Stream) StartTime() int64 {
	return int64(s.start_time)
}

// AVCodecParser ...
func (s *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(s.parser))
}

// LastInPacketBuffer ...
func (s *Stream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(s.last_in_packet_buffer))
}

// DtsMisordered ...
func (s *Stream) DtsMisordered() uint8 {
	return uint8(s.dts_misordered)
}

// DtsOrdered ...
func (s *Stream) DtsOrdered() uint8 {
	return uint8(s.dts_ordered)
}

// PtsReorderErrorCount ...
func (s *Stream) PtsReorderErrorCount() uint8 {
	return uint8(s.pts_reorder_error_count[0])
}

// IndexEntriesAllocatedSize ...
func (s *Stream) IndexEntriesAllocatedSize() uint {
	return uint(s.index_entries_allocated_size)
}

// Free ...
func (s *Stream) Free() {
	C.av_freep(unsafe.Pointer(s))
}
