package main

import (
	"cn.lqservice.qddxCourse/api"
	"cn.lqservice.qddxCourse/core"
	"cn.lqservice.qddxCourse/log"
	"go.uber.org/zap"
	"time"
)

var token = "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxMTA2NTIwMjAyMzA0NjMxIiwiaWQiOiJCNzY1MkZDRS03RkM1LTQyQTgtQTI4MS1DQjJCOTk3QUEyN0QiLCJleHAiOjE2Nzg2OTAyMTgsImNyZWF0ZWQiOjE2Nzg2MDM4MTgwMzF9.069OtaiyxBME79LHhDZ7BTfgSr1VVI-eQeey-wq0dIpxVP1vvanGCeWByGuBRcCy54OQBbiE394BvhaWH_MEZA"

func main() {
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(log.ZapLogger)
	logger := log.Logger

	api.SetAccessToken(&token)

	courses, err := core.GetAllCourses(false)
	if err != nil {
		logger.Errorf("Get all unpassed courses failed! Error -> %v\n", err)
		return
	}
	passedCourse, err := core.GetAllCourses(true)
	if err != nil {
		logger.Errorf("Get all passed courses failed! Error -> %v\n", err)
		return
	}
	totalCourses := append(courses, passedCourse...)
	// 启用监听
	core.StartWatching()
	// 遍历所有课程
	for _, course := range totalCourses {
		if course.Progress == "100" {
			logger.Infoln("[跳过] 已经完成课程 -> ", course.Name)
			continue
		}
		logger.Infof("[进入] 模拟学习课程 -> [%s].", course.Name)
		modules, err := core.GetModulesOfCourse(course.OpenId)
		if err != nil {
			logger.Errorf("Get course [%s](ID[%s]) modules failed!\n %v", course.Name, course.OpenId, err)
			return
		}
		if modules == nil {
			continue
		}
		// 遍历所有章节
		for _, module := range modules {
			if module.Percent >= 100 {
				logger.Infof("[跳过] 已完成章节 -> [%s]", module.Title)
				continue
			}
			logger.Infof("[进入] 章节 -> [%s](已学习-%d%%), 尝试模拟学习...", module.Title, module.Percent)
			// 遍历所有小节
			for _, topic := range module.Topics {
				if topic.Percent >= 100 {
					logger.Infof("[跳过] 已完成小节 -> [%s]", topic.Name)
					continue
				}
				logger.Infof("[进入] 模拟学习小节 -> [%s]", topic.Name)
				for _, cell := range topic.Cells {
					core.LearnCell(&cell, course.OpenId)
					time.Sleep(time.Second * 1)
				}
			}
		}
	}
	// 关闭监听 - 阻塞 main routine 防止任务 routine 未完成就被终止
	core.StopWatching()
}
