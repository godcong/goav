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
	Context C.struct_AVCodecContext
)

func (c *Context) AvCodecGetPktTimebase() Rational {
	return Rational(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(c)))
}

// AvCodecGetPktTimebase2 returns the timebase rational number as numerator and denominator
func (c *Context) AvCodecGetPktTimebase2() Rational {
	return c.AvCodecGetPktTimebase()
}

func (c *Context) AvCodecSetPktTimebase(r Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(c), (C.struct_AVRational)(r))
}

func (c *Context) AvCodecGetCodecDescriptor() *Descriptor {
	return (*Descriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetCodecDescriptor(d *Descriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(c), (*C.struct_AVCodecDescriptor)(d))
}

func (c *Context) AvCodecGetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(c), C.int(i))
}

func (c *Context) AvCodecGetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(c), C.int(i))
}

func (c *Context) AvCodecGetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(c)))
}

func (c *Context) AvCodecSetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(c), (*C.uint16_t)(t))
}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (c *Context) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(c)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (c *Context) AvcodecGetContextDefaults3(c *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(c), (*C.struct_AVCodec)(c)))
}

//Copy the settings of the source Context into the destination Context.
func (c *Context) AvcodecCopyContext(ctxt2 *Context) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(c), (*C.struct_AVCodecContext)(ctxt2)))
}

//Initialize the Context to use the given Codec
func (c *Context) AvcodecOpen2(c *Codec, d **Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(c), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (c *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(c)))
}

//The default callback for Context.get_buffer2().
func (c *Context) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (c *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(c), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (c *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(c), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (c *Context) AvcodecDecodeAudio4(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
func (c *Context) AvcodecDecodeVideo2(p *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode a subtitle message.
func (c *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(c), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Encode a frame of audio.
func (c *Context) AvcodecEncodeAudio2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//Encode a frame of video
func (c *Context) AvcodecEncodeVideo2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

func (c *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(c), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (c *Context) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(c), (*C.enum_AVPixelFormat)(f)))
}

//Reset the internal decoder state / flush internal buffers.
func (c *Context) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(c))
}

//Return audio frame duration.
func (c *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(c), C.int(f)))
}

func (c *Context) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(c)))
}

//Parse a packet.
func (c *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(c), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (c *Context) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(c), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
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

func (c *Context) SetTimebase(num1 int, den1 int) {
	c.time_base.num = C.int(num1)
	c.time_base.den = C.int(den1)
}

func (c *Context) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	c.width = C.int(width)
	c.height = C.int(height)
	// c.bit_rate = 1000000
	c.gop_size = C.int(gopSize)
	// c.max_b_frames = 2
	if hasBframes {
		c.has_b_frames = 1
	} else {
		c.has_b_frames = 0
	}
	// c.extradata = nil
	// c.extradata_size = 0
	// c.channels = 0
	c.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(c.priv_data, "preset", "ultrafast", 0)
}

func (c *Context) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	c.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

func (c *Context) AvcodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(c), (*C.struct_AVPacket)(packet)))
}

func (c *Context) AvcodecReceiveFrame(frame *Frame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(c), (*C.struct_AVFrame)(frame)))
}

const (
	AvseekFlagBackward = 1 ///< seek backward
	AvseekFlagByte     = 2 ///< seeking based on position in bytes
	AvseekFlagAny      = 4 ///< seek to any frame, even non-keyframes
	AvseekFlagFrame    = 8 ///< seeking based on frame number
)

func (c *Context) AvFormatGetProbeScore() int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(c)))
}

func (c *Context) AvFormatGetVideoCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(c)))
}

func (c *Context) AvFormatSetVideoCodec(c *AvCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(c))
}

func (c *Context) AvFormatGetAudioCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(c)))
}

func (c *Context) AvFormatSetAudioCodec(c *AvCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(c))
}

func (c *Context) AvFormatGetSubtitleCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(c)))
}

func (c *Context) AvFormatSetSubtitleCodec(c *AvCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(c))
}

func (c *Context) AvFormatGetMetadataHeaderPadding() int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(c)))
}

func (c *Context) AvFormatSetMetadataHeaderPadding(c int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(c), C.int(c))
}

func (c *Context) AvFormatGetOpaque() {
	C.av_format_get_opaque((*C.struct_AVFormatContext)(c))
}

func (c *Context) AvFormatSetOpaque(o int) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(c), unsafe.Pointer(&o))
}

//This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func (c *Context) AvFormatInjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(c))
}

//Returns the method used to set ctx->duration.
func (c *Context) AvFmtCtxGetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(c)))
}

//Free an Context and all its streams.
func (c *Context) AvformatFreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(c))
}

//Add a new stream to a media file.
func (c *Context) AvformatNewStream(c *AvCodec) *Stream {
	return (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(c)))
}

func (c *Context) AvNewProgram(id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(c), C.int(id)))
}

//Read packets of a media file to get stream information.
func (c *Context) AvformatFindStreamInfo(d **Dictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Find the programs which belong to a given stream.
func (c *Context) AvFindProgramFromStream(l *AvProgram, su int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(c), (*C.struct_AVProgram)(l), C.int(su)))
}

//Find the "best" stream in the file.
func AvFindBestStream(ic *Context, t MediaType, ws, rs int, c **AvCodec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

//Return the next frame of a stream.
func (c *Context) AvReadFrame(pkt *Packet) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(c)), toCPacket(pkt)))
}

//Seek to the keyframe at timestamp.
func (c *Context) AvSeekFrame(st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(c), C.int(st), C.int64_t(t), C.int(f)))
}

// AvSeekFrameTime seeks to a specified time location.
// |timebase| is codec specific and can be obtained by calling AvCodecGetPktTimebase2
func (c *Context) AvSeekFrameTime(st int, at time.Duration, timebase Rational) int {
	t2 := C.double(C.double(at.Seconds())*C.double(timebase.Den())) / (C.double(timebase.Num()))
	// log.Printf("Seeking to time :%v TimebaseTime:%v ActualTimebase:%v", at, t2, timebase)
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(c), C.int(st), C.int64_t(t2), AvseekFlagBackward))
}

//Seek to timestamp ts.
func (c *Context) AvformatSeekFile(si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(c), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

//Start playing a network-based stream (e.g.
func (c *Context) AvReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(c)))
}

//Pause a network-based stream (e.g.
func (c *Context) AvReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(c)))
}

//Close an opened input Context.
func (c *Context) AvformatCloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&c)))
}

//Allocate the stream private data and write the stream header to an output media file.
func (c *Context) AvformatWriteHeader(o **Dictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(c), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//Write a packet to an output media file.
func (c *Context) AvWriteFrame(pkt *Packet) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(c), toCPacket(pkt)))
}

//Write a packet to an output media file ensuring correct interleaving.
func (c *Context) AvInterleavedWriteFrame(pkt *Packet) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(c), toCPacket(pkt)))
}

//Write a uncoded frame to an output media file.
func (c *Context) AvWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(c), C.int(si), (*C.struct_AVFrame)(f)))
}

//Write a uncoded frame to an output media file.
func (c *Context) AvInterleavedWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(c), C.int(si), (*C.struct_AVFrame)(f)))
}

//Test whether a muxer supports uncoded frame.
func (c *Context) AvWriteUncodedFrameQuery(si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(c), C.int(si)))
}

//Write the stream trailer to an output media file and free the file private data.
func (c *Context) AvWriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(c)))
}

//Get timing information for the data currently output.
func (c *Context) AvGetOutputTimestamp(st int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(c), C.int(st), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

func (c *Context) AvFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(c)))
}

//Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (c *Context) AvDumpFormat(i int, url string, io int) {
	Curl := C.CString(url)
	defer C.free(unsafe.Pointer(Curl))

	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(c)), C.int(i), Curl, C.int(io))
}

//Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func (c *Context) AvGuessSampleAspectRatio(st *Stream, fr *Frame) Rational {
	return newRational(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(c), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//Guess the frame rate, based on both the container and codec information.
func (c *Context) AvGuessFrameRate(st *Stream, fr *Frame) Rational {
	return newRational(C.av_guess_frame_rate((*C.struct_AVFormatContext)(c), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//Check if the stream st contained in s is matched by the stream specifier spec.
func (c *Context) AvformatMatchStreamSpecifier(st *Stream, spec string) int {
	Cspec := C.CString(spec)
	defer C.free(unsafe.Pointer(Cspec))

	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(c), (*C.struct_AVStream)(st), Cspec))
}

func (c *Context) AvformatQueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(c)))
}

func (c *Context) AvformatNewStream2(c *AvCodec) *Stream {
	stream := (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(c)))
	stream.codec.pix_fmt = int32(AV_PIX_FMT_YUV)
	stream.codec.width = 640
	stream.codec.height = 480
	stream.time_base.num = 1
	stream.time_base.num = 25
	return stream
}

// //av_format_control_message av_format_get_control_message_cb (const Context *s)
// func (s *Context) AvFormatControlMessage() C.av_format_get_control_message_cb {
// 	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s))
// }

// //void av_format_set_control_message_cb (Context *s, av_format_control_message callback)
// func (s *Context) AvFormatSetControlMessageCb(c AvFormatControlMessage) C.av_format_get_control_message_cb {
// 	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), (C.struct_AVFormatControlMessage)(c))
// }

// //AvCodec * av_format_get_data_codec (const Context *s)
// func (s *Context)AvFormatGetDataCodec() *AvCodec {
// 	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
// }

// //void av_format_set_data_codec (Context *s, AvCodec *c)
// func (s *Context)AvFormatSetDataCodec( c *AvCodec) {
// 	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
// }

//Allocate an empty Context.
func SwsAllocContext() *Context {
	return (*Context)(C.sws_alloc_context())
}

//Initialize the swscaler context sws_context.
func SwsInitContext(ctxt *Context, sf, df *Filter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctxt), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

//Free the swscaler context swsContext.
func SwsFreecontext(ctxt *Context) {
	C.sws_freeContext((*C.struct_SwsContext)(ctxt))
}

//Allocate and return an Context.
func SwsGetcontext(sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *int) *Context {
	return (*Context)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(unsafe.Pointer(p))))
}

//Check if context can be reused, otherwise reallocate a new one.
func SwsGetcachedcontext(ctxt *Context, sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *float64) *Context {
	return (*Context)(C.sws_getCachedContext((*C.struct_SwsContext)(ctxt), C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(p)))
}
