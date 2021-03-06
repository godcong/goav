package main

import (
	"log"
)

func main() {

	// Register all formats and codecs
	goav.AvRegisterAll()
	goav.AVCodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", goav.AVFilterVersion())
	log.Printf("AvDevice Version:\t%v", goav.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", goav.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", goav.AvutilVersion())
	log.Printf("AVCodec Version:\t%v", goav.AVCodecVersion())
	log.Printf("Resample Version:\t%v", goav.SwresampleLicense())

}
