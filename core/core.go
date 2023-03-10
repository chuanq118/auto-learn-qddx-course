package core

// Limiter 此处设置最大并发(即最多同时学习的课程数目) 默认为 16.
var Limiter = make(chan struct{}, 1<<4)

func LearnCell(cell *Cell) {
	// 如果队列满员,在此处应能阻塞整个外部循环~
	Limiter <- struct{}{}
	defer func() {
		// 我们在闭包中确定 limiter 必定要能够释放出来
		<-Limiter
	}()

	// TODO 处理核心逻辑
	if cell.IsLearn == 1 {
		return
	}
	logger.Infof("[]")

}
