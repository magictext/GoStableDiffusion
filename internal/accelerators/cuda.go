//go:build cuda
// +build cuda

package accelerators

/*
#cgo LDFLAGS: -L../../stable-diffusion.cpp/build/ggml/src/ggml-cuda -Wl,-Bstatic -lggml-cuda
#cgo LDFLAGS: -L/opt/cuda/lib64 -L/usr/local/cuda/targets/x86_64-linux/lib/ -Wl,-Bdynamic -lcuda -lcudart -lcublas
*/
import "C"

var HasCuda = true
