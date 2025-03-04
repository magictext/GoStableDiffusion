<p align="center">
  <img src="./assets/thumbnail.png" width="360x">
</p>

# GoStableDiffusion

CGO bindings for the awesome [stable-diffusion.cpp](https://github.com/leejet/stable-diffusion.cpp) project

> [!WARNING]
> This project is in an early stage and currently only works on Linux

## Setup

Get the package
```shell
$ go get -u github.com/binozo/gostablediffusion
```

Compile `stable-diffusion.cpp` bindings
```shell
# cd into the package location
cd $GOPATH/pkg/mod/github.com/binozo/gostablediffusion
go generate
```

> [!NOTE]
> `go generate` executes the [build.sh](./build.sh) script on unix or [build.bat](./build.bat) on windows.
> This script precompiles `stable-diffusion.cpp` for you.

### Hardware acceleration

This project currently supports cuda, vulkan and metal as hardware accelerators.

To enable one of these, simply set one of the following:
- `CUDA=1`
- `VULKAN=1`
- `METAL=1`

before running `go generate`. Example:

```shell
...
$ CUDA=1 go generate
```

Now you are ready to go! 🚀

> [!TIP]
> If you need extra help you can look at the [Dockerfile](./Dockerfile), [Dockerfile.cuda](./Dockerfile.cuda) or [Dockerfile.vulkan](./Dockerfile.vulkan) files.
> Or just create an issue.

## Example

[main.go](./cmd/main/main.go)
```go
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
	fmt.Println("Running with Cuda:", sd.HasCuda())
	fmt.Println("Running with Vulkan:", sd.HasVulkan())

	ctx, err := sd.NewDefault("models/sd-v1-4.ckpt")
	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	sd.SetLogCallback(func(level sd.LogLevel, text string, data unsafe.Pointer) {
		switch level {
		case sd.LogDebug:
			fmt.Print("DEBUG:", text)
		case sd.LogInfo:
			fmt.Print("INFO:", text)
		case sd.LogError:
			fmt.Print("ERROR:", text)
		default:
			fmt.Print("???:", text)
		}
	})

	sd.SetProgressCallback(func(step, steps int, time time.Duration, data unsafe.Pointer) {
		fmt.Printf("PROGRESS: Completed step %d of %d in %0.2fs\n", step, steps, time.Seconds())
	})

	params := sd.NewDefaultParams()
	params.Prompt = "A beautiful and sunny landscape with lots of dogs"

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


```