// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

// Multiple encoders have the same ID and are able to encode compatible streams.
const (
	AVCodecID012v            = int(C.AV_CODEC_ID_012V)
	AVCodecID4xm             = int(C.AV_CODEC_ID_4XM)
	AVCodecID8bps            = int(C.AV_CODEC_ID_8BPS)
	AVCodecID8svxExp         = int(C.AV_CODEC_ID_8SVX_EXP)
	AVCodecID8svxFib         = int(C.AV_CODEC_ID_8SVX_FIB)
	AVCodecIDA64Multi        = int(C.AV_CODEC_ID_A64_MULTI)
	AVCodecIDA64Multi5       = int(C.AV_CODEC_ID_A64_MULTI5)
	AVCodecIDAac             = int(C.AV_CODEC_ID_AAC)
	AVCodecIDAacLatm         = int(C.AV_CODEC_ID_AAC_LATM)
	AVCodecIDAasc            = int(C.AV_CODEC_ID_AASC)
	AVCodecIDAc3             = int(C.AV_CODEC_ID_AC3)
	AVCodecIDAdpcm4xm        = int(C.AV_CODEC_ID_ADPCM_4XM)
	AVCodecIDAdpcmAdx        = int(C.AV_CODEC_ID_ADPCM_ADX)
	AVCodecIDAdpcmAfc        = int(C.AV_CODEC_ID_ADPCM_AFC)
	AVCodecIDAdpcmCt         = int(C.AV_CODEC_ID_ADPCM_CT)
	AVCodecIDAdpcmDtk        = int(C.AV_CODEC_ID_ADPCM_DTK)
	AVCodecIDAdpcmEa         = int(C.AV_CODEC_ID_ADPCM_EA)
	AVCodecIDAdpcmEaMaxisXa  = int(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	AVCodecIDAdpcmEaR1       = int(C.AV_CODEC_ID_ADPCM_EA_R1)
	AVCodecIDAdpcmEaR2       = int(C.AV_CODEC_ID_ADPCM_EA_R2)
	AVCodecIDAdpcmEaR3       = int(C.AV_CODEC_ID_ADPCM_EA_R3)
	AVCodecIDAdpcmEaXas      = int(C.AV_CODEC_ID_ADPCM_EA_XAS)
	AVCodecIDAdpcmG722       = int(C.AV_CODEC_ID_ADPCM_G722)
	AVCodecIDAdpcmG726       = int(C.AV_CODEC_ID_ADPCM_G726)
	AVCodecIDAdpcmG726le     = int(C.AV_CODEC_ID_ADPCM_G726LE)
	AVCodecIDAdpcmImaAmv     = int(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	AVCodecIDAdpcmImaApc     = int(C.AV_CODEC_ID_ADPCM_IMA_APC)
	AVCodecIDAdpcmImaDk3     = int(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	AVCodecIDAdpcmImaDk4     = int(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	AVCodecIDAdpcmImaEaEacs  = int(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	AVCodecIDAdpcmImaEaSead  = int(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	AVCodecIDAdpcmImaIss     = int(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	AVCodecIDAdpcmImaOki     = int(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	AVCodecIDAdpcmImaQt      = int(C.AV_CODEC_ID_ADPCM_IMA_QT)
	AVCodecIDAdpcmImaRad     = int(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	AVCodecIDAdpcmImaSmjpeg  = int(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	AVCodecIDAdpcmImaWav     = int(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	AVCodecIDAdpcmImaWs      = int(C.AV_CODEC_ID_ADPCM_IMA_WS)
	AVCodecIDAdpcmMs         = int(C.AV_CODEC_ID_ADPCM_MS)
	AVCodecIDAdpcmSbpro2     = int(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	AVCodecIDAdpcmSbpro3     = int(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	AVCodecIDAdpcmSbpro4     = int(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	AVCodecIDAdpcmSwf        = int(C.AV_CODEC_ID_ADPCM_SWF)
	AVCodecIDAdpcmThp        = int(C.AV_CODEC_ID_ADPCM_THP)
	AVCodecIDAdpcmVima       = int(C.AV_CODEC_ID_ADPCM_VIMA)
	AVCodecIDAdpcmXa         = int(C.AV_CODEC_ID_ADPCM_XA)
	AVCodecIDAdpcmYamaha     = int(C.AV_CODEC_ID_ADPCM_YAMAHA)
	AVCodecIDAic             = int(C.AV_CODEC_ID_AIC)
	AVCodecIDAlac            = int(C.AV_CODEC_ID_ALAC)
	AVCodecIDAliasPix        = int(C.AV_CODEC_ID_ALIAS_PIX)
	AVCodecIDAmrNb           = int(C.AV_CODEC_ID_AMR_NB)
	AVCodecIDAmrWb           = int(C.AV_CODEC_ID_AMR_WB)
	AVCodecIDAmv             = int(C.AV_CODEC_ID_AMV)
	AVCodecIDAnm             = int(C.AV_CODEC_ID_ANM)
	AVCodecIDAnsi            = int(C.AV_CODEC_ID_ANSI)
	AVCodecIDApe             = int(C.AV_CODEC_ID_APE)
	AVCodecIDAss             = int(C.AV_CODEC_ID_ASS)
	AVCodecIDAsv1            = int(C.AV_CODEC_ID_ASV1)
	AVCodecIDAsv2            = int(C.AV_CODEC_ID_ASV2)
	AVCodecIDAtrac1          = int(C.AV_CODEC_ID_ATRAC1)
	AVCodecIDAtrac3          = int(C.AV_CODEC_ID_ATRAC3)
	AVCodecIDAtrac3p         = int(C.AV_CODEC_ID_ATRAC3P)
	AVCodecIDAura            = int(C.AV_CODEC_ID_AURA)
	AVCodecIDAura2           = int(C.AV_CODEC_ID_AURA2)
	AVCodecIDAvrn            = int(C.AV_CODEC_ID_AVRN)
	AVCodecIDAvrp            = int(C.AV_CODEC_ID_AVRP)
	AVCodecIDAvs             = int(C.AV_CODEC_ID_AVS)
	AVCodecIDAvui            = int(C.AV_CODEC_ID_AVUI)
	AVCodecIDAyuv            = int(C.AV_CODEC_ID_AYUV)
	AVCodecIDBethsoftvid     = int(C.AV_CODEC_ID_BETHSOFTVID)
	AVCodecIDBfi             = int(C.AV_CODEC_ID_BFI)
	AVCodecIDBinkaudioDct    = int(C.AV_CODEC_ID_BINKAUDIO_DCT)
	AVCodecIDBinkaudioRdft   = int(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	AVCodecIDBinkvideo       = int(C.AV_CODEC_ID_BINKVIDEO)
	AVCodecIDBintext         = int(C.AV_CODEC_ID_BINTEXT)
	AVCodecIDBinData         = int(C.AV_CODEC_ID_BIN_DATA)
	AVCodecIDBmp             = int(C.AV_CODEC_ID_BMP)
	AVCodecIDBmvAudio        = int(C.AV_CODEC_ID_BMV_AUDIO)
	AVCodecIDBmvVideo        = int(C.AV_CODEC_ID_BMV_VIDEO)
	AVCodecIDBrenderPix      = int(C.AV_CODEC_ID_BRENDER_PIX)
	AVCodecIDC93             = int(C.AV_CODEC_ID_C93)
	AVCodecIDCavs            = int(C.AV_CODEC_ID_CAVS)
	AVCodecIDCdgraphics      = int(C.AV_CODEC_ID_CDGRAPHICS)
	AVCodecIDCdxl            = int(C.AV_CODEC_ID_CDXL)
	AVCodecIDCelt            = int(C.AV_CODEC_ID_CELT)
	AVCodecIDCinepak         = int(C.AV_CODEC_ID_CINEPAK)
	AVCodecIDCljr            = int(C.AV_CODEC_ID_CLJR)
	AVCodecIDCllc            = int(C.AV_CODEC_ID_CLLC)
	AVCodecIDCmv             = int(C.AV_CODEC_ID_CMV)
	AVCodecIDComfortNoise    = int(C.AV_CODEC_ID_COMFORT_NOISE)
	AVCodecIDCook            = int(C.AV_CODEC_ID_COOK)
	AVCodecIDCpia            = int(C.AV_CODEC_ID_CPIA)
	AVCodecIDCscd            = int(C.AV_CODEC_ID_CSCD)
	AVCodecIDCyuv            = int(C.AV_CODEC_ID_CYUV)
	AVCodecIDDfa             = int(C.AV_CODEC_ID_DFA)
	AVCodecIDDirac           = int(C.AV_CODEC_ID_DIRAC)
	AVCodecIDDnxhd           = int(C.AV_CODEC_ID_DNXHD)
	AVCodecIDDpx             = int(C.AV_CODEC_ID_DPX)
	AVCodecIDDsdLsbf         = int(C.AV_CODEC_ID_DSD_LSBF)
	AVCodecIDDsdLsbfPlanar   = int(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	AVCodecIDDsdMsbf         = int(C.AV_CODEC_ID_DSD_MSBF)
	AVCodecIDDsdMsbfPlanar   = int(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	AVCodecIDDsicinaudio     = int(C.AV_CODEC_ID_DSICINAUDIO)
	AVCodecIDDsicinvideo     = int(C.AV_CODEC_ID_DSICINVIDEO)
	AVCodecIDDts             = int(C.AV_CODEC_ID_DTS)
	AVCodecIDDvaudio         = int(C.AV_CODEC_ID_DVAUDIO)
	AVCodecIDDvbSubtitle     = int(C.AV_CODEC_ID_DVB_SUBTITLE)
	AVCodecIDDvbTeletext     = int(C.AV_CODEC_ID_DVB_TELETEXT)
	AVCodecIDDvdNav          = int(C.AV_CODEC_ID_DVD_NAV)
	AVCodecIDDvdSubtitle     = int(C.AV_CODEC_ID_DVD_SUBTITLE)
	AVCodecIDDvvideo         = int(C.AV_CODEC_ID_DVVIDEO)
	AVCodecIDDxa             = int(C.AV_CODEC_ID_DXA)
	AVCodecIDDxtory          = int(C.AV_CODEC_ID_DXTORY)
	AVCodecIDEac3            = int(C.AV_CODEC_ID_EAC3)
	AVCodecIDEia608          = int(C.AV_CODEC_ID_EIA_608)
	AVCodecIDEscape124       = int(C.AV_CODEC_ID_ESCAPE124)
	AVCodecIDEscape130       = int(C.AV_CODEC_ID_ESCAPE130)
	AVCodecIDEvrc            = int(C.AV_CODEC_ID_EVRC)
	AVCodecIDExr             = int(C.AV_CODEC_ID_EXR)
	AVCodecIDFfmetadata      = int(C.AV_CODEC_ID_FFMETADATA)
	AVCodecIDFfv1            = int(C.AV_CODEC_ID_FFV1)
	AVCodecIDFfvhuff         = int(C.AV_CODEC_ID_FFVHUFF)
	AVCodecIDFfwavesynth     = int(C.AV_CODEC_ID_FFWAVESYNTH)
	AVCodecIDFic             = int(C.AV_CODEC_ID_FIC)
	AVCodecIDFirstAudio      = int(C.AV_CODEC_ID_FIRST_AUDIO)
	AVCodecIDFirstSubtitle   = int(C.AV_CODEC_ID_FIRST_SUBTITLE)
	AVCodecIDFirstUnknown    = int(C.AV_CODEC_ID_FIRST_UNKNOWN)
	AVCodecIDFlac            = int(C.AV_CODEC_ID_FLAC)
	AVCodecIDFlashsv         = int(C.AV_CODEC_ID_FLASHSV)
	AVCodecIDFlashsv2        = int(C.AV_CODEC_ID_FLASHSV2)
	AVCodecIDFlic            = int(C.AV_CODEC_ID_FLIC)
	AVCodecIDFlv1            = int(C.AV_CODEC_ID_FLV1)
	AVCodecIDFraps           = int(C.AV_CODEC_ID_FRAPS)
	AVCodecIDFrwu            = int(C.AV_CODEC_ID_FRWU)
	AVCodecIDG2m             = int(C.AV_CODEC_ID_G2M)
	AVCodecIDG7231           = int(C.AV_CODEC_ID_G723_1)
	AVCodecIDG729            = int(C.AV_CODEC_ID_G729)
	AVCodecIDGif             = int(C.AV_CODEC_ID_GIF)
	AVCodecIDGsm             = int(C.AV_CODEC_ID_GSM)
	AVCodecIDGsmMs           = int(C.AV_CODEC_ID_GSM_MS)
	AVCodecIDH261            = int(C.AV_CODEC_ID_H261)
	AVCodecIDH263            = int(C.AV_CODEC_ID_H263)
	AVCodecIDH263i           = int(C.AV_CODEC_ID_H263I)
	AVCodecIDH263p           = int(C.AV_CODEC_ID_H263P)
	AVCodecIDH264            = int(C.AV_CODEC_ID_H264)
	AVCodecIDH265            = int(C.AV_CODEC_ID_H265)
	AVCodecIDHdmvPgsSubtitle = int(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	AVCodecIDHevc            = int(C.AV_CODEC_ID_HEVC)
	AVCodecIDHnm4Video       = int(C.AV_CODEC_ID_HNM4_VIDEO)
	AVCodecIDHuffyuv         = int(C.AV_CODEC_ID_HUFFYUV)
	AVCodecIDIac             = int(C.AV_CODEC_ID_IAC)
	AVCodecIDIdcin           = int(C.AV_CODEC_ID_IDCIN)
	AVCodecIDIdf             = int(C.AV_CODEC_ID_IDF)
	AVCodecIDIffByterun1     = int(C.AV_CODEC_ID_IFF_BYTERUN1)
	AVCodecIDIffIlbm         = int(C.AV_CODEC_ID_IFF_ILBM)
	AVCodecIDIlbc            = int(C.AV_CODEC_ID_ILBC)
	AVCodecIDImc             = int(C.AV_CODEC_ID_IMC)
	AVCodecIDIndeo2          = int(C.AV_CODEC_ID_INDEO2)
	AVCodecIDIndeo3          = int(C.AV_CODEC_ID_INDEO3)
	AVCodecIDIndeo4          = int(C.AV_CODEC_ID_INDEO4)
	AVCodecIDIndeo5          = int(C.AV_CODEC_ID_INDEO5)
	AVCodecIDInterplayDpcm   = int(C.AV_CODEC_ID_INTERPLAY_DPCM)
	AVCodecIDInterplayVideo  = int(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	AVCodecIDJacosub         = int(C.AV_CODEC_ID_JACOSUB)
	AVCodecIDJpeg2000        = int(C.AV_CODEC_ID_JPEG2000)
	AVCodecIDJpegls          = int(C.AV_CODEC_ID_JPEGLS)
	AVCodecIDJv              = int(C.AV_CODEC_ID_JV)
	AVCodecIDKgv1            = int(C.AV_CODEC_ID_KGV1)
	AVCodecIDKmvc            = int(C.AV_CODEC_ID_KMVC)
	AVCodecIDLagarith        = int(C.AV_CODEC_ID_LAGARITH)
	AVCodecIDLjpeg           = int(C.AV_CODEC_ID_LJPEG)
	AVCodecIDLoco            = int(C.AV_CODEC_ID_LOCO)
	AVCodecIDMace3           = int(C.AV_CODEC_ID_MACE3)
	AVCodecIDMace6           = int(C.AV_CODEC_ID_MACE6)
	AVCodecIDMad             = int(C.AV_CODEC_ID_MAD)
	AVCodecIDMdec            = int(C.AV_CODEC_ID_MDEC)
	AVCodecIDMetasound       = int(C.AV_CODEC_ID_METASOUND)
	AVCodecIDMicrodvd        = int(C.AV_CODEC_ID_MICRODVD)
	AVCodecIDMimic           = int(C.AV_CODEC_ID_MIMIC)
	AVCodecIDMjpeg           = int(C.AV_CODEC_ID_MJPEG)
	AVCodecIDMjpegb          = int(C.AV_CODEC_ID_MJPEGB)
	AVCodecIDMlp             = int(C.AV_CODEC_ID_MLP)
	AVCodecIDMmvideo         = int(C.AV_CODEC_ID_MMVIDEO)
	AVCodecIDMotionpixels    = int(C.AV_CODEC_ID_MOTIONPIXELS)
	AVCodecIDMovText         = int(C.AV_CODEC_ID_MOV_TEXT)
	AVCodecIDMp1             = int(C.AV_CODEC_ID_MP1)
	AVCodecIDMp2             = int(C.AV_CODEC_ID_MP2)
	AVCodecIDMp3             = int(C.AV_CODEC_ID_MP3)
	AVCodecIDMp3adu          = int(C.AV_CODEC_ID_MP3ADU)
	AVCodecIDMp3on4          = int(C.AV_CODEC_ID_MP3ON4)
	AVCodecIDMp4als          = int(C.AV_CODEC_ID_MP4ALS)
	AVCodecIDMpeg1video      = int(C.AV_CODEC_ID_MPEG1VIDEO)
	AVCodecIDMpeg2ts         = int(C.AV_CODEC_ID_MPEG2TS)
	AVCodecIDMpeg2video      = int(C.AV_CODEC_ID_MPEG2VIDEO)
	//AVCodecIDMpeg2videoXvmc  = int(C.AV_CODEC_ID_MPEG2VIDEO_XVMC)
	AVCodecIDMpeg4           = int(C.AV_CODEC_ID_MPEG4)
	AVCodecIDMpeg4systems    = int(C.AV_CODEC_ID_MPEG4SYSTEMS)
	AVCodecIDMpl2            = int(C.AV_CODEC_ID_MPL2)
	AVCodecIDMsa1            = int(C.AV_CODEC_ID_MSA1)
	AVCodecIDMsmpeg4v1       = int(C.AV_CODEC_ID_MSMPEG4V1)
	AVCodecIDMsmpeg4v2       = int(C.AV_CODEC_ID_MSMPEG4V2)
	AVCodecIDMsmpeg4v3       = int(C.AV_CODEC_ID_MSMPEG4V3)
	AVCodecIDMsrle           = int(C.AV_CODEC_ID_MSRLE)
	AVCodecIDMss1            = int(C.AV_CODEC_ID_MSS1)
	AVCodecIDMss2            = int(C.AV_CODEC_ID_MSS2)
	AVCodecIDMsvideo1        = int(C.AV_CODEC_ID_MSVIDEO1)
	AVCodecIDMszh            = int(C.AV_CODEC_ID_MSZH)
	AVCodecIDMts2            = int(C.AV_CODEC_ID_MTS2)
	AVCodecIDMusepack7       = int(C.AV_CODEC_ID_MUSEPACK7)
	AVCodecIDMusepack8       = int(C.AV_CODEC_ID_MUSEPACK8)
	AVCodecIDMvc1            = int(C.AV_CODEC_ID_MVC1)
	AVCodecIDMvc2            = int(C.AV_CODEC_ID_MVC2)
	AVCodecIDMxpeg           = int(C.AV_CODEC_ID_MXPEG)
	AVCodecIDNellymoser      = int(C.AV_CODEC_ID_NELLYMOSER)
	AVCodecIDNone            = int(C.AV_CODEC_ID_NONE)
	AVCodecIDNuv             = int(C.AV_CODEC_ID_NUV)
	AVCodecIDOn2avc          = int(C.AV_CODEC_ID_ON2AVC)
	AVCodecIDOpus            = int(C.AV_CODEC_ID_OPUS)
	AVCodecIDOtf             = int(C.AV_CODEC_ID_OTF)
	AVCodecIDPafAudio        = int(C.AV_CODEC_ID_PAF_AUDIO)
	AVCodecIDPafVideo        = int(C.AV_CODEC_ID_PAF_VIDEO)
	AVCodecIDPam             = int(C.AV_CODEC_ID_PAM)
	AVCodecIDPbm             = int(C.AV_CODEC_ID_PBM)
	AVCodecIDPcmAlaw         = int(C.AV_CODEC_ID_PCM_ALAW)
	AVCodecIDPcmBluray       = int(C.AV_CODEC_ID_PCM_BLURAY)
	AVCodecIDPcmDvd          = int(C.AV_CODEC_ID_PCM_DVD)
	AVCodecIDPcmF32be        = int(C.AV_CODEC_ID_PCM_F32BE)
	AVCodecIDPcmF32le        = int(C.AV_CODEC_ID_PCM_F32LE)
	AVCodecIDPcmF64be        = int(C.AV_CODEC_ID_PCM_F64BE)
	AVCodecIDPcmF64le        = int(C.AV_CODEC_ID_PCM_F64LE)
	AVCodecIDPcmLxf          = int(C.AV_CODEC_ID_PCM_LXF)
	AVCodecIDPcmMulaw        = int(C.AV_CODEC_ID_PCM_MULAW)
	AVCodecIDPcmS16be        = int(C.AV_CODEC_ID_PCM_S16BE)
	AVCodecIDPcmS16bePlanar  = int(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	AVCodecIDPcmS16le        = int(C.AV_CODEC_ID_PCM_S16LE)
	AVCodecIDPcmS16lePlanar  = int(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	AVCodecIDPcmS24be        = int(C.AV_CODEC_ID_PCM_S24BE)
	AVCodecIDPcmS24daud      = int(C.AV_CODEC_ID_PCM_S24DAUD)
	AVCodecIDPcmS24le        = int(C.AV_CODEC_ID_PCM_S24LE)
	AVCodecIDPcmS24lePlanar  = int(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	AVCodecIDPcmS32be        = int(C.AV_CODEC_ID_PCM_S32BE)
	AVCodecIDPcmS32le        = int(C.AV_CODEC_ID_PCM_S32LE)
	AVCodecIDPcmS32lePlanar  = int(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	AVCodecIDPcmS8           = int(C.AV_CODEC_ID_PCM_S8)
	AVCodecIDPcmS8Planar     = int(C.AV_CODEC_ID_PCM_S8_PLANAR)
	AVCodecIDPcmU16be        = int(C.AV_CODEC_ID_PCM_U16BE)
	AVCodecIDPcmU16le        = int(C.AV_CODEC_ID_PCM_U16LE)
	AVCodecIDPcmU24be        = int(C.AV_CODEC_ID_PCM_U24BE)
	AVCodecIDPcmU24le        = int(C.AV_CODEC_ID_PCM_U24LE)
	AVCodecIDPcmU32be        = int(C.AV_CODEC_ID_PCM_U32BE)
	AVCodecIDPcmU32le        = int(C.AV_CODEC_ID_PCM_U32LE)
	AVCodecIDPcmU8           = int(C.AV_CODEC_ID_PCM_U8)
	AVCodecIDPcmZork         = int(C.AV_CODEC_ID_PCM_ZORK)
	AVCodecIDPcx             = int(C.AV_CODEC_ID_PCX)
	AVCodecIDPgm             = int(C.AV_CODEC_ID_PGM)
	AVCodecIDPgmyuv          = int(C.AV_CODEC_ID_PGMYUV)
	AVCodecIDPictor          = int(C.AV_CODEC_ID_PICTOR)
	AVCodecIDPjs             = int(C.AV_CODEC_ID_PJS)
	AVCodecIDPng             = int(C.AV_CODEC_ID_PNG)
	AVCodecIDPpm             = int(C.AV_CODEC_ID_PPM)
	AVCodecIDProbe           = int(C.AV_CODEC_ID_PROBE)
	AVCodecIDProres          = int(C.AV_CODEC_ID_PRORES)
	AVCodecIDPtx             = int(C.AV_CODEC_ID_PTX)
	AVCodecIDQcelp           = int(C.AV_CODEC_ID_QCELP)
	AVCodecIDQdm2            = int(C.AV_CODEC_ID_QDM2)
	AVCodecIDQdmc            = int(C.AV_CODEC_ID_QDMC)
	AVCodecIDQdraw           = int(C.AV_CODEC_ID_QDRAW)
	AVCodecIDQpeg            = int(C.AV_CODEC_ID_QPEG)
	AVCodecIDQtrle           = int(C.AV_CODEC_ID_QTRLE)
	AVCodecIDR10k            = int(C.AV_CODEC_ID_R10K)
	AVCodecIDR210            = int(C.AV_CODEC_ID_R210)
	AVCodecIDRalf            = int(C.AV_CODEC_ID_RALF)
	AVCodecIDRawvideo        = int(C.AV_CODEC_ID_RAWVIDEO)
	AVCodecIDRa144           = int(C.AV_CODEC_ID_RA_144)
	AVCodecIDRa288           = int(C.AV_CODEC_ID_RA_288)
	AVCodecIDRealtext        = int(C.AV_CODEC_ID_REALTEXT)
	AVCodecIDRl2             = int(C.AV_CODEC_ID_RL2)
	AVCodecIDRoq             = int(C.AV_CODEC_ID_ROQ)
	AVCodecIDRoqDpcm         = int(C.AV_CODEC_ID_ROQ_DPCM)
	AVCodecIDRpza            = int(C.AV_CODEC_ID_RPZA)
	AVCodecIDRv10            = int(C.AV_CODEC_ID_RV10)
	AVCodecIDRv20            = int(C.AV_CODEC_ID_RV20)
	AVCodecIDRv30            = int(C.AV_CODEC_ID_RV30)
	AVCodecIDRv40            = int(C.AV_CODEC_ID_RV40)
	AVCodecIDS302m           = int(C.AV_CODEC_ID_S302M)
	AVCodecIDSami            = int(C.AV_CODEC_ID_SAMI)
	AVCodecIDSanm            = int(C.AV_CODEC_ID_SANM)
	AVCodecIDSgi             = int(C.AV_CODEC_ID_SGI)
	AVCodecIDSgirle          = int(C.AV_CODEC_ID_SGIRLE)
	AVCodecIDShorten         = int(C.AV_CODEC_ID_SHORTEN)
	AVCodecIDSipr            = int(C.AV_CODEC_ID_SIPR)
	AVCodecIDSmackaudio      = int(C.AV_CODEC_ID_SMACKAUDIO)
	AVCodecIDSmackvideo      = int(C.AV_CODEC_ID_SMACKVIDEO)
	AVCodecIDSmc             = int(C.AV_CODEC_ID_SMC)
	AVCodecIDSmpteKlv        = int(C.AV_CODEC_ID_SMPTE_KLV)
	AVCodecIDSmv             = int(C.AV_CODEC_ID_SMV)
	AVCodecIDSmvjpeg         = int(C.AV_CODEC_ID_SMVJPEG)
	AVCodecIDSnow            = int(C.AV_CODEC_ID_SNOW)
	AVCodecIDSolDpcm         = int(C.AV_CODEC_ID_SOL_DPCM)
	AVCodecIDSonic           = int(C.AV_CODEC_ID_SONIC)
	AVCodecIDSonicLs         = int(C.AV_CODEC_ID_SONIC_LS)
	AVCodecIDSp5x            = int(C.AV_CODEC_ID_SP5X)
	AVCodecIDSpeex           = int(C.AV_CODEC_ID_SPEEX)
	AVCodecIDSrt             = int(C.AV_CODEC_ID_SRT)
	AVCodecIDSsa             = int(C.AV_CODEC_ID_SSA)
	AVCodecIDSubrip          = int(C.AV_CODEC_ID_SUBRIP)
	AVCodecIDSubviewer       = int(C.AV_CODEC_ID_SUBVIEWER)
	AVCodecIDSubviewer1      = int(C.AV_CODEC_ID_SUBVIEWER1)
	AVCodecIDSunrast         = int(C.AV_CODEC_ID_SUNRAST)
	AVCodecIDSvq1            = int(C.AV_CODEC_ID_SVQ1)
	AVCodecIDSvq3            = int(C.AV_CODEC_ID_SVQ3)
	AVCodecIDTak             = int(C.AV_CODEC_ID_TAK)
	AVCodecIDTarga           = int(C.AV_CODEC_ID_TARGA)
	AVCodecIDTargaY216       = int(C.AV_CODEC_ID_TARGA_Y216)
	AVCodecIDText            = int(C.AV_CODEC_ID_TEXT)
	AVCodecIDTgq             = int(C.AV_CODEC_ID_TGQ)
	AVCodecIDTgv             = int(C.AV_CODEC_ID_TGV)
	AVCodecIDTheora          = int(C.AV_CODEC_ID_THEORA)
	AVCodecIDThp             = int(C.AV_CODEC_ID_THP)
	AVCodecIDTiertexseqvideo = int(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	AVCodecIDTiff            = int(C.AV_CODEC_ID_TIFF)
	AVCodecIDTimedID3        = int(C.AV_CODEC_ID_TIMED_ID3)
	AVCodecIDTmv             = int(C.AV_CODEC_ID_TMV)
	AVCodecIDTqi             = int(C.AV_CODEC_ID_TQI)
	AVCodecIDTruehd          = int(C.AV_CODEC_ID_TRUEHD)
	AVCodecIDTruemotion1     = int(C.AV_CODEC_ID_TRUEMOTION1)
	AVCodecIDTruemotion2     = int(C.AV_CODEC_ID_TRUEMOTION2)
	AVCodecIDTruespeech      = int(C.AV_CODEC_ID_TRUESPEECH)
	AVCodecIDTscc            = int(C.AV_CODEC_ID_TSCC)
	AVCodecIDTscc2           = int(C.AV_CODEC_ID_TSCC2)
	AVCodecIDTta             = int(C.AV_CODEC_ID_TTA)
	AVCodecIDTtf             = int(C.AV_CODEC_ID_TTF)
	AVCodecIDTwinvq          = int(C.AV_CODEC_ID_TWINVQ)
	AVCodecIDTxd             = int(C.AV_CODEC_ID_TXD)
	AVCodecIDUlti            = int(C.AV_CODEC_ID_ULTI)
	AVCodecIDUtvideo         = int(C.AV_CODEC_ID_UTVIDEO)
	AVCodecIDV210            = int(C.AV_CODEC_ID_V210)
	AVCodecIDV210x           = int(C.AV_CODEC_ID_V210X)
	AVCodecIDV308            = int(C.AV_CODEC_ID_V308)
	AVCodecIDV408            = int(C.AV_CODEC_ID_V408)
	AVCodecIDV410            = int(C.AV_CODEC_ID_V410)
	AVCodecIDVb              = int(C.AV_CODEC_ID_VB)
	AVCodecIDVble            = int(C.AV_CODEC_ID_VBLE)
	AVCodecIDVc1             = int(C.AV_CODEC_ID_VC1)
	AVCodecIDVc1image        = int(C.AV_CODEC_ID_VC1IMAGE)
	AVCodecIDVcr1            = int(C.AV_CODEC_ID_VCR1)
	//AVCodecIDVima            = int(C.AV_CODEC_ID_VIMA)
	AVCodecIDVixl     = int(C.AV_CODEC_ID_VIXL)
	AVCodecIDVmdaudio = int(C.AV_CODEC_ID_VMDAUDIO)
	AVCodecIDVmdvideo = int(C.AV_CODEC_ID_VMDVIDEO)
	AVCodecIDVmnc     = int(C.AV_CODEC_ID_VMNC)
	AVCodecIDVorbis   = int(C.AV_CODEC_ID_VORBIS)
	//AVCodecIDVoxware         = int(C.AV_CODEC_ID_VOXWARE)
	AVCodecIDVp3          = int(C.AV_CODEC_ID_VP3)
	AVCodecIDVp5          = int(C.AV_CODEC_ID_VP5)
	AVCodecIDVp6          = int(C.AV_CODEC_ID_VP6)
	AVCodecIDVp6a         = int(C.AV_CODEC_ID_VP6A)
	AVCodecIDVp6f         = int(C.AV_CODEC_ID_VP6F)
	AVCodecIDVp7          = int(C.AV_CODEC_ID_VP7)
	AVCodecIDVp8          = int(C.AV_CODEC_ID_VP8)
	AVCodecIDVp9          = int(C.AV_CODEC_ID_VP9)
	AVCodecIDVplayer      = int(C.AV_CODEC_ID_VPLAYER)
	AVCodecIDWavpack      = int(C.AV_CODEC_ID_WAVPACK)
	AVCodecIDWebp         = int(C.AV_CODEC_ID_WEBP)
	AVCodecIDWebvtt       = int(C.AV_CODEC_ID_WEBVTT)
	AVCodecIDWestwoodSnd1 = int(C.AV_CODEC_ID_WESTWOOD_SND1)
	AVCodecIDWmalossless  = int(C.AV_CODEC_ID_WMALOSSLESS)
	AVCodecIDWmapro       = int(C.AV_CODEC_ID_WMAPRO)
	AVCodecIDWmav1        = int(C.AV_CODEC_ID_WMAV1)
	AVCodecIDWmav2        = int(C.AV_CODEC_ID_WMAV2)
	AVCodecIDWmavoice     = int(C.AV_CODEC_ID_WMAVOICE)
	AVCodecIDWmv1         = int(C.AV_CODEC_ID_WMV1)
	AVCodecIDWmv2         = int(C.AV_CODEC_ID_WMV2)
	AVCodecIDWmv3         = int(C.AV_CODEC_ID_WMV3)
	AVCodecIDWmv3image    = int(C.AV_CODEC_ID_WMV3IMAGE)
	AVCodecIDWnv1         = int(C.AV_CODEC_ID_WNV1)
	AVCodecIDWsVqa        = int(C.AV_CODEC_ID_WS_VQA)
	AVCodecIDXanDpcm      = int(C.AV_CODEC_ID_XAN_DPCM)
	AVCodecIDXanWc3       = int(C.AV_CODEC_ID_XAN_WC3)
	AVCodecIDXanWc4       = int(C.AV_CODEC_ID_XAN_WC4)
	AVCodecIDXbin         = int(C.AV_CODEC_ID_XBIN)
	AVCodecIDXbm          = int(C.AV_CODEC_ID_XBM)
	AVCodecIDXface        = int(C.AV_CODEC_ID_XFACE)
	AVCodecIDXsub         = int(C.AV_CODEC_ID_XSUB)
	AVCodecIDXwd          = int(C.AV_CODEC_ID_XWD)
	AVCodecIDY41p         = int(C.AV_CODEC_ID_Y41P)
	AVCodecIDYop          = int(C.AV_CODEC_ID_YOP)
	AVCodecIDYuv4         = int(C.AV_CODEC_ID_YUV4)
	AVCodecIDZerocodec    = int(C.AV_CODEC_ID_ZEROCODEC)
	AVCodecIDZlib         = int(C.AV_CODEC_ID_ZLIB)
	AVCodecIDZmbv         = int(C.AV_CODEC_ID_ZMBV)
)
