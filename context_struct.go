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
func (ctx *AVCodecContext) ActiveThreadType() int {
	return int(ctx.active_thread_type)
}

// BFrameStrategy ...
func (ctx *AVCodecContext) BFrameStrategy() int {
	return int(ctx.b_frame_strategy)
}

// BQuantFactor ...
func (ctx *AVCodecContext) BQuantFactor() float64 {
	return float64(ctx.b_quant_factor)
}

// BQuantOffset ...
func (ctx *AVCodecContext) BQuantOffset() float64 {
	return float64(ctx.b_quant_offset)
}

// BSensitivity ...
func (ctx *AVCodecContext) BSensitivity() int {
	return int(ctx.b_sensitivity)
}

// BidirRefine ...
func (ctx *AVCodecContext) BidirRefine() int {
	return int(ctx.bidir_refine)
}

// BitRate ...
func (ctx *AVCodecContext) BitRate() int {
	return int(ctx.bit_rate)
}

// BitRateTolerance ...
func (ctx *AVCodecContext) BitRateTolerance() int {
	return int(ctx.bit_rate_tolerance)
}

// BitsPerCodedSample ...
func (ctx *AVCodecContext) BitsPerCodedSample() int {
	return int(ctx.bits_per_coded_sample)
}

// BitsPerRawSample ...
func (ctx *AVCodecContext) BitsPerRawSample() int {
	return int(ctx.bits_per_raw_sample)
}

// BlockAlign ...
func (ctx *AVCodecContext) BlockAlign() int {
	return int(ctx.block_align)
}

// BrdScale ...
func (ctx *AVCodecContext) BrdScale() int {
	return int(ctx.brd_scale)
}

// Channels ...
func (ctx *AVCodecContext) Channels() int {
	return int(ctx.channels)
}

// Chromaoffset ...
func (ctx *AVCodecContext) Chromaoffset() int {
	return int(ctx.chromaoffset)
}

// CodedHeight ...
func (ctx *AVCodecContext) CodedHeight() int {
	return int(ctx.coded_height)
}

// CodedWidth ...
func (ctx *AVCodecContext) CodedWidth() int {
	return int(ctx.coded_width)
}

// CoderType ...
func (ctx *AVCodecContext) CoderType() int {
	return int(ctx.coder_type)
}

// CompressionLevel ...
func (ctx *AVCodecContext) CompressionLevel() int {
	return int(ctx.compression_level)
}

// ContextModel ...
func (ctx *AVCodecContext) ContextModel() int {
	return int(ctx.context_model)
}

// Cutoff ...
func (ctx *AVCodecContext) Cutoff() int {
	return int(ctx.cutoff)
}

// DarkMasking ...
func (ctx *AVCodecContext) DarkMasking() float64 {
	return float64(ctx.dark_masking)
}

// DctAlgo ...
func (ctx *AVCodecContext) DctAlgo() int {
	return int(ctx.dct_algo)
}

// Debug ...
func (ctx *AVCodecContext) Debug() int {
	return int(ctx.debug)
}

// DebugMv ...
func (ctx *AVCodecContext) DebugMv() int {
	return int(ctx.debug_mv)
}

// Delay ...
func (ctx *AVCodecContext) Delay() int {
	return int(ctx.delay)
}

// DiaSize ...
func (ctx *AVCodecContext) DiaSize() int {
	return int(ctx.dia_size)
}

// ErrRecognition ...
func (ctx *AVCodecContext) ErrRecognition() int {
	return int(ctx.err_recognition)
}

// ErrorConcealment ...
func (ctx *AVCodecContext) ErrorConcealment() int {
	return int(ctx.error_concealment)
}

// ExtradataSize ...
func (ctx *AVCodecContext) ExtradataSize() int {
	return int(ctx.extradata_size)
}

// Flags ...
func (ctx *AVCodecContext) Flags() int {
	return int(ctx.flags)
}

// Flags2 ...
func (ctx *AVCodecContext) Flags2() int {
	return int(ctx.flags2)
}

// FrameBits ...
func (ctx *AVCodecContext) FrameBits() int {
	return int(ctx.frame_bits)
}

// FrameNumber ...
func (ctx *AVCodecContext) FrameNumber() int {
	return int(ctx.frame_number)
}

// FrameSize ...
func (ctx *AVCodecContext) FrameSize() int {
	return int(ctx.frame_size)
}

// FrameSkipCmp ...
func (ctx *AVCodecContext) FrameSkipCmp() int {
	return int(ctx.frame_skip_cmp)
}

// FrameSkipExp ...
func (ctx *AVCodecContext) FrameSkipExp() int {
	return int(ctx.frame_skip_exp)
}

// FrameSkipFactor ...
func (ctx *AVCodecContext) FrameSkipFactor() int {
	return int(ctx.frame_skip_factor)
}

// FrameSkipThreshold ...
func (ctx *AVCodecContext) FrameSkipThreshold() int {
	return int(ctx.frame_skip_threshold)
}

// GlobalQuality ...
func (ctx *AVCodecContext) GlobalQuality() int {
	return int(ctx.global_quality)
}

// GopSize ...
func (ctx *AVCodecContext) GopSize() int {
	return int(ctx.gop_size)
}

// HasBFrames ...
func (ctx *AVCodecContext) HasBFrames() int {
	return int(ctx.has_b_frames)
}

// HeaderBits ...
func (ctx *AVCodecContext) HeaderBits() int {
	return int(ctx.header_bits)
}

// Height ...
func (ctx *AVCodecContext) Height() int {
	return int(ctx.height)
}

// ICount ...
func (ctx *AVCodecContext) ICount() int {
	return int(ctx.i_count)
}

// IQuantFactor ...
func (ctx *AVCodecContext) IQuantFactor() float64 {
	return float64(ctx.i_quant_factor)
}

// IQuantOffset ...
func (ctx *AVCodecContext) IQuantOffset() float64 {
	return float64(ctx.i_quant_offset)
}

// ITexBits ...
func (ctx *AVCodecContext) ITexBits() int {
	return int(ctx.i_tex_bits)
}

// IdctAlgo ...
func (ctx *AVCodecContext) IdctAlgo() int {
	return int(ctx.idct_algo)
}

// IldctCmp ...
func (ctx *AVCodecContext) IldctCmp() int {
	return int(ctx.ildct_cmp)
}

// IntraDcPrecision ...
func (ctx *AVCodecContext) IntraDcPrecision() int {
	return int(ctx.intra_dc_precision)
}

// KeyintMin ...
func (ctx *AVCodecContext) KeyintMin() int {
	return int(ctx.keyint_min)
}

// LastPredictorCount ...
func (ctx *AVCodecContext) LastPredictorCount() int {
	return int(ctx.last_predictor_count)
}

// Level ...
func (ctx *AVCodecContext) Level() int {
	return int(ctx.level)
}

// LogLevelOffset ...
func (ctx *AVCodecContext) LogLevelOffset() int {
	return int(ctx.log_level_offset)
}

// Lowres ...
func (ctx *AVCodecContext) Lowres() int {
	return int(ctx.lowres)
}

// LumiMasking ...
func (ctx *AVCodecContext) LumiMasking() float64 {
	return float64(ctx.lumi_masking)
}

// MaxBFrames ...
func (ctx *AVCodecContext) MaxBFrames() int {
	return int(ctx.max_b_frames)
}

// MaxPredictionOrder ...
func (ctx *AVCodecContext) MaxPredictionOrder() int {
	return int(ctx.max_prediction_order)
}

// MaxQdiff ...
func (ctx *AVCodecContext) MaxQdiff() int {
	return int(ctx.max_qdiff)
}

// MbCmp ...
func (ctx *AVCodecContext) MbCmp() int {
	return int(ctx.mb_cmp)
}

// MbDecision ...
func (ctx *AVCodecContext) MbDecision() int {
	return int(ctx.mb_decision)
}

// MbLmax ...
func (ctx *AVCodecContext) MbLmax() int {
	return int(ctx.mb_lmax)
}

// MbLmin ...
func (ctx *AVCodecContext) MbLmin() int {
	return int(ctx.mb_lmin)
}

// MeCmp ...
func (ctx *AVCodecContext) MeCmp() int {
	return int(ctx.me_cmp)
}

// MePenaltyCompensation ...
func (ctx *AVCodecContext) MePenaltyCompensation() int {
	return int(ctx.me_penalty_compensation)
}

// MePreCmp ...
func (ctx *AVCodecContext) MePreCmp() int {
	return int(ctx.me_pre_cmp)
}

// MeRange ...
func (ctx *AVCodecContext) MeRange() int {
	return int(ctx.me_range)
}

// MeSubCmp ...
func (ctx *AVCodecContext) MeSubCmp() int {
	return int(ctx.me_sub_cmp)
}

// MeSubpelQuality ...
func (ctx *AVCodecContext) MeSubpelQuality() int {
	return int(ctx.me_subpel_quality)
}

// MinPredictionOrder ...
func (ctx *AVCodecContext) MinPredictionOrder() int {
	return int(ctx.min_prediction_order)
}

// MiscBits ...
func (ctx *AVCodecContext) MiscBits() int {
	return int(ctx.misc_bits)
}

// MpegQuant ...
func (ctx *AVCodecContext) MpegQuant() int {
	return int(ctx.mpeg_quant)
}

// Mv0Threshold ...
func (ctx *AVCodecContext) Mv0Threshold() int {
	return int(ctx.mv0_threshold)
}

// MvBits ...
func (ctx *AVCodecContext) MvBits() int {
	return int(ctx.mv_bits)
}

// NoiseReduction ...
func (ctx *AVCodecContext) NoiseReduction() int {
	return int(ctx.noise_reduction)
}

// NsseWeight ...
func (ctx *AVCodecContext) NsseWeight() int {
	return int(ctx.nsse_weight)
}

// PCount ...
func (ctx *AVCodecContext) PCount() int {
	return int(ctx.p_count)
}

// PMasking ...
func (ctx *AVCodecContext) PMasking() float64 {
	return float64(ctx.p_masking)
}

// PTexBits ...
func (ctx *AVCodecContext) PTexBits() int {
	return int(ctx.p_tex_bits)
}

// PreDiaSize ...
func (ctx *AVCodecContext) PreDiaSize() int {
	return int(ctx.pre_dia_size)
}

// PreMe ...
func (ctx *AVCodecContext) PreMe() int {
	return int(ctx.pre_me)
}

// PredictionMethod ...
func (ctx *AVCodecContext) PredictionMethod() int {
	return int(ctx.prediction_method)
}

// Profile ...
func (ctx *AVCodecContext) Profile() int {
	return int(ctx.profile)
}

// Qblur ...
func (ctx *AVCodecContext) Qblur() float64 {
	return float64(ctx.qblur)
}

// Qcompress ...
func (ctx *AVCodecContext) Qcompress() float64 {
	return float64(ctx.qcompress)
}

// Qmax ...
func (ctx *AVCodecContext) Qmax() int {
	return int(ctx.qmax)
}

// Qmin ...
func (ctx *AVCodecContext) Qmin() int {
	return int(ctx.qmin)
}

// RcBufferSize ...
func (ctx *AVCodecContext) RcBufferSize() int {
	return int(ctx.rc_buffer_size)
}

// RcInitialBufferOccupancy ...
func (ctx *AVCodecContext) RcInitialBufferOccupancy() int {
	return int(ctx.rc_initial_buffer_occupancy)
}

// RcMaxAvailableVbvUse ...
func (ctx *AVCodecContext) RcMaxAvailableVbvUse() float64 {
	return float64(ctx.rc_max_available_vbv_use)
}

// RcMaxRate ...
func (ctx *AVCodecContext) RcMaxRate() int {
	return int(ctx.rc_max_rate)
}

// RcMinRate ...
func (ctx *AVCodecContext) RcMinRate() int {
	return int(ctx.rc_min_rate)
}

// RcMinVbvOverflowUse ...
func (ctx *AVCodecContext) RcMinVbvOverflowUse() float64 {
	return float64(ctx.rc_min_vbv_overflow_use)
}

// RcOverrideCount ...
func (ctx *AVCodecContext) RcOverrideCount() int {
	return int(ctx.rc_override_count)
}

// RefcountedFrames ...
func (ctx *AVCodecContext) RefcountedFrames() int {
	return int(ctx.refcounted_frames)
}

// Refs ...
func (ctx *AVCodecContext) Refs() int {
	return int(ctx.refs)
}

// RtpPayloadSize ...
func (ctx *AVCodecContext) RtpPayloadSize() int {
	return int(ctx.rtp_payload_size)
}

// SampleRate ...
func (ctx *AVCodecContext) SampleRate() int {
	return int(ctx.sample_rate)
}

// ScenechangeThreshold ...
func (ctx *AVCodecContext) ScenechangeThreshold() int {
	return int(ctx.scenechange_threshold)
}

// SeekPreroll ...
func (ctx *AVCodecContext) SeekPreroll() int {
	return int(ctx.seek_preroll)
}

// SideDataOnlyPackets ...
func (ctx *AVCodecContext) SideDataOnlyPackets() int {
	return int(ctx.side_data_only_packets)
}

// SkipAlpha ...
func (ctx *AVCodecContext) SkipAlpha() int {
	return int(ctx.skip_alpha)
}

// SkipBottom ...
func (ctx *AVCodecContext) SkipBottom() int {
	return int(ctx.skip_bottom)
}

// SkipCount ...
func (ctx *AVCodecContext) SkipCount() int {
	return int(ctx.skip_count)
}

// SkipTop ...
func (ctx *AVCodecContext) SkipTop() int {
	return int(ctx.skip_top)
}

// SliceCount ...
func (ctx *AVCodecContext) SliceCount() int {
	return int(ctx.slice_count)
}

// SliceFlags ...
func (ctx *AVCodecContext) SliceFlags() int {
	return int(ctx.slice_flags)
}

// Slices ...
func (ctx *AVCodecContext) Slices() int {
	return int(ctx.slices)
}

// SpatialCplxMasking ...
func (ctx *AVCodecContext) SpatialCplxMasking() float64 {
	return float64(ctx.spatial_cplx_masking)
}

// StrictStdCompliance ...
func (ctx *AVCodecContext) StrictStdCompliance() int {
	return int(ctx.strict_std_compliance)
}

// SubCharencMode ...
func (ctx *AVCodecContext) SubCharencMode() int {
	return int(ctx.sub_charenc_mode)
}

// SubtitleHeaderSize ...
func (ctx *AVCodecContext) SubtitleHeaderSize() int {
	return int(ctx.subtitle_header_size)
}

// TemporalCplxMasking ...
func (ctx *AVCodecContext) TemporalCplxMasking() float64 {
	return float64(ctx.temporal_cplx_masking)
}

// ThreadCount ...
func (ctx *AVCodecContext) ThreadCount() int {
	return int(ctx.thread_count)
}

// ThreadSafeCallbacks ...
func (ctx *AVCodecContext) ThreadSafeCallbacks() int {
	return int(ctx.thread_safe_callbacks)
}

// ThreadType ...
func (ctx *AVCodecContext) ThreadType() int {
	return int(ctx.thread_type)
}

// TicksPerFrame ...
func (ctx *AVCodecContext) TicksPerFrame() int {
	return int(ctx.ticks_per_frame)
}

// Trellis ...
func (ctx *AVCodecContext) Trellis() int {
	return int(ctx.trellis)
}

// Width ...
func (ctx *AVCodecContext) Width() int {
	return int(ctx.width)
}

// WorkaroundBugs ...
func (ctx *AVCodecContext) WorkaroundBugs() int {
	return int(ctx.workaround_bugs)
}

// AudioServiceType ...
func (ctx *AVCodecContext) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(ctx.audio_service_type)
}

// ChromaSampleLocation ...
func (ctx *AVCodecContext) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(ctx.chroma_sample_location)
}

// CodecDescriptor ...
func (ctx *AVCodecContext) CodecDescriptor() *Descriptor {
	return (*Descriptor)(ctx.codec_descriptor)
}

// CodecID ...
func (ctx *AVCodecContext) CodecID() CodecID {
	return (CodecID)(ctx.codec_id)
}

// CodecType ...
func (ctx *AVCodecContext) CodecType() MediaType {
	return (MediaType)(ctx.codec_type)
}

// ColorPrimaries ...
func (ctx *AVCodecContext) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(ctx.color_primaries)
}

// ColorRange ...
func (ctx *AVCodecContext) ColorRange() AvColorRange {
	return (AvColorRange)(ctx.color_range)
}

// ColorTrc ...
func (ctx *AVCodecContext) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(ctx.color_trc)
}

// Colorspace ...
func (ctx *AVCodecContext) Colorspace() AvColorSpace {
	return (AvColorSpace)(ctx.colorspace)
}

// FieldOrder ...
func (ctx *AVCodecContext) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(ctx.field_order)
}

// PixFmt ...
func (ctx *AVCodecContext) PixFmt() PixelFormat {
	return (PixelFormat)(ctx.pix_fmt)
}

// RequestSampleFmt ...
func (ctx *AVCodecContext) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.request_sample_fmt)
}

// SampleFmt ...
func (ctx *AVCodecContext) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.sample_fmt)
}

// SkipFrame ...
func (ctx *AVCodecContext) SkipFrame() AvDiscard {
	return (AvDiscard)(ctx.skip_frame)
}

// SkipIdct ...
func (ctx *AVCodecContext) SkipIdct() AvDiscard {
	return (AvDiscard)(ctx.skip_idct)
}

// SkipLoopFilter ...
func (ctx *AVCodecContext) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(ctx.skip_loop_filter)
}

// Chapters ...
func (ctx *AVFormatContext) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(ctx.chapters))
}

// AudioCodec ...
func (ctx *AVFormatContext) AudioCodec() *AVCodec {
	return (*AVCodec)(unsafe.Pointer(ctx.audio_codec))
}

// SubtitleCodec ...
func (ctx *AVFormatContext) SubtitleCodec() *AVCodec {
	return (*AVCodec)(unsafe.Pointer(ctx.subtitle_codec))
}

// VideoCodec ...
func (ctx *AVFormatContext) VideoCodec() *AVCodec {
	return (*AvCodec)(unsafe.Pointer(ctx.video_codec))
}

// Metadata ...
func (ctx *AVFormatContext) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(ctx.metadata))
}

// Internal ...
func (ctx *AVFormatContext) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(ctx.internal))
}

// Pb ...
func (ctx *AVFormatContext) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(ctx.pb))
}

// InterruptCallback ...
func (ctx *AVFormatContext) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(ctx.interrupt_callback)
}

// Programs ...
func (ctx *AVFormatContext) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctx.programs)),
		Len:  int(ctx.NbPrograms()),
		Cap:  int(ctx.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

// Streams ...
func (ctx *AVFormatContext) Streams() []*Stream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctx.streams)),
		Len:  int(ctx.NbStreams()),
		Cap:  int(ctx.NbStreams()),
	}

	return *((*[]*Stream)(unsafe.Pointer(&header)))
}

// Filename ...
func (ctx *AVFormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctx.filename[0])))
}

// AudioCodecID ...
func (ctx *AVFormatContext) AudioCodecID() CodecID {
	return CodecID(ctx.audio_codec_id)
}

// SubtitleCodecID ...
func (ctx *AVFormatContext) SubtitleCodecID() CodecID {
	return CodecID(ctx.subtitle_codec_id)
}

// VideoCodecID ...
func (ctx *AVFormatContext) VideoCodecID() CodecID {
	return CodecID(ctx.video_codec_id)
}

// DurationEstimationMethod ...
func (ctx *AVFormatContext) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(ctx.duration_estimation_method)
}

// AudioPreload ...
func (ctx *AVFormatContext) AudioPreload() int {
	return int(ctx.audio_preload)
}

// AvioFlags ...
func (ctx *AVFormatContext) AvioFlags() int {
	return int(ctx.avio_flags)
}

// AvoidNegativeTs ...
func (ctx *AVFormatContext) AvoidNegativeTs() int {
	return int(ctx.avoid_negative_ts)
}

// BitRate ...
func (ctx *AVFormatContext) BitRate() int {
	return int(ctx.bit_rate)
}

// CtxFlags ...
func (ctx *AVFormatContext) CtxFlags() int {
	return int(ctx.ctx_flags)
}

// Debug ...
func (ctx *AVFormatContext) Debug() int {
	return int(ctx.debug)
}

// ErrorRecognition ...
func (ctx *AVFormatContext) ErrorRecognition() int {
	return int(ctx.error_recognition)
}

// EventFlags ...
func (ctx *AVFormatContext) EventFlags() int {
	return int(ctx.event_flags)
}

// Flags ...
func (ctx *AVFormatContext) Flags() int {
	return int(ctx.flags)
}

// FlushPackets ...
func (ctx *AVFormatContext) FlushPackets() int {
	return int(ctx.flush_packets)
}

// FormatProbesize ...
func (ctx *AVFormatContext) FormatProbesize() int {
	return int(ctx.format_probesize)
}

// FpsProbeSize ...
func (ctx *AVFormatContext) FpsProbeSize() int {
	return int(ctx.fps_probe_size)
}

// IoRepositioned ...
func (ctx *AVFormatContext) IoRepositioned() int {
	return int(ctx.io_repositioned)
}

// Keylen ...
func (ctx *AVFormatContext) Keylen() int {
	return int(ctx.keylen)
}

// MaxChunkDuration ...
func (ctx *AVFormatContext) MaxChunkDuration() int {
	return int(ctx.max_chunk_duration)
}

// MaxChunkSize ...
func (ctx *AVFormatContext) MaxChunkSize() int {
	return int(ctx.max_chunk_size)
}

// MaxDelay ...
func (ctx *AVFormatContext) MaxDelay() int {
	return int(ctx.max_delay)
}

// MaxTsProbe ...
func (ctx *AVFormatContext) MaxTsProbe() int {
	return int(ctx.max_ts_probe)
}

// MetadataHeaderPadding ...
func (ctx *AVFormatContext) MetadataHeaderPadding() int {
	return int(ctx.metadata_header_padding)
}

// ProbeScore ...
func (ctx *AVFormatContext) ProbeScore() int {
	return int(ctx.probe_score)
}

// Seek2any ...
func (ctx *AVFormatContext) Seek2any() int {
	return int(ctx.seek2any)
}

// StrictStdCompliance ...
func (ctx *AVFormatContext) StrictStdCompliance() int {
	return int(ctx.strict_std_compliance)
}

// TsID ...
func (ctx *AVFormatContext) TsID() int {
	return int(ctx.ts_id)
}

// UseWallclockAsTimestamps ...
func (ctx *AVFormatContext) UseWallclockAsTimestamps() int {
	return int(ctx.use_wallclock_as_timestamps)
}

// Duration ...
func (ctx *AVFormatContext) Duration() int64 {
	return int64(ctx.duration)
}

// MaxAnalyzeDuration2 ...
func (ctx *AVFormatContext) MaxAnalyzeDuration2() int64 {
	return int64(ctx.max_analyze_duration)
}

// MaxInterleaveDelta ...
func (ctx *AVFormatContext) MaxInterleaveDelta() int64 {
	return int64(ctx.max_interleave_delta)
}

// OutputTsOffset ...
func (ctx *AVFormatContext) OutputTsOffset() int64 {
	return int64(ctx.output_ts_offset)
}

// Probesize2 ...
func (ctx *AVFormatContext) Probesize2() int64 {
	return int64(ctx.probesize)
}

// SkipInitialBytes ...
func (ctx *AVFormatContext) SkipInitialBytes() int64 {
	return int64(ctx.skip_initial_bytes)
}

// StartTime ...
func (ctx *AVFormatContext) StartTime() int64 {
	return int64(ctx.start_time)
}

// StartTimeRealtime ...
func (ctx *AVFormatContext) StartTimeRealtime() int64 {
	return int64(ctx.start_time_realtime)
}

// Iformat ...
func (ctx *AVFormatContext) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(ctx.iformat))
}

// Oformat ...
func (ctx *AVFormatContext) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(ctx.oformat))
}

// CorrectTsOverflow ...
func (ctx *AVFormatContext) CorrectTsOverflow() int {
	return int(ctx.correct_ts_overflow)
}

// MaxIndexSize ...
func (ctx *AVFormatContext) MaxIndexSize() uint {
	return uint(ctx.max_index_size)
}

// MaxPictureBuffer ...
func (ctx *AVFormatContext) MaxPictureBuffer() uint {
	return uint(ctx.max_picture_buffer)
}

// NbChapters ...
func (ctx *AVFormatContext) NbChapters() uint {
	return uint(ctx.nb_chapters)
}

// NbPrograms ...
func (ctx *AVFormatContext) NbPrograms() uint {
	return uint(ctx.nb_programs)
}

// NbStreams ...
func (ctx *AVFormatContext) NbStreams() uint {
	return uint(ctx.nb_streams)
}

// PacketSize ...
func (ctx *AVFormatContext) PacketSize() uint {
	return uint(ctx.packet_size)
}

// Probesize ...
func (ctx *AVFormatContext) Probesize() uint {
	return uint(ctx.probesize)
}

// SetPb ...
func (ctx *AVFormatContext) SetPb(pb *AvIOContext) {
	ctx.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

// Pb2 ...
func (ctx *AVFormatContext) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&ctx.pb))
}
