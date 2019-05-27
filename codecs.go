// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package goav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

// Multiple encoders have the same ID and are able to encode compatible streams.
const (
	AvCodecId012v               = int(C.AV_CODEC_ID_012V)
	AvCodecId4xm                = int(C.AV_CODEC_ID_4XM)
	AvCodecId8bps               = int(C.AV_CODEC_ID_8BPS)
	AvCodecId8svxExp            = int(C.AV_CODEC_ID_8SVX_EXP)
	AvCodecId8svxFib            = int(C.AV_CODEC_ID_8SVX_FIB)
	AvCodecIdA64Multi           = int(C.AV_CODEC_ID_A64_MULTI)
	AvCodecIdA64Multi5          = int(C.AV_CODEC_ID_A64_MULTI5)
	AvCodecIdAac                = int(C.AV_CODEC_ID_AAC)
	AvCodecIdAacLatm            = int(C.AV_CODEC_ID_AAC_LATM)
	AvCodecIdAasc               = int(C.AV_CODEC_ID_AASC)
	AvCodecIdAc3                = int(C.AV_CODEC_ID_AC3)
	AvCodecIdAdpcm4xm           = int(C.AV_CODEC_ID_ADPCM_4XM)
	AvCodecIdAdpcmAdx           = int(C.AV_CODEC_ID_ADPCM_ADX)
	AvCodecIdAdpcmAfc           = int(C.AV_CODEC_ID_ADPCM_AFC)
	AvCodecIdAdpcmCt            = int(C.AV_CODEC_ID_ADPCM_CT)
	AvCodecIdAdpcmDtk           = int(C.AV_CODEC_ID_ADPCM_DTK)
	AvCodecIdAdpcmEa            = int(C.AV_CODEC_ID_ADPCM_EA)
	AvCodecIdAdpcmEaMaxisXa     = int(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	AvCodecIdAdpcmEaR1          = int(C.AV_CODEC_ID_ADPCM_EA_R1)
	AvCodecIdAdpcmEaR2          = int(C.AV_CODEC_ID_ADPCM_EA_R2)
	AvCodecIdAdpcmEaR3          = int(C.AV_CODEC_ID_ADPCM_EA_R3)
	AvCodecIdAdpcmEaXas         = int(C.AV_CODEC_ID_ADPCM_EA_XAS)
	AvCodecIdAdpcmG722          = int(C.AV_CODEC_ID_ADPCM_G722)
	AvCodecIdAdpcmG726          = int(C.AV_CODEC_ID_ADPCM_G726)
	AvCodecIdAdpcmG726le        = int(C.AV_CODEC_ID_ADPCM_G726LE)
	AvCodecIdAdpcmImaAmv        = int(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	AvCodecIdAdpcmImaApc        = int(C.AV_CODEC_ID_ADPCM_IMA_APC)
	AvCodecIdAdpcmImaDk3        = int(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	AvCodecIdAdpcmImaDk4        = int(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	AvCodecIdAdpcmImaEaEacs     = int(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	AvCodecIdAdpcmImaEaSead     = int(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	AvCodecIdAdpcmImaIss        = int(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	AvCodecIdAdpcmImaOki        = int(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	AvCodecIdAdpcmImaQt         = int(C.AV_CODEC_ID_ADPCM_IMA_QT)
	AvCodecIdAdpcmImaRad        = int(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	AvCodecIdAdpcmImaSmjpeg     = int(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	AvCodecIdAdpcmImaWav        = int(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	AvCodecIdAdpcmImaWs         = int(C.AV_CODEC_ID_ADPCM_IMA_WS)
	AvCodecIdAdpcmMs            = int(C.AV_CODEC_ID_ADPCM_MS)
	AvCodecIdAdpcmSbpro2        = int(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	AvCodecIdAdpcmSbpro3        = int(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	AvCodecIdAdpcmSbpro4        = int(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	AvCodecIdAdpcmSwf           = int(C.AV_CODEC_ID_ADPCM_SWF)
	AvCodecIdAdpcmThp           = int(C.AV_CODEC_ID_ADPCM_THP)
	AvCodecIdAdpcmVima          = int(C.AV_CODEC_ID_ADPCM_VIMA)
	AvCodecIdAdpcmXa            = int(C.AV_CODEC_ID_ADPCM_XA)
	AvCodecIdAdpcmYamaha        = int(C.AV_CODEC_ID_ADPCM_YAMAHA)
	AvCodecIdAic                = int(C.AV_CODEC_ID_AIC)
	AvCodecIdAlac               = int(C.AV_CODEC_ID_ALAC)
	AvCodecIdAliasPix           = int(C.AV_CODEC_ID_ALIAS_PIX)
	AvCodecIdAmrNb              = int(C.AV_CODEC_ID_AMR_NB)
	AvCodecIdAmrWb              = int(C.AV_CODEC_ID_AMR_WB)
	AvCodecIdAmv                = int(C.AV_CODEC_ID_AMV)
	AvCodecIdAnm                = int(C.AV_CODEC_ID_ANM)
	AvCodecIdAnsi               = int(C.AV_CODEC_ID_ANSI)
	AvCodecIdApe                = int(C.AV_CODEC_ID_APE)
	AvCodecIdAss                = int(C.AV_CODEC_ID_ASS)
	AvCodecIdAsv1               = int(C.AV_CODEC_ID_ASV1)
	AvCodecIdAsv2               = int(C.AV_CODEC_ID_ASV2)
	AvCodecIdAtrac1             = int(C.AV_CODEC_ID_ATRAC1)
	AvCodecIdAtrac3             = int(C.AV_CODEC_ID_ATRAC3)
	AvCodecIdAtrac3p            = int(C.AV_CODEC_ID_ATRAC3P)
	AvCodecIdAura               = int(C.AV_CODEC_ID_AURA)
	AvCodecIdAura2              = int(C.AV_CODEC_ID_AURA2)
	AvCodecIdAvrn               = int(C.AV_CODEC_ID_AVRN)
	AvCodecIdAvrp               = int(C.AV_CODEC_ID_AVRP)
	AvCodecIdAvs                = int(C.AV_CODEC_ID_AVS)
	AvCodecIdAvui               = int(C.AV_CODEC_ID_AVUI)
	AvCodecIdAyuv               = int(C.AV_CODEC_ID_AYUV)
	AvCodecIdBethsoftvid        = int(C.AV_CODEC_ID_BETHSOFTVID)
	AvCodecIdBfi                = int(C.AV_CODEC_ID_BFI)
	AvCodecIdBinkaudioDct       = int(C.AV_CODEC_ID_BINKAUDIO_DCT)
	AvCodecIdBinkaudioRdft      = int(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	AvCodecIdBinkvideo          = int(C.AV_CODEC_ID_BINKVIDEO)
	AvCodecIdBintext            = int(C.AV_CODEC_ID_BINTEXT)
	AvCodecIdBinData            = int(C.AV_CODEC_ID_BIN_DATA)
	AvCodecIdBmp                = int(C.AV_CODEC_ID_BMP)
	AvCodecIdBmvAudio           = int(C.AV_CODEC_ID_BMV_AUDIO)
	AvCodecIdBmvVideo           = int(C.AV_CODEC_ID_BMV_VIDEO)
	AvCodecIdBrenderPix         = int(C.AV_CODEC_ID_BRENDER_PIX)
	AvCodecIdC93                = int(C.AV_CODEC_ID_C93)
	AvCodecIdCavs               = int(C.AV_CODEC_ID_CAVS)
	AvCodecIdCdgraphics         = int(C.AV_CODEC_ID_CDGRAPHICS)
	AvCodecIdCdxl               = int(C.AV_CODEC_ID_CDXL)
	AvCodecIdCelt               = int(C.AV_CODEC_ID_CELT)
	AvCodecIdCinepak            = int(C.AV_CODEC_ID_CINEPAK)
	AvCodecIdCljr               = int(C.AV_CODEC_ID_CLJR)
	AvCodecIdCllc               = int(C.AV_CODEC_ID_CLLC)
	AvCodecIdCmv                = int(C.AV_CODEC_ID_CMV)
	AvCodecIdComfortNoise       = int(C.AV_CODEC_ID_COMFORT_NOISE)
	AvCodecIdCook               = int(C.AV_CODEC_ID_COOK)
	AvCodecIdCpia               = int(C.AV_CODEC_ID_CPIA)
	AvCodecIdCscd               = int(C.AV_CODEC_ID_CSCD)
	AvCodecIdCyuv               = int(C.AV_CODEC_ID_CYUV)
	AvCodecIdDfa                = int(C.AV_CODEC_ID_DFA)
	AvCodecIdDirac              = int(C.AV_CODEC_ID_DIRAC)
	AvCodecIdDnxhd              = int(C.AV_CODEC_ID_DNXHD)
	AvCodecIdDpx                = int(C.AV_CODEC_ID_DPX)
	AvCodecIdDsdLsbf            = int(C.AV_CODEC_ID_DSD_LSBF)
	AvCodecIdDsdLsbfPlanar      = int(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	AvCodecIdDsdMsbf            = int(C.AV_CODEC_ID_DSD_MSBF)
	AvCodecIdDsdMsbfPlanar      = int(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	AvCodecIdDsicinaudio        = int(C.AV_CODEC_ID_DSICINAUDIO)
	AvCodecIdDsicinvideo        = int(C.AV_CODEC_ID_DSICINVIDEO)
	AvCodecIdDts                = int(C.AV_CODEC_ID_DTS)
	AvCodecIdDvaudio            = int(C.AV_CODEC_ID_DVAUDIO)
	AvCodecIdDvbSubtitle        = int(C.AV_CODEC_ID_DVB_SUBTITLE)
	AvCodecIdDvbTeletext        = int(C.AV_CODEC_ID_DVB_TELETEXT)
	AvCodecIdDvdNav             = int(C.AV_CODEC_ID_DVD_NAV)
	AvCodecIdDvdSubtitle        = int(C.AV_CODEC_ID_DVD_SUBTITLE)
	AvCodecIdDvvideo            = int(C.AV_CODEC_ID_DVVIDEO)
	AvCodecIdDxa                = int(C.AV_CODEC_ID_DXA)
	AvCodecIdDxtory             = int(C.AV_CODEC_ID_DXTORY)
	AvCodecIdEac3               = int(C.AV_CODEC_ID_EAC3)
	AvCodecIdEia608             = int(C.AV_CODEC_ID_EIA_608)
	AvCodecIdEscape124          = int(C.AV_CODEC_ID_ESCAPE124)
	AvCodecIdEscape130          = int(C.AV_CODEC_ID_ESCAPE130)
	AvCodecIdEvrc               = int(C.AV_CODEC_ID_EVRC)
	AvCodecIdExr                = int(C.AV_CODEC_ID_EXR)
	AvCodecIdFfmetadata         = int(C.AV_CODEC_ID_FFMETADATA)
	AvCodecIdFfv1               = int(C.AV_CODEC_ID_FFV1)
	AvCodecIdFfvhuff            = int(C.AV_CODEC_ID_FFVHUFF)
	AvCodecIdFfwavesynth        = int(C.AV_CODEC_ID_FFWAVESYNTH)
	AvCodecIdFic                = int(C.AV_CODEC_ID_FIC)
	AvCodecIdFirstAudio         = int(C.AV_CODEC_ID_FIRST_AUDIO)
	AvCodecIdFirstSubtitle      = int(C.AV_CODEC_ID_FIRST_SUBTITLE)
	AvCodecIdFirstUnknown       = int(C.AV_CODEC_ID_FIRST_UNKNOWN)
	AvCodecIdFlac               = int(C.AV_CODEC_ID_FLAC)
	AvCodecIdFlashsv            = int(C.AV_CODEC_ID_FLASHSV)
	AvCodecIdFlashsv2           = int(C.AV_CODEC_ID_FLASHSV2)
	AvCodecIdFlic               = int(C.AV_CODEC_ID_FLIC)
	AvCodecIdFlv1               = int(C.AV_CODEC_ID_FLV1)
	AvCodecIdFraps              = int(C.AV_CODEC_ID_FRAPS)
	AvCodecIdFrwu               = int(C.AV_CODEC_ID_FRWU)
	AvCodecIdG2m                = int(C.AV_CODEC_ID_G2M)
	AvCodecIdG7231              = int(C.AV_CODEC_ID_G723_1)
	AvCodecIdG729               = int(C.AV_CODEC_ID_G729)
	AvCodecIdGif                = int(C.AV_CODEC_ID_GIF)
	AvCodecIdGsm                = int(C.AV_CODEC_ID_GSM)
	AvCodecIdGsmMs              = int(C.AV_CODEC_ID_GSM_MS)
	AvCodecIdH261               = int(C.AV_CODEC_ID_H261)
	AvCodecIdH263               = int(C.AV_CODEC_ID_H263)
	AvCodecIdH263i              = int(C.AV_CODEC_ID_H263I)
	AvCodecIdH263p              = int(C.AV_CODEC_ID_H263P)
	AvCodecIdH264               = int(C.AV_CODEC_ID_H264)
	AvCodecIdH265               = int(C.AV_CODEC_ID_H265)
	AvCodecIdHdmvPgsSubtitle    = int(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	AvCodecIdHevc               = int(C.AV_CODEC_ID_HEVC)
	AvCodecIdHnm4Video          = int(C.AV_CODEC_ID_HNM4_VIDEO)
	AvCodecIdHuffyuv            = int(C.AV_CODEC_ID_HUFFYUV)
	AvCodecIdIac                = int(C.AV_CODEC_ID_IAC)
	AvCodecIdIdcin              = int(C.AV_CODEC_ID_IDCIN)
	AvCodecIdIdf                = int(C.AV_CODEC_ID_IDF)
	AvCodecIdIffByterun1        = int(C.AV_CODEC_ID_IFF_BYTERUN1)
	AvCodecIdIffIlbm            = int(C.AV_CODEC_ID_IFF_ILBM)
	AvCodecIdIlbc               = int(C.AV_CODEC_ID_ILBC)
	AvCodecIdImc                = int(C.AV_CODEC_ID_IMC)
	AvCodecIdIndeo2             = int(C.AV_CODEC_ID_INDEO2)
	AvCodecIdIndeo3             = int(C.AV_CODEC_ID_INDEO3)
	AvCodecIdIndeo4             = int(C.AV_CODEC_ID_INDEO4)
	AvCodecIdIndeo5             = int(C.AV_CODEC_ID_INDEO5)
	AvCodecIdInterplayDpcm      = int(C.AV_CODEC_ID_INTERPLAY_DPCM)
	AvCodecIdInterplayVideo     = int(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	AvCodecIdJacosub            = int(C.AV_CODEC_ID_JACOSUB)
	AvCodecIdJpeg2000           = int(C.AV_CODEC_ID_JPEG2000)
	AvCodecIdJpegls             = int(C.AV_CODEC_ID_JPEGLS)
	AvCodecIdJv                 = int(C.AV_CODEC_ID_JV)
	AvCodecIdKgv1               = int(C.AV_CODEC_ID_KGV1)
	AvCodecIdKmvc               = int(C.AV_CODEC_ID_KMVC)
	AvCodecIdLagarith           = int(C.AV_CODEC_ID_LAGARITH)
	AvCodecIdLjpeg              = int(C.AV_CODEC_ID_LJPEG)
	AvCodecIdLoco               = int(C.AV_CODEC_ID_LOCO)
	AvCodecIdMace3              = int(C.AV_CODEC_ID_MACE3)
	AvCodecIdMace6              = int(C.AV_CODEC_ID_MACE6)
	AvCodecIdMad                = int(C.AV_CODEC_ID_MAD)
	AvCodecIdMdec               = int(C.AV_CODEC_ID_MDEC)
	AvCodecIdMetasound          = int(C.AV_CODEC_ID_METASOUND)
	AvCodecIdMicrodvd           = int(C.AV_CODEC_ID_MICRODVD)
	AvCodecIdMimic              = int(C.AV_CODEC_ID_MIMIC)
	AvCodecIdMjpeg              = int(C.AV_CODEC_ID_MJPEG)
	AvCodecIdMjpegb             = int(C.AV_CODEC_ID_MJPEGB)
	AvCodecIdMlp                = int(C.AV_CODEC_ID_MLP)
	AvCodecIdMmvideo            = int(C.AV_CODEC_ID_MMVIDEO)
	AvCodecIdMotionpixels       = int(C.AV_CODEC_ID_MOTIONPIXELS)
	AvCodecIdMovText            = int(C.AV_CODEC_ID_MOV_TEXT)
	AvCodecIdMp1                = int(C.AV_CODEC_ID_MP1)
	AvCodecIdMp2                = int(C.AV_CODEC_ID_MP2)
	AvCodecIdMp3                = int(C.AV_CODEC_ID_MP3)
	AvCodecIdMp3adu             = int(C.AV_CODEC_ID_MP3ADU)
	AvCodecIdMp3on4             = int(C.AV_CODEC_ID_MP3ON4)
	AvCodecIdMp4als             = int(C.AV_CODEC_ID_MP4ALS)
	AvCodecIdMpeg1video         = int(C.AV_CODEC_ID_MPEG1VIDEO)
	AvCodecIdMpeg2ts            = int(C.AV_CODEC_ID_MPEG2TS)
	AvCodecIdMpeg2video         = int(C.AV_CODEC_ID_MPEG2VIDEO)
	AvCodecIdMpeg2videoXvmc     = int(C.AV_CODEC_ID_MPEG2VIDEO_XVMC)
	AvCodecIdMpeg4              = int(C.AV_CODEC_ID_MPEG4)
	AvCodecIdMpeg4systems       = int(C.AV_CODEC_ID_MPEG4SYSTEMS)
	AvCodecIdMpl2               = int(C.AV_CODEC_ID_MPL2)
	AvCodecIdMsa1               = int(C.AV_CODEC_ID_MSA1)
	AvCodecIdMsmpeg4v1          = int(C.AV_CODEC_ID_MSMPEG4V1)
	AvCodecIdMsmpeg4v2          = int(C.AV_CODEC_ID_MSMPEG4V2)
	AvCodecIdMsmpeg4v3          = int(C.AV_CODEC_ID_MSMPEG4V3)
	AvCodecIdMsrle              = int(C.AV_CODEC_ID_MSRLE)
	AvCodecIdMss1               = int(C.AV_CODEC_ID_MSS1)
	AvCodecIdMss2               = int(C.AV_CODEC_ID_MSS2)
	AvCodecIdMsvideo1           = int(C.AV_CODEC_ID_MSVIDEO1)
	AvCodecIdMszh               = int(C.AV_CODEC_ID_MSZH)
	AvCodecIdMts2               = int(C.AV_CODEC_ID_MTS2)
	AvCodecIdMusepack7          = int(C.AV_CODEC_ID_MUSEPACK7)
	AvCodecIdMusepack8          = int(C.AV_CODEC_ID_MUSEPACK8)
	AvCodecIdMvc1               = int(C.AV_CODEC_ID_MVC1)
	AvCodecIdMvc2               = int(C.AV_CODEC_ID_MVC2)
	AvCodecIdMxpeg              = int(C.AV_CODEC_ID_MXPEG)
	AvCodecIdNellymoser         = int(C.AV_CODEC_ID_NELLYMOSER)
	AvCodecIdNone               = int(C.AV_CODEC_ID_NONE)
	AvCodecIdNuv                = int(C.AV_CODEC_ID_NUV)
	AvCodecIdOn2avc             = int(C.AV_CODEC_ID_ON2AVC)
	AvCodecIdOpus               = int(C.AV_CODEC_ID_OPUS)
	AvCodecIdOtf                = int(C.AV_CODEC_ID_OTF)
	AvCodecIdPafAudio           = int(C.AV_CODEC_ID_PAF_AUDIO)
	AvCodecIdPafVideo           = int(C.AV_CODEC_ID_PAF_VIDEO)
	AvCodecIdPam                = int(C.AV_CODEC_ID_PAM)
	AvCodecIdPbm                = int(C.AV_CODEC_ID_PBM)
	AvCodecIdPcmAlaw            = int(C.AV_CODEC_ID_PCM_ALAW)
	AvCodecIdPcmBluray          = int(C.AV_CODEC_ID_PCM_BLURAY)
	AvCodecIdPcmDvd             = int(C.AV_CODEC_ID_PCM_DVD)
	AvCodecIdPcmF32be           = int(C.AV_CODEC_ID_PCM_F32BE)
	AvCodecIdPcmF32le           = int(C.AV_CODEC_ID_PCM_F32LE)
	AvCodecIdPcmF64be           = int(C.AV_CODEC_ID_PCM_F64BE)
	AvCodecIdPcmF64le           = int(C.AV_CODEC_ID_PCM_F64LE)
	AvCodecIdPcmLxf             = int(C.AV_CODEC_ID_PCM_LXF)
	AvCodecIdPcmMulaw           = int(C.AV_CODEC_ID_PCM_MULAW)
	AvCodecIdPcmS16be           = int(C.AV_CODEC_ID_PCM_S16BE)
	AvCodecIdPcmS16bePlanar     = int(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	AvCodecIdPcmS16le           = int(C.AV_CODEC_ID_PCM_S16LE)
	AvCodecIdPcmS16lePlanar     = int(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	AvCodecIdPcmS24be           = int(C.AV_CODEC_ID_PCM_S24BE)
	AvCodecIdPcmS24daud         = int(C.AV_CODEC_ID_PCM_S24DAUD)
	AvCodecIdPcmS24le           = int(C.AV_CODEC_ID_PCM_S24LE)
	AvCodecIdPcmS24lePlanar     = int(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	AvCodecIdPcmS32be           = int(C.AV_CODEC_ID_PCM_S32BE)
	AvCodecIdPcmS32le           = int(C.AV_CODEC_ID_PCM_S32LE)
	AvCodecIdPcmS32lePlanar     = int(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	AvCodecIdPcmS8              = int(C.AV_CODEC_ID_PCM_S8)
	AvCodecIdPcmS8Planar        = int(C.AV_CODEC_ID_PCM_S8_PLANAR)
	AvCodecIdPcmU16be           = int(C.AV_CODEC_ID_PCM_U16BE)
	AvCodecIdPcmU16le           = int(C.AV_CODEC_ID_PCM_U16LE)
	AvCodecIdPcmU24be           = int(C.AV_CODEC_ID_PCM_U24BE)
	AvCodecIdPcmU24le           = int(C.AV_CODEC_ID_PCM_U24LE)
	AvCodecIdPcmU32be           = int(C.AV_CODEC_ID_PCM_U32BE)
	AvCodecIdPcmU32le           = int(C.AV_CODEC_ID_PCM_U32LE)
	AvCodecIdPcmU8              = int(C.AV_CODEC_ID_PCM_U8)
	AvCodecIdPcmZork            = int(C.AV_CODEC_ID_PCM_ZORK)
	AvCodecIdPcx                = int(C.AV_CODEC_ID_PCX)
	AvCodecIdPgm                = int(C.AV_CODEC_ID_PGM)
	AvCodecIdPgmyuv             = int(C.AV_CODEC_ID_PGMYUV)
	AvCodecIdPictor             = int(C.AV_CODEC_ID_PICTOR)
	AvCodecIdPjs                = int(C.AV_CODEC_ID_PJS)
	AvCodecIdPng                = int(C.AV_CODEC_ID_PNG)
	AvCodecIdPpm                = int(C.AV_CODEC_ID_PPM)
	AvCodecIdProbe              = int(C.AV_CODEC_ID_PROBE)
	AvCodecIdProres             = int(C.AV_CODEC_ID_PRORES)
	AvCodecIdPtx                = int(C.AV_CODEC_ID_PTX)
	AvCodecIdQcelp              = int(C.AV_CODEC_ID_QCELP)
	AvCodecIdQdm2               = int(C.AV_CODEC_ID_QDM2)
	AvCodecIdQdmc               = int(C.AV_CODEC_ID_QDMC)
	AvCodecIdQdraw              = int(C.AV_CODEC_ID_QDRAW)
	AvCodecIdQpeg               = int(C.AV_CODEC_ID_QPEG)
	AvCodecIdQtrle              = int(C.AV_CODEC_ID_QTRLE)
	AvCodecIdR10k               = int(C.AV_CODEC_ID_R10K)
	AvCodecIdR210               = int(C.AV_CODEC_ID_R210)
	AvCodecIdRalf               = int(C.AV_CODEC_ID_RALF)
	AvCodecIdRawvideo           = int(C.AV_CODEC_ID_RAWVIDEO)
	AvCodecIdRa144              = int(C.AV_CODEC_ID_RA_144)
	AvCodecIdRa288              = int(C.AV_CODEC_ID_RA_288)
	AvCodecIdRealtext           = int(C.AV_CODEC_ID_REALTEXT)
	AvCodecIdRl2                = int(C.AV_CODEC_ID_RL2)
	AvCodecIdRoq                = int(C.AV_CODEC_ID_ROQ)
	AvCodecIdRoqDpcm            = int(C.AV_CODEC_ID_ROQ_DPCM)
	AvCodecIdRpza               = int(C.AV_CODEC_ID_RPZA)
	AvCodecIdRv10               = int(C.AV_CODEC_ID_RV10)
	AvCodecIdRv20               = int(C.AV_CODEC_ID_RV20)
	AvCodecIdRv30               = int(C.AV_CODEC_ID_RV30)
	AvCodecIdRv40               = int(C.AV_CODEC_ID_RV40)
	AvCodecIdS302m              = int(C.AV_CODEC_ID_S302M)
	AvCodecIdSami               = int(C.AV_CODEC_ID_SAMI)
	AvCodecIdSanm               = int(C.AV_CODEC_ID_SANM)
	AvCodecIdSgi                = int(C.AV_CODEC_ID_SGI)
	AV_CODEC_ID_SGIRLE          = int(C.AV_CODEC_ID_SGIRLE)
	AV_CODEC_ID_SHORTEN         = int(C.AV_CODEC_ID_SHORTEN)
	AV_CODEC_ID_SIPR            = int(C.AV_CODEC_ID_SIPR)
	AV_CODEC_ID_SMACKAUDIO      = int(C.AV_CODEC_ID_SMACKAUDIO)
	AV_CODEC_ID_SMACKVIDEO      = int(C.AV_CODEC_ID_SMACKVIDEO)
	AV_CODEC_ID_SMC             = int(C.AV_CODEC_ID_SMC)
	AV_CODEC_ID_SMPTE_KLV       = int(C.AV_CODEC_ID_SMPTE_KLV)
	AV_CODEC_ID_SMV             = int(C.AV_CODEC_ID_SMV)
	AV_CODEC_ID_SMVJPEG         = int(C.AV_CODEC_ID_SMVJPEG)
	AV_CODEC_ID_SNOW            = int(C.AV_CODEC_ID_SNOW)
	AV_CODEC_ID_SOL_DPCM        = int(C.AV_CODEC_ID_SOL_DPCM)
	AV_CODEC_ID_SONIC           = int(C.AV_CODEC_ID_SONIC)
	AV_CODEC_ID_SONIC_LS        = int(C.AV_CODEC_ID_SONIC_LS)
	AV_CODEC_ID_SP5X            = int(C.AV_CODEC_ID_SP5X)
	AV_CODEC_ID_SPEEX           = int(C.AV_CODEC_ID_SPEEX)
	AV_CODEC_ID_SRT             = int(C.AV_CODEC_ID_SRT)
	AV_CODEC_ID_SSA             = int(C.AV_CODEC_ID_SSA)
	AV_CODEC_ID_SUBRIP          = int(C.AV_CODEC_ID_SUBRIP)
	AV_CODEC_ID_SUBVIEWER       = int(C.AV_CODEC_ID_SUBVIEWER)
	AV_CODEC_ID_SUBVIEWER1      = int(C.AV_CODEC_ID_SUBVIEWER1)
	AV_CODEC_ID_SUNRAST         = int(C.AV_CODEC_ID_SUNRAST)
	AV_CODEC_ID_SVQ1            = int(C.AV_CODEC_ID_SVQ1)
	AV_CODEC_ID_SVQ3            = int(C.AV_CODEC_ID_SVQ3)
	AV_CODEC_ID_TAK             = int(C.AV_CODEC_ID_TAK)
	AV_CODEC_ID_TARGA           = int(C.AV_CODEC_ID_TARGA)
	AV_CODEC_ID_TARGA_Y216      = int(C.AV_CODEC_ID_TARGA_Y216)
	AV_CODEC_ID_TEXT            = int(C.AV_CODEC_ID_TEXT)
	AV_CODEC_ID_TGQ             = int(C.AV_CODEC_ID_TGQ)
	AV_CODEC_ID_TGV             = int(C.AV_CODEC_ID_TGV)
	AV_CODEC_ID_THEORA          = int(C.AV_CODEC_ID_THEORA)
	AV_CODEC_ID_THP             = int(C.AV_CODEC_ID_THP)
	AV_CODEC_ID_TIERTEXSEQVIDEO = int(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	AV_CODEC_ID_TIFF            = int(C.AV_CODEC_ID_TIFF)
	AV_CODEC_ID_TIMED_ID3       = int(C.AV_CODEC_ID_TIMED_ID3)
	AV_CODEC_ID_TMV             = int(C.AV_CODEC_ID_TMV)
	AV_CODEC_ID_TQI             = int(C.AV_CODEC_ID_TQI)
	AV_CODEC_ID_TRUEHD          = int(C.AV_CODEC_ID_TRUEHD)
	AV_CODEC_ID_TRUEMOTION1     = int(C.AV_CODEC_ID_TRUEMOTION1)
	AV_CODEC_ID_TRUEMOTION2     = int(C.AV_CODEC_ID_TRUEMOTION2)
	AV_CODEC_ID_TRUESPEECH      = int(C.AV_CODEC_ID_TRUESPEECH)
	AV_CODEC_ID_TSCC            = int(C.AV_CODEC_ID_TSCC)
	AV_CODEC_ID_TSCC2           = int(C.AV_CODEC_ID_TSCC2)
	AV_CODEC_ID_TTA             = int(C.AV_CODEC_ID_TTA)
	AV_CODEC_ID_TTF             = int(C.AV_CODEC_ID_TTF)
	AV_CODEC_ID_TWINVQ          = int(C.AV_CODEC_ID_TWINVQ)
	AV_CODEC_ID_TXD             = int(C.AV_CODEC_ID_TXD)
	AV_CODEC_ID_ULTI            = int(C.AV_CODEC_ID_ULTI)
	AV_CODEC_ID_UTVIDEO         = int(C.AV_CODEC_ID_UTVIDEO)
	AV_CODEC_ID_V210            = int(C.AV_CODEC_ID_V210)
	AV_CODEC_ID_V210X           = int(C.AV_CODEC_ID_V210X)
	AV_CODEC_ID_V308            = int(C.AV_CODEC_ID_V308)
	AV_CODEC_ID_V408            = int(C.AV_CODEC_ID_V408)
	AV_CODEC_ID_V410            = int(C.AV_CODEC_ID_V410)
	AV_CODEC_ID_VB              = int(C.AV_CODEC_ID_VB)
	AV_CODEC_ID_VBLE            = int(C.AV_CODEC_ID_VBLE)
	AV_CODEC_ID_VC1             = int(C.AV_CODEC_ID_VC1)
	AV_CODEC_ID_VC1IMAGE        = int(C.AV_CODEC_ID_VC1IMAGE)
	AV_CODEC_ID_VCR1            = int(C.AV_CODEC_ID_VCR1)
	// AV_CODEC_ID_VIMA             = int(C.AV_CODEC_ID_VIMA)
	AV_CODEC_ID_VIXL     = int(C.AV_CODEC_ID_VIXL)
	AV_CODEC_ID_VMDAUDIO = int(C.AV_CODEC_ID_VMDAUDIO)
	AV_CODEC_ID_VMDVIDEO = int(C.AV_CODEC_ID_VMDVIDEO)
	AV_CODEC_ID_VMNC     = int(C.AV_CODEC_ID_VMNC)
	AV_CODEC_ID_VORBIS   = int(C.AV_CODEC_ID_VORBIS)
	// AV_CODEC_ID_VOXWARE          = int(C.AV_CODEC_ID_VOXWARE)
	AV_CODEC_ID_VP3           = int(C.AV_CODEC_ID_VP3)
	AV_CODEC_ID_VP5           = int(C.AV_CODEC_ID_VP5)
	AV_CODEC_ID_VP6           = int(C.AV_CODEC_ID_VP6)
	AV_CODEC_ID_VP6A          = int(C.AV_CODEC_ID_VP6A)
	AV_CODEC_ID_VP6F          = int(C.AV_CODEC_ID_VP6F)
	AV_CODEC_ID_VP7           = int(C.AV_CODEC_ID_VP7)
	AV_CODEC_ID_VP8           = int(C.AV_CODEC_ID_VP8)
	AV_CODEC_ID_VP9           = int(C.AV_CODEC_ID_VP9)
	AV_CODEC_ID_VPLAYER       = int(C.AV_CODEC_ID_VPLAYER)
	AV_CODEC_ID_WAVPACK       = int(C.AV_CODEC_ID_WAVPACK)
	AV_CODEC_ID_WEBP          = int(C.AV_CODEC_ID_WEBP)
	AV_CODEC_ID_WEBVTT        = int(C.AV_CODEC_ID_WEBVTT)
	AV_CODEC_ID_WESTWOOD_SND1 = int(C.AV_CODEC_ID_WESTWOOD_SND1)
	AV_CODEC_ID_WMALOSSLESS   = int(C.AV_CODEC_ID_WMALOSSLESS)
	AV_CODEC_ID_WMAPRO        = int(C.AV_CODEC_ID_WMAPRO)
	AV_CODEC_ID_WMAV1         = int(C.AV_CODEC_ID_WMAV1)
	AV_CODEC_ID_WMAV2         = int(C.AV_CODEC_ID_WMAV2)
	AV_CODEC_ID_WMAVOICE      = int(C.AV_CODEC_ID_WMAVOICE)
	AV_CODEC_ID_WMV1          = int(C.AV_CODEC_ID_WMV1)
	AV_CODEC_ID_WMV2          = int(C.AV_CODEC_ID_WMV2)
	AV_CODEC_ID_WMV3          = int(C.AV_CODEC_ID_WMV3)
	AV_CODEC_ID_WMV3IMAGE     = int(C.AV_CODEC_ID_WMV3IMAGE)
	AV_CODEC_ID_WNV1          = int(C.AV_CODEC_ID_WNV1)
	AV_CODEC_ID_WS_VQA        = int(C.AV_CODEC_ID_WS_VQA)
	AV_CODEC_ID_XAN_DPCM      = int(C.AV_CODEC_ID_XAN_DPCM)
	AV_CODEC_ID_XAN_WC3       = int(C.AV_CODEC_ID_XAN_WC3)
	AV_CODEC_ID_XAN_WC4       = int(C.AV_CODEC_ID_XAN_WC4)
	AV_CODEC_ID_XBIN          = int(C.AV_CODEC_ID_XBIN)
	AV_CODEC_ID_XBM           = int(C.AV_CODEC_ID_XBM)
	AV_CODEC_ID_XFACE         = int(C.AV_CODEC_ID_XFACE)
	AV_CODEC_ID_XSUB          = int(C.AV_CODEC_ID_XSUB)
	AV_CODEC_ID_XWD           = int(C.AV_CODEC_ID_XWD)
	AV_CODEC_ID_Y41P          = int(C.AV_CODEC_ID_Y41P)
	AV_CODEC_ID_YOP           = int(C.AV_CODEC_ID_YOP)
	AV_CODEC_ID_YUV4          = int(C.AV_CODEC_ID_YUV4)
	AV_CODEC_ID_ZEROCODEC     = int(C.AV_CODEC_ID_ZEROCODEC)
	AV_CODEC_ID_ZLIB          = int(C.AV_CODEC_ID_ZLIB)
	AV_CODEC_ID_ZMBV          = int(C.AV_CODEC_ID_ZMBV)
)
