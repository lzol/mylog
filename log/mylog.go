package log

import (
	"encoding/json"
	"github.com/arthurkiller/rollingWriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var Log *zap.Logger

//函数指针，指向Logger.Debug
//var Info func(msg string, fields ...zapcore.Field)
//var Debug func(msg string, fields ...zapcore.Field)
//var Warn func(msg string, fields ...zapcore.Field)
//var Error func(msg string, fields ...zapcore.Field)

//Fatal和Panic打印日志后会退出程序，慎用
//var Fatal func(msg string, fields ...zapcore.Field)
//var Panic func(msg string, fields ...zapcore.Field)

func InitLogger(file string) {
	cfg := initLog(file)

	Log, _ = cfg.Build()

	var encoder zapcore.Encoder

	if strings.EqualFold(cfg.Encoding, "json") {
		encoder = zapcore.NewJSONEncoder(cfg.EncoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	}
	rollingWriter := initRollingLog(file)
	rollingCore := zapcore.NewCore(
		encoder,
		zapcore.Lock(zapcore.AddSync(rollingWriter)),
		cfg.Level,
	)
	core := zapcore.NewTee(Log.Core(), rollingCore)
	Log = zap.New(core, zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel))
	defer Log.Sync()

}

func initLog(file string) zap.Config {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("read Log config file failed!")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("read Log config file failed!")
	}
	var cfg zap.Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//Log, err = cfg.Build()
	if err != nil {
		log.Fatal("init Log error: ", err)
	}
	return cfg
}

func initRollingLog(file string) io.Writer {
	var rollingConfig rollingwriter.Config
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("read Log config file failed!")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("read Log config file failed!")
	}
	if err := json.Unmarshal(b, &rollingConfig); err != nil {
		panic(err)
	}
	writer, err := rollingwriter.NewWriterFromConfig(&rollingConfig)
	if err != nil {
		log.Fatal("read rolling config file failed!")
	}

	return writer
}
