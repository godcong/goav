// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package goav contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package goav

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"unsafe"
)

// Codec ...
type (
	Codec C.struct_AVCodec

	Descriptor    C.struct_AVCodecDescriptor
	Parser        C.struct_AVCodecParser
	ParserContext C.struct_AVCodecParserContext
	//Dictionary             C.struct_AVDictionary
	//Frame                  C.struct_AVFrame
	//MediaType              C.enum_AVMediaType
	Packet                 C.struct_AVPacket
	BitStreamFilter        C.struct_AVBitStreamFilter
	BitStreamFilterContext C.struct_AVBitStreamFilterContext
	//Rational               C.struct_AVRational
	//Class                         C.struct_AVClass
	AVCodecParameters C.struct_AVCodecParameters
	AVHWAccel         C.struct_AVHWAccel
	AvPacketSideData  C.struct_AVPacketSideData
	AvPanScan         C.struct_AVPanScan
	Picture           C.struct_AVPicture
	AvProfile         C.struct_AVProfile
	AvSubtitle        C.struct_AVSubtitle
	AvSubtitleRect    C.struct_AVSubtitleRect
	RcOverride        C.struct_RcOverride
	//AvBufferRef                   C.struct_AVBufferRef
	AvAudioServiceType            C.enum_AVAudioServiceType
	AvChromaLocation              C.enum_AVChromaLocation
	CodecID                       C.enum_AVCodecID
	AvColorPrimaries              C.enum_AVColorPrimaries
	AvColorRange                  C.enum_AVColorRange
	AvColorSpace                  C.enum_AVColorSpace
	AvColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	AvDiscard                     C.enum_AVDiscard
	AvFieldOrder                  C.enum_AVFieldOrder
	AvPacketSideDataType          C.enum_AVPacketSideDataType
	//PixelFormat                   C.enum_AVPixelFormat
	AvSampleFormat C.enum_AVSampleFormat
)

// AVCodecGetID ...
func (cp *AVCodecParameters) AVCodecGetID() CodecID {
	return *((*CodecID)(unsafe.Pointer(&cp.codec_id)))
}

// AVCodecGetType ...
func (cp *AVCodecParameters) AVCodecGetType() MediaType {
	return *((*MediaType)(unsafe.Pointer(&cp.codec_type)))
}

// AVCodecGetWidth ...
func (cp *AVCodecParameters) AVCodecGetWidth() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.width))))
}

// AVCodecGetHeight ...
func (cp *AVCodecParameters) AVCodecGetHeight() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.height))))
}

// AVCodecGetChannels ...
func (cp *AVCodecParameters) AVCodecGetChannels() int {
	return *((*int)(unsafe.Pointer(&cp.channels)))
}

// AVCodecGetSampleRate ...
func (cp *AVCodecParameters) AVCodecGetSampleRate() int {
	return *((*int)(unsafe.Pointer(&cp.sample_rate)))
}

// AVCodecGetMaxLowres ...
func (c *Codec) AVCodecGetMaxLowres() int {
	return int(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

// AVCodecNext If c is NULL, returns the first registered codec, if c is non-NULL,
func (c *Codec) AVCodecNext() *Codec {
	return (*Codec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

// Register the codec codec and initialize libavcodec.
func (c *Codec) AVCodecRegister() {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

//Return a name for the specified profile, if available.
func (c *Codec) AvGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

//Allocate an Context and set its fields to default values.
func (c *Codec) AVCodecAllocContext3() *AVCodecContext {
	return (*AVCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

// AVCodecIsEncoder ...
func (c *Codec) AVCodecIsEncoder() int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c)))
}

// AVCodecIsDecoder ...
func (c *Codec) AVCodecIsDecoder() int {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(c)))
}

//AvFastPaddedMalloc Same behaviour av_fast_malloc but the buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
func AvFastPaddedMalloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//AVCodecVersion Return the LIBAvCODEC_VERSION_INT constant.
func AVCodecVersion() uint {
	return uint(C.avcodec_version())
}

//AVCodecConfiguration Return the libavcodec build-time configuration.
func AVCodecConfiguration() string {
	return C.GoString(C.avcodec_configuration())

}

//AVCodecLicense Return the libavcodec license.
func AVCodecLicense() string {
	return C.GoString(C.avcodec_license())
}

//AVCodecRegisterAll Register all the codecs, parsers and bitstream filters which were enabled at configuration time.
func AVCodecRegisterAll() {
	C.av_register_all()
	C.avcodec_register_all()
	// C.av_log_set_level(0xffff)
}

//AVCodecGetClass Get the Class for Context.
func AVCodecGetClass() *Class {
	return (*Class)(C.avcodec_get_class())
}

//AVCodecGetFrameClass Get the Class for Frame.
func AVCodecGetFrameClass() *Class {
	return (*Class)(C.avcodec_get_frame_class())
}

//AVCodecGetSubtitleRectClass Get the Class for AvSubtitleRect.
func AVCodecGetSubtitleRectClass() *Class {
	return (*Class)(C.avcodec_get_subtitle_rect_class())
}

//AvsubtitleFree Free all allocated data in the given subtitle struct.
func AvsubtitleFree(s *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// AvPacketAlloc ...
func AvPacketAlloc() *Packet {
	return (*Packet)(C.av_packet_alloc())
}

//Pack a dictionary for use in side_data.
func AvPacketPackDictionary(d *Dictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

//Unpack a dictionary from side_data.
func AvPacketUnpackDictionary(d *uint8, s int, dt **Dictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

//Find a registered decoder with a matching codec ID.
func AVCodecFindDecoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AVCodecIterate ...
func AVCodecIterate(p *unsafe.Pointer) *Codec {
	return (*Codec)(C.av_codec_iterate(p))
}

//Find a registered decoder with the specified name.
func AVCodecFindDecoderByName(n string) *Codec {
	return (*Codec)(C.avcodec_find_decoder_by_name(C.CString(n)))
}

//Converts AvChromaLocation to swscale x/y chroma position.
func AVCodecEnumToChromaPos(x, y *int, l AvChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

//Converts swscale x/y chroma position to AvChromaLocation.
func AVCodecChromaPosToEnum(x, y int) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

//Find a registered encoder with a matching codec ID.
func AVCodecFindEncoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

//Find a registered encoder with the specified name.
func AVCodecFindEncoderByName(c string) *Codec {
	return (*Codec)(C.avcodec_find_encoder_by_name(C.CString(c)))
}

//Put a string representing the codec tag codec_tag in buf.
func AvGetCodecTagString(b string, bf uintptr, c uint) uintptr {
	return uintptr(C.av_get_codec_tag_string(C.CString(b), C.size_t(bf), C.uint(c)))
}

// AVCodecString ...
func AVCodecString(b string, bs int, ctxt *AVCodecContext, e int) {
	C.avcodec_string(C.CString(b), C.int(bs), (*C.struct_AVCodecContext)(ctxt), C.int(e))
}

//AVCodecFillAudioFrame Fill Frame audio data and linesize pointers.
func AVCodecFillAudioFrame(f *Frame, c int, s AvSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

//Return codec bits per sample.
func AvGetBitsPerSample(c CodecID) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

//Return the PCM codec associated with a sample format.
func AvGetPcmCodec(f AvSampleFormat, b int) CodecID {
	return (CodecID)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

//AVGetExactBitsPerSample Return codec bits per sample.
func AVGetExactBitsPerSample(c CodecID) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

//AVFastPaddedMallocz Same behaviour av_fast_padded_malloc except that buffer will always be 0-initialized after call.
func AVFastPaddedMallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//AVXiphlacing Encode extradata length to a buffer.
func AVXiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

//AVHWAccel If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
func (a *AVHWAccel) AVHWAccelNext() *AVHWAccel {
	return (*AVHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

//AVCodecGetType Get the type of the given codec.
func AVCodecGetType(c CodecID) MediaType {
	return (MediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

//AVCodecGetName Get the name of a codec.
func AVCodecGetName(d CodecID) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

//AVCodecDescriptorGet const Descriptor *avcodec_descriptor_get (enum CodecID id)
func AVCodecDescriptorGet(id CodecID) *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

//AVCodecDescriptorNext Iterate over all codec descriptors known to libavcodec.
func (d *Descriptor) AVCodecDescriptorNext() *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

// AVCodecDescriptorGetByName ...
func AVCodecDescriptorGetByName(n string) *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_get_by_name(C.CString(n)))
}

// Pts ...
func (f *Frame) Pts() int64 {
	return int64(f.pts)
}
