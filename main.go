package main

import (
	"fmt"
	"github.com/arthurkiller/rollingWriter"
	"go.uber.org/zap"
	"mylog/log"
)

func main() {

	log.InitLogger("./conf/log_conf.json", "./conf/rolling_conf.json")
	s := []string{
		"hello info",
		"hello error",
		"hello debug",
		"hello fatal",
	}
	fmt.Println(rollingwriter.VolumeRolling)
	for a := 1; a < 10000; a++ {
		log.Info("info:", zap.String("s", s[0]))
		log.Error("info:", zap.String("s", s[1]))
		log.Debug("info:", zap.String("s", s[2]))
		//log.Panic("info:", zap.String("s", s[3]))
	}

	//log.InitRollingLog("./conf/rolling_conf.json")
}
