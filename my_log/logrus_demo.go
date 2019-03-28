package main

import (
	"github.com/sirupsen/logrus"
)

var datetimeLayout = "2006-01-02 15:04:05"

func main() {
	// logrus整合logrus_mate的file hook
	log := logrus.StandardLogger()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: datetimeLayout, ForceColors: true})
	logLevel, err := logrus.ParseLevel("debug")
	if err == nil {
		log.SetLevel(logLevel)
	}

	// 基于lumberjack
	//fileHook := &lumberjack.Logger{
	//    Filename: "log/logrus_demo.log",
	//    MaxSize:  1,
	//    MaxAge:   15,
	//}
	//log.SetOutput(fileHook)

	// 基于logrus_mate
	//errLogConfigStr := `{filename:"log/logrus_demo_err.log","level": 0}`
	//errHook, err := logrus_file.NewFileHook(config.NewHOCONConfiguration(configuration.ParseString(errLogConfigStr)))
	//if err == nil {
	//   log.AddHook(errHook)
	//}
	//warnLogConfigStr := `{filename:"log/logrus_demo_warn.log","level": 1}`
	//warnHook, err := logrus_file.NewFileHook(config.NewHOCONConfiguration(configuration.ParseString(warnLogConfigStr)))
	//if err == nil {
	//   log.AddHook(warnHook)
	//}
	//infoLogConfigStr := `{filename:"log/logrus_demo_info.log","level": 2}`
	//infoHook, err := logrus_file.NewFileHook(config.NewHOCONConfiguration(configuration.ParseString(infoLogConfigStr)))
	//if err == nil {
	//   log.AddHook(infoHook)
	//}
	//debugLogConfigStr := `{"filename":"log/logrus_demo_debug.log","level": 3, "max-size": 102400}`
	//debugHook, err := logrus_file.NewFileHook(config.NewHOCONConfiguration(configuration.ParseString(debugLogConfigStr)))
	//if err == nil {
	//   log.AddHook(debugHook)
	//}
	log.Error("Err!!")
	log.Warn("Warn!")
	log.Info("Info.")
	log.Debug("Debug.")
}
