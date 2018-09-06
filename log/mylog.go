package log

import (
	"encoding/json"
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

//函数指针，指向Logger.Debug
var Info func(msg string, fields ...zapcore.Field)
var Debug func(msg string, fields ...zapcore.Field)
var Warn func(msg string, fields ...zapcore.Field)
var Error func(msg string, fields ...zapcore.Field)
var Fatal func(msg string, fields ...zapcore.Field)

func InitLogger() {
	// 日志地址 "out.log" 自定义
	lp := "/Volumes/Share/GOPATH/src/mylog/test.log"
	// 日志级别 DEBUG,ERROR, INFO
	lv := "DEBUG"
	// 是否 DEBUG

	initLogger(lp, lv, true)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	Info = logger.Info
	Debug = logger.Debug
	Warn = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal
}

func initLogger(lp string, lv string, isDebug bool) {
	var js string
	if isDebug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, lv)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, lv, lp, lp)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}
