package api

type ListCourseReqBody struct {
	IsPass     int    `json:"isPass"`
	Order      string `json:"order"`
	OrderField string `json:"orderField"`
	PageNum    int    `json:"pageNum"`
	PageSize   int    `json:"pageSize"`
}
