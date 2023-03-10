package core

import (
	"cn.lqservice.qddxCourse/api"
	"cn.lqservice.qddxCourse/log"
)

var logger = log.Logger

// GetAllCourses 解析出所有的课程信息
func GetAllCourses(pass bool) (courses []*Course, err error) {
	logger.Infoln("try fetch all pass = ", pass, "courses...")
	passNum := 0
	if pass {
		passNum = 1
	}
	for pageNum := 1; ; pageNum++ {
		logger.Infoln("send request to page", pageNum)
		// 获取当前学习课程
		courseList, err := api.ReqCourseList(&api.ListCourseReqBody{
			IsPass:     passNum,
			Order:      "",
			OrderField: "",
			PageNum:    pageNum,
			PageSize:   api.PageSize,
		})
		if err != nil {
			logger.Errorf("Get course list fail in page-%d!\n", pageNum)
			return nil, err
		}
		data := courseList["data"].(map[string]any)
		total := data["total"].(int)
		list := data["list"].([]map[string]any)
		for _, jo := range list {
			courses = append(courses, &Course{
				OpenId:    jo["id"].(string),
				Name:      jo["name"].(string),
				Progress:  jo["progress"].(string),
				Teachers:  jo["teacherName"].(string),
				Term:      jo["term"].(int),
				Thumbnail: jo["thumbnail"].(string),
				IsPass:    jo["isPass"].(int),
			})
		}
		// 循环仅在总课程数小于已获取数目时结束
		if total <= pageNum*api.PageSize {
			break
		}
	}
	return
}

// GetModulesOfCourse 获取课程的所有章节信息
func GetModulesOfCourse(courseOpenId string) ([]*Module, error) {
	logger.Infoln("Try fetch all course directory...")
	dir, err := api.ReqCourseDirectory(courseOpenId)
	if err != nil {
		return nil, err
	}
	logger.Infoln("parsing the course directory json...")
	moduleList := dir["data"].(map[string]any)["moduleList"].([]map[string]any)
	var modules []*Module
	for _, module := range moduleList {
		var topics []Topic
		for _, topic := range module["topics"].([]map[string]any) {
			var cells []Cell
			for _, cell := range topic["cells"].([]map[string]any) {
				cells = append(cells, Cell{
					Type:        cell["type"].(int),
					IsLearn:     cell["isLearn"].(int),
					Id:          cell["id"].(string),
					Process:     cell["process"].(int),
					VideoLength: cell["videoLength"].(int),
					SubName:     cell["subName"].(string),
					Name:        cell["name"].(string),
				})
			}
			topics = append(topics, Topic{
				Name:     topic["name"].(string),
				Percent:  topic["percent"].(int),
				Duration: topic["studyTime"].(int),
				Id:       topic["name"].(string),
				Cells:    cells,
			})
		}
		modules = append(modules, &Module{
			Title:    module["name"].(string),
			Percent:  module["percent"].(int),
			Duration: module["moduleStudyTime"].(int),
			Topics:   topics,
		})
	}
	return modules, nil
}
