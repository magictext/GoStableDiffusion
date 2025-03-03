package sd

import "github.com/binozo/gostablediffusion/internal/accelerators"

func HasCuda() bool {
	return accelerators.HasCuda
}

func HasVulkan() bool {
	return accelerators.HasVulkan
}
