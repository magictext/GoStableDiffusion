package main

import (
	"fmt"
	"github.com/binozo/gostablediffusion/pkg/sd"
	"image/png"
	"os"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("Warming up")

	sd.SetLogCallback(func(level sd.LogLevel, text string, data unsafe.Pointer) {
		switch level {
		case sd.LogDebug:
			fmt.Print("DEBUG:", text)
		case sd.LogInfo:
			fmt.Print("INFO:", text)
		case sd.LogWarn:
			fmt.Print("WARN:", text)
		case sd.LogError:
			fmt.Print("ERROR:", text)
		default:
			fmt.Print("???:", text)
		}
	})

	sd.SetProgressCallback(func(step, steps int, time time.Duration, data unsafe.Pointer) {
		fmt.Printf("PROGRESS: Completed step %d of %d in %0.2fs\n", step, steps, time.Seconds())
	})

	ctx, err := sd.New().
		SetDiffusionModel("models/flux1-dev-q8_0.gguf").
		SetClipL("models/clip_l.safetensors").
		SetVaePath("models/ae.safetensors").
		SetT5xxlPath("models/t5xxl_fp16.safetensors").
		UseFlashAttn().
		Load()

	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	params := sd.NewDefaultParams()
	params.SampleMethod = sd.Euler
	params.Height = 512
	params.Width = 512
	params.CfgScale = 1
	params.Prompt = "a lovely cat holding a sign says 'flux.cpp"

	fmt.Println("Running inference")
	result := ctx.Text2Img(params)

	fmt.Println("Writing result to output.png")
	targetFile, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer targetFile.Close()
	if err := png.Encode(targetFile, result.Image()); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
