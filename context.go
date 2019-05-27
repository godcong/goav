// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat libavcodec libswscale libavutil
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libswscale/swscale.h>
import "C"
import (
	"time"
	"unsafe"
)

type (
	AVCodecContext C.struct_AVCodecContext
)

func (ctx *AVCodecContext) AvCodecGetPktTimebase() Rational {
	return Rational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctx)))
}

// AvCodecGetPktTimebase2 returns the timebase rational number as numerator and denominator
func (ctx *AVCodecContext) AvCodecGetPktTimebase2() Rational {
	return ctx.AvCodecGetPktTimebase()
}

func (ctx *AVCodecContext) AvCodecSetPktTimebase(r Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctx), (C.struct_AVRational)(r))
}

func (ctx *AVCodecContext) AvCodecGetCodecDescriptor() *Descriptor {
	return (*Descriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctx)))
}

func (ctx *AVCodecContext) AvCodecSetCodecDescriptor(d *Descriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodecDescriptor)(d))
}

func (ctx *AVCodecContext) AvCodecGetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctx)))
}

func (ctx *AVCodecContext) AvCodecSetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctx), C.int(i))
}

func (ctx *AVCodecContext) AvCodecGetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctx)))
}

func (ctx *AVCodecContext) AvCodecSetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctx), C.int(i))
}

func (ctx *AVCodecContext) AvCodecGetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctx)))
}

func (ctx *AVCodecContext) AvCodecSetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctx), (*C.uint16_t)(t))
}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctx *AVCodecContext) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(ctx)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctx *AVCodecContext) AvcodecGetContextDefaults3(codec *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodec)(codec)))
}

//Copy the settings of the source Context into the destination Context.
func (ctx *AVCodecContext) AvcodecCopyContext(ctxt2 *AVCodecContext) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodecContext)(ctxt2)))
}

//Initialize the Context to use the given Codec
func (ctx *AVCodecContext) AvcodecOpen2(codec *Codec, d **Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodec)(codec), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (ctx *AVCodecContext) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctx)))
}

//The default callback for Context.get_buffer2().
func (ctx *AVCodecContext) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctx *AVCodecContext) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctx), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctx *AVCodecContext) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctx), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (ctx *AVCodecContext) AvcodecDecodeAudio4(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
func (ctx *AVCodecContext) AvcodecDecodeVideo2(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode a subtitle message.
func (ctx *AVCodecContext) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Encode a frame of audio.
func (ctx *AVCodecContext) AvcodecEncodeAudio2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//Encode a frame of video
func (ctx *AVCodecContext) AvcodecEncodeVideo2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

func (ctx *AVCodecContext) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctx), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctx *AVCodecContext) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctx), (*C.enum_AVPixelFormat)(f)))
}

//Reset the internal decoder state / flush internal buffers.
func (ctx *AVCodecContext) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctx))
}

//Return audio frame duration.
func (ctx *AVCodecContext) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctx), C.int(f)))
}

func (ctx *AVCodecContext) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctx)))
}

//Parse a packet.
func (ctx *AVCodecContext) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctx), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (ctx *AVCodecContext) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctx), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (p *Parser) AvParserNext() *Parser {
	return (*Parser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

func (p *Parser) AvRegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}

func (ctx *AVCodecContext) SetTimebase(num1 int, den1 int) {
	ctx.time_base.num = C.int(num1)
	ctx.time_base.den = C.int(den1)
}

func (ctx *AVCodecContext) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
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

func (ctx *AVCodecContext) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	ctx.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

func (ctx *AVCodecContext) AvcodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(packet)))
}

func (ctx *AVCodecContext) AvcodecReceiveFrame(f *Frame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f)))
}

const (
	AvseekFlagBackward = 1 ///< seek backward
	AvseekFlagByte     = 2 ///< seeking based on position in bytes
	AvseekFlagAny      = 4 ///< seek to any frame, even non-keyframes
	AvseekFlagFrame    = 8 ///< seeking based on frame number
)

func (ctx *AVFormatContext) AvFormatGetProbeScore() int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(ctx)))
}

func (ctx *AVFormatContext) AvFormatGetVideoCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(ctx)))
}

func (ctx *AVFormatContext) AvFormatSetVideoCodec(c *AvCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c))
}

func (ctx *AVFormatContext) AvFormatGetAudioCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(ctx)))
}

func (ctx *AVFormatContext) AvFormatSetAudioCodec(c *AvCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c))
}

func (ctx *AVFormatContext) AvFormatGetSubtitleCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(ctx)))
}

func (ctx *AVFormatContext) AvFormatSetSubtitleCodec(c *AvCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c))
}

func (ctx *AVFormatContext) AvFormatGetMetadataHeaderPadding() int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(ctx)))
}

func (ctx *AVFormatContext) AvFormatSetMetadataHeaderPadding(c int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(ctx), C.int(c))
}

func (ctx *AVFormatContext) AvFormatGetOpaque() {
	C.av_format_get_opaque((*C.struct_AVFormatContext)(ctx))
}

func (ctx *AVFormatContext) AvFormatSetOpaque(o int) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(ctx), unsafe.Pointer(&o))
}

//This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func (ctx *AVFormatContext) AvFormatInjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(ctx))
}

//Returns the method used to set ctx->duration.
func (ctx *AVFormatContext) AvFmtCtxGetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(ctx)))
}

//Free an Context and all its streams.
func (ctx *AVFormatContext) AvformatFreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(ctx))
}

//Add a new stream to a media file.
func (ctx *AVFormatContext) AvformatNewStream(c *AvCodec) *Stream {
	return (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c)))
}

func (ctx *AVFormatContext) AvNewProgram(id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(ctx), C.int(id)))
}

//Read packets of a media file to get stream information.
func (ctx *AVFormatContext) AvformatFindStreamInfo(d **Dictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Find the programs which belong to a given stream.
func (ctx *AVFormatContext) AvFindProgramFromStream(l *AvProgram, su int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(ctx), (*C.struct_AVProgram)(l), C.int(su)))
}

//Find the "best" stream in the file.
func AvFindBestStream(ic *AVFormatContext, t MediaType, ws, rs int, c **AvCodec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

//Return the next frame of a stream.
func (ctx *AVFormatContext) AvReadFrame(pkt *Packet) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)), toCPacket(pkt)))
}

//Seek to the keyframe at timestamp.
func (ctx *AVFormatContext) AvSeekFrame(st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(ctx), C.int(st), C.int64_t(t), C.int(f)))
}

// AvSeekFrameTime seeks to a specified time location.
// |timebase| is codec specific and can be obtained by calling AvCodecGetPktTimebase2
func (ctx *AVFormatContext) AvSeekFrameTime(st int, at time.Duration, timebase Rational) int {
	t2 := C.double(C.double(at.Seconds())*C.double(timebase.Den())) / (C.double(timebase.Num()))
	// log.Printf("Seeking to time :%v TimebaseTime:%v ActualTimebase:%v", at, t2, timebase)
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(ctx), C.int(st), C.int64_t(t2), AvseekFlagBackward))
}

//Seek to timestamp ts.
func (ctx *AVFormatContext) AvformatSeekFile(si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(ctx), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

//Start playing a network-based stream (e.g.
func (ctx *AVFormatContext) AvReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(ctx)))
}

//Pause a network-based stream (e.g.
func (ctx *AVFormatContext) AvReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(ctx)))
}

//Close an opened input Context.
func (ctx *AVFormatContext) AvformatCloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)))
}

//Allocate the stream private data and write the stream header to an output media file.
func (ctx *AVFormatContext) AvformatWriteHeader(o **Dictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//Write a packet to an output media file.
func (ctx *AVFormatContext) AvWriteFrame(pkt *Packet) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(ctx), toCPacket(pkt)))
}

//Write a packet to an output media file ensuring correct interleaving.
func (ctx *AVFormatContext) AvInterleavedWriteFrame(pkt *Packet) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(ctx), toCPacket(pkt)))
}

//Write a uncoded frame to an output media file.
func (ctx *AVFormatContext) AvWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(ctx), C.int(si), (*C.struct_AVFrame)(f)))
}

//Write a uncoded frame to an output media file.
func (ctx *AVFormatContext) AvInterleavedWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(ctx), C.int(si), (*C.struct_AVFrame)(f)))
}

//Test whether a muxer supports uncoded frame.
func (ctx *AVFormatContext) AvWriteUncodedFrameQuery(si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(ctx), C.int(si)))
}

//Write the stream trailer to an output media file and free the file private data.
func (ctx *AVFormatContext) AvWriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(ctx)))
}

//Get timing information for the data currently output.
func (ctx *AVFormatContext) AvGetOutputTimestamp(st int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(ctx), C.int(st), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

func (ctx *AVFormatContext) AvFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(ctx)))
}

//Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (ctx *AVFormatContext) AvDumpFormat(i int, url string, io int) {
	Curl := C.CString(url)
	defer C.free(unsafe.Pointer(Curl))

	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)), C.int(i), Curl, C.int(io))
}

//Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func (ctx *AVFormatContext) AvGuessSampleAspectRatio(st *Stream, f *Frame) Rational {
	return newRational(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(f)))
}

//Guess the frame rate, based on both the container and codec information.
func (ctx *AVFormatContext) AvGuessFrameRate(st *Stream, f *Frame) Rational {
	return newRational(C.av_guess_frame_rate((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(f)))
}

//Check if the stream st contained in s is matched by the stream specifier spec.
func (ctx *AVFormatContext) AvformatMatchStreamSpecifier(st *Stream, spec string) int {
	Cspec := C.CString(spec)
	defer C.free(unsafe.Pointer(Cspec))

	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), Cspec))
}

func (ctx *AVFormatContext) AvformatQueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(ctx)))
}

func (ctx *AVFormatContext) AvformatNewStream2(c *AvCodec) *Stream {
	stream := (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(ctx), (*C.struct_AVCodec)(c)))
	stream.codec.pix_fmt = int32(AV_PIX_FMT_YUV)
	stream.codec.width = 640
	stream.codec.height = 480
	stream.time_base.num = 1
	stream.time_base.num = 25
	return stream
}

// //av_format_control_message av_format_get_control_message_cb (const Context *s)
// func (s *AVFormatContext) AvFormatControlMessage() C.av_format_get_control_message_cb {
// 	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s))
// }

// //void av_format_set_control_message_cb (Context *s, av_format_control_message callback)
// func (s *AVFormatContext) AvFormatSetControlMessageCb(c AvFormatControlMessage) C.av_format_get_control_message_cb {
// 	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), (C.struct_AVFormatControlMessage)(c))
// }

// //AvCodec * av_format_get_data_codec (const Context *s)
// func (s *AVFormatContext)AvFormatGetDataCodec() *AvCodec {
// 	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
// }

// //void av_format_set_data_codec (Context *s, AvCodec *c)
// func (s *AVFormatContext)AvFormatSetDataCodec( c *AvCodec) {
// 	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
// }

//Allocate an empty Context.
func SwsAllocContext() *SwsContext {
	return (*SwsContext)(C.sws_alloc_context())
}

//Initialize the swscaler context sws_context.
func SwsInitContext(ctx *SwsContext, sf, df *Filter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctx), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

//Free the swscaler context swsContext.
func SwsFreecontext(ctx *SwsContext) {
	C.sws_freeContext((*C.struct_SwsContext)(ctx))
}

//Allocate and return an Context.
func SwsGetcontext(sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *int) *SwsContext {
	return (*SwsContext)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(unsafe.Pointer(p))))
}

//Check if context can be reused, otherwise reallocate a new one.
func SwsGetcachedcontext(ctx *SwsContext, sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *float64) *SwsContext {
	return (*SwsContext)(C.sws_getCachedContext((*C.struct_SwsContext)(ctx), C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(p)))
}

//Initialize context after user parameters have been set.
func (s *SwrContext) SwrInit() int {
	return int(C.swr_init((*C.struct_SwrContext)(s)))
}

//Check whether an swr context has been initialized or not.
func (s *SwrContext) SwrIsInitialized() int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(s)))
}

//Allocate SwrContext if needed and set/reset common parameters.
func (s *SwrContext) SwrAllocSetOpts(ocl int64, osf AvSampleFormat, osr int, icl int64, isf AvSampleFormat, isr, lo, lc int) *SwrContext {
	return (*SwrContext)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

//SwrContext destructor functions. Free the given SwrContext and set the pointer to NULL.
func (s *SwrContext) SwrFree() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(s)))
}

//Closes the context so that swr_is_initialized() returns 0.
func (s *SwrContext) SwrClose() {
	C.swr_close((*C.struct_SwrContext)(s))
}

//Core conversion functions. Convert audio
func (s *SwrContext) SwrConvert(out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(s), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

//Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func (s *SwrContext) SwrNextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(s), C.int64_t(pts)))
}

//Low-level option setting functions
//These functons provide a means to set low-level options that is not possible with the AvOption API.
//Activate resampling compensation ("soft" compensation).
func (s *SwrContext) SwrSetCompensation(sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(s), C.int(sd), C.int(cd)))
}

//Set a customized input channel mapping.
func (s *SwrContext) SwrSetChannelMapping(cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(unsafe.Pointer(cm))))
}

//Set a customized remix matrix.
func (s *SwrContext) SwrSetMatrix(m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

//Sample handling functions. Drops the specified number of output samples.
func (s *SwrContext) SwrDropOutput(c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(s), C.int(c)))
}

//Injects the specified number of silence samples.
func (s *SwrContext) SwrInjectSilence(c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(s), C.int(c)))
}

//Gets the delay the next input sample will experience relative to the next output sample.
func (s *SwrContext) SwrGetDelay(b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(s), C.int64_t(b)))
}

//Frame based API. Convert the samples in the input Frame and write them to the output Frame.
func (s *SwrContext) SwrConvertFrame(o, i *Frame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

//Configure or reconfigure the SwrContext using the information provided by the AvFrames.
func (s *SwrContext) SwrConfigFrame(o, i *Frame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}
