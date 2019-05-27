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
func (s *AVStream) CodecParameters() *AvCodecParameters {
	return (*AvCodecParameters)(unsafe.Pointer(s.codecpar))
}

// Codec ...
func (s *AVStream) Codec() *AVCodecContext {
	return (*AVCodecContext)(unsafe.Pointer(s.codec))
}

// Metadata ...
func (s *AVStream) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(s.metadata))
}

// IndexEntries ...
func (s *AVStream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(s.index_entries))
}

// AttachedPic ...
func (s *AVStream) AttachedPic() AVPacket {
	return *fromCPacket(&s.attached_pic)
}

// SideData ...
func (s *AVStream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(s.side_data))
}

// ProbeData ...
func (s *AVStream) ProbeData() AVProbeData {
	return AVProbeData(s.probe_data)
}

// AvgFrameRate ...
func (s *AVStream) AvgFrameRate() AVRational {
	return newRational(s.avg_frame_rate)
}

// RFrameRate ...
func (s *AVStream) RFrameRate() AVRational {
	return newRational(s.r_frame_rate)
}

// SampleAspectRatio ...
func (s *AVStream) SampleAspectRatio() AVRational {
	return newRational(s.sample_aspect_ratio)
}

// TimeBase ...
func (s *AVStream) TimeBase() AVRational {
	return newRational(s.time_base)
}

// Discard ...
func (s *AVStream) Discard() AVDiscard {
	return AVDiscard(s.discard)
}

// NeedParsing ...
func (s *AVStream) NeedParsing() AVStreamParseType {
	return AVStreamParseType(s.need_parsing)
}

// CodecInfoNbFrames ...
func (s *AVStream) CodecInfoNbFrames() int {
	return int(s.codec_info_nb_frames)
}

// Disposition ...
func (s *AVStream) Disposition() int {
	return int(s.disposition)
}

// EventFlags ...
func (s *AVStream) EventFlags() int {
	return int(s.event_flags)
}

// ID ...
func (s *AVStream) ID() int {
	return int(s.id)
}

// Index ...
func (s *AVStream) Index() int {
	return int(s.index)
}

// InjectGlobalSideData ...
func (s *AVStream) InjectGlobalSideData() int {
	return int(s.inject_global_side_data)
}

// LastIPDuration ...
func (s *AVStream) LastIPDuration() int {
	return int(s.last_IP_duration)
}

// NbDecodedFrames ...
func (s *AVStream) NbDecodedFrames() int {
	return int(s.nb_decoded_frames)
}

// NbIndexEntries ...
func (s *AVStream) NbIndexEntries() int {
	return int(s.nb_index_entries)
}

// NbSideData ...
func (s *AVStream) NbSideData() int {
	return int(s.nb_side_data)
}

// ProbePackets ...
func (s *AVStream) ProbePackets() int {
	return int(s.probe_packets)
}

// PtsWrapBehavior ...
func (s *AVStream) PtsWrapBehavior() int {
	return int(s.pts_wrap_behavior)
}

// RequestProbe ...
func (s *AVStream) RequestProbe() int {
	return int(s.request_probe)
}

// SkipSamples ...
func (s *AVStream) SkipSamples() int {
	return int(s.skip_samples)
}

// SkipToKeyframe ...
func (s *AVStream) SkipToKeyframe() int {
	return int(s.skip_to_keyframe)
}

// StreamIdentifier ...
func (s *AVStream) StreamIdentifier() int {
	return int(s.stream_identifier)
}

// UpdateInitialDurationsDone ...
func (s *AVStream) UpdateInitialDurationsDone() int {
	return int(s.update_initial_durations_done)
}

// CurDts ...
func (s *AVStream) CurDts() int64 {
	return int64(s.cur_dts)
}

// Duration ...
func (s *AVStream) Duration() int64 {
	return int64(s.duration)
}

// FirstDts ...
func (s *AVStream) FirstDts() int64 {
	return int64(s.first_dts)
}

// InterleaverChunkDuration ...
func (s *AVStream) InterleaverChunkDuration() int64 {
	return int64(s.interleaver_chunk_duration)
}

// InterleaverChunkSize ...
func (s *AVStream) InterleaverChunkSize() int64 {
	return int64(s.interleaver_chunk_size)
}

// LastDtsForOrderCheck ...
func (s *AVStream) LastDtsForOrderCheck() int64 {
	return int64(s.last_dts_for_order_check)
}

// LastIPPts ...
func (s *AVStream) LastIPPts() int64 {
	return int64(s.last_IP_pts)
}

// MuxTsOffset ...
func (s *AVStream) MuxTsOffset() int64 {
	return int64(s.mux_ts_offset)
}

// NbFrames ...
func (s *AVStream) NbFrames() int64 {
	return int64(s.nb_frames)
}

// PtsBuffer ...
func (s *AVStream) PtsBuffer() int64 {
	return int64(s.pts_buffer[0])
}

// PtsReorderError ...
func (s *AVStream) PtsReorderError() int64 {
	return int64(s.pts_reorder_error[0])
}

// PtsWrapReference ...
func (s *AVStream) PtsWrapReference() int64 {
	return int64(s.pts_wrap_reference)
}

// StartTime ...
func (s *AVStream) StartTime() int64 {
	return int64(s.start_time)
}

// AVCodecParser ...
func (s *AVStream) Parser() *AVCodecParserContext {
	return (*AVCodecParserContext)(unsafe.Pointer(s.parser))
}

// LastInPacketBuffer ...
func (s *AVStream) LastInPacketBuffer() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(s.last_in_packet_buffer))
}

// DtsMisordered ...
func (s *AVStream) DtsMisordered() uint8 {
	return uint8(s.dts_misordered)
}

// DtsOrdered ...
func (s *AVStream) DtsOrdered() uint8 {
	return uint8(s.dts_ordered)
}

// PtsReorderErrorCount ...
func (s *AVStream) PtsReorderErrorCount() uint8 {
	return uint8(s.pts_reorder_error_count[0])
}

// IndexEntriesAllocatedSize ...
func (s *AVStream) IndexEntriesAllocatedSize() uint {
	return uint(s.index_entries_allocated_size)
}

// Free ...
func (s *AVStream) Free() {
	C.av_freep(unsafe.Pointer(s))
}
