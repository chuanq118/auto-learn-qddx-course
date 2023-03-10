package core

type Course struct {
	OpenId    string `json:"openId"`
	Name      string `json:"name"`
	Progress  string `json:"progress"`
	Teachers  string `json:"teachers"`
	Term      int    `json:"term"`
	Thumbnail string `json:"thumbnail"`
	IsPass    int    `json:"isPass"`
}

type Module struct {
	Title    string  `json:"title"`
	Percent  int     `json:"percent"`
	Duration int     `json:"duration"`
	Topics   []Topic `json:"topics"`
}

type Topic struct {
	Name     string `json:"name"`
	Percent  int    `json:"percent"`
	Duration int    `json:"duration"`
	Id       string `json:"id"`
	Cells    []Cell `json:"cells"`
}

type Cell struct {
	Type        int    `json:"type"`
	IsLearn     int    `json:"isLearn"`
	Id          string `json:"id"`
	Process     int    `json:"process"`
	VideoLength int    `json:"videoLength"`
	SubName     string `json:"subName"`
	Name        string `json:"name"`
}
