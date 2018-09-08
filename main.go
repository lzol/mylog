package main

import (
	"fmt"
	"github.com/arthurkiller/rollingWriter"
	"go.uber.org/zap"
	"mylog/log"
)

func main() {

	log.InitLogger("./conf/log_conf.json")
	s := []string{
		"hello info",
		"hello error",
		"hello debug",
		"hello fatal",
	}
	fmt.Println(rollingwriter.VolumeRolling)
	for {
		log.Info("info:", zap.String("s", s[0]))
		log.Error("info:", zap.String("s", s[1]))
		log.Debug("info:", zap.String("s", s[2]))
		//log.Panic("info:", zap.String("s", s[3]))
	}

}
