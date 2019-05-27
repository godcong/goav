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

func (ctx *AVCodecContext) ActiveThreadType() int {
	return int(ctx.active_thread_type)
}

func (ctx *AVCodecContext) BFrameStrategy() int {
	return int(ctx.b_frame_strategy)
}

func (ctx *AVCodecContext) BQuantFactor() float64 {
	return float64(ctx.b_quant_factor)
}

func (ctx *AVCodecContext) BQuantOffset() float64 {
	return float64(ctx.b_quant_offset)
}

func (ctx *AVCodecContext) BSensitivity() int {
	return int(ctx.b_sensitivity)
}

func (ctx *AVCodecContext) BidirRefine() int {
	return int(ctx.bidir_refine)
}

func (ctx *AVCodecContext) BitRate() int {
	return int(ctx.bit_rate)
}

func (ctx *AVCodecContext) BitRateTolerance() int {
	return int(ctx.bit_rate_tolerance)
}

func (ctx *AVCodecContext) BitsPerCodedSample() int {
	return int(ctx.bits_per_coded_sample)
}

func (ctx *AVCodecContext) BitsPerRawSample() int {
	return int(ctx.bits_per_raw_sample)
}

func (ctx *AVCodecContext) BlockAlign() int {
	return int(ctx.block_align)
}

func (ctx *AVCodecContext) BrdScale() int {
	return int(ctx.brd_scale)
}

func (ctx *AVCodecContext) Channels() int {
	return int(ctx.channels)
}

func (ctx *AVCodecContext) Chromaoffset() int {
	return int(ctx.chromaoffset)
}

func (ctx *AVCodecContext) CodedHeight() int {
	return int(ctx.coded_height)
}

func (ctx *AVCodecContext) CodedWidth() int {
	return int(ctx.coded_width)
}

func (ctx *AVCodecContext) CoderType() int {
	return int(ctx.coder_type)
}

func (ctx *AVCodecContext) CompressionLevel() int {
	return int(ctx.compression_level)
}

func (ctx *AVCodecContext) ContextModel() int {
	return int(ctx.context_model)
}

func (ctx *AVCodecContext) Cutoff() int {
	return int(ctx.cutoff)
}

func (ctx *AVCodecContext) DarkMasking() float64 {
	return float64(ctx.dark_masking)
}

func (ctx *AVCodecContext) DctAlgo() int {
	return int(ctx.dct_algo)
}

func (ctx *AVCodecContext) Debug() int {
	return int(ctx.debug)
}

func (ctx *AVCodecContext) DebugMv() int {
	return int(ctx.debug_mv)
}

func (ctx *AVCodecContext) Delay() int {
	return int(ctx.delay)
}

func (ctx *AVCodecContext) DiaSize() int {
	return int(ctx.dia_size)
}

func (ctx *AVCodecContext) ErrRecognition() int {
	return int(ctx.err_recognition)
}

func (ctx *AVCodecContext) ErrorConcealment() int {
	return int(ctx.error_concealment)
}

func (ctx *AVCodecContext) ExtradataSize() int {
	return int(ctx.extradata_size)
}

func (ctx *AVCodecContext) Flags() int {
	return int(ctx.flags)
}

func (ctx *AVCodecContext) Flags2() int {
	return int(ctx.flags2)
}

func (ctx *AVCodecContext) FrameBits() int {
	return int(ctx.frame_bits)
}

func (ctx *AVCodecContext) FrameNumber() int {
	return int(ctx.frame_number)
}

func (ctx *AVCodecContext) FrameSize() int {
	return int(ctx.frame_size)
}

func (ctx *AVCodecContext) FrameSkipCmp() int {
	return int(ctx.frame_skip_cmp)
}

func (ctx *AVCodecContext) FrameSkipExp() int {
	return int(ctx.frame_skip_exp)
}

func (ctx *AVCodecContext) FrameSkipFactor() int {
	return int(ctx.frame_skip_factor)
}

func (ctx *AVCodecContext) FrameSkipThreshold() int {
	return int(ctx.frame_skip_threshold)
}

func (ctx *AVCodecContext) GlobalQuality() int {
	return int(ctx.global_quality)
}

func (ctx *AVCodecContext) GopSize() int {
	return int(ctx.gop_size)
}

func (ctx *AVCodecContext) HasBFrames() int {
	return int(ctx.has_b_frames)
}

func (ctx *AVCodecContext) HeaderBits() int {
	return int(ctx.header_bits)
}

func (ctx *AVCodecContext) Height() int {
	return int(ctx.height)
}

func (ctx *AVCodecContext) ICount() int {
	return int(ctx.i_count)
}

func (ctx *AVCodecContext) IQuantFactor() float64 {
	return float64(ctx.i_quant_factor)
}

func (ctx *AVCodecContext) IQuantOffset() float64 {
	return float64(ctx.i_quant_offset)
}

func (ctx *AVCodecContext) ITexBits() int {
	return int(ctx.i_tex_bits)
}

func (ctx *AVCodecContext) IdctAlgo() int {
	return int(ctx.idct_algo)
}

func (ctx *AVCodecContext) IldctCmp() int {
	return int(ctx.ildct_cmp)
}

func (ctx *AVCodecContext) IntraDcPrecision() int {
	return int(ctx.intra_dc_precision)
}

func (ctx *AVCodecContext) KeyintMin() int {
	return int(ctx.keyint_min)
}

func (ctx *AVCodecContext) LastPredictorCount() int {
	return int(ctx.last_predictor_count)
}

func (ctx *AVCodecContext) Level() int {
	return int(ctx.level)
}

func (ctx *AVCodecContext) LogLevelOffset() int {
	return int(ctx.log_level_offset)
}

func (ctx *AVCodecContext) Lowres() int {
	return int(ctx.lowres)
}

func (ctx *AVCodecContext) LumiMasking() float64 {
	return float64(ctx.lumi_masking)
}

func (ctx *AVCodecContext) MaxBFrames() int {
	return int(ctx.max_b_frames)
}

func (ctx *AVCodecContext) MaxPredictionOrder() int {
	return int(ctx.max_prediction_order)
}

func (ctx *AVCodecContext) MaxQdiff() int {
	return int(ctx.max_qdiff)
}

func (ctx *AVCodecContext) MbCmp() int {
	return int(ctx.mb_cmp)
}

func (ctx *AVCodecContext) MbDecision() int {
	return int(ctx.mb_decision)
}

func (ctx *AVCodecContext) MbLmax() int {
	return int(ctx.mb_lmax)
}

func (ctx *AVCodecContext) MbLmin() int {
	return int(ctx.mb_lmin)
}

func (ctx *AVCodecContext) MeCmp() int {
	return int(ctx.me_cmp)
}

func (ctx *AVCodecContext) MePenaltyCompensation() int {
	return int(ctx.me_penalty_compensation)
}

func (ctx *AVCodecContext) MePreCmp() int {
	return int(ctx.me_pre_cmp)
}

func (ctx *AVCodecContext) MeRange() int {
	return int(ctx.me_range)
}

func (ctx *AVCodecContext) MeSubCmp() int {
	return int(ctx.me_sub_cmp)
}

func (ctx *AVCodecContext) MeSubpelQuality() int {
	return int(ctx.me_subpel_quality)
}

func (ctx *AVCodecContext) MinPredictionOrder() int {
	return int(ctx.min_prediction_order)
}

func (ctx *AVCodecContext) MiscBits() int {
	return int(ctx.misc_bits)
}

func (ctx *AVCodecContext) MpegQuant() int {
	return int(ctx.mpeg_quant)
}

func (ctx *AVCodecContext) Mv0Threshold() int {
	return int(ctx.mv0_threshold)
}

func (ctx *AVCodecContext) MvBits() int {
	return int(ctx.mv_bits)
}

func (ctx *AVCodecContext) NoiseReduction() int {
	return int(ctx.noise_reduction)
}

func (ctx *AVCodecContext) NsseWeight() int {
	return int(ctx.nsse_weight)
}

func (ctx *AVCodecContext) PCount() int {
	return int(ctx.p_count)
}

func (ctx *AVCodecContext) PMasking() float64 {
	return float64(ctx.p_masking)
}

func (ctx *AVCodecContext) PTexBits() int {
	return int(ctx.p_tex_bits)
}

func (ctx *AVCodecContext) PreDiaSize() int {
	return int(ctx.pre_dia_size)
}

func (ctx *AVCodecContext) PreMe() int {
	return int(ctx.pre_me)
}

func (ctx *AVCodecContext) PredictionMethod() int {
	return int(ctx.prediction_method)
}

func (ctx *AVCodecContext) Profile() int {
	return int(ctx.profile)
}

func (ctx *AVCodecContext) Qblur() float64 {
	return float64(ctx.qblur)
}

func (ctx *AVCodecContext) Qcompress() float64 {
	return float64(ctx.qcompress)
}

func (ctx *AVCodecContext) Qmax() int {
	return int(ctx.qmax)
}

func (ctx *AVCodecContext) Qmin() int {
	return int(ctx.qmin)
}

func (ctx *AVCodecContext) RcBufferSize() int {
	return int(ctx.rc_buffer_size)
}

func (ctx *AVCodecContext) RcInitialBufferOccupancy() int {
	return int(ctx.rc_initial_buffer_occupancy)
}

func (ctx *AVCodecContext) RcMaxAvailableVbvUse() float64 {
	return float64(ctx.rc_max_available_vbv_use)
}

func (ctx *AVCodecContext) RcMaxRate() int {
	return int(ctx.rc_max_rate)
}

func (ctx *AVCodecContext) RcMinRate() int {
	return int(ctx.rc_min_rate)
}

func (ctx *AVCodecContext) RcMinVbvOverflowUse() float64 {
	return float64(ctx.rc_min_vbv_overflow_use)
}

func (ctx *AVCodecContext) RcOverrideCount() int {
	return int(ctx.rc_override_count)
}

func (ctx *AVCodecContext) RefcountedFrames() int {
	return int(ctx.refcounted_frames)
}

func (ctx *AVCodecContext) Refs() int {
	return int(ctx.refs)
}

func (ctx *AVCodecContext) RtpPayloadSize() int {
	return int(ctx.rtp_payload_size)
}

func (ctx *AVCodecContext) SampleRate() int {
	return int(ctx.sample_rate)
}

func (ctx *AVCodecContext) ScenechangeThreshold() int {
	return int(ctx.scenechange_threshold)
}

func (ctx *AVCodecContext) SeekPreroll() int {
	return int(ctx.seek_preroll)
}

func (ctx *AVCodecContext) SideDataOnlyPackets() int {
	return int(ctx.side_data_only_packets)
}

func (ctx *AVCodecContext) SkipAlpha() int {
	return int(ctx.skip_alpha)
}

func (ctx *AVCodecContext) SkipBottom() int {
	return int(ctx.skip_bottom)
}

func (ctx *AVCodecContext) SkipCount() int {
	return int(ctx.skip_count)
}

func (ctx *AVCodecContext) SkipTop() int {
	return int(ctx.skip_top)
}

func (ctx *AVCodecContext) SliceCount() int {
	return int(ctx.slice_count)
}

func (ctx *AVCodecContext) SliceFlags() int {
	return int(ctx.slice_flags)
}

func (ctx *AVCodecContext) Slices() int {
	return int(ctx.slices)
}

func (ctx *AVCodecContext) SpatialCplxMasking() float64 {
	return float64(ctx.spatial_cplx_masking)
}

func (ctx *AVCodecContext) StrictStdCompliance() int {
	return int(ctx.strict_std_compliance)
}

func (ctx *AVCodecContext) SubCharencMode() int {
	return int(ctx.sub_charenc_mode)
}

func (ctx *AVCodecContext) SubtitleHeaderSize() int {
	return int(ctx.subtitle_header_size)
}

func (ctx *AVCodecContext) TemporalCplxMasking() float64 {
	return float64(ctx.temporal_cplx_masking)
}

func (ctx *AVCodecContext) ThreadCount() int {
	return int(ctx.thread_count)
}

func (ctx *AVCodecContext) ThreadSafeCallbacks() int {
	return int(ctx.thread_safe_callbacks)
}

func (ctx *AVCodecContext) ThreadType() int {
	return int(ctx.thread_type)
}

func (ctx *AVCodecContext) TicksPerFrame() int {
	return int(ctx.ticks_per_frame)
}

func (ctx *AVCodecContext) Trellis() int {
	return int(ctx.trellis)
}

func (ctx *AVCodecContext) Width() int {
	return int(ctx.width)
}

func (ctx *AVCodecContext) WorkaroundBugs() int {
	return int(ctx.workaround_bugs)
}

func (ctx *AVCodecContext) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(ctx.audio_service_type)
}

func (ctx *AVCodecContext) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(ctx.chroma_sample_location)
}

func (ctx *AVCodecContext) CodecDescriptor() *Descriptor {
	return (*Descriptor)(ctx.codec_descriptor)
}

func (ctx *AVCodecContext) CodecId() CodecId {
	return (CodecId)(ctx.codec_id)
}

func (ctx *AVCodecContext) CodecType() MediaType {
	return (MediaType)(ctx.codec_type)
}

func (ctx *AVCodecContext) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(ctx.color_primaries)
}

func (ctx *AVCodecContext) ColorRange() AvColorRange {
	return (AvColorRange)(ctx.color_range)
}

func (ctx *AVCodecContext) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(ctx.color_trc)
}

func (ctx *AVCodecContext) Colorspace() AvColorSpace {
	return (AvColorSpace)(ctx.colorspace)
}

func (ctx *AVCodecContext) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(ctx.field_order)
}

func (ctx *AVCodecContext) PixFmt() PixelFormat {
	return (PixelFormat)(ctx.pix_fmt)
}

func (ctx *AVCodecContext) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.request_sample_fmt)
}

func (ctx *AVCodecContext) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.sample_fmt)
}

func (ctx *AVCodecContext) SkipFrame() AvDiscard {
	return (AvDiscard)(ctx.skip_frame)
}

func (ctx *AVCodecContext) SkipIdct() AvDiscard {
	return (AvDiscard)(ctx.skip_idct)
}

func (ctx *AVCodecContext) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(ctx.skip_loop_filter)
}

func (ctx *AVFormatContext) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(ctx.chapters))
}

func (ctx *AVFormatContext) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctx.audio_codec))
}

func (ctx *AVFormatContext) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctx.subtitle_codec))
}

func (ctx *AVFormatContext) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctx.video_codec))
}

func (ctx *AVFormatContext) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(ctx.metadata))
}

func (ctx *AVFormatContext) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(ctx.internal))
}

func (ctx *AVFormatContext) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(ctx.pb))
}

func (ctx *AVFormatContext) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(ctx.interrupt_callback)
}

func (ctx *AVFormatContext) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctx.programs)),
		Len:  int(ctx.NbPrograms()),
		Cap:  int(ctx.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

func (ctx *AVFormatContext) Streams() []*Stream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctx.streams)),
		Len:  int(ctx.NbStreams()),
		Cap:  int(ctx.NbStreams()),
	}

	return *((*[]*Stream)(unsafe.Pointer(&header)))
}

func (ctx *AVFormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctx.filename[0])))
}

// func (ctxt *AVFormatContext) CodecWhitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *AVFormatContext) FormatWhitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

func (ctx *AVFormatContext) AudioCodecId() CodecId {
	return CodecId(ctx.audio_codec_id)
}

func (ctx *AVFormatContext) SubtitleCodecId() CodecId {
	return CodecId(ctx.subtitle_codec_id)
}

func (ctx *AVFormatContext) VideoCodecId() CodecId {
	return CodecId(ctx.video_codec_id)
}

func (ctx *AVFormatContext) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(ctx.duration_estimation_method)
}

func (ctx *AVFormatContext) AudioPreload() int {
	return int(ctx.audio_preload)
}

func (ctx *AVFormatContext) AvioFlags() int {
	return int(ctx.avio_flags)
}

func (ctx *AVFormatContext) AvoidNegativeTs() int {
	return int(ctx.avoid_negative_ts)
}

func (ctx *AVFormatContext) BitRate() int {
	return int(ctx.bit_rate)
}

func (ctx *AVFormatContext) CtxFlags() int {
	return int(ctx.ctx_flags)
}

func (ctx *AVFormatContext) Debug() int {
	return int(ctx.debug)
}

func (ctx *AVFormatContext) ErrorRecognition() int {
	return int(ctx.error_recognition)
}

func (ctx *AVFormatContext) EventFlags() int {
	return int(ctx.event_flags)
}

func (ctx *AVFormatContext) Flags() int {
	return int(ctx.flags)
}

func (ctx *AVFormatContext) FlushPackets() int {
	return int(ctx.flush_packets)
}

func (ctx *AVFormatContext) FormatProbesize() int {
	return int(ctx.format_probesize)
}

func (ctx *AVFormatContext) FpsProbeSize() int {
	return int(ctx.fps_probe_size)
}

func (ctx *AVFormatContext) IoRepositioned() int {
	return int(ctx.io_repositioned)
}

func (ctx *AVFormatContext) Keylen() int {
	return int(ctx.keylen)
}

func (ctx *AVFormatContext) MaxChunkDuration() int {
	return int(ctx.max_chunk_duration)
}

func (ctx *AVFormatContext) MaxChunkSize() int {
	return int(ctx.max_chunk_size)
}

func (ctx *AVFormatContext) MaxDelay() int {
	return int(ctx.max_delay)
}

func (ctx *AVFormatContext) MaxTsProbe() int {
	return int(ctx.max_ts_probe)
}

func (ctx *AVFormatContext) MetadataHeaderPadding() int {
	return int(ctx.metadata_header_padding)
}

func (ctx *AVFormatContext) ProbeScore() int {
	return int(ctx.probe_score)
}

func (ctx *AVFormatContext) Seek2any() int {
	return int(ctx.seek2any)
}

func (ctx *AVFormatContext) StrictStdCompliance() int {
	return int(ctx.strict_std_compliance)
}

func (ctx *AVFormatContext) TsId() int {
	return int(ctx.ts_id)
}

func (ctx *AVFormatContext) UseWallclockAsTimestamps() int {
	return int(ctx.use_wallclock_as_timestamps)
}

func (ctx *AVFormatContext) Duration() int64 {
	return int64(ctx.duration)
}

func (ctx *AVFormatContext) MaxAnalyzeDuration2() int64 {
	return int64(ctx.max_analyze_duration)
}

func (ctx *AVFormatContext) MaxInterleaveDelta() int64 {
	return int64(ctx.max_interleave_delta)
}

func (ctx *AVFormatContext) OutputTsOffset() int64 {
	return int64(ctx.output_ts_offset)
}

func (ctx *AVFormatContext) Probesize2() int64 {
	return int64(ctx.probesize)
}

func (ctx *AVFormatContext) SkipInitialBytes() int64 {
	return int64(ctx.skip_initial_bytes)
}

func (ctx *AVFormatContext) StartTime() int64 {
	return int64(ctx.start_time)
}

func (ctx *AVFormatContext) StartTimeRealtime() int64 {
	return int64(ctx.start_time_realtime)
}

func (ctx *AVFormatContext) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(ctx.iformat))
}

func (ctx *AVFormatContext) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(ctx.oformat))
}

// func (ctxt *AVFormatContext) DumpSeparator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

func (ctx *AVFormatContext) CorrectTsOverflow() int {
	return int(ctx.correct_ts_overflow)
}

func (ctx *AVFormatContext) MaxIndexSize() uint {
	return uint(ctx.max_index_size)
}

func (ctx *AVFormatContext) MaxPictureBuffer() uint {
	return uint(ctx.max_picture_buffer)
}

func (ctx *AVFormatContext) NbChapters() uint {
	return uint(ctx.nb_chapters)
}

func (ctx *AVFormatContext) NbPrograms() uint {
	return uint(ctx.nb_programs)
}

func (ctx *AVFormatContext) NbStreams() uint {
	return uint(ctx.nb_streams)
}

func (ctx *AVFormatContext) PacketSize() uint {
	return uint(ctx.packet_size)
}

func (ctx *AVFormatContext) Probesize() uint {
	return uint(ctx.probesize)
}

func (ctx *AVFormatContext) SetPb(pb *AvIOContext) {
	ctx.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

func (ctx *AVFormatContext) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&ctx.pb))
}
