package sd

// This is really weird, somehow I get linking errors if I remove this here. This is the same as in internal/sd/c.go

/*
#cgo CFLAGS: -I../../stable-diffusion.cpp/
#cgo LDFLAGS: -Wl,-Bstatic -L../../stable-diffusion.cpp/build/ggml/src -lggml -lggml-cpu -lggml-base
#cgo LDFLAGS: -Wl,-Bdynamic -lstdc++ -lm -fopenmp -L../../stable-diffusion.cpp/build -lstable-diffusion
#include "../../stable-diffusion.cpp/stable-diffusion.h"
*/
import "C"

func init() {

}
