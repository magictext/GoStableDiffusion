package sd

type builder struct {
	params ContextParams
}

func New() *builder {
	return &builder{
		params: GetDefaultContextParams(),
	}
}

func (b *builder) SetModel(model string) *builder {
	b.params.ModelPath = model
	return b
}

func (b *builder) SetDiffusionModel(diffusionModelPath string) *builder {
	b.params.DiffusionModelPath = diffusionModelPath
	return b
}

func (b *builder) SetClipL(clipLPath string) *builder {
	b.params.ClipLPath = clipLPath
	return b
}

func (b *builder) SetClipG(clipGPath string) *builder {
	b.params.ClipGPath = clipGPath
	return b
}

func (b *builder) SetVaePath(vaePath string) *builder {
	b.params.VaePath = vaePath
	return b
}

func (b *builder) SetT5xxlPath(t5xxlPath string) *builder {
	b.params.T5xxlPath = t5xxlPath
	return b
}

func (b *builder) SetLoRaDir(loraDir string) *builder {
	b.params.LoraModelDir = loraDir
	return b
}

func (b *builder) UseFlashAttn() *builder {
	b.params.DiffusionFlashAttn = true
	return b
}

func (b *builder) Load() (*Context, error) {
	if err := b.params.validate(); err != nil {
		return nil, err
	}

	return NewContext(b.params)
}
