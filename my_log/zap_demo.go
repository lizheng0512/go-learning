package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func main() {
	// zap是uber公司的高性能日志框架
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006-01-02 15:04:05.000")))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	//
	//config := zap.Config{
	//    Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
	//    Encoding:         "console",
	//    EncoderConfig:    encoderConfig,
	//    OutputPaths:      []string{"stdout", "log/zap_demo.log"},
	//    ErrorOutputPaths: []string{"stderr", "log/zap_demo_err.log"},
	//}
	//
	//zapcore.NewMultiWriteSyncer()
	//
	//log, _ := config.Build()

	fileHook := &lumberjack.Logger{
		Filename:  "log/zap_demo.log",
		MaxSize:   1,
		MaxAge:    15,
		LocalTime: true,
	}

	errFileHook := &lumberjack.Logger{
		Filename:  "log/zap_demo_err.log",
		MaxSize:   1,
		MaxAge:    15,
		LocalTime: true,
	}

	//encoder := zapcore.NewJSONEncoder(encoderConfig)
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	//debugLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	errorLevel := zap.NewAtomicLevelAt(zap.ErrorLevel)
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	tee := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(fileHook), lowPriority),
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), lowPriority),
		zapcore.NewCore(encoder, zapcore.AddSync(errFileHook), highPriority),
		zapcore.NewCore(encoder, zapcore.Lock(os.Stderr), highPriority),
	)
	log := zap.New(tee, zap.AddStacktrace(errorLevel))
	sugar := log.Sugar()

	sugar.Debugw("debug.")
	sugar.Info("info.")
	sugar.Warn("warn!")
	sugar.Error("err!!")
	//
	//for {
	//    log.Debug("debug.", zap.String("testField", "testField"))
	//    log.Info("info.")
	//    log.Warn("warn!")
	//    log.Error("err!!")
	//    time.Sleep(time.Millisecond)
	//}

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println(err.(error).Error())
			sugar.Errorf("出错啦!%s", err)
		}
	}()
	arr := []int{0, 1}
	fmt.Println(arr[2])

	sugar.Sync()
	log.Sync()

}
