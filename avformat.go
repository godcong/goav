// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package goav provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package goav

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"unsafe"
)

// AvProbeData ...
type (
	AvProbeData     C.struct_AVProbeData
	InputFormat     C.struct_AVInputFormat
	OutputFormat    C.struct_AVOutputFormat
	AVFormatContext C.struct_AVFormatContext
	//Frame                      C.struct_AVFrame
	CodecContext       C.struct_AVCodecContext
	AvIndexEntry       C.struct_AVIndexEntry
	Stream             C.struct_AVStream
	AvProgram          C.struct_AVProgram
	AvChapter          C.struct_AVChapter
	AvPacketList       C.struct_AVPacketList
	CodecParserContext C.struct_AVCodecParserContext
	AvIOContext        C.struct_AVIOContext
	AVCodec            C.struct_AVCodec
	AVCodecTag         C.struct_AVCodecTag
	//Class                      C.struct_AVClass
	AvFormatInternal C.struct_AVFormatInternal
	AvIOInterruptCB  C.struct_AVIOInterruptCB
	//AvPacketSideData           C.struct_AVPacketSideData
	FFFrac            C.struct_FFFrac
	AvStreamParseType C.enum_AVStreamParseType
	//AVDiscard                  C.enum_AVDiscard
	//MediaType                  C.enum_AVMediaType
	AvDurationEstimationMethod C.enum_AVDurationEstimationMethod
	//AVPacketSideDataType       C.enum_AVPacketSideDataType
	//AVCodecID                    C.enum_AVCodecID
)

//Allocate and read the payload of a packet and initialize its fields with default values.
func (ctxt *AvIOContext) AvGetPacket(pkt *AVPacket, s int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt), toCPacket(pkt), C.int(s)))
}

//Read data and append it to the current content of the AVPacket.
func (ctxt *AvIOContext) AvAppendPacket(pkt *AVPacket, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt), toCPacket(pkt), C.int(s)))
}

// Close ...
func (ctxt *AvIOContext) Close() error {
	return ErrorFromCode(int(C.avio_close((*C.AVIOContext)(unsafe.Pointer(ctxt)))))
}

// AVRegisterInputFormat ...
func (f *InputFormat) AVRegisterInputFormat() {
	C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

// AVRegisterOutputFormat ...
func (f *OutputFormat) AVRegisterOutputFormat() {
	C.av_register_output_format((*C.struct_AVOutputFormat)(f))
}

//AVIFormatNext If f is NULL, returns the first registered input format, if f is non-NULL, returns the next registered input format after f or NULL if f is the last one.
func (f *InputFormat) AVIFormatNext() *InputFormat {
	return (*InputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

//AVOFormatNext If f is NULL, returns the first registered output format, if f is non-NULL, returns the next registered output format after f or NULL if f is the last one.
func (f *OutputFormat) AVOFormatNext() *OutputFormat {
	return (*OutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

//AVFormatVersion Return the LIBAvFORMAT_VERSION_INT constant.
func AVFormatVersion() uint {
	return uint(C.avformat_version())
}

//AVFormatConfiguration Return the libavformat build-time configuration.
func AVFormatConfiguration() string {
	return C.GoString(C.avformat_configuration())
}

//AVFormatLicense Return the libavformat license.
func AVFormatLicense() string {
	return C.GoString(C.avformat_license())
}

//AVRegisterAll Initialize libavformat and register all the muxers, demuxers and protocols.
func AVRegisterAll() {
	C.av_register_all()
}

//AVFormatNetworkInit Do global initialization of network components.
func AVFormatNetworkInit() int {
	return int(C.avformat_network_init())
}

//AVFormatNetworkDeinit Undo the initialization done by avformat_network_init.
func AVFormatNetworkDeinit() int {
	return int(C.avformat_network_deinit())
}

//AVFormatAllocContext Allocate an Context.
func AVFormatAllocContext() *AVFormatContext {
	return (*AVFormatContext)(C.avformat_alloc_context())
}

//AVFormatGetClass Get the Class for Context.
func AVFormatGetClass() *Class {
	return (*Class)(C.avformat_get_class())
}

//AVStreamGetSideData Get side information from stream.
func (s *Stream) AVStreamGetSideData(t AVPacketSideDataType, z int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

//AVFormatAllocOutputContext2 Allocate an Context for an output format.
func AVFormatAllocOutputContext2(ctx **AVFormatContext, o *OutputFormat, fo, fi string) int {
	CformatName := C.CString(fo)
	defer C.free(unsafe.Pointer(CformatName))

	CfileName := C.CString(fi)
	defer C.free(unsafe.Pointer(CfileName))

	return int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)), (*C.struct_AVOutputFormat)(o), CformatName, CfileName))
}

//AVFindInputFormat Find InputFormat based on the short name of the input format.
func AVFindInputFormat(s string) *InputFormat {
	CshortName := C.CString(s)
	defer C.free(unsafe.Pointer(CshortName))

	return (*InputFormat)(C.av_find_input_format(CshortName))
}

//AVProbeInputFormat Guess the file format.
func AVProbeInputFormat(pd *AvProbeData, i int) *InputFormat {
	return (*InputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(i)))
}

//AVProbeInputFormat2 Guess the file format.
func AVProbeInputFormat2(pd *AvProbeData, o int, sm *int) *InputFormat {
	return (*InputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sm))))
}

//AVProbeInputFormat3 Guess the file format.
func AVProbeInputFormat3(pd *AvProbeData, o int, sl *int) *InputFormat {
	return (*InputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sl))))
}

//AVProbeInputBuffer2 Probe a bytestream to determine the input format.
func AVProbeInputBuffer2(pb *AvIOContext, f **InputFormat, fi string, l int, o, m uint) int {
	Curl := C.CString(fi)
	defer C.free(unsafe.Pointer(Curl))

	return int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), Curl, unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//AVProbeInputBuffer Like av_probe_input_buffer2() but returns 0 on success.
func AVProbeInputBuffer(pb *AvIOContext, f **InputFormat, fi string, l int, o, m uint) int {
	Curl := C.CString(fi)
	defer C.free(unsafe.Pointer(Curl))

	return int(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), Curl, unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//AVFormatOpenInput Open an input stream and read the header.
func AVFormatOpenInput(ps **AVFormatContext, fi string, fmt *InputFormat, d **Dictionary) int {
	Cfi := C.CString(fi)
	defer C.free(unsafe.Pointer(Cfi))

	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), Cfi, (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AVGuessFormat Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func AVGuessFormat(sn, f, mt string) *OutputFormat {
	CshortName := C.CString(sn)
	defer C.free(unsafe.Pointer(CshortName))

	CfileName := C.CString(f)
	defer C.free(unsafe.Pointer(CfileName))

	CmimeType := C.CString(mt)
	defer C.free(unsafe.Pointer(CmimeType))

	return (*OutputFormat)(C.av_guess_format(CshortName, CfileName, CmimeType))
}

//AVGuessCodec Guess the codec ID based upon muxer and filename.
func AVGuessCodec(fmt *OutputFormat, sn, f, mt string, t MediaType) AVCodecID {
	CshortName := C.CString(sn)
	defer C.free(unsafe.Pointer(CshortName))

	CfileName := C.CString(f)
	defer C.free(unsafe.Pointer(CfileName))

	CmimeType := C.CString(mt)
	defer C.free(unsafe.Pointer(CmimeType))

	return (AVCodecID)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), CshortName, CfileName, CmimeType, (C.enum_AVMediaType)(t)))
}

//AVHexDump Send a nice hexadecimal dump of a buffer to the specified file stream.
func AVHexDump(f *File, b *uint8, s int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(b), C.int(s))
}

//AVHexDumpLog Send a nice hexadecimal dump of a buffer to the log.
func AVHexDumpLog(a, l int, b *uint8, s int) {
	C.av_hex_dump_log(unsafe.Pointer(&a), C.int(l), (*C.uint8_t)(b), C.int(s))
}

//AVPktDump2 Send a nice dump of a packet to the specified file stream.
func AVPktDump2(f *File, pkt *AVPacket, dp int, st *Stream) {
	C.av_pkt_dump2((*C.FILE)(f), toCPacket(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//AVPktDumpLog2 Send a nice dump of a packet to the log.
func AVPktDumpLog2(a int, l int, pkt *AVPacket, dp int, st *Stream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), toCPacket(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//AVCodecGetID enum AVCodecID av_codec_get_id (const struct AVCodecTag *const *tags, unsigned int tag)
//Get the AVCodecID for the given codec tag tag.
func AVCodecGetID(t **AVCodecTag, tag uint) AVCodecID {
	return (AVCodecID)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

//AVCodecGetTag Get the codec tag for the given codec id id.
func AVCodecGetTag(t **AVCodecTag, id AVCodecID) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

//AVCodecGetTag2 Get the codec tag for the given codec id.
func AVCodecGetTag2(t **AVCodecTag, id AVCodecID, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

//AVIndexSearchTimestamp Get the index for a specific timestamp.
func AVIndexSearchTimestamp(st *Stream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

//AVAddIndexEntry Add an index entry into a sorted list.
func AVAddIndexEntry(st *Stream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

//AvURLSplit Split a URL string into components.
func AvURLSplit(protoSize, authorizationSize, hostnameSize int, pp *int, pathSize int, url string) (proto, authorization, hostname, path string) {
	Cproto := (*C.char)(C.malloc(C.sizeof_char * C.ulong(protoSize)))
	defer C.free(unsafe.Pointer(Cproto))

	Cauthorization := (*C.char)(C.malloc(C.sizeof_char * C.ulong(authorizationSize)))
	defer C.free(unsafe.Pointer(Cauthorization))

	Chostname := (*C.char)(C.malloc(C.sizeof_char * C.ulong(hostnameSize)))
	defer C.free(unsafe.Pointer(Chostname))

	Cpath := (*C.char)(C.malloc(C.sizeof_char * C.ulong(pathSize)))
	defer C.free(unsafe.Pointer(Cpath))

	Curl := C.CString(url)
	defer C.free(unsafe.Pointer(Curl))

	C.av_url_split(
		Cproto, C.int(protoSize),
		Cauthorization, C.int(authorizationSize),
		Chostname, C.int(hostnameSize),
		(*C.int)(unsafe.Pointer(pp)),
		Cpath, C.int(pathSize),
		Curl,
	)

	return C.GoString(Cproto), C.GoString(Cauthorization), C.GoString(Chostname), C.GoString(Cpath)
}

//AvGetFrameFilename int av_get_frame_filename (char *buf, int buf_size, const char *path, int number)
//Return in 'buf' the path with 'd' replaced by a number.
func AvGetFrameFilename(bufSize int, path string, number int) (int, string) {
	Cbuf := (*C.char)(C.malloc(C.sizeof_char * C.ulong(bufSize)))
	defer C.free(unsafe.Pointer(Cbuf))

	Cpath := C.CString(path)
	defer C.free(unsafe.Pointer(Cpath))

	ret := int(C.av_get_frame_filename(Cbuf, C.int(bufSize), Cpath, C.int(number)))

	return ret, C.GoString(Cbuf)
}

//AvFilenameNumberTest Check whether filename actually is a numbered sequence generator.
func AvFilenameNumberTest(filename string) int {
	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	return int(C.av_filename_number_test(Cfilename))
}

//AvSdpCreate Generate an SDP for an RTP session.
func AvSdpCreate(ac **AVFormatContext, nFiles int, bufSize int) (int, string) {
	Cbuf := (*C.char)(C.malloc(C.sizeof_char * C.ulong(bufSize)))
	defer C.free(unsafe.Pointer(Cbuf))

	ret := int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nFiles), Cbuf, C.int(bufSize)))

	return ret, C.GoString(Cbuf)
}

//AvMatchExt int av_match_ext (const char *filename, const char *extensions)
//Return a positive value if the given filename has one of the given extensions, 0 otherwise.
func AvMatchExt(filename, extensions string) int {
	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	Cextensions := C.CString(extensions)
	defer C.free(unsafe.Pointer(Cextensions))

	return int(C.av_match_ext(Cfilename, Cextensions))
}

//AVFormatQueryCodec Test if the given container can store a codec.
func AVFormatQueryCodec(o *OutputFormat, cd AVCodecID, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
}

// AVFormatGetRiffVideoTags ...
func AVFormatGetRiffVideoTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_riff_video_tags())
}

//AVFormatGetRiffAudioTags struct AVCodecTag * avformat_get_riff_audio_tags (void)
func AVFormatGetRiffAudioTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_riff_audio_tags())
}

// AVFormatGetMovVideoTags ...
func AVFormatGetMovVideoTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_mov_video_tags())
}

// AVFormatGetMovAudioTags ...
func AVFormatGetMovAudioTags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_mov_audio_tags())
}

// AvIOOpen ...
func AvIOOpen(url string, flags int) (res *AvIOContext, err error) {
	urlStr := C.CString(url)
	defer C.free(unsafe.Pointer(urlStr))

	err = ErrorFromCode(int(C.avio_open((**C.AVIOContext)(unsafe.Pointer(&res)), urlStr, C.int(flags))))

	return
}
