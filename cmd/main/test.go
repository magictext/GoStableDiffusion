package main

import (
	"errors"
	"fmt"
	"github.com/binozo/gostablediffusion/pkg/sd"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"unsafe"
)

const modelUrl = "https://huggingface.co/CompVis/stable-diffusion-v-1-4-original/resolve/main/sd-v1-4.ckpt"

func main() {
	fmt.Println("Checking for ai model")
	if _, err := os.Stat("models/sd-v1-4.ckpt"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Downloading sd model from", modelUrl)
		if err = downloadModel(); err != nil {
			log.Fatal(err)
		}
	}
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
		SetModel("models/sd-v1-4.ckpt").
		UseFlashAttn().
		Load()

	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	params := sd.NewDefaultParams()
	params.CfgScale = 5
	params.SampleSteps = 30
	params.SampleMethod = sd.Euler
	params.Height = 768
	params.Width = 768
	params.Seed = 42
	params.Prompt = "fantasy medieval village world inside a glass sphere , high detail, fantasy, realistic, light effect, hyper detail, volumetric lighting, cinematic, macro, depth of field, blur, red light and clouds from the back, highly detailed epic cinematic concept art cg render made in maya, blender and photoshop, octane render, excellent composition, dynamic dramatic cinematic lighting, aesthetic, very inspirational, world inside a glass sphere by james gurney by artgerm with james jean, joe fenton and tristan eaton by ross tran, fine details, 4k resolution"

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

func downloadModel() error {
	out, err := os.Create("models/sd-v1-4.ckpt")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(modelUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
