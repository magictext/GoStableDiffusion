package sd

import "github.com/Binozo/GoStableDiffusion/internal/sd"

type Params struct {
	Prompt, NegativePrompt       string
	ClipSkip                     int
	CfgScale                     float32
	Guidance                     float32
	Eta                          float32
	Width, Height                int
	SampleMethod                 sd.SampleMethod
	SampleSteps                  int
	Seed                         int64
	BatchCount                   int
	ControlImage                 *sd.Image
	ControlStrength              float32
	StyleStrength                float32
	NormalizeInput               bool
	InputIdImagesPath            string
	SkipLayers                   []int
	SkipLayersCount              uint64
	SlgScale                     float32
	SkipLayerStart, SkipLayerEnd float32
}

func NewDefaultParams() Params {
	return Params{
		ClipSkip:        -1,
		CfgScale:        7,
		Guidance:        3.5,
		Width:           512,
		Height:          512,
		SampleMethod:    EulerA,
		SampleSteps:     20,
		Seed:            42,
		BatchCount:      1,
		ControlStrength: 0.9,
		StyleStrength:   20,
		SkipLayers:      []int{7, 8, 9},
		SkipLayersCount: 3,
		SkipLayerStart:  0.01,
		SkipLayerEnd:    0.2,
	}
}
