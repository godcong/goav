// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"reflect"
	"unsafe"
)

func (c *Context) ActiveThreadType() int {
	return int(c.active_thread_type)
}

func (c *Context) BFrameStrategy() int {
	return int(c.b_frame_strategy)
}

func (c *Context) BQuantFactor() float64 {
	return float64(c.b_quant_factor)
}

func (c *Context) BQuantOffset() float64 {
	return float64(c.b_quant_offset)
}

func (c *Context) BSensitivity() int {
	return int(c.b_sensitivity)
}

func (c *Context) BidirRefine() int {
	return int(c.bidir_refine)
}

func (c *Context) BitRate() int {
	return int(c.bit_rate)
}

func (c *Context) BitRateTolerance() int {
	return int(c.bit_rate_tolerance)
}

func (c *Context) BitsPerCodedSample() int {
	return int(c.bits_per_coded_sample)
}

func (c *Context) BitsPerRawSample() int {
	return int(c.bits_per_raw_sample)
}

func (c *Context) BlockAlign() int {
	return int(c.block_align)
}

func (c *Context) BrdScale() int {
	return int(c.brd_scale)
}

func (c *Context) Channels() int {
	return int(c.channels)
}

func (c *Context) Chromaoffset() int {
	return int(c.chromaoffset)
}

func (c *Context) CodedHeight() int {
	return int(c.coded_height)
}

func (c *Context) CodedWidth() int {
	return int(c.coded_width)
}

func (c *Context) CoderType() int {
	return int(c.coder_type)
}

func (c *Context) CompressionLevel() int {
	return int(c.compression_level)
}

func (c *Context) ContextModel() int {
	return int(c.context_model)
}

func (c *Context) Cutoff() int {
	return int(c.cutoff)
}

func (c *Context) DarkMasking() float64 {
	return float64(c.dark_masking)
}

func (c *Context) DctAlgo() int {
	return int(c.dct_algo)
}

//func (ctxt *Context) Debug() int {
//	return int(ctxt.debug)
//}

func (c *Context) DebugMv() int {
	return int(c.debug_mv)
}

func (c *Context) Delay() int {
	return int(c.delay)
}

func (c *Context) DiaSize() int {
	return int(c.dia_size)
}

func (c *Context) ErrRecognition() int {
	return int(c.err_recognition)
}

func (c *Context) ErrorConcealment() int {
	return int(c.error_concealment)
}

func (c *Context) ExtradataSize() int {
	return int(c.extradata_size)
}

//func (ctxt *Context) Flags() int {
//	return int(ctxt.flags)
//}

func (c *Context) Flags2() int {
	return int(c.flags2)
}

func (c *Context) FrameBits() int {
	return int(c.frame_bits)
}

func (c *Context) FrameNumber() int {
	return int(c.frame_number)
}

func (c *Context) FrameSize() int {
	return int(c.frame_size)
}

func (c *Context) FrameSkipCmp() int {
	return int(c.frame_skip_cmp)
}

func (c *Context) FrameSkipExp() int {
	return int(c.frame_skip_exp)
}

func (c *Context) FrameSkipFactor() int {
	return int(c.frame_skip_factor)
}

func (c *Context) FrameSkipThreshold() int {
	return int(c.frame_skip_threshold)
}

func (c *Context) GlobalQuality() int {
	return int(c.global_quality)
}

func (c *Context) GopSize() int {
	return int(c.gop_size)
}

func (c *Context) HasBFrames() int {
	return int(c.has_b_frames)
}

func (c *Context) HeaderBits() int {
	return int(c.header_bits)
}

func (c *Context) Height() int {
	return int(c.height)
}

func (c *Context) ICount() int {
	return int(c.i_count)
}

func (c *Context) IQuantFactor() float64 {
	return float64(c.i_quant_factor)
}

func (c *Context) IQuantOffset() float64 {
	return float64(c.i_quant_offset)
}

func (c *Context) ITexBits() int {
	return int(c.i_tex_bits)
}

func (c *Context) IdctAlgo() int {
	return int(c.idct_algo)
}

func (c *Context) IldctCmp() int {
	return int(c.ildct_cmp)
}

func (c *Context) IntraDcPrecision() int {
	return int(c.intra_dc_precision)
}

func (c *Context) KeyintMin() int {
	return int(c.keyint_min)
}

func (c *Context) LastPredictorCount() int {
	return int(c.last_predictor_count)
}

func (c *Context) Level() int {
	return int(c.level)
}

func (c *Context) LogLevelOffset() int {
	return int(c.log_level_offset)
}

func (c *Context) Lowres() int {
	return int(c.lowres)
}

func (c *Context) LumiMasking() float64 {
	return float64(c.lumi_masking)
}

func (c *Context) MaxBFrames() int {
	return int(c.max_b_frames)
}

func (c *Context) MaxPredictionOrder() int {
	return int(c.max_prediction_order)
}

func (c *Context) MaxQdiff() int {
	return int(c.max_qdiff)
}

func (c *Context) MbCmp() int {
	return int(c.mb_cmp)
}

func (c *Context) MbDecision() int {
	return int(c.mb_decision)
}

func (c *Context) MbLmax() int {
	return int(c.mb_lmax)
}

func (c *Context) MbLmin() int {
	return int(c.mb_lmin)
}

func (c *Context) MeCmp() int {
	return int(c.me_cmp)
}

func (c *Context) MePenaltyCompensation() int {
	return int(c.me_penalty_compensation)
}

func (c *Context) MePreCmp() int {
	return int(c.me_pre_cmp)
}

func (c *Context) MeRange() int {
	return int(c.me_range)
}

func (c *Context) MeSubCmp() int {
	return int(c.me_sub_cmp)
}

func (c *Context) MeSubpelQuality() int {
	return int(c.me_subpel_quality)
}

func (c *Context) MinPredictionOrder() int {
	return int(c.min_prediction_order)
}

func (c *Context) MiscBits() int {
	return int(c.misc_bits)
}

func (c *Context) MpegQuant() int {
	return int(c.mpeg_quant)
}

func (c *Context) Mv0Threshold() int {
	return int(c.mv0_threshold)
}

func (c *Context) MvBits() int {
	return int(c.mv_bits)
}

func (c *Context) NoiseReduction() int {
	return int(c.noise_reduction)
}

func (c *Context) NsseWeight() int {
	return int(c.nsse_weight)
}

func (c *Context) PCount() int {
	return int(c.p_count)
}

func (c *Context) PMasking() float64 {
	return float64(c.p_masking)
}

func (c *Context) PTexBits() int {
	return int(c.p_tex_bits)
}

func (c *Context) PreDiaSize() int {
	return int(c.pre_dia_size)
}

func (c *Context) PreMe() int {
	return int(c.pre_me)
}

func (c *Context) PredictionMethod() int {
	return int(c.prediction_method)
}

func (c *Context) Profile() int {
	return int(c.profile)
}

func (c *Context) Qblur() float64 {
	return float64(c.qblur)
}

func (c *Context) Qcompress() float64 {
	return float64(c.qcompress)
}

func (c *Context) Qmax() int {
	return int(c.qmax)
}

func (c *Context) Qmin() int {
	return int(c.qmin)
}

func (c *Context) RcBufferSize() int {
	return int(c.rc_buffer_size)
}

func (c *Context) RcInitialBufferOccupancy() int {
	return int(c.rc_initial_buffer_occupancy)
}

func (c *Context) RcMaxAvailableVbvUse() float64 {
	return float64(c.rc_max_available_vbv_use)
}

func (c *Context) RcMaxRate() int {
	return int(c.rc_max_rate)
}

func (c *Context) RcMinRate() int {
	return int(c.rc_min_rate)
}

func (c *Context) RcMinVbvOverflowUse() float64 {
	return float64(c.rc_min_vbv_overflow_use)
}

func (c *Context) RcOverrideCount() int {
	return int(c.rc_override_count)
}

func (c *Context) RefcountedFrames() int {
	return int(c.refcounted_frames)
}

func (c *Context) Refs() int {
	return int(c.refs)
}

func (c *Context) RtpPayloadSize() int {
	return int(c.rtp_payload_size)
}

func (c *Context) SampleRate() int {
	return int(c.sample_rate)
}

func (c *Context) ScenechangeThreshold() int {
	return int(c.scenechange_threshold)
}

func (c *Context) SeekPreroll() int {
	return int(c.seek_preroll)
}

func (c *Context) SideDataOnlyPackets() int {
	return int(c.side_data_only_packets)
}

func (c *Context) SkipAlpha() int {
	return int(c.skip_alpha)
}

func (c *Context) SkipBottom() int {
	return int(c.skip_bottom)
}

func (c *Context) SkipCount() int {
	return int(c.skip_count)
}

func (c *Context) SkipTop() int {
	return int(c.skip_top)
}

func (c *Context) SliceCount() int {
	return int(c.slice_count)
}

func (c *Context) SliceFlags() int {
	return int(c.slice_flags)
}

func (c *Context) Slices() int {
	return int(c.slices)
}

func (c *Context) SpatialCplxMasking() float64 {
	return float64(c.spatial_cplx_masking)
}

//func (ctxt *Context) StrictStdCompliance() int {
//	return int(ctxt.strict_std_compliance)
//}

func (c *Context) SubCharencMode() int {
	return int(c.sub_charenc_mode)
}

func (c *Context) SubtitleHeaderSize() int {
	return int(c.subtitle_header_size)
}

func (c *Context) TemporalCplxMasking() float64 {
	return float64(c.temporal_cplx_masking)
}

func (c *Context) ThreadCount() int {
	return int(c.thread_count)
}

func (c *Context) ThreadSafeCallbacks() int {
	return int(c.thread_safe_callbacks)
}

func (c *Context) ThreadType() int {
	return int(c.thread_type)
}

func (c *Context) TicksPerFrame() int {
	return int(c.ticks_per_frame)
}

func (c *Context) Trellis() int {
	return int(c.trellis)
}

func (c *Context) Width() int {
	return int(c.width)
}

func (c *Context) WorkaroundBugs() int {
	return int(c.workaround_bugs)
}

func (c *Context) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(c.audio_service_type)
}

func (c *Context) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(c.chroma_sample_location)
}

func (c *Context) CodecDescriptor() *Descriptor {
	return (*Descriptor)(c.codec_descriptor)
}

func (c *Context) CodecId() CodecId {
	return (CodecId)(c.codec_id)
}

func (c *Context) CodecType() MediaType {
	return (MediaType)(c.codec_type)
}

func (c *Context) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(c.color_primaries)
}

func (c *Context) ColorRange() AvColorRange {
	return (AvColorRange)(c.color_range)
}

func (c *Context) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(c.color_trc)
}

func (c *Context) Colorspace() AvColorSpace {
	return (AvColorSpace)(c.colorspace)
}

func (c *Context) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(c.field_order)
}

func (c *Context) PixFmt() PixelFormat {
	return (PixelFormat)(c.pix_fmt)
}

func (c *Context) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(c.request_sample_fmt)
}

func (c *Context) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(c.sample_fmt)
}

func (c *Context) SkipFrame() AvDiscard {
	return (AvDiscard)(c.skip_frame)
}

func (c *Context) SkipIdct() AvDiscard {
	return (AvDiscard)(c.skip_idct)
}

func (c *Context) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(c.skip_loop_filter)
}

func (c *Context) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(c.chapters))
}

func (c *Context) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(c.audio_codec))
}

func (c *Context) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(c.subtitle_codec))
}

func (c *Context) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(c.video_codec))
}

func (c *Context) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(c.metadata))
}

func (c *Context) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(c.internal))
}

func (c *Context) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(c.pb))
}

func (c *Context) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(c.interrupt_callback)
}

func (c *Context) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.programs)),
		Len:  int(c.NbPrograms()),
		Cap:  int(c.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

func (c *Context) Streams() []*Stream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.streams)),
		Len:  int(c.NbStreams()),
		Cap:  int(c.NbStreams()),
	}

	return *((*[]*Stream)(unsafe.Pointer(&header)))
}

func (c *Context) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&c.filename[0])))
}

// func (ctxt *Context) CodecWhitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *Context) FormatWhitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

func (c *Context) AudioCodecId() CodecId {
	return CodecId(c.audio_codec_id)
}

func (c *Context) SubtitleCodecId() CodecId {
	return CodecId(c.subtitle_codec_id)
}

func (c *Context) VideoCodecId() CodecId {
	return CodecId(c.video_codec_id)
}

func (c *Context) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(c.duration_estimation_method)
}

func (c *Context) AudioPreload() int {
	return int(c.audio_preload)
}

func (c *Context) AvioFlags() int {
	return int(c.avio_flags)
}

func (c *Context) AvoidNegativeTs() int {
	return int(c.avoid_negative_ts)
}

//func (ctxt *Context) BitRate() int {
//	return int(ctxt.bit_rate)
//}

func (c *Context) CtxFlags() int {
	return int(c.ctx_flags)
}

func (c *Context) Debug() int {
	return int(c.debug)
}

func (c *Context) ErrorRecognition() int {
	return int(c.error_recognition)
}

func (c *Context) EventFlags() int {
	return int(c.event_flags)
}

func (c *Context) Flags() int {
	return int(c.flags)
}

func (c *Context) FlushPackets() int {
	return int(c.flush_packets)
}

func (c *Context) FormatProbesize() int {
	return int(c.format_probesize)
}

func (c *Context) FpsProbeSize() int {
	return int(c.fps_probe_size)
}

func (c *Context) IoRepositioned() int {
	return int(c.io_repositioned)
}

func (c *Context) Keylen() int {
	return int(c.keylen)
}

func (c *Context) MaxChunkDuration() int {
	return int(c.max_chunk_duration)
}

func (c *Context) MaxChunkSize() int {
	return int(c.max_chunk_size)
}

func (c *Context) MaxDelay() int {
	return int(c.max_delay)
}

func (c *Context) MaxTsProbe() int {
	return int(c.max_ts_probe)
}

func (c *Context) MetadataHeaderPadding() int {
	return int(c.metadata_header_padding)
}

func (c *Context) ProbeScore() int {
	return int(c.probe_score)
}

func (c *Context) Seek2any() int {
	return int(c.seek2any)
}

func (c *Context) StrictStdCompliance() int {
	return int(c.strict_std_compliance)
}

func (c *Context) TsId() int {
	return int(c.ts_id)
}

func (c *Context) UseWallclockAsTimestamps() int {
	return int(c.use_wallclock_as_timestamps)
}

func (c *Context) Duration() int64 {
	return int64(c.duration)
}

func (c *Context) MaxAnalyzeDuration2() int64 {
	return int64(c.max_analyze_duration)
}

func (c *Context) MaxInterleaveDelta() int64 {
	return int64(c.max_interleave_delta)
}

func (c *Context) OutputTsOffset() int64 {
	return int64(c.output_ts_offset)
}

func (c *Context) Probesize2() int64 {
	return int64(c.probesize)
}

func (c *Context) SkipInitialBytes() int64 {
	return int64(c.skip_initial_bytes)
}

func (c *Context) StartTime() int64 {
	return int64(c.start_time)
}

func (c *Context) StartTimeRealtime() int64 {
	return int64(c.start_time_realtime)
}

func (c *Context) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(c.iformat))
}

func (c *Context) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(c.oformat))
}

// func (ctxt *Context) DumpSeparator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

func (c *Context) CorrectTsOverflow() int {
	return int(c.correct_ts_overflow)
}

func (c *Context) MaxIndexSize() uint {
	return uint(c.max_index_size)
}

func (c *Context) MaxPictureBuffer() uint {
	return uint(c.max_picture_buffer)
}

func (c *Context) NbChapters() uint {
	return uint(c.nb_chapters)
}

func (c *Context) NbPrograms() uint {
	return uint(c.nb_programs)
}

func (c *Context) NbStreams() uint {
	return uint(c.nb_streams)
}

func (c *Context) PacketSize() uint {
	return uint(c.packet_size)
}

func (c *Context) Probesize() uint {
	return uint(c.probesize)
}

func (c *Context) SetPb(pb *AvIOContext) {
	c.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

func (c *Context) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&c.pb))
}
