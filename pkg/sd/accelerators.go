package sd

import "GoStableDiffusion/internal/accelerators"

func HasCuda() bool {
	return accelerators.HasCuda
}
