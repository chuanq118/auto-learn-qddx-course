package core

import (
	"cn.lqservice.qddxCourse/api"
	"cn.lqservice.qddxCourse/util"
	"sync"
	"time"
)

var ticker = time.NewTicker(time.Second * 3)

// 简单的任务存储器, key 为秒级时间戳, 一个 key 内包含一个任务列表
var taskMap = struct {
	sync.RWMutex
	m map[int64][]*api.LeaveCellLogReqBody
}{m: make(map[int64][]*api.LeaveCellLogReqBody, 0)}

func RegisterJob(timestamp int64, leaveCellLog *api.LeaveCellLogReqBody) {
	taskMap.Lock()
	defer taskMap.Unlock()
	taskMap.m[timestamp] = append(taskMap.m[timestamp], leaveCellLog)
}

func StartWatching() {
	go start()
}

func StopWatching() {
	for len(taskMap.m) > 0 {
		time.Sleep(time.Second * 1)
	}
}

func start() {
	var timing time.Time
	for {
		timing = <-ticker.C
		logger.Infof("[%s] 检查并调用任务", timing.Format("2006-01-02 15:04:05"))
		timeline := timing.Unix()
		for ts, _ := range taskMap.m {
			// 判断处于当前时间段的任务
			if ts >= timeline && ts < (timeline+3) {
				for _, leaveCellBody := range taskMap.m[ts] {
					response, err := api.ReqLeaveCellLog(leaveCellBody)
					if err != nil {
						logger.Errorf("Send leave cell log failed for cell log id [%s]\n%v\n", taskMap.m[ts], err)
						return
					}
					logger.Infof("Send leave successfully!\n%s\n", util.ToJsonString(response))
				}
				delete(taskMap.m, ts)
			}
		}
	}
}
