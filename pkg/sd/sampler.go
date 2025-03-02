package sd

import "GoStableDiffusion/internal/sd"

const (
	EulerA       sd.SampleMethod = sd.EulerA
	Euler        sd.SampleMethod = sd.Euler
	Heun         sd.SampleMethod = sd.Heun
	Dpm2         sd.SampleMethod = sd.Dpm2
	DpmPP2S      sd.SampleMethod = sd.DpmPP2SA
	DpmPP2M      sd.SampleMethod = sd.DpmPP2M
	DpmPP2Mv2    sd.SampleMethod = sd.DpmPP2Mv2
	IpnDm        sd.SampleMethod = sd.IpnDm
	IpnDmV       sd.SampleMethod = sd.IpnDmV
	Lcm          sd.SampleMethod = sd.Lcm
	DdimTrailing sd.SampleMethod = sd.DdimTrailing
	Tcd          sd.SampleMethod = sd.Tcd
)
