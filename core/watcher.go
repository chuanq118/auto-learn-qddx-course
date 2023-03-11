package core

import (
	"time"
)

var ticker = time.NewTicker(time.Second * 3)

func Register() {

}

func Start() {
	go start()
}

func start() {
	var timing time.Time
	for {
		timing = <-ticker.C
		logger.Infof("[%s] 检查并调用任务\n", timing.Format("2006-01-02 15:04:05"))

	}
}
