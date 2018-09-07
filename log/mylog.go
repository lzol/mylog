package log

import (
	"encoding/json"
	"github.com/arthurkiller/rollingWriter"
	"io"
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

func InitLogger(file string, rollingFile string) {
	initLog(file)
	log.SetFlags(log.Lmicroseconds | log.Llongfile | log.LstdFlags)

	//writer := InitRollingLog(rollingFile)
	log.SetOutput(testWriter())
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

func testWriter() io.Writer {
	config := rollingwriter.Config{
		LogPath:       "./",           //日志路径
		TimeTagFormat: "060102150405", //时间格式串
		FileName:      "test",         //日志文件名
		MaxRemain:     5,              //配置日志最大存留数

		// 目前有2中滚动策略: 按照时间滚动按照大小滚动
		// - 时间滚动: 配置策略如同 crontable, 例如,每天0:0切分, 则配置 0 0 0 * * *
		// - 大小滚动: 配置单个日志文件(未压缩)的滚动大小门限, 入1G, 500M
		RollingPolicy:      rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
		RollingTimePattern: "* * * * * *",               //配置时间滚动策略
		RollingVolumeSize:  "5M",                        //配置截断文件下限大小
		Compress:           true,                        //配置是否压缩存储

		// writer 支持3种方式:
		// - 无保护的 writer: 不提供并发安全保障
		// - lock 保护的 writer: 提供由 mutex 保护的并发安全保障
		// - 异步 writer: 异步 write, 并发安全. 异步开启后忽略 Lock 选项
		Asynchronous: false, //配置是否异步写
		Lock:         true,  //配置是否同步加锁写
	}

	// 创建一个 writer
	writer, err := rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		// 应该处理错误
		panic(err)
	}
	return writer
}

func InitRollingLog(file string) io.Writer {
	var rollingConfig rollingwriter.Config
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("read log config file failed!")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("read log config file failed!")
	}
	if err := json.Unmarshal(b, &rollingConfig); err != nil {
		panic(err)
	}
	writer, err := rollingwriter.NewWriterFromConfig(&rollingConfig)
	if err != nil {
		log.Fatal("read rolling config file failed!")
	}

	//logger = logger.WithOptions(zap.WrapCore(func(zapcore.Core) zapcore.Core {
	//	return zapcore.NewCore(
	//		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	//		zapcore.Lock(zapcore.AddSync(writer)),
	//		zapcore.DebugLevel,
	//	)
	//}))
	//logger.Info("sfdsdfsdfsdf")
	return writer
}
