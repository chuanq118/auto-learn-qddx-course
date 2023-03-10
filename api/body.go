package api

const PageSize = 10

type ListCourseReqBody struct {
	IsPass     int    `json:"isPass"`
	Order      string `json:"order"`
	OrderField string `json:"orderField"`
	PageNum    int    `json:"pageNum"`
	PageSize   int    `json:"pageSize"`
}

type AddQueueReqBody struct {
	CellId       string `json:"cellId"`
	CourseOpenId string `json:"courseOpenId"`
}

type LeaveCellLogReqBody struct {
	Id           string `json:"id"`
	StopSeconds  int    `json:"stopSeconds"`
	VideoEndTime int    `json:"videoEndTime"`
}
