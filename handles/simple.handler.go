package handles

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SimpleHandle struct {
	path         string
	level        string
	callerEnable bool

	writer *zap.Logger
}

func MakeSimpleHandle(path, level string, enable bool) *SimpleHandle {
	handler := &SimpleHandle{
		path: path, level: level, callerEnable: enable,
	}
	handler.init()
	return handler
}

func (handle *SimpleHandle) init() {
	file, _ := os.OpenFile(handle.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	core := zapcore.NewCore(
		jsonEncoder(),
		zapcore.WriteSyncer(file),
		level(handle.level),
	)
	handle.writer = zap.New(
		core,
		zap.WithCaller(handle.callerEnable),
		zap.AddCallerSkip(callerSkipOffset),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}

func (handle SimpleHandle) Debug(msg string, fields ...zap.Field) {
	handle.writer.Debug(msg, fields...)
}

func (handle SimpleHandle) Debugf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Debugf(msg, fields...)
}

func (handle SimpleHandle) Info(msg string, fields ...zap.Field) {
	handle.writer.Info(msg, fields...)
}

func (handle SimpleHandle) Infof(msg string, fields ...interface{}) {
	handle.writer.Sugar().Infof(msg, fields...)
}

func (handle SimpleHandle) Warn(msg string, fields ...zap.Field) {
	handle.writer.Warn(msg, fields...)
}

func (handle SimpleHandle) Warnf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Warnf(msg, fields...)
}

func (handle SimpleHandle) Error(msg string, fields ...zap.Field) {
	handle.writer.Error(msg, fields...)
}

func (handle SimpleHandle) Errorf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Errorf(msg, fields...)
}

func (handle SimpleHandle) DPanic(msg string, fields ...zap.Field) {
	handle.writer.DPanic(msg, fields...)
}

func (handle SimpleHandle) DPanicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().DPanicf(msg, fields...)
}

func (handle SimpleHandle) Panic(msg string, fields ...zap.Field) {
	handle.writer.Panic(msg, fields...)
}

func (handle SimpleHandle) Panicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Panicf(msg, fields...)
}

func (handle SimpleHandle) Fatal(msg string, fields ...zap.Field) {
	handle.writer.Fatal(msg, fields...)
}

func (handle SimpleHandle) Fatalf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Fatalf(msg, fields...)
}
