package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
import "C"
import (
	"image"
	"unsafe"
)

type Image struct {
	width, height int
	channel       int
	data          []uint8
	internal      *C.sd_image_t
}

func (c *Ctx) Txt2Img(
	prompt, negativePrompt string,
	clipSkip int,
	cfgScale, guidance, eta float32,
	width, height int,
	sampler SampleMethod,
	sampleSteps int,
	seed int64,
	batchCount int,
	controlCond *Image,
	controlStrength, styleStrength float32,
	normalizeInput bool,
	inputIdImagesPath string,
	skipLayers []int,
	skipLayersCount uint64,
	slgScale, skipLayerStart, skipLayerEnd float32,
) Image {
	cSkipLayers := (*C.int)(unsafe.Pointer(&skipLayers[0]))
	cSkipLayersLen := C.size_t(skipLayersCount)
	var controlImage *C.sd_image_t
	if controlCond != nil {
		controlImage = controlCond.internal
	}

	generatedImage := C.txt2img(c.internal,
		C.CString(prompt),
		C.CString(negativePrompt),
		C.int(clipSkip),
		C.float(cfgScale),
		C.float(guidance),
		C.float(eta),
		C.int(width),
		C.int(height),
		sampler.internal(),
		C.int(sampleSteps),
		C.int64_t(seed),
		C.int(batchCount),
		controlImage,
		C.float(controlStrength),
		C.float(styleStrength),
		C.bool(normalizeInput),
		C.CString(inputIdImagesPath),
		cSkipLayers,
		cSkipLayersLen,
		C.float(slgScale),
		C.float(skipLayerStart),
		C.float(skipLayerEnd),
	)

	generatedImageSize := uint32(generatedImage.width) * uint32(generatedImage.height) * uint32(generatedImage.channel)

	return Image{
		internal: generatedImage,
		width:    width,
		height:   height,
		channel:  int(generatedImage.channel),
		data:     unsafe.Slice((*uint8)(unsafe.Pointer(generatedImage.data)), generatedImageSize),
	}
}

func (i *Image) Image() *image.RGBA {
	return parseImage(i)
}

func parseImage(img *Image) *image.RGBA {
	rect := image.Rect(0, 0, img.width, img.height)
	return &image.RGBA{
		Pix:    convertRGBtoRGBA(img.data, img.width, img.height),
		Rect:   rect,
		Stride: img.width * 4, // By default 3 channels are getting generated but for RGBA we need 4
	}
}

func convertRGBtoRGBA(data []uint8, width, height int) []uint8 {
	rgbaData := make([]uint8, width*height*4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			offsetRGB := (y*width + x) * 3
			offsetRGBA := (y*width + x) * 4
			rgbaData[offsetRGBA] = data[offsetRGB]     // R
			rgbaData[offsetRGBA+1] = data[offsetRGB+1] // G
			rgbaData[offsetRGBA+2] = data[offsetRGB+2] // B
			rgbaData[offsetRGBA+3] = 255               // Alpha
		}
	}
	return rgbaData
}
