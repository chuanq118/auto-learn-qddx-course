package core

import (
	"cn.lqservice.qddxCourse/api"
	"strconv"
	"strings"
	"time"
)

const leaveDuration = int64(60)

//// Limiter 此处设置最大并发(即最多同时学习的课程数目) 默认为 64.
//var Limiter = make(chan struct{}, 1 << 6)
//
//func LearnCell(cell *Cell, courseOpenId string) {
//	// 如果队列满员,在此处应能阻塞整个外部循环~
//	Limiter <- struct{}{}
//	go func(limiter <-chan struct{}, cell *Cell) {
//		doLearnCell(cell, courseOpenId)
//		defer func() {
//			<-limiter
//		}()
//	}(Limiter, cell)
//}

func LearnCell(cell *Cell, courseOpenId string) {
	doLearnCell(cell, courseOpenId)
}

func doLearnCell(cell *Cell, courseOpenId string) {

	if cell.IsLearn == 1 {
		logger.Infof("[跳过] 已学习单元 %s(%s)", cell.Name, cell.SubName)
		return
	}

	logger.Infof("[进入] 模拟学习 %s -> %s", cell.SubName, cell.Name)
	// 判断是否的视频或者文档类型
	if cell.Type == 1 || cell.Type == 2 {
		detail, err := api.ReqCellDetail(cell.Id)
		if err != nil {
			logger.Errorf("Get cell [%s](%s) detail failed!\n %v\n", cell.Name, cell.Id, err)
			return
		}
		data := detail["data"].(map[string]any)
		logId := data["cellLogId"].(string)
		courseOpenId := data["courseOpenId"].(string)
		_, err = api.ReqAddQueueList(&api.AddQueueReqBody{
			CellId:       cell.Id,
			CourseOpenId: courseOpenId,
		})
		if err != nil {
			logger.Errorf("add queue list failed! For course-[%s] cell-[%s].\n %v\n", courseOpenId, cell.Id, err)
			return
		}
		// 注册延时任务
		if cell.Type == 1 {
			// 获取视频时长(秒)
			vLen := parseVideoLength(data)
			if vLen == 0 {
				vLen = 10 * leaveDuration
			}
			registerVideoJobs(vLen, logId)
		} else {
			registerDocxJob(logId)
		}
	} else {
		logger.Errorf("Unsupported cell type -> [%d], sub-name -> [%s],"+
			" name -> [%s]\n", cell.Type, cell.SubName, cell.Name)
		return
	}

}

func registerVideoJobs(videoDuration int64, logId string) {
	ut := time.Now().Unix()
	var endDur int64
	shouldBreak := false
	for i := int64(1); ; i++ {
		if shouldBreak {
			break
		}
		endDur = i * leaveDuration
		if endDur > videoDuration {
			timing := ut + endDur
			logger.Infoln("Register [video] leave log task at", timing, "for", logId)
			RegisterJob(timing, &api.LeaveCellLogReqBody{
				Id:           logId,
				StopSeconds:  0,
				VideoEndTime: endDur,
			})
			shouldBreak = true
		}
	}
}

func registerDocxJob(logId string) {
	timing := time.Now().Unix() + 60
	logger.Infoln("Register [docx] leave log task at", timing, "for", logId)
	RegisterJob(timing, &api.LeaveCellLogReqBody{
		Id:           logId,
		StopSeconds:  0,
		VideoEndTime: leaveDuration,
	})
}

// 00:08:43.8400000
func parseVideoLength(data map[string]any) int64 {
	defer func() {
		if p := recover(); p != nil {
			logger.Errorf("Parse video length from json failed! -> %v", p)
		}
	}()
	s := data["filePreviewInfo"].(map[string]any)["fileStatus"].(map[string]any)["args"].(map[string]any)["duration"].(string)
	pieces := strings.Split(strings.Split(s, ".")[0], ":")
	if len(pieces) == 3 {
		s1, err := strconv.Atoi(pieces[0])
		s2, err := strconv.Atoi(pieces[1])
		s3, err := strconv.Atoi(pieces[2])
		if err == nil {
			return int64(s1*3600 + s2*60 + s3)
		}
		logger.Errorln("Parsing time string to number failed!", err)
	}
	return 0
}
