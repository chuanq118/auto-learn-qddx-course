package main

import (
	"cn.lqservice.qddxCourse/api"
	"cn.lqservice.qddxCourse/core"
	"cn.lqservice.qddxCourse/log"
	"go.uber.org/zap"
)

var token = ""

func main() {
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(log.ZapLogger)
	logger := log.Logger

	api.SetAccessToken(&token)

	courses, err := core.GetAllCourses(false)
	if err != nil {
		logger.Errorf("Get all courses failed! Error -> %v\n", err)
		return
	}

	// 遍历所有课程
	for _, course := range courses {
		if course.Progress == "100" {
			logger.Infoln("[跳过] 已经完成课程 -> ", course.Name)
			continue
		}
		logger.Infof("[进入] 模拟学习课程 -> [%s].\n", course.Name)
		modules, err := core.GetModulesOfCourse(course.OpenId)
		if err != nil {
			logger.Errorf("Get course [%s](ID[%s]) modules failed!\n %v\n", course.Name, course.OpenId, err)
			return
		}
		// 遍历所有章节
		for _, module := range modules {
			if module.Percent >= 100 {
				logger.Infof("[跳过] 已完成章节 -> [%s]\n", module.Title)
				continue
			}
			logger.Infof("[进入] 章节 -> [%s](已学习-%d%%), 尝试模拟学习...\n", module.Title, module.Percent)
			// 遍历所有小节
			for _, topic := range module.Topics {
				if topic.Percent >= 100 {
					logger.Infof("[跳过] 已完成小节 -> [%s]\n", topic.Name)
					continue
				}
				logger.Infof("[进入] 模拟学习小节 -> [%s]\n", topic.Name)
				for _, cell := range topic.Cells {
					core.LearnCell(&cell, course.OpenId)
				}
			}
		}
	}

}
