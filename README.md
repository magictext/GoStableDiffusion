<p align="center">
  <img src="./assets/thumbnail.png" width="360x">
</p>

# GoStableDiffusion

CGO bindings for the awesome [stable-diffusion.cpp](https://github.com/leejet/stable-diffusion.cpp) project

> [!WARNING]
> This project is in an early stage

## Setup

### Compilation

You have to compile the shared `stable-diffusion.cpp` library first. You can do this in two ways:

- Compile it yourself using the [official manual](https://github.com/leejet/stable-diffusion.cpp?tab=readme-ov-file#build)
- Let this project compile it yourself

#### Compiling using `go:generate`

Clone this project
```shell
$ git clone https://github.com/Binozo/GoStableDiffusion
$ cd GoStableDiffusion
$ go generate
```

> [!NOTE]
> `go generate` executes the [build.sh](./build.sh) script on unix or [build.bat](./build.bat) on windows.

The `go:generate` script currently supports cuda, vulkan and metal as hardware accelerators.

To enable one of these, simply set one of the following:
- `CUDA=1`
- `VULKAN=1`
- `METAL=1`

before running `go generate`. Example:

```shell
...
$ CUDA=1 go generate
```

> [!TIP]
> If you need extra help you can look at the [Dockerfile](./Dockerfile), [Dockerfile.cuda](./Dockerfile.cuda) or [Dockerfile.vulkan](./Dockerfile.vulkan) files.
> Or just create an issue.

After successful compilation you will find your built shared library at `stable-diffusion.cpp/build/bin/libstable-diffusion<.so/.dylib/.dll>`.

You will have to install it yourself. On linux you could copy it to `/usr/local/lib/`. (Better tutorial coming soon)


### Go

Get the package
```shell
$ go get -u github.com/binozo/gostablediffusion
```

Now you are ready to go! 🚀

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