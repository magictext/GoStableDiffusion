<p align="center">
  <img src="./assets/thumbnail.png" width="360x">
</p>

# GoStableDiffusion

CGO bindings for the awesome [stable-diffusion.cpp](https://github.com/leejet/stable-diffusion.cpp) project

## Setup

Get the package
```shell
$ go get -u github.com/Binozo/GoStableDiffusion
```

Compile `stable-diffusion.cpp` bindings
```shell
# cd into the package location
cd $GOPATH/src/github.com/Binozo/GoStableDiffusion
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
> If you need extra help you can look at the [Dockerfile](./Dockerfile) or [Dockerfile.cuda](./Dockerfile.cuda) files.
> Or just create an issue.