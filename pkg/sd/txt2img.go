package sd

import (
	"github.com/Binozo/GoStableDiffusion/internal/sd"
)

func (c *Context) Text2Img(params Params) sd.Image {
	return c.internal.Txt2Img(
		params.Prompt,
		params.NegativePrompt,
		params.ClipSkip,
		params.CfgScale,
		params.Guidance,
		params.Eta,
		params.Width,
		params.Height,
		params.SampleMethod,
		params.SampleSteps,
		params.Seed,
		params.BatchCount,
		params.ControlImage,
		params.ControlStrength,
		params.StyleStrength,
		params.NormalizeInput,
		params.InputIdImagesPath,
		params.SkipLayers,
		params.SkipLayersCount,
		params.SlgScale,
		params.SkipLayerStart,
		params.SkipLayerEnd,
	)
}
