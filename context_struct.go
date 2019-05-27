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

// ActiveThreadType ...
func (ctx *CodecContext) ActiveThreadType() int {
	return int(ctx.active_thread_type)
}

// BFrameStrategy ...
func (ctx *CodecContext) BFrameStrategy() int {
	return int(ctx.b_frame_strategy)
}

// BQuantFactor ...
func (ctx *CodecContext) BQuantFactor() float64 {
	return float64(ctx.b_quant_factor)
}

// BQuantOffset ...
func (ctx *CodecContext) BQuantOffset() float64 {
	return float64(ctx.b_quant_offset)
}

// BSensitivity ...
func (ctx *CodecContext) BSensitivity() int {
	return int(ctx.b_sensitivity)
}

// BidirRefine ...
func (ctx *CodecContext) BidirRefine() int {
	return int(ctx.bidir_refine)
}

// BitRate ...
func (ctx *CodecContext) BitRate() int {
	return int(ctx.bit_rate)
}

// BitRateTolerance ...
func (ctx *CodecContext) BitRateTolerance() int {
	return int(ctx.bit_rate_tolerance)
}

// BitsPerCodedSample ...
func (ctx *CodecContext) BitsPerCodedSample() int {
	return int(ctx.bits_per_coded_sample)
}

// BitsPerRawSample ...
func (ctx *CodecContext) BitsPerRawSample() int {
	return int(ctx.bits_per_raw_sample)
}

// BlockAlign ...
func (ctx *CodecContext) BlockAlign() int {
	return int(ctx.block_align)
}

// BrdScale ...
func (ctx *CodecContext) BrdScale() int {
	return int(ctx.brd_scale)
}

// Channels ...
func (ctx *CodecContext) Channels() int {
	return int(ctx.channels)
}

// Chromaoffset ...
func (ctx *CodecContext) Chromaoffset() int {
	return int(ctx.chromaoffset)
}

// CodedHeight ...
func (ctx *CodecContext) CodedHeight() int {
	return int(ctx.coded_height)
}

// CodedWidth ...
func (ctx *CodecContext) CodedWidth() int {
	return int(ctx.coded_width)
}

// CoderType ...
func (ctx *CodecContext) CoderType() int {
	return int(ctx.coder_type)
}

// CompressionLevel ...
func (ctx *CodecContext) CompressionLevel() int {
	return int(ctx.compression_level)
}

// ContextModel ...
func (ctx *CodecContext) ContextModel() int {
	return int(ctx.context_model)
}

// Cutoff ...
func (ctx *CodecContext) Cutoff() int {
	return int(ctx.cutoff)
}

// DarkMasking ...
func (ctx *CodecContext) DarkMasking() float64 {
	return float64(ctx.dark_masking)
}

// DctAlgo ...
func (ctx *CodecContext) DctAlgo() int {
	return int(ctx.dct_algo)
}

// Debug ...
func (ctx *CodecContext) Debug() int {
	return int(ctx.debug)
}

// DebugMv ...
func (ctx *CodecContext) DebugMv() int {
	return int(ctx.debug_mv)
}

// Delay ...
func (ctx *CodecContext) Delay() int {
	return int(ctx.delay)
}

// DiaSize ...
func (ctx *CodecContext) DiaSize() int {
	return int(ctx.dia_size)
}

// ErrRecognition ...
func (ctx *CodecContext) ErrRecognition() int {
	return int(ctx.err_recognition)
}

// ErrorConcealment ...
func (ctx *CodecContext) ErrorConcealment() int {
	return int(ctx.error_concealment)
}

// ExtradataSize ...
func (ctx *CodecContext) ExtradataSize() int {
	return int(ctx.extradata_size)
}

// Flags ...
func (ctx *CodecContext) Flags() int {
	return int(ctx.flags)
}

// Flags2 ...
func (ctx *CodecContext) Flags2() int {
	return int(ctx.flags2)
}

// FrameBits ...
func (ctx *CodecContext) FrameBits() int {
	return int(ctx.frame_bits)
}

// FrameNumber ...
func (ctx *CodecContext) FrameNumber() int {
	return int(ctx.frame_number)
}

// FrameSize ...
func (ctx *CodecContext) FrameSize() int {
	return int(ctx.frame_size)
}

// FrameSkipCmp ...
func (ctx *CodecContext) FrameSkipCmp() int {
	return int(ctx.frame_skip_cmp)
}

// FrameSkipExp ...
func (ctx *CodecContext) FrameSkipExp() int {
	return int(ctx.frame_skip_exp)
}

// FrameSkipFactor ...
func (ctx *CodecContext) FrameSkipFactor() int {
	return int(ctx.frame_skip_factor)
}

// FrameSkipThreshold ...
func (ctx *CodecContext) FrameSkipThreshold() int {
	return int(ctx.frame_skip_threshold)
}

// GlobalQuality ...
func (ctx *CodecContext) GlobalQuality() int {
	return int(ctx.global_quality)
}

// GopSize ...
func (ctx *CodecContext) GopSize() int {
	return int(ctx.gop_size)
}

// HasBFrames ...
func (ctx *CodecContext) HasBFrames() int {
	return int(ctx.has_b_frames)
}

// HeaderBits ...
func (ctx *CodecContext) HeaderBits() int {
	return int(ctx.header_bits)
}

// Height ...
func (ctx *CodecContext) Height() int {
	return int(ctx.height)
}

// ICount ...
func (ctx *CodecContext) ICount() int {
	return int(ctx.i_count)
}

// IQuantFactor ...
func (ctx *CodecContext) IQuantFactor() float64 {
	return float64(ctx.i_quant_factor)
}

// IQuantOffset ...
func (ctx *CodecContext) IQuantOffset() float64 {
	return float64(ctx.i_quant_offset)
}

// ITexBits ...
func (ctx *CodecContext) ITexBits() int {
	return int(ctx.i_tex_bits)
}

// IdctAlgo ...
func (ctx *CodecContext) IdctAlgo() int {
	return int(ctx.idct_algo)
}

// IldctCmp ...
func (ctx *CodecContext) IldctCmp() int {
	return int(ctx.ildct_cmp)
}

// IntraDcPrecision ...
func (ctx *CodecContext) IntraDcPrecision() int {
	return int(ctx.intra_dc_precision)
}

// KeyintMin ...
func (ctx *CodecContext) KeyintMin() int {
	return int(ctx.keyint_min)
}

// LastPredictorCount ...
func (ctx *CodecContext) LastPredictorCount() int {
	return int(ctx.last_predictor_count)
}

// Level ...
func (ctx *CodecContext) Level() int {
	return int(ctx.level)
}

// LogLevelOffset ...
func (ctx *CodecContext) LogLevelOffset() int {
	return int(ctx.log_level_offset)
}

// Lowres ...
func (ctx *CodecContext) Lowres() int {
	return int(ctx.lowres)
}

// LumiMasking ...
func (ctx *CodecContext) LumiMasking() float64 {
	return float64(ctx.lumi_masking)
}

// MaxBFrames ...
func (ctx *CodecContext) MaxBFrames() int {
	return int(ctx.max_b_frames)
}

// MaxPredictionOrder ...
func (ctx *CodecContext) MaxPredictionOrder() int {
	return int(ctx.max_prediction_order)
}

// MaxQdiff ...
func (ctx *CodecContext) MaxQdiff() int {
	return int(ctx.max_qdiff)
}

// MbCmp ...
func (ctx *CodecContext) MbCmp() int {
	return int(ctx.mb_cmp)
}

// MbDecision ...
func (ctx *CodecContext) MbDecision() int {
	return int(ctx.mb_decision)
}

// MbLmax ...
func (ctx *CodecContext) MbLmax() int {
	return int(ctx.mb_lmax)
}

// MbLmin ...
func (ctx *CodecContext) MbLmin() int {
	return int(ctx.mb_lmin)
}

// MeCmp ...
func (ctx *CodecContext) MeCmp() int {
	return int(ctx.me_cmp)
}

// MePenaltyCompensation ...
func (ctx *CodecContext) MePenaltyCompensation() int {
	return int(ctx.me_penalty_compensation)
}

// MePreCmp ...
func (ctx *CodecContext) MePreCmp() int {
	return int(ctx.me_pre_cmp)
}

// MeRange ...
func (ctx *CodecContext) MeRange() int {
	return int(ctx.me_range)
}

// MeSubCmp ...
func (ctx *CodecContext) MeSubCmp() int {
	return int(ctx.me_sub_cmp)
}

// MeSubpelQuality ...
func (ctx *CodecContext) MeSubpelQuality() int {
	return int(ctx.me_subpel_quality)
}

// MinPredictionOrder ...
func (ctx *CodecContext) MinPredictionOrder() int {
	return int(ctx.min_prediction_order)
}

// MiscBits ...
func (ctx *CodecContext) MiscBits() int {
	return int(ctx.misc_bits)
}

// MpegQuant ...
func (ctx *CodecContext) MpegQuant() int {
	return int(ctx.mpeg_quant)
}

// Mv0Threshold ...
func (ctx *CodecContext) Mv0Threshold() int {
	return int(ctx.mv0_threshold)
}

// MvBits ...
func (ctx *CodecContext) MvBits() int {
	return int(ctx.mv_bits)
}

// NoiseReduction ...
func (ctx *CodecContext) NoiseReduction() int {
	return int(ctx.noise_reduction)
}

// NsseWeight ...
func (ctx *CodecContext) NsseWeight() int {
	return int(ctx.nsse_weight)
}

// PCount ...
func (ctx *CodecContext) PCount() int {
	return int(ctx.p_count)
}

// PMasking ...
func (ctx *CodecContext) PMasking() float64 {
	return float64(ctx.p_masking)
}

// PTexBits ...
func (ctx *CodecContext) PTexBits() int {
	return int(ctx.p_tex_bits)
}

// PreDiaSize ...
func (ctx *CodecContext) PreDiaSize() int {
	return int(ctx.pre_dia_size)
}

// PreMe ...
func (ctx *CodecContext) PreMe() int {
	return int(ctx.pre_me)
}

// PredictionMethod ...
func (ctx *CodecContext) PredictionMethod() int {
	return int(ctx.prediction_method)
}

// Profile ...
func (ctx *CodecContext) Profile() int {
	return int(ctx.profile)
}

// Qblur ...
func (ctx *CodecContext) Qblur() float64 {
	return float64(ctx.qblur)
}

// Qcompress ...
func (ctx *CodecContext) Qcompress() float64 {
	return float64(ctx.qcompress)
}

// Qmax ...
func (ctx *CodecContext) Qmax() int {
	return int(ctx.qmax)
}

// Qmin ...
func (ctx *CodecContext) Qmin() int {
	return int(ctx.qmin)
}

// RcBufferSize ...
func (ctx *CodecContext) RcBufferSize() int {
	return int(ctx.rc_buffer_size)
}

// RcInitialBufferOccupancy ...
func (ctx *CodecContext) RcInitialBufferOccupancy() int {
	return int(ctx.rc_initial_buffer_occupancy)
}

// RcMaxAvailableVbvUse ...
func (ctx *CodecContext) RcMaxAvailableVbvUse() float64 {
	return float64(ctx.rc_max_available_vbv_use)
}

// RcMaxRate ...
func (ctx *CodecContext) RcMaxRate() int {
	return int(ctx.rc_max_rate)
}

// RcMinRate ...
func (ctx *CodecContext) RcMinRate() int {
	return int(ctx.rc_min_rate)
}

// RcMinVbvOverflowUse ...
func (ctx *CodecContext) RcMinVbvOverflowUse() float64 {
	return float64(ctx.rc_min_vbv_overflow_use)
}

// RcOverrideCount ...
func (ctx *CodecContext) RcOverrideCount() int {
	return int(ctx.rc_override_count)
}

// RefcountedFrames ...
func (ctx *CodecContext) RefcountedFrames() int {
	return int(ctx.refcounted_frames)
}

// Refs ...
func (ctx *CodecContext) Refs() int {
	return int(ctx.refs)
}

// RtpPayloadSize ...
func (ctx *CodecContext) RtpPayloadSize() int {
	return int(ctx.rtp_payload_size)
}

// SampleRate ...
func (ctx *CodecContext) SampleRate() int {
	return int(ctx.sample_rate)
}

// ScenechangeThreshold ...
func (ctx *CodecContext) ScenechangeThreshold() int {
	return int(ctx.scenechange_threshold)
}

// SeekPreroll ...
func (ctx *CodecContext) SeekPreroll() int {
	return int(ctx.seek_preroll)
}

// SideDataOnlyPackets ...
func (ctx *CodecContext) SideDataOnlyPackets() int {
	return int(ctx.side_data_only_packets)
}

// SkipAlpha ...
func (ctx *CodecContext) SkipAlpha() int {
	return int(ctx.skip_alpha)
}

// SkipBottom ...
func (ctx *CodecContext) SkipBottom() int {
	return int(ctx.skip_bottom)
}

// SkipCount ...
func (ctx *CodecContext) SkipCount() int {
	return int(ctx.skip_count)
}

// SkipTop ...
func (ctx *CodecContext) SkipTop() int {
	return int(ctx.skip_top)
}

// SliceCount ...
func (ctx *CodecContext) SliceCount() int {
	return int(ctx.slice_count)
}

// SliceFlags ...
func (ctx *CodecContext) SliceFlags() int {
	return int(ctx.slice_flags)
}

// Slices ...
func (ctx *CodecContext) Slices() int {
	return int(ctx.slices)
}

// SpatialCplxMasking ...
func (ctx *CodecContext) SpatialCplxMasking() float64 {
	return float64(ctx.spatial_cplx_masking)
}

// StrictStdCompliance ...
func (ctx *CodecContext) StrictStdCompliance() int {
	return int(ctx.strict_std_compliance)
}

// SubCharencMode ...
func (ctx *CodecContext) SubCharencMode() int {
	return int(ctx.sub_charenc_mode)
}

// SubtitleHeaderSize ...
func (ctx *CodecContext) SubtitleHeaderSize() int {
	return int(ctx.subtitle_header_size)
}

// TemporalCplxMasking ...
func (ctx *CodecContext) TemporalCplxMasking() float64 {
	return float64(ctx.temporal_cplx_masking)
}

// ThreadCount ...
func (ctx *CodecContext) ThreadCount() int {
	return int(ctx.thread_count)
}

// ThreadSafeCallbacks ...
func (ctx *CodecContext) ThreadSafeCallbacks() int {
	return int(ctx.thread_safe_callbacks)
}

// ThreadType ...
func (ctx *CodecContext) ThreadType() int {
	return int(ctx.thread_type)
}

// TicksPerFrame ...
func (ctx *CodecContext) TicksPerFrame() int {
	return int(ctx.ticks_per_frame)
}

// Trellis ...
func (ctx *CodecContext) Trellis() int {
	return int(ctx.trellis)
}

// Width ...
func (ctx *CodecContext) Width() int {
	return int(ctx.width)
}

// WorkaroundBugs ...
func (ctx *CodecContext) WorkaroundBugs() int {
	return int(ctx.workaround_bugs)
}

// AudioServiceType ...
func (ctx *CodecContext) AudioServiceType() AudioServiceType {
	return (AudioServiceType)(ctx.audio_service_type)
}

// ChromaSampleLocation ...
func (ctx *CodecContext) ChromaSampleLocation() ChromaLocation {
	return (ChromaLocation)(ctx.chroma_sample_location)
}

// CodecDescriptor ...
func (ctx *CodecContext) CodecDescriptor() *CodecDescriptor {
	return (*CodecDescriptor)(ctx.codec_descriptor)
}

// CodecID ...
func (ctx *CodecContext) CodecID() CodecID {
	return (CodecID)(ctx.codec_id)
}

// CodecType ...
func (ctx *CodecContext) CodecType() MediaType {
	return (MediaType)(ctx.codec_type)
}

// ColorPrimaries ...
func (ctx *CodecContext) ColorPrimaries() ColorPrimaries {
	return (ColorPrimaries)(ctx.color_primaries)
}

// ColorRange ...
func (ctx *CodecContext) ColorRange() ColorRange {
	return (ColorRange)(ctx.color_range)
}

// ColorTrc ...
func (ctx *CodecContext) ColorTrc() ColorTransferCharacteristic {
	return (ColorTransferCharacteristic)(ctx.color_trc)
}

// Colorspace ...
func (ctx *CodecContext) Colorspace() ColorSpace {
	return (ColorSpace)(ctx.colorspace)
}

// FieldOrder ...
func (ctx *CodecContext) FieldOrder() FieldOrder {
	return (FieldOrder)(ctx.field_order)
}

// PixFmt ...
func (ctx *CodecContext) PixFmt() PixelFormat {
	return (PixelFormat)(ctx.pix_fmt)
}

// RequestSampleFmt ...
func (ctx *CodecContext) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.request_sample_fmt)
}

// SampleFmt ...
func (ctx *CodecContext) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.sample_fmt)
}

// SkipFrame ...
func (ctx *CodecContext) SkipFrame() Discard {
	return (Discard)(ctx.skip_frame)
}

// SkipIdct ...
func (ctx *CodecContext) SkipIdct() Discard {
	return (Discard)(ctx.skip_idct)
}

// SkipLoopFilter ...
func (ctx *CodecContext) SkipLoopFilter() Discard {
	return (Discard)(ctx.skip_loop_filter)
}

// Chapters ...
func (ctx *FormatContext) Chapters() **Chapter {
	return (**Chapter)(unsafe.Pointer(ctx.chapters))
}

// AudioCodec ...
func (ctx *FormatContext) AudioCodec() *Codec {
	return (*Codec)(unsafe.Pointer(ctx.audio_codec))
}

// SubtitleCodec ...
func (ctx *FormatContext) SubtitleCodec() *Codec {
	return (*Codec)(unsafe.Pointer(ctx.subtitle_codec))
}

// VideoCodec ...
func (ctx *FormatContext) VideoCodec() *Codec {
	return (*Codec)(unsafe.Pointer(ctx.video_codec))
}

// Metadata ...
func (ctx *FormatContext) Metadata() *AVDictionary {
	return (*AVDictionary)(unsafe.Pointer(ctx.metadata))
}

// Internal ...
func (ctx *FormatContext) Internal() *FormatInternal {
	return (*FormatInternal)(unsafe.Pointer(ctx.internal))
}

// Pb ...
func (ctx *FormatContext) Pb() *IOContext {
	return (*IOContext)(unsafe.Pointer(ctx.pb))
}

// InterruptCallback ...
func (ctx *FormatContext) InterruptCallback() IOInterruptCB {
	return IOInterruptCB(ctx.interrupt_callback)
}

// Programs ...
func (ctx *FormatContext) Programs() []*Program {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctx.programs)),
		Len:  int(ctx.NbPrograms()),
		Cap:  int(ctx.NbPrograms()),
	}

	return *((*[]*Program)(unsafe.Pointer(&header)))
}

// Streams ...
func (ctx *FormatContext) Streams() []*Stream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctx.streams)),
		Len:  int(ctx.NbStreams()),
		Cap:  int(ctx.NbStreams()),
	}

	return *((*[]*Stream)(unsafe.Pointer(&header)))
}

// Filename ...
func (ctx *FormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctx.filename[0])))
}

// AudioCodecID ...
func (ctx *FormatContext) AudioCodecID() CodecID {
	return CodecID(ctx.audio_codec_id)
}

// SubtitleCodecID ...
func (ctx *FormatContext) SubtitleCodecID() CodecID {
	return CodecID(ctx.subtitle_codec_id)
}

// VideoCodecID ...
func (ctx *FormatContext) VideoCodecID() CodecID {
	return CodecID(ctx.video_codec_id)
}

// DurationEstimationMethod ...
func (ctx *FormatContext) DurationEstimationMethod() DurationEstimationMethod {
	return DurationEstimationMethod(ctx.duration_estimation_method)
}

// AudioPreload ...
func (ctx *FormatContext) AudioPreload() int {
	return int(ctx.audio_preload)
}

// AvioFlags ...
func (ctx *FormatContext) AvioFlags() int {
	return int(ctx.avio_flags)
}

// AvoidNegativeTs ...
func (ctx *FormatContext) AvoidNegativeTs() int {
	return int(ctx.avoid_negative_ts)
}

// BitRate ...
func (ctx *FormatContext) BitRate() int {
	return int(ctx.bit_rate)
}

// CtxFlags ...
func (ctx *FormatContext) CtxFlags() int {
	return int(ctx.ctx_flags)
}

// Debug ...
func (ctx *FormatContext) Debug() int {
	return int(ctx.debug)
}

// ErrorRecognition ...
func (ctx *FormatContext) ErrorRecognition() int {
	return int(ctx.error_recognition)
}

// EventFlags ...
func (ctx *FormatContext) EventFlags() int {
	return int(ctx.event_flags)
}

// Flags ...
func (ctx *FormatContext) Flags() int {
	return int(ctx.flags)
}

// FlushPackets ...
func (ctx *FormatContext) FlushPackets() int {
	return int(ctx.flush_packets)
}

// FormatProbesize ...
func (ctx *FormatContext) FormatProbesize() int {
	return int(ctx.format_probesize)
}

// FpsProbeSize ...
func (ctx *FormatContext) FpsProbeSize() int {
	return int(ctx.fps_probe_size)
}

// IoRepositioned ...
func (ctx *FormatContext) IoRepositioned() int {
	return int(ctx.io_repositioned)
}

// Keylen ...
func (ctx *FormatContext) Keylen() int {
	return int(ctx.keylen)
}

// MaxChunkDuration ...
func (ctx *FormatContext) MaxChunkDuration() int {
	return int(ctx.max_chunk_duration)
}

// MaxChunkSize ...
func (ctx *FormatContext) MaxChunkSize() int {
	return int(ctx.max_chunk_size)
}

// MaxDelay ...
func (ctx *FormatContext) MaxDelay() int {
	return int(ctx.max_delay)
}

// MaxTsProbe ...
func (ctx *FormatContext) MaxTsProbe() int {
	return int(ctx.max_ts_probe)
}

// MetadataHeaderPadding ...
func (ctx *FormatContext) MetadataHeaderPadding() int {
	return int(ctx.metadata_header_padding)
}

// ProbeScore ...
func (ctx *FormatContext) ProbeScore() int {
	return int(ctx.probe_score)
}

// Seek2any ...
func (ctx *FormatContext) Seek2any() int {
	return int(ctx.seek2any)
}

// StrictStdCompliance ...
func (ctx *FormatContext) StrictStdCompliance() int {
	return int(ctx.strict_std_compliance)
}

// TsID ...
func (ctx *FormatContext) TsID() int {
	return int(ctx.ts_id)
}

// UseWallclockAsTimestamps ...
func (ctx *FormatContext) UseWallclockAsTimestamps() int {
	return int(ctx.use_wallclock_as_timestamps)
}

// Duration ...
func (ctx *FormatContext) Duration() int64 {
	return int64(ctx.duration)
}

// MaxAnalyzeDuration2 ...
func (ctx *FormatContext) MaxAnalyzeDuration2() int64 {
	return int64(ctx.max_analyze_duration)
}

// MaxInterleaveDelta ...
func (ctx *FormatContext) MaxInterleaveDelta() int64 {
	return int64(ctx.max_interleave_delta)
}

// OutputTsOffset ...
func (ctx *FormatContext) OutputTsOffset() int64 {
	return int64(ctx.output_ts_offset)
}

// Probesize2 ...
func (ctx *FormatContext) Probesize2() int64 {
	return int64(ctx.probesize)
}

// SkipInitialBytes ...
func (ctx *FormatContext) SkipInitialBytes() int64 {
	return int64(ctx.skip_initial_bytes)
}

// StartTime ...
func (ctx *FormatContext) StartTime() int64 {
	return int64(ctx.start_time)
}

// StartTimeRealtime ...
func (ctx *FormatContext) StartTimeRealtime() int64 {
	return int64(ctx.start_time_realtime)
}

// Iformat ...
func (ctx *FormatContext) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(ctx.iformat))
}

// Oformat ...
func (ctx *FormatContext) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(ctx.oformat))
}

// CorrectTsOverflow ...
func (ctx *FormatContext) CorrectTsOverflow() int {
	return int(ctx.correct_ts_overflow)
}

// MaxIndexSize ...
func (ctx *FormatContext) MaxIndexSize() uint {
	return uint(ctx.max_index_size)
}

// MaxPictureBuffer ...
func (ctx *FormatContext) MaxPictureBuffer() uint {
	return uint(ctx.max_picture_buffer)
}

// NbChapters ...
func (ctx *FormatContext) NbChapters() uint {
	return uint(ctx.nb_chapters)
}

// NbPrograms ...
func (ctx *FormatContext) NbPrograms() uint {
	return uint(ctx.nb_programs)
}

// NbStreams ...
func (ctx *FormatContext) NbStreams() uint {
	return uint(ctx.nb_streams)
}

// PacketSize ...
func (ctx *FormatContext) PacketSize() uint {
	return uint(ctx.packet_size)
}

// Probesize ...
func (ctx *FormatContext) Probesize() uint {
	return uint(ctx.probesize)
}

// SetPb ...
func (ctx *FormatContext) SetPb(pb *IOContext) {
	ctx.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

// Pb2 ...
func (ctx *FormatContext) Pb2() **IOContext {
	return (**IOContext)(unsafe.Pointer(&ctx.pb))
}
