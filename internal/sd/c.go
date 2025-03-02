package sd

/*
#cgo CFLAGS: -I../../stable-diffusion.cpp/
#cgo LDFLAGS: -Wl,-Bstatic -L../../stable-diffusion.cpp/build/ggml/src -lggml -lggml-cpu -lggml-base
#cgo LDFLAGS: -Wl,-Bdynamic -lstdc++ -lm -fopenmp -L../../stable-diffusion.cpp/build -lstable-diffusion
#include "../../stable-diffusion.cpp/stable-diffusion.h"
*/
import "C"

func GetSystemInfo() string {
	return C.GoString(C.sd_get_system_info())
}
