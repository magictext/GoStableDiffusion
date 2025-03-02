package sd

import "github.com/Binozo/GoStableDiffusion/internal/accelerators"

func HasCuda() bool {
	return accelerators.HasCuda
}
