package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
import "C"

type Ctx struct {
	internal *C.sd_ctx_t
}

func NewSdContext(
	modelPath, clipLPath, clipGPath, t5xxlPath, diffusionModelPath, vaePath, taeSdPath, controlNetPath, loraModelDir, embedDir, stackedIdEmbedDir string,
	vaeDecodingOnly, vaeTiling, freeParamsImmediately bool,
	nThreads int,
	wType Type,
	rngType Rng,
	schedule Scheduler,
	keepClipOnCpu, keepControlNetCpu, keepVaeOnCpu, diffusionFlashAttn bool,
) *Ctx {
	cInternal := C.new_sd_ctx(
		C.CString(modelPath),
		C.CString(clipLPath),
		C.CString(clipGPath),
		C.CString(t5xxlPath),
		C.CString(diffusionModelPath),
		C.CString(vaePath),
		C.CString(taeSdPath),
		C.CString(controlNetPath),
		C.CString(loraModelDir),
		C.CString(embedDir),
		C.CString(stackedIdEmbedDir),
		C.bool(vaeDecodingOnly),
		C.bool(vaeTiling),
		C.bool(freeParamsImmediately),
		C.int(nThreads),
		wType.internal(),
		rngType.internal(),
		schedule.internal(),
		C.bool(keepClipOnCpu),
		C.bool(keepControlNetCpu),
		C.bool(keepVaeOnCpu),
		C.bool(diffusionFlashAttn),
	)

	return &Ctx{
		internal: cInternal,
	}
}

func (c *Ctx) Free() {
	C.free_sd_ctx(c.internal)
}
