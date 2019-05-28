// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat libavcodec libswresample libswscale libavutil
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libswscale/swscale.h>
//#include <libswresample/swresample.h>
import "C"
import (
	"time"
	"unsafe"
)

// AVCodecGetPktTimebase ...
func (ctx *CodecContext) AVCodecGetPktTimebase() Rational {
	return Rational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctx)))
}

// AVCodecGetPktTimebase2 returns the timebase rational number as numerator and denominator
func (ctx *CodecContext) AVCodecGetPktTimebase2() Rational {
	return ctx.AVCodecGetPktTimebase()
}

// AVCodecSetPktTimebase ...
func (ctx *CodecContext) AVCodecSetPktTimebase(r Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctx), (C.struct_AVRational)(r))
}

// AVCodecGetCodecDescriptor ...
func (ctx *CodecContext) AVCodecGetCodecDescriptor() *CodecDescriptor {
	return (*CodecDescriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctx)))
}

// AVCodecSetCodecDescriptor ...
func (ctx *CodecContext) AVCodecSetCodecDescriptor(d *CodecDescriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodecDescriptor)(d))
}

// AVCodecGetLowres ...
func (ctx *CodecContext) AVCodecGetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctx)))
}

// AVCodecSetLowres ...
func (ctx *CodecContext) AVCodecSetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctx), C.int(i))
}

// AVCodecGetSeekPreroll ...
func (ctx *CodecContext) AVCodecGetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctx)))
}

// AVCodecSetSeekPreroll ...
func (ctx *CodecContext) AVCodecSetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctx), C.int(i))
}

// AVCodecGetChromaIntraMatrix ...
func (ctx *CodecContext) AVCodecGetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctx)))
}

// AVCodecSetChromaIntraMatrix ...
func (ctx *CodecContext) AVCodecSetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctx), (*C.uint16_t)(t))
}

//AVCodecFreeContext Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctx *CodecContext) AVCodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(ctx)))
}

//AVCodecGetContextDefaults3 Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctx *CodecContext) AVCodecGetContextDefaults3(codec *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodec)(codec)))
}

//AVCodecCopyContext Copy the settings of the source Context into the destination Context.
func (ctx *CodecContext) AVCodecCopyContext(ctxt2 *CodecContext) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodecContext)(ctxt2)))
}

//AVCodecOpen2 Initialize the Context to use the given Codec
func (ctx *CodecContext) AVCodecOpen2(codec *Codec, d **AVDictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodec)(codec), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AVCodecClose Close a given Context and free all the data associated with it (but not the Context itself).
func (ctx *CodecContext) AVCodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctx)))
}

//AVCodecDefaultGetBuffer2 The default callback for Context.get_buffer2().
func (ctx *CodecContext) AVCodecDefaultGetBuffer2(f *AVFrame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), C.int(l)))
}

//AVCodecAlignDimensions Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctx *CodecContext) AVCodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctx), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//AVCodecAlignDimensions2 Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctx *CodecContext) AVCodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctx), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//AVCodecDecodeAudio4 Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (ctx *CodecContext) AVCodecDecodeAudio4(f *AVFrame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//AVCodecDecodeVideo2 Decode the video frame of size avpkt->size from avpkt->data into picture.
func (ctx *CodecContext) AVCodecDecodeVideo2(f *AVFrame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//AVCodecDecodeSubtitle2 Decode a subtitle message.
func (ctx *CodecContext) AVCodecDecodeSubtitle2(s *Subtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//AVCodecEncodeAudio2 Encode a frame of audio.
func (ctx *CodecContext) AVCodecEncodeAudio2(p *Packet, f *AVFrame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//AVCodecEncodeVideo2 Encode a frame of video
func (ctx *CodecContext) AVCodecEncodeVideo2(p *Packet, f *AVFrame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

// AVCodecEncodeSubtitle ...
func (ctx *CodecContext) AVCodecEncodeSubtitle(b *uint8, bs int, s *Subtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctx), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

// AVCodecDefaultGetFormat ...
func (ctx *CodecContext) AVCodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctx), (*C.enum_AVPixelFormat)(f)))
}

//AVCodecFlushBuffers Reset the internal decoder state / flush internal buffers.
func (ctx *CodecContext) AVCodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctx))
}

//AVGetAudioFrameDuration Return audio frame duration.
func (ctx *CodecContext) AVGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctx), C.int(f)))
}

// AVCodecIsOpen ...
func (ctx *CodecContext) AVCodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctx)))
}

//AVParserParse2 Parse a packet.
func (ctx *CodecContext) AVParserParse2(ctxtp *CodecParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctx), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

// AvParserChange ...
func (ctx *CodecContext) AvParserChange(ctxtp *CodecParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctx), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

// AvParserInit ...
func AvParserInit(c int) *CodecParserContext {
	return (*CodecParserContext)(C.av_parser_init(C.int(c)))
}

// AvParserClose ...
func AvParserClose(ctxtp *CodecParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

// AvParserNext ...
func (p *CodecParser) AvParserNext() *CodecParser {
	return (*CodecParser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

// AvRegisterCodecParser ...
func (p *CodecParser) AvRegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}

// SetTimebase ...
func (ctx *CodecContext) SetTimebase(num1 int, den1 int) {
	ctx.time_base.num = C.int(num1)
	ctx.time_base.den = C.int(den1)
}

// SetEncodeParams2 ...
func (ctx *CodecContext) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	ctx.width = C.int(width)
	ctx.height = C.int(height)
	// c.bit_rate = 1000000
	ctx.gop_size = C.int(gopSize)
	// c.max_b_frames = 2
	if hasBframes {
		ctx.has_b_frames = 1
	} else {
		ctx.has_b_frames = 0
	}
	// c.extradata = nil
	// c.extradata_size = 0
	// c.channels = 0
	ctx.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(c.priv_data, "preset", "ultrafast", 0)
}

// SetEncodeParams ...
func (ctx *CodecContext) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	ctx.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

// AVCodecSendPacket ...
func (ctx *CodecContext) AVCodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(packet)))
}

// AVCodecReceiveFrame ...
func (ctx *CodecContext) AVCodecReceiveFrame(f *AVFrame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f)))
}

// AvseekFlagBackward ...
const (
	AvseekFlagBackward = 1 ///< seek backward
	AvseekFlagByte     = 2 ///< seeking based on position in bytes
	AvseekFlagAny      = 4 ///< seek to any frame, even non-keyframes
	AvseekFlagFrame    = 8 ///< seeking based on frame number
)

// AvFormatGetProbeScore ...
func (ctx *FormatContext) AvFormatGetProbeScore() int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(ctx)))
}

// AvFormatGetVideoCodec ...
func (ctx *FormatContext) AvFormatGetVideoCodec() *Codec {
	return (*Codec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(ctx)))
}

// AvFormatSetVideoCodec ...
func (ctx *FormatContext) AvFormatSetVideoCodec(c *Codec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c))
}

// AvFormatGetAudioCodec ...
func (ctx *FormatContext) AvFormatGetAudioCodec() *Codec {
	return (*Codec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(ctx)))
}

// AvFormatSetAudioCodec ...
func (ctx *FormatContext) AvFormatSetAudioCodec(c *Codec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c))
}

// AvFormatGetSubtitleCodec ...
func (ctx *FormatContext) AvFormatGetSubtitleCodec() *Codec {
	return (*Codec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(ctx)))
}

// AvFormatSetSubtitleCodec ...
func (ctx *FormatContext) AvFormatSetSubtitleCodec(c *Codec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c))
}

// AvFormatGetMetadataHeaderPadding ...
func (ctx *FormatContext) AvFormatGetMetadataHeaderPadding() int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(ctx)))
}

// AvFormatSetMetadataHeaderPadding ...
func (ctx *FormatContext) AvFormatSetMetadataHeaderPadding(c int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(ctx), C.int(c))
}

// AvFormatGetOpaque ...
func (ctx *FormatContext) AvFormatGetOpaque() {
	C.av_format_get_opaque((*C.struct_AVFormatContext)(ctx))
}

// AvFormatSetOpaque ...
func (ctx *FormatContext) AvFormatSetOpaque(o int) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(ctx), unsafe.Pointer(&o))
}

//AVFormatInjectGlobalSideData This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func (ctx *FormatContext) AVFormatInjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(ctx))
}

//AVFmtCtxGetDurationEstimationMethod Returns the method used to set ctx->duration.
func (ctx *FormatContext) AVFmtCtxGetDurationEstimationMethod() DurationEstimationMethod {
	return (DurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(ctx)))
}

//AVFormatFreeContext Free an Context and all its streams.
func (ctx *FormatContext) AVFormatFreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(ctx))
}

//AVFormatNewStream Add a new stream to a media file.
func (ctx *FormatContext) AVFormatNewStream(c *Codec) *Stream {
	return (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c)))
}

// AVNewProgram ...
func (ctx *FormatContext) AVNewProgram(id int) *Program {
	return (*Program)(C.av_new_program((*C.struct_AVFormatContext)(ctx), C.int(id)))
}

//AVFormatFindStreamInfo Read packets of a media file to get stream information.
func (ctx *FormatContext) AVFormatFindStreamInfo(d **AVDictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AVFindProgramFromStream Find the programs which belong to a given stream.
func (ctx *FormatContext) AVFindProgramFromStream(l *Program, su int) *Program {
	return (*Program)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(ctx), (*C.struct_AVProgram)(l), C.int(su)))
}

//AVFindBestStream Find the "best" stream in the file.
func AVFindBestStream(ic *FormatContext, t MediaType, ws, rs int, c **Codec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

//AVReadFrame Return the next frame of a stream.
func (ctx *FormatContext) AVReadFrame(pkt *Packet) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)), toCPacket(pkt)))
}

//AVSeekFrame Seek to the keyframe at timestamp.
func (ctx *FormatContext) AVSeekFrame(st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(ctx), C.int(st), C.int64_t(t), C.int(f)))
}

// AVSeekFrameTime seeks to a specified time location.
// |timebase| is codec specific and can be obtained by calling AvCodecGetPktTimebase2
func (ctx *FormatContext) AVSeekFrameTime(st int, at time.Duration, timebase Rational) int {
	t2 := C.double(C.double(at.Seconds())*C.double(timebase.Den())) / (C.double(timebase.Num()))
	// log.Printf("Seeking to time :%v TimebaseTime:%v ActualTimebase:%v", at, t2, timebase)
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(ctx), C.int(st), C.int64_t(t2), AvseekFlagBackward))
}

//AVFormatSeekFile Seek to timestamp ts.
func (ctx *FormatContext) AVFormatSeekFile(si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(ctx), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

//AVReadPlay Start playing a network-based stream (e.g.
func (ctx *FormatContext) AVReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(ctx)))
}

//AVReadPause Pause a network-based stream (e.g.
func (ctx *FormatContext) AVReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(ctx)))
}

//AVFormatCloseInput Close an opened input Context.
func (ctx *FormatContext) AVFormatCloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)))
}

//AVFormatWriteHeader Allocate the stream private data and write the stream header to an output media file.
func (ctx *FormatContext) AVFormatWriteHeader(o **AVDictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//AVWriteFrame Write a packet to an output media file.
func (ctx *FormatContext) AVWriteFrame(pkt *Packet) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(ctx), toCPacket(pkt)))
}

//AVInterleavedWriteFrame Write a packet to an output media file ensuring correct interleaving.
func (ctx *FormatContext) AVInterleavedWriteFrame(pkt *Packet) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(ctx), toCPacket(pkt)))
}

//AVWriteUncodedFrame Write a uncoded frame to an output media file.
func (ctx *FormatContext) AVWriteUncodedFrame(si int, f *AVFrame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(ctx), C.int(si), (*C.struct_AVFrame)(f)))
}

//AVInterleavedWriteUncodedFrame Write a uncoded frame to an output media file.
func (ctx *FormatContext) AVInterleavedWriteUncodedFrame(si int, f *AVFrame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(ctx), C.int(si), (*C.struct_AVFrame)(f)))
}

//AVWriteUncodedFrameQuery Test whether a muxer supports uncoded frame.
func (ctx *FormatContext) AVWriteUncodedFrameQuery(si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(ctx), C.int(si)))
}

//AVWriteTrailer Write the stream trailer to an output media file and free the file private data.
func (ctx *FormatContext) AVWriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(ctx)))
}

//AVGetOutputTimestamp Get timing information for the data currently output.
func (ctx *FormatContext) AVGetOutputTimestamp(st int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(ctx), C.int(st), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

// AVFindDefaultStreamIndex ...
func (ctx *FormatContext) AVFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(ctx)))
}

//AVDumpFormat Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (ctx *FormatContext) AVDumpFormat(i int, url string, io int) {
	Curl := C.CString(url)
	defer C.free(unsafe.Pointer(Curl))

	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)), C.int(i), Curl, C.int(io))
}

//AVGuessSampleAspectRatio Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func (ctx *FormatContext) AVGuessSampleAspectRatio(st *Stream, f *AVFrame) Rational {
	return newRational(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(f)))
}

//AVGuessFrameRate Guess the frame rate, based on both the container and codec information.
func (ctx *FormatContext) AVGuessFrameRate(st *Stream, f *AVFrame) Rational {
	return newRational(C.av_guess_frame_rate((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(f)))
}

//AVFormatMatchStreamSpecifier Check if the stream st contained in s is matched by the stream specifier spec.
func (ctx *FormatContext) AVFormatMatchStreamSpecifier(st *Stream, spec string) int {
	Cspec := C.CString(spec)
	defer C.free(unsafe.Pointer(Cspec))

	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), Cspec))
}

// AVFormatQueueAttachedPictures ...
func (ctx *FormatContext) AVFormatQueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(ctx)))
}

// AVFormatNewStream2 ...
func (ctx *FormatContext) AVFormatNewStream2(c *Codec) *Stream {
	stream := (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c)))
	stream.codec.pix_fmt = int32(AvPixFmtYuv)
	stream.codec.width = 640
	stream.codec.height = 480
	stream.time_base.num = 1
	stream.time_base.num = 25
	return stream
}

//SwsAllocContext Allocate an empty Context.
func SwsAllocContext() *SwsContext {
	return (*SwsContext)(C.sws_alloc_context())
}

//SwsInitContext Initialize the swscaler context sws_context.
func SwsInitContext(ctx *SwsContext, sf, df *SwsFilter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctx), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

//SwsFreecontext Free the swscaler context swsContext.
func SwsFreecontext(ctx *SwsContext) {
	C.sws_freeContext((*C.struct_SwsContext)(ctx))
}

//SwsGetcontext Allocate and return an Context.
func SwsGetcontext(sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *SwsFilter, p *int) *SwsContext {
	return (*SwsContext)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(unsafe.Pointer(p))))
}

//SwsGetCachedContext Check if context can be reused, otherwise reallocate a new one.
func SwsGetCachedContext(ctx *SwsContext, sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *SwsFilter, p *float64) *SwsContext {
	return (*SwsContext)(C.sws_getCachedContext((*C.struct_SwsContext)(ctx), C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(p)))
}

//SwrInit Initialize context after user parameters have been set.
func (s *SwrContext) SwrInit() int {
	return int(C.swr_init((*C.struct_SwrContext)(s)))
}

//SwrIsInitialized Check whether an swr context has been initialized or not.
func (s *SwrContext) SwrIsInitialized() int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(s)))
}

//SwrAllocSetOpts Allocate SwrContext if needed and set/reset common parameters.
func (s *SwrContext) SwrAllocSetOpts(ocl int64, osf SampleFormat, osr int, icl int64, isf SampleFormat, isr, lo, lc int) *SwrContext {
	return (*SwrContext)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

//SwrFree SwrContext destructor functions. Free the given SwrContext and set the pointer to NULL.
func (s *SwrContext) SwrFree() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(s)))
}

//SwrClose Closes the context so that swr_is_initialized() returns 0.
func (s *SwrContext) SwrClose() {
	C.swr_close((*C.struct_SwrContext)(s))
}

//SwrConvert Core conversion functions. Convert audio
func (s *SwrContext) SwrConvert(out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(s), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

//SwrNextPts Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func (s *SwrContext) SwrNextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(s), C.int64_t(pts)))
}

//SwrSetCompensation Low-level option setting functions
//These functons provide a means to set low-level options that is not possible with the AvOption API.
//Activate resampling compensation ("soft" compensation).
func (s *SwrContext) SwrSetCompensation(sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(s), C.int(sd), C.int(cd)))
}

//SwrSetChannelMapping Set a customized input channel mapping.
func (s *SwrContext) SwrSetChannelMapping(cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(unsafe.Pointer(cm))))
}

//SwrSetMatrix Set a customized remix matrix.
func (s *SwrContext) SwrSetMatrix(m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

//SwrDropOutput Sample handling functions. Drops the specified number of output samples.
func (s *SwrContext) SwrDropOutput(c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(s), C.int(c)))
}

//SwrInjectSilence Injects the specified number of silence samples.
func (s *SwrContext) SwrInjectSilence(c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(s), C.int(c)))
}

//SwrGetDelay Gets the delay the next input sample will experience relative to the next output sample.
func (s *SwrContext) SwrGetDelay(b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(s), C.int64_t(b)))
}

//SwrConvertFrame SwrConvertFrame AVFrame based API. Convert the samples in the input AVFrame and write them to the output AVFrame.
func (s *SwrContext) SwrConvertFrame(o, i *AVFrame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

//SwrConfigFrame Configure or reconfigure the SwrContext using the information provided by the AvFrames.
func (s *SwrContext) SwrConfigFrame(o, i *AVFrame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}
