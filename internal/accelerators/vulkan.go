//go:build vulkan
// +build vulkan

package accelerators

/*
#cgo LDFLAGS: -L../../stable-diffusion.cpp/build/ggml/src/ggml-vulkan -Wl,-Bstatic -lggml-vulkan
#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu -L/lib/x86_64-linux-gnu -L/usr/lib -Wl,-Bdynamic -lvulkan
*/
import "C"

var HasVulkan = true
