package sd

import "GoStableDiffusion/internal/sd"

type Context struct {
	internal *sd.Ctx
}

func NewContext(params ContextParams) (*Context, error) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	cCtx := sd.NewSdContext(
		params.ModelPath,
		params.ClipLPath,
		params.ClipGPath,
		params.T5xxlPath,
		params.DiffusionModelPath,
		params.VaePath,
		params.TaeSdPath,
		params.ControlNetPath,
		params.LoraModelDir,
		params.EmbedDir,
		params.StackedIdEmbedDir,
		params.VaeDecodeOnly,
		params.VaeTiling,
		params.FreeParamsImmediately,
		params.NThreads,
		params.WType,
		params.RngType,
		params.Schedule,
		params.KeepClipOnCpu,
		params.KeepControlNetCpu,
		params.KeepVaeOnCpu,
		params.DiffusionFlashAttn,
	)

	return &Context{
		internal: cCtx,
	}, nil
}

func (c *Context) Free() {
	c.internal.Free()
}
