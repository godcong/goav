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
type Codec C.struct_AVCodec

// CodecTag ...
type CodecTag C.struct_AVCodecTag

// CodecDescriptor ...
type CodecDescriptor C.struct_AVCodecDescriptor

// CodecParser ...
type CodecParser C.struct_AVCodecParser

// CodecParserContext ...
type CodecParserContext C.struct_AVCodecParserContext

// Packet ...
type Packet C.struct_AVPacket

// BitStreamFilter ...
type BitStreamFilter C.struct_AVBitStreamFilter

// BitStreamFilterContext ...
type BitStreamFilterContext C.struct_AVBitStreamFilterContext

// CodecParameters ...
type CodecParameters C.struct_AVCodecParameters

// HWAccel ...
type HWAccel C.struct_AVHWAccel

// PacketSideData ...
type PacketSideData C.struct_AVPacketSideData

// PanScan ...
type PanScan C.struct_AVPanScan

// Picture ...
type Picture C.struct_AVPicture

// Profile ...
type Profile C.struct_AVProfile

// Subtitle ...
type Subtitle C.struct_AVSubtitle

// SubtitleRect ...
type SubtitleRect C.struct_AVSubtitleRect

// RcOverride ...
type RcOverride C.struct_RcOverride

// AudioServiceType ...
type AudioServiceType C.enum_AVAudioServiceType

// ColorPrimaries ...
type ColorPrimaries C.enum_AVColorPrimaries

// ColorRange ...
type ColorRange C.enum_AVColorRange

// ColorSpace ...
type ColorSpace C.enum_AVColorSpace

// ColorTransferCharacteristic ...
type ColorTransferCharacteristic C.enum_AVColorTransferCharacteristic

// Discard ...
type Discard C.enum_AVDiscard

// FieldOrder ...
type FieldOrder C.enum_AVFieldOrder

// SampleFormat ...
type SampleFormat C.enum_AVSampleFormat

// SampleFormatNone ...
const SampleFormatNone SampleFormat = C.AV_SAMPLE_FMT_NONE

// SampleFormatU8 ...
const SampleFormatU8 SampleFormat = C.AV_SAMPLE_FMT_U8

// SampleFormatS16 ...
const SampleFormatS16 SampleFormat = C.AV_SAMPLE_FMT_S16

// SampleFormatS32 ...
const SampleFormatS32 SampleFormat = C.AV_SAMPLE_FMT_S32

// SampleFormatFLT ...
const SampleFormatFLT SampleFormat = C.AV_SAMPLE_FMT_FLT

// SampleFormatDBL ...
const SampleFormatDBL SampleFormat = C.AV_SAMPLE_FMT_DBL

// SampleFormatU8P ...
const SampleFormatU8P SampleFormat = C.AV_SAMPLE_FMT_U8P

// SampleFormatS16P ...
const SampleFormatS16P SampleFormat = C.AV_SAMPLE_FMT_S16P

// SampleFormatS32P ...
const SampleFormatS32P SampleFormat = C.AV_SAMPLE_FMT_S32P

// SampleFormatFLTP ...
const SampleFormatFLTP SampleFormat = C.AV_SAMPLE_FMT_FLTP

// SampleFormatDBLP ...
const SampleFormatDBLP SampleFormat = C.AV_SAMPLE_FMT_DBLP

// Capabilities ...
type Capabilities int

// CapabilityDrawHorizBand ...
const CapabilityDrawHorizBand Capabilities = C.CODEC_CAP_DRAW_HORIZ_BAND

// CapabilityDR1 ...
const CapabilityDR1 Capabilities = C.CODEC_CAP_DR1

// CapabilityTruncated ...
const CapabilityTruncated Capabilities = C.CODEC_CAP_TRUNCATED

// CapabilityHWAccel ...
const CapabilityHWAccel Capabilities = C.GO_CODEC_CAP_HWACCEL

// CapabilityDelay ...
const CapabilityDelay Capabilities = C.CODEC_CAP_DELAY

// CapabilitySmallLastFrame ...
const CapabilitySmallLastFrame Capabilities = C.CODEC_CAP_SMALL_LAST_FRAME

// CapabilityHWAccelVDPAU ...
const CapabilityHWAccelVDPAU Capabilities = C.GO_CODEC_CAP_HWACCEL_VDPAU

// CapabilitySubframes ...
const CapabilitySubframes Capabilities = C.CODEC_CAP_SUBFRAMES

// CapabilityExperimental ...
const CapabilityExperimental Capabilities = C.CODEC_CAP_EXPERIMENTAL

// CapabilityChannelConf ...
const CapabilityChannelConf Capabilities = C.CODEC_CAP_CHANNEL_CONF

// CapabilityFrameThreads ...
const CapabilityFrameThreads Capabilities = C.CODEC_CAP_FRAME_THREADS

// CapabilitySliceThreads ...
const CapabilitySliceThreads Capabilities = C.CODEC_CAP_SLICE_THREADS

// CapabilityParamChange ...
const CapabilityParamChange Capabilities = C.CODEC_CAP_PARAM_CHANGE

// CapabilityAutoThreads ...
const CapabilityAutoThreads Capabilities = C.CODEC_CAP_AUTO_THREADS

// CapabilityVariableFrameSize ...
const CapabilityVariableFrameSize Capabilities = C.CODEC_CAP_VARIABLE_FRAME_SIZE

// CapabilityIntraOnly ...
const CapabilityIntraOnly Capabilities = C.CODEC_CAP_INTRA_ONLY

// CapabilityLossless ...
const CapabilityLossless Capabilities = C.CODEC_CAP_LOSSLESS

// Compliance ...
type Compliance int

// ComplianceVeryStrict ...
const ComplianceVeryStrict Compliance = C.FF_COMPLIANCE_VERY_STRICT

// ComplianceStrict ...
const ComplianceStrict Compliance = C.FF_COMPLIANCE_STRICT

// ComplianceNormal ...
const ComplianceNormal Compliance = C.FF_COMPLIANCE_NORMAL

// ComplianceUnofficial ...
const ComplianceUnofficial Compliance = C.FF_COMPLIANCE_UNOFFICIAL

// ComplianceExperimental ...
const ComplianceExperimental Compliance = C.FF_COMPLIANCE_EXPERIMENTAL

// PacketSideDataType ...
type PacketSideDataType C.enum_AVPacketSideDataType

// PacketSideDataPalette ...
const PacketSideDataPalette PacketSideDataType = C.AV_PKT_DATA_PALETTE

// PacketSideDataNewExtraData ...
const PacketSideDataNewExtraData PacketSideDataType = C.AV_PKT_DATA_NEW_EXTRADATA

// PacketSideDataParamChange ...
const PacketSideDataParamChange PacketSideDataType = C.AV_PKT_DATA_PARAM_CHANGE

// PacketSideDataH263MBInfo ...
const PacketSideDataH263MBInfo PacketSideDataType = C.AV_PKT_DATA_H263_MB_INFO

// PacketSideDataReplayGain ...
const PacketSideDataReplayGain PacketSideDataType = C.AV_PKT_DATA_REPLAYGAIN

// PacketSideDataDisplayMatrix ...
const PacketSideDataDisplayMatrix PacketSideDataType = C.AV_PKT_DATA_DISPLAYMATRIX

// PacketSideDataStereo3D ...
const PacketSideDataStereo3D PacketSideDataType = C.AV_PKT_DATA_STEREO3D

// PacketSideDataAudioServiceType ...
const PacketSideDataAudioServiceType PacketSideDataType = C.AV_PKT_DATA_AUDIO_SERVICE_TYPE

// PacketSideDataSkipSamples ...
const PacketSideDataSkipSamples PacketSideDataType = C.AV_PKT_DATA_SKIP_SAMPLES

// PacketSideDataJPDualMono ...
const PacketSideDataJPDualMono PacketSideDataType = C.AV_PKT_DATA_JP_DUALMONO

// PacketSideDataStringsMetaData ...
const PacketSideDataStringsMetaData PacketSideDataType = C.AV_PKT_DATA_STRINGS_METADATA

// PacketSideDataSubtitlePosition ...
const PacketSideDataSubtitlePosition PacketSideDataType = C.AV_PKT_DATA_SUBTITLE_POSITION

// PacketSideDataMatroskaBlockAdditional ...
const PacketSideDataMatroskaBlockAdditional PacketSideDataType = C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL

// PacketSideDataWebVTTIdentifier ...
const PacketSideDataWebVTTIdentifier PacketSideDataType = C.AV_PKT_DATA_WEBVTT_IDENTIFIER

// PacketSideDataWebVTTSettings ...
const PacketSideDataWebVTTSettings PacketSideDataType = C.AV_PKT_DATA_WEBVTT_SETTINGS

// PacketSideDataMetaDataUpdate ...
const PacketSideDataMetaDataUpdate PacketSideDataType = C.AV_PKT_DATA_METADATA_UPDATE

// PacketFlags ...
type PacketFlags int

// PacketFlagKey ...
const PacketFlagKey PacketFlags = C.AV_PKT_FLAG_KEY

// PacketFlagCorrupt ...
const PacketFlagCorrupt PacketFlags = C.AV_PKT_FLAG_CORRUPT

// DCTAlgorithm ...
type DCTAlgorithm int

// DCTAlgorithmAuto ...
const DCTAlgorithmAuto DCTAlgorithm = C.FF_DCT_AUTO

// DCTAlgorithmFastInt ...
const DCTAlgorithmFastInt DCTAlgorithm = C.FF_DCT_FASTINT

// DCTAlgorithmInt ...
const DCTAlgorithmInt DCTAlgorithm = C.GO_FF_DCT_INT

// DCTAlgorithmMMX ...
const DCTAlgorithmMMX DCTAlgorithm = C.FF_DCT_MMX

// DCTAlgorithmAltiVec ...
const DCTAlgorithmAltiVec DCTAlgorithm = C.FF_DCT_ALTIVEC

// DCTAlgorithmFAAN ...
const DCTAlgorithmFAAN DCTAlgorithm = C.FF_DCT_FAAN

// IDCTAlgorithm ...
type IDCTAlgorithm int

// IDCTAlgorithmAuto ...
const IDCTAlgorithmAuto IDCTAlgorithm = C.FF_IDCT_AUTO

// IDCTAlgorithmInt ...
const IDCTAlgorithmInt IDCTAlgorithm = C.FF_IDCT_INT

// IDCTAlgorithmSimple ...
const IDCTAlgorithmSimple IDCTAlgorithm = C.FF_IDCT_SIMPLE

// IDCTAlgorithmSimpleMMX ...
const IDCTAlgorithmSimpleMMX IDCTAlgorithm = C.FF_IDCT_SIMPLEMMX

// IDCTAlgorithmARM ...
const IDCTAlgorithmARM IDCTAlgorithm = C.FF_IDCT_ARM

// IDCTAlgorithmAltiVec ...
const IDCTAlgorithmAltiVec IDCTAlgorithm = C.FF_IDCT_ALTIVEC

// IDCTAlgorithmSH4 ...
const IDCTAlgorithmSH4 IDCTAlgorithm = C.GO_FF_IDCT_SH4

// IDCTAlgorithmSimpleARM ...
const IDCTAlgorithmSimpleARM IDCTAlgorithm = C.FF_IDCT_SIMPLEARM

// IDCTAlgorithmIPP ...
const IDCTAlgorithmIPP IDCTAlgorithm = C.GO_FF_IDCT_IPP

// IDCTAlgorithmXvid ...
const IDCTAlgorithmXvid IDCTAlgorithm = C.FF_IDCT_XVID

// IDCTAlgorithmXvidMMX ...
const IDCTAlgorithmXvidMMX IDCTAlgorithm = C.GO_FF_IDCT_XVIDMMX

// IDCTAlgorithmSimpleARMv5TE ...
const IDCTAlgorithmSimpleARMv5TE IDCTAlgorithm = C.FF_IDCT_SIMPLEARMV5TE

// IDCTAlgorithmSimpleARMv6 ...
const IDCTAlgorithmSimpleARMv6 IDCTAlgorithm = C.FF_IDCT_SIMPLEARMV6

// IDCTAlgorithmSimpleVis ...
const IDCTAlgorithmSimpleVis IDCTAlgorithm = C.GO_FF_IDCT_SIMPLEVIS

// IDCTAlgorithmFAAN ...
const IDCTAlgorithmFAAN IDCTAlgorithm = C.FF_IDCT_FAAN

// IDCTAlgorithmSimpleNEON ...
const IDCTAlgorithmSimpleNEON IDCTAlgorithm = C.FF_IDCT_SIMPLENEON

// IDCTAlgorithmSimpleAlpha ...
const IDCTAlgorithmSimpleAlpha IDCTAlgorithm = C.GO_FF_IDCT_SIMPLEALPHA

// IDCTAlgorithmSimpleAuto ...
const IDCTAlgorithmSimpleAuto IDCTAlgorithm = C.FF_IDCT_SIMPLEAUTO

// ThreadType ...
type ThreadType int

// ThreadTypeFrame ...
const ThreadTypeFrame ThreadType = C.FF_THREAD_FRAME

// ThreadTypeSlice ...
const ThreadTypeSlice ThreadType = C.FF_THREAD_SLICE

// ProfileUnknown ...
const (
	ProfileUnknown int = C.FF_PROFILE_UNKNOWN
)

// SubtitlesEncodingMode ...
type SubtitlesEncodingMode int

// SubtitlesEncodingModeDoNothing ...
const SubtitlesEncodingModeDoNothing SubtitlesEncodingMode = C.FF_SUB_CHARENC_MODE_DO_NOTHING

// SubtitlesEncodingModeAutomatic ...
const SubtitlesEncodingModeAutomatic SubtitlesEncodingMode = C.FF_SUB_CHARENC_MODE_AUTOMATIC

// SubtitlesEncodingModePreDecoder ...
const SubtitlesEncodingModePreDecoder SubtitlesEncodingMode = C.FF_SUB_CHARENC_MODE_PRE_DECODER

// CodecProps ...
type CodecProps int

// CodecPropIntraOnly ...
const CodecPropIntraOnly CodecProps = C.AV_CODEC_PROP_INTRA_ONLY

// CodecPropLossy ...
const CodecPropLossy CodecProps = C.AV_CODEC_PROP_LOSSY

// CodecPropLossless ...
const CodecPropLossless CodecProps = C.AV_CODEC_PROP_LOSSLESS

// CodecPropReorder ...
const CodecPropReorder CodecProps = C.AV_CODEC_PROP_REORDER

// CodecPropBitmapSub ...
const CodecPropBitmapSub CodecProps = C.AV_CODEC_PROP_BITMAP_SUB

// CodecPropTextSub ...
const CodecPropTextSub CodecProps = C.AV_CODEC_PROP_TEXT_SUB

// AVCodecGetID ...
func (cp *CodecParameters) AVCodecGetID() CodecID {
	return *((*CodecID)(unsafe.Pointer(&cp.codec_id)))
}

// AVCodecGetType ...
func (cp *CodecParameters) AVCodecGetType() MediaType {
	return *((*MediaType)(unsafe.Pointer(&cp.codec_type)))
}

// AVCodecGetWidth ...
func (cp *CodecParameters) AVCodecGetWidth() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.width))))
}

// AVCodecGetHeight ...
func (cp *CodecParameters) AVCodecGetHeight() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.height))))
}

// AVCodecGetChannels ...
func (cp *CodecParameters) AVCodecGetChannels() int {
	return *((*int)(unsafe.Pointer(&cp.channels)))
}

// AVCodecGetSampleRate ...
func (cp *CodecParameters) AVCodecGetSampleRate() int {
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

//AVCodecRegister Register the codec codec and initialize libavcodec.
func (c *Codec) AVCodecRegister() {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

//AVGetProfileName Return a name for the specified profile, if available.
func (c *Codec) AVGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

//AVCodecAllocContext3 Allocate an Context and set its fields to default values.
func (c *Codec) AVCodecAllocContext3() *CodecContext {
	return (*CodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
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

//AVCodecGetFrameClass Get the Class for AVFrame.
func AVCodecGetFrameClass() *Class {
	return (*Class)(C.avcodec_get_frame_class())
}

//AVCodecGetSubtitleRectClass Get the Class for AvSubtitleRect.
func AVCodecGetSubtitleRectClass() *Class {
	return (*Class)(C.avcodec_get_subtitle_rect_class())
}

//AvsubtitleFree Free all allocated data in the given subtitle struct.
func AvsubtitleFree(s *Subtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// AvPacketAlloc ...
func AvPacketAlloc() *Packet {
	return (*Packet)(C.av_packet_alloc())
}

//AVPacketPackDictionary Pack a dictionary for use in side_data.
func AVPacketPackDictionary(d *AVDictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

//AVPacketUnpackDictionary Unpack a dictionary from side_data.
func AVPacketUnpackDictionary(d *uint8, s int, dt **AVDictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

//AVCodecFindDecoder Find a registered decoder with a matching codec ID.
func AVCodecFindDecoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AVCodecIterate ...
func AVCodecIterate(p *unsafe.Pointer) *Codec {
	return (*Codec)(C.av_codec_iterate(p))
}

//AVCodecFindDecoderByName Find a registered decoder with the specified name.
func AVCodecFindDecoderByName(n string) *Codec {
	return (*Codec)(C.avcodec_find_decoder_by_name(C.CString(n)))
}

//AVCodecEnumToChromaPos Converts ChromaLocation to swscale x/y chroma position.
func AVCodecEnumToChromaPos(x, y *int, l ChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

//AVCodecChromaPosToEnum Converts swscale x/y chroma position to ChromaLocation.
func AVCodecChromaPosToEnum(x, y int) ChromaLocation {
	return (ChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

//AVCodecFindEncoder Find a registered encoder with a matching codec ID.
func AVCodecFindEncoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

//AVCodecFindEncoderByName Find a registered encoder with the specified name.
func AVCodecFindEncoderByName(c string) *Codec {
	return (*Codec)(C.avcodec_find_encoder_by_name(C.CString(c)))
}

//AVGetCodecTagString Put a string representing the codec tag codec_tag in buf.
func AVGetCodecTagString(b string, bf uintptr, c uint) uintptr {
	return uintptr(C.av_get_codec_tag_string(C.CString(b), C.size_t(bf), C.uint(c)))
}

// AVCodecString ...
func AVCodecString(b string, bs int, ctxt *CodecContext, e int) {
	C.avcodec_string(C.CString(b), C.int(bs), (*C.struct_AVCodecContext)(ctxt), C.int(e))
}

//AVCodecFillAudioFrame Fill AVFrame audio data and linesize pointers.
func AVCodecFillAudioFrame(f *AVFrame, c int, s SampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

//AVGetBitsPerSample Return codec bits per sample.
func AVGetBitsPerSample(c CodecID) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

//AVGetPcmCodec Return the PCM codec associated with a sample format.
func AVGetPcmCodec(f SampleFormat, b int) CodecID {
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

//AVHWAccelNext If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
func (a *HWAccel) AVHWAccelNext() *HWAccel {
	return (*HWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

//AVCodecGetType Get the type of the given codec.
func AVCodecGetType(c CodecID) MediaType {
	return (MediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

//AVCodecGetName Get the name of a codec.
func AVCodecGetName(d CodecID) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

//AVCodecDescriptorGet const CodecDescriptor *avcodec_descriptor_get (enum CodecID id)
func AVCodecDescriptorGet(id CodecID) *CodecDescriptor {
	return (*CodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

//AVCodecDescriptorNext Iterate over all codec descriptors known to libavcodec.
func (d *CodecDescriptor) AVCodecDescriptorNext() *CodecDescriptor {
	return (*CodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

// AVCodecDescriptorGetByName ...
func AVCodecDescriptorGetByName(n string) *CodecDescriptor {
	return (*CodecDescriptor)(C.avcodec_descriptor_get_by_name(C.CString(n)))
}

// Pts ...
func (f *AVFrame) Pts() int64 {
	return int64(f.pts)
}
