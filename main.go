package main

import (
	"fmt"
	"github.com/arthurkiller/rollingWriter"
	"go.uber.org/zap"
	. "mylog/log"
)

func main() {

	InitLogger("./conf/log_conf.json")
	s := []string{
		"hello info",
		"hello error",
		"hello debug",
		"hello fatal",
	}
	fmt.Println(rollingwriter.VolumeRolling)
	for {
		Log.Info("info:", zap.String("s", s[0]))
		Log.Error("info:", zap.String("s", s[1]))
		Log.Debug("info:", zap.String("s", s[2]))
		//log.Panic("info:", zap.String("s", s[3]))
	}

}
