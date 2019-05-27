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

// AVCodec ...
type AVCodec C.struct_AVCodec

// AVCodecTag ...
type AVCodecTag C.struct_AVCodecTag

// AVCodecDescriptor ...
type AVCodecDescriptor C.struct_AVCodecDescriptor

// AVCodecParser ...
type AVCodecParser C.struct_AVCodecParser

// AVCodecParserContext ...
type AVCodecParserContext C.struct_AVCodecParserContext

// AVPacket ...
type AVPacket C.struct_AVPacket

// AVBitStreamFilter ...
type AVBitStreamFilter C.struct_AVBitStreamFilter

// AVBitStreamFilterContext ...
type AVBitStreamFilterContext C.struct_AVBitStreamFilterContext

// AVCodecParameters ...
type AVCodecParameters C.struct_AVCodecParameters

// AVHWAccel ...
type AVHWAccel C.struct_AVHWAccel

// AVPacketSideData ...
type AVPacketSideData C.struct_AVPacketSideData

// AVPanScan ...
type AVPanScan C.struct_AVPanScan

// AVPicture ...
type AVPicture C.struct_AVPicture

// AVProfile ...
type AVProfile C.struct_AVProfile

// AVSubtitle ...
type AVSubtitle C.struct_AVSubtitle

// AVSubtitleRect ...
type AVSubtitleRect C.struct_AVSubtitleRect

// RcOverride ...
type RcOverride C.struct_RcOverride

// AVAudioServiceType ...
type AVAudioServiceType C.enum_AVAudioServiceType

// AVChromaLocation ...
type AVChromaLocation C.enum_AVChromaLocation

// AVCodecID ...
type AVCodecID C.enum_AVCodecID

// AVColorPrimaries ...
type AVColorPrimaries C.enum_AVColorPrimaries

// AVColorRange ...
type AVColorRange C.enum_AVColorRange

// AVColorSpace ...
type AVColorSpace C.enum_AVColorSpace

// AVColorTransferCharacteristic ...
type AVColorTransferCharacteristic C.enum_AVColorTransferCharacteristic

// AVDiscard ...
type AVDiscard C.enum_AVDiscard

// AVFieldOrder ...
type AVFieldOrder C.enum_AVFieldOrder

// AVPacketSideDataType ...
type AVPacketSideDataType C.enum_AVPacketSideDataType

// AVSampleFormat ...
type AVSampleFormat C.enum_AVSampleFormat

// AVCodecGetID ...
func (cp *AVCodecParameters) AVCodecGetID() AVCodecID {
	return *((*AVCodecID)(unsafe.Pointer(&cp.codec_id)))
}

// AVCodecGetType ...
func (cp *AVCodecParameters) AVCodecGetType() AVMediaType {
	return *((*AVMediaType)(unsafe.Pointer(&cp.codec_type)))
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
func (c *AVCodec) AVCodecGetMaxLowres() int {
	return int(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

// AVCodecNext If c is NULL, returns the first registered codec, if c is non-NULL,
func (c *AVCodec) AVCodecNext() *AVCodec {
	return (*AVCodec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

//AVCodecRegister Register the codec codec and initialize libavcodec.
func (c *AVCodec) AVCodecRegister() {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

//AVGetProfileName Return a name for the specified profile, if available.
func (c *AVCodec) AVGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

//AVCodecAllocContext3 Allocate an Context and set its fields to default values.
func (c *AVCodec) AVCodecAllocContext3() *AVCodecContext {
	return (*AVCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

// AVCodecIsEncoder ...
func (c *AVCodec) AVCodecIsEncoder() int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c)))
}

// AVCodecIsDecoder ...
func (c *AVCodec) AVCodecIsDecoder() int {
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
func AvsubtitleFree(s *AVSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// AvPacketAlloc ...
func AvPacketAlloc() *AVPacket {
	return (*AVPacket)(C.av_packet_alloc())
}

//AVPacketPackDictionary Pack a dictionary for use in side_data.
func AVPacketPackDictionary(d *Dictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

//AVPacketUnpackDictionary Unpack a dictionary from side_data.
func AVPacketUnpackDictionary(d *uint8, s int, dt **Dictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

//AVCodecFindDecoder Find a registered decoder with a matching codec ID.
func AVCodecFindDecoder(id AVCodecID) *AVCodec {
	return (*AVCodec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AVCodecIterate ...
func AVCodecIterate(p *unsafe.Pointer) *AVCodec {
	return (*AVCodec)(C.av_codec_iterate(p))
}

//AVCodecFindDecoderByName Find a registered decoder with the specified name.
func AVCodecFindDecoderByName(n string) *AVCodec {
	return (*AVCodec)(C.avcodec_find_decoder_by_name(C.CString(n)))
}

//AVCodecEnumToChromaPos Converts AVChromaLocation to swscale x/y chroma position.
func AVCodecEnumToChromaPos(x, y *int, l AVChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

//AVCodecChromaPosToEnum Converts swscale x/y chroma position to AVChromaLocation.
func AVCodecChromaPosToEnum(x, y int) AVChromaLocation {
	return (AVChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

//AVCodecFindEncoder Find a registered encoder with a matching codec ID.
func AVCodecFindEncoder(id AVCodecID) *AVCodec {
	return (*AVCodec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

//AVCodecFindEncoderByName Find a registered encoder with the specified name.
func AVCodecFindEncoderByName(c string) *AVCodec {
	return (*AVCodec)(C.avcodec_find_encoder_by_name(C.CString(c)))
}

//AVGetCodecTagString Put a string representing the codec tag codec_tag in buf.
func AVGetCodecTagString(b string, bf uintptr, c uint) uintptr {
	return uintptr(C.av_get_codec_tag_string(C.CString(b), C.size_t(bf), C.uint(c)))
}

// AVCodecString ...
func AVCodecString(b string, bs int, ctxt *AVCodecContext, e int) {
	C.avcodec_string(C.CString(b), C.int(bs), (*C.struct_AVCodecContext)(ctxt), C.int(e))
}

//AVCodecFillAudioFrame Fill Frame audio data and linesize pointers.
func AVCodecFillAudioFrame(f *Frame, c int, s AVSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

//AVGetBitsPerSample Return codec bits per sample.
func AVGetBitsPerSample(c AVCodecID) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

//AVGetPcmCodec Return the PCM codec associated with a sample format.
func AVGetPcmCodec(f AVSampleFormat, b int) AVCodecID {
	return (AVCodecID)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

//AVGetExactBitsPerSample Return codec bits per sample.
func AVGetExactBitsPerSample(c AVCodecID) int {
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

//AVHWAccelNext If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
func (a *AVHWAccel) AVHWAccelNext() *AVHWAccel {
	return (*AVHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

//AVCodecGetType Get the type of the given codec.
func AVCodecGetType(c AVCodecID) AVMediaType {
	return (AVMediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

//AVCodecGetName Get the name of a codec.
func AVCodecGetName(d AVCodecID) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

//AVCodecDescriptorGet const AVCodecDescriptor *avcodec_descriptor_get (enum AVCodecID id)
func AVCodecDescriptorGet(id AVCodecID) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

//AVCodecDescriptorNext Iterate over all codec descriptors known to libavcodec.
func (d *AVCodecDescriptor) AVCodecDescriptorNext() *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

// AVCodecDescriptorGetByName ...
func AVCodecDescriptorGetByName(n string) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_get_by_name(C.CString(n)))
}

// Pts ...
func (f *Frame) Pts() int64 {
	return int64(f.pts)
}
