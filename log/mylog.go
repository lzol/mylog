package log

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

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
var Panic func(msg string, fields ...zapcore.Field)

func InitLogger(file string) {

	initLog(file)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	Info = logger.Info
	Debug = logger.Debug
	Warn = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal
	Panic = logger.Panic
}

func initLog(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("read log config file failed!")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("read log config file failed!")
	}
	var cfg zap.Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}
