package core

import (
	"cn.lqservice.qddxCourse/api"
	"cn.lqservice.qddxCourse/log"
	"cn.lqservice.qddxCourse/util"
)

var logger = log.Logger

// GetAllCourses 解析出所有的课程信息
func GetAllCourses(pass bool) (courses []*Course, err error) {
	logger.Infoln("try fetch all pass =", pass, "courses...")
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
		total := int(data["total"].(float64))
		list := data["list"].([]any)
		for _, v := range list {
			jo := v.(map[string]any)
			// json -> 默认为 string,如果该 json 字段不存在则 为 nil
			teacherName := ""
			if jo["teacherName"] != nil {
				teacherName = jo["teacherName"].(string)
			}
			courses = append(courses, &Course{
				OpenId:    jo["id"].(string),
				Name:      jo["name"].(string),
				Progress:  jo["progress"].(string),
				Teachers:  teacherName,
				Term:      int(jo["term"].(float64)),
				Thumbnail: jo["thumbnail"].(string),
				IsPass:    int(jo["isPass"].(float64)),
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
	logger.Infoln("Try fetch all course", courseOpenId, "directory...")
	dir, err := api.ReqCourseDirectory(courseOpenId)
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			logger.Errorln("Unexpected error occurred when trying get module of course! Error ->", p)
			logger.Errorln(util.ToJsonString(dir))
		}
	}()
	logger.Infoln("parsing the course directory json...")
	moduleList := dir["data"].(map[string]any)["moduleList"].([]any)
	var modules []*Module
	for _, moduleV := range moduleList {
		module := moduleV.(map[string]any)
		var topics []Topic
		for _, topicV := range module["topics"].([]any) {
			topic := topicV.(map[string]any)
			var cells []Cell
			for _, cellV := range topic["cells"].([]any) {
				cell := cellV.(map[string]any)
				if cell["type"] == nil {
					logger.Warnf("Non type cell ->\n%s\n", util.ToJsonString(cell))
					continue
				}
				cells = append(cells, Cell{
					Type:        int(cell["type"].(float64)),
					IsLearn:     int(cell["isLearn"].(float64)),
					Id:          cell["id"].(string),
					Process:     int(cell["process"].(float64)),
					VideoLength: int(cell["videoLength"].(float64)),
					SubName:     cell["subName"].(string),
					Name:        cell["name"].(string),
				})
			}
			topics = append(topics, Topic{
				Name:     topic["name"].(string),
				Percent:  int(topic["percent"].(float64)),
				Duration: int(topic["studyTime"].(float64)),
				Id:       topic["name"].(string),
				Cells:    cells,
			})
		}
		modules = append(modules, &Module{
			Title:    module["name"].(string),
			Percent:  int(module["percent"].(float64)),
			Duration: int(module["moduleStudyTime"].(float64)),
			Topics:   topics,
		})
	}
	return modules, nil
}
