package core

import "cn.lqservice.qddxCourse/api"

// Limiter 此处设置最大并发(即最多同时学习的课程数目) 默认为 16.
var Limiter = make(chan struct{}, 1<<4)

func LearnCell(cell *Cell, courseOpenId string) {
	// 如果队列满员,在此处应能阻塞整个外部循环~
	Limiter <- struct{}{}
	go func(limiter <-chan struct{}, cell *Cell) {
		doLearnCell(cell, courseOpenId)
		defer func() {
			<-limiter
		}()
	}(Limiter, cell)
}

func doLearnCell(cell *Cell, courseOpenId string) {
	// TODO 处理核心逻辑
	if cell.IsLearn == 1 {
		logger.Infof("[跳过] 已学习单元 %s(%s) \n", cell.Name, cell.SubName)
		return
	}

	logger.Infof("[进入] 模拟学习 %s -> %s \n", cell.SubName, cell.Name)

	if cell.Type == 1 {
		doLearnVideo(cell, courseOpenId)
	} else if cell.Type == 2 {
		doLearnDocx(cell, courseOpenId)
	} else {
		logger.Errorf("Unsupported cell type -> [%d], sub-name -> [%s],"+
			" name -> [%s]\n", cell.Type, cell.SubName, cell.Name)
		return
	}

}

func doLearnDocx(cell *Cell, courseOpenId string) {
	_, err := api.ReqAddQueueList(&api.AddQueueReqBody{
		CellId:       cell.Id,
		CourseOpenId: courseOpenId,
	})
	if err != nil {
		logger.Errorf("add queue list failed! For course-[%s] cell-[%s].\n %v\n", courseOpenId, cell.Id, err)
		return
	}

}

func doLearnVideo(cell *Cell, courseOpenId string) {

}
