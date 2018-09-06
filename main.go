package main

import (
	"go.uber.org/zap"
	"mylog/log"
)

func main() {
	log.InitLogger()
	s := []string{
		"hello info",
		"hello error",
		"hello debug",
		"hello fatal",
	}
	log.Info("info:", zap.String("s", s[0]))
	log.Error("info:", zap.String("s", s[1]))
	log.Debug("info:", zap.String("s", s[2]))
	log.Fatal("info:", zap.String("s", s[3]))
}
