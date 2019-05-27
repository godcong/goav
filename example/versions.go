package main

import (
	"log"


)

func main() {

	// Register all formats and codecs
	goav.AvRegisterAll()
	goav.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", goav.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", goav.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", goav.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", goav.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", goav.AvcodecVersion())
	log.Printf("Resample Version:\t%v", goav.SwresampleLicense())

}
