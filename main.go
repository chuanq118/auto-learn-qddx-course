package main

import (
	"cn.lqservice.qddxCourse/log"
	"go.uber.org/zap"
)

func main() {
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(log.ZapLogger)
	logger := log.Logger

	logger.Infoln("aaa", "bbb")
}
