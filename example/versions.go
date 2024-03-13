package main

import (
	"log"

	"github.com/wwllww/goav/avcodec"
	"github.com/wwllww/goav/avdevice"
	"github.com/wwllww/goav/avfilter"
	"github.com/wwllww/goav/avformat"
	"github.com/wwllww/goav/avutil"
	"github.com/wwllww/goav/swresample"
	"github.com/wwllww/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
