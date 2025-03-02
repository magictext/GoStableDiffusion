package sd

import (
	"GoStableDiffusion/internal/sd"
	"unsafe"
)

type LogLevel sd.LogLevel

const (
	LogDebug LogLevel = LogLevel(sd.LogDebug)
	LogInfo  LogLevel = LogLevel(sd.LogInfo)
	LogWarn  LogLevel = LogLevel(sd.LogWarn)
	LogError LogLevel = LogLevel(sd.LogError)
)

type LogCallback func(level LogLevel, text string, data unsafe.Pointer)

func SetLogCallback(callback LogCallback) {
	sd.SetLogCallback(func(level sd.LogLevel, text string, data unsafe.Pointer) {
		callback(LogLevel(level), text, data)
	})
}

func SetProgressCallback(callback sd.ProgressCallback) {
	sd.SetProgressCallback(callback)
}
