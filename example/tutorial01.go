package main

// tutorial01.c
// Code based on a tutorial at http://dranger.com/ffmpeg/tutorial01.html

// A small sample program that shows how to use libavformat and libavcodec to
// read video from a file.
//
// Use
//
// gcc -o tutorial01 tutorial01.c -lavformat -lavcodec -lswscale -lz
//
// to build (assuming libavformat and libavcodec are correctly installed
// your system).
//
// Run using
//
// tutorial01 myvideofile.mpg
//
// to write the first five frames from "myvideofile.mpg" to disk in PPM
// format.
import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/godcong/goav"
)

// SaveFrame writes a single frame to disk as a PPM file
func SaveFrame(frame *goav.AVFrame, width, height, frameNumber int) {
	// Open file
	fileName := fmt.Sprintf("frame%d.ppm", frameNumber)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error Reading")
	}
	defer file.Close()

	// Write header
	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
	file.Write([]byte(header))

	// Write pixel data
	for y := 0; y < height; y++ {
		data0 := goav.Data(frame)[0]
		buf := make([]byte, width*3)
		startPos := uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(goav.Linesize(frame)[0])
		for i := 0; i < width*3; i++ {
			element := *(*uint8)(unsafe.Pointer(startPos + uintptr(i)))
			buf[i] = element
		}
		file.Write(buf)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie file")
		os.Exit(1)
	}

	// Open video file
	pFormatContext := goav.AVFormatAllocContext()
	if goav.AVFormatOpenInput(&pFormatContext, os.Args[1], nil, nil) != 0 {
		fmt.Printf("Unable to open file %s\n", os.Args[1])
		os.Exit(1)
	}

	// Retrieve stream information
	if pFormatContext.AVFormatFindStreamInfo(nil) < 0 {
		fmt.Println("Couldn't find stream information")
		os.Exit(1)
	}

	// Dump information about file onto standard error
	pFormatContext.AvDumpFormat(0, os.Args[1], 0)

	// Find the first video stream
	for i := 0; i < int(pFormatContext.NbStreams()); i++ {
		switch pFormatContext.Streams()[i].CodecParameters().AVCodecGetType() {
		case goav.AVMediaTypeVideo:

			// Get a pointer to the codec context for the video stream
			pCodecCtxOrig := pFormatContext.Streams()[i].Codec()
			// Find the decoder for the video stream
			pCodec := goav.AVCodecFindDecoder(goav.CodecID(pCodecCtxOrig.GetCodecID()))
			if pCodec == nil {
				fmt.Println("Unsupported codec!")
				os.Exit(1)
			}
			// Copy context
			pCodecCtx := pCodec.AVCodecAllocContext3()
			if pCodecCtx.AVCodecCopyContext((*goav.CodecContext)(unsafe.Pointer(pCodecCtxOrig))) != 0 {
				fmt.Println("Couldn't copy codec context")
				os.Exit(1)
			}

			// Open codec
			if pCodecCtx.AVCodecOpen2(pCodec, nil) < 0 {
				fmt.Println("Could not open codec")
				os.Exit(1)
			}

			// Allocate video frame
			pFrame := goav.AVFrameAlloc()

			// Allocate an AVFrame structure
			pFrameRGB := goav.AVFrameAlloc()
			if pFrameRGB == nil {
				fmt.Println("Unable to allocate RGB AVFrame")
				os.Exit(1)
			}

			// Determine required buffer size and allocate buffer
			numBytes := uintptr(goav.AVPictureGetSize(goav.AvPixFmtRgb24, pCodecCtx.Width(),
				pCodecCtx.Height()))
			buffer := goav.AVMalloc(numBytes)

			// Assign appropriate parts of buffer to image planes in pFrameRGB
			// Note that pFrameRGB is an AVFrame, but AVFrame is a superset
			// of Picture
			avp := (*goav.Picture)(unsafe.Pointer(pFrameRGB))
			avp.AvpictureFill((*uint8)(buffer), goav.AvPixFmtRgb24, pCodecCtx.Width(), pCodecCtx.Height())

			// initialize SWS context for software scaling
			swsCtx := goav.SwsGetContext(
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				(goav.PixelFormat)(pCodecCtx.PixFmt()),
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				goav.AvPixFmtRgb24,
				goav.SwsBilinear,
				nil,
				nil,
				nil,
			)

			// Read frames and save first five frames to disk
			frameNumber := 1
			packet := goav.AvPacketAlloc()
			for pFormatContext.AvReadFrame(packet) >= 0 {
				// Is this a packet from the video stream?
				if packet.StreamIndex() == i {
					// Decode video frame
					response := pCodecCtx.AVCodecSendPacket(packet)
					if response < 0 {
						fmt.Printf("Error while sending a packet to the decoder: %s\n", goav.ErrorFromCode(response))
					}
					for response >= 0 {
						response = pCodecCtx.AVCodecReceiveFrame((*goav.AVFrame)(unsafe.Pointer(pFrame)))
						if response == goav.AvErrorEAGAIN || response == goav.AvErrorEOF {
							break
						} else if response < 0 {
							fmt.Printf("Error while receiving a frame from the decoder: %s\n", goav.ErrorFromCode(response))
							return
						}

						if frameNumber <= 5 {
							// Convert the image from its native format to RGB
							goav.SwsScale2(swsCtx, goav.Data(pFrame),
								goav.Linesize(pFrame), 0, pCodecCtx.Height(),
								goav.Data(pFrameRGB), goav.Linesize(pFrameRGB))

							// Save the frame to disk
							fmt.Printf("Writing frame %d\n", frameNumber)
							SaveFrame(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(), frameNumber)
						} else {
							return
						}
						frameNumber++
					}
				}

				// Free the packet that was allocated by av_read_frame
				packet.AvFreePacket()
			}

			// Free the RGB image
			goav.AVFree(buffer)
			goav.AvFrameFree(pFrameRGB)

			// Free the YUV frame
			goav.AvFrameFree(pFrame)

			// Close the codecs
			pCodecCtx.AVCodecClose()
			(*goav.CodecContext)(unsafe.Pointer(pCodecCtxOrig)).AVCodecClose()

			// Close the video file
			pFormatContext.AVFormatCloseInput()

			// Stop after saving frames of first video straem
			break

		default:
			fmt.Println("Didn't find a video stream")
			os.Exit(1)
		}
	}
}
