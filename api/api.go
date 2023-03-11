package api

import (
	"bytes"
	"cn.lqservice.qddxCourse/log"
	"cn.lqservice.qddxCourse/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	uaKey          = "User-Agent"
	contentTypeKey = "Content-Type"
	accessTokenKey = "access-token"
	referKey       = "Referer"
	originKey      = "Origin"
)

const (
	uaValEdge           = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	applicationJsonUtf8 = "application/json;charset=UTF-8"
	refererVal          = "http://student.jxjyzx.qdu.edu.cn/"
	originVal           = "http://student.jxjyzx.qdu.edu.cn"
)

var HttpClient = &http.Client{Timeout: (1 << 4) * time.Second}

var accessToken = ""

const (
	listCourseURL      = "http://api.jxjyzx.qdu.edu.cn/LearningSpace/list"
	courseDirectoryURL = "http://api.jxjyzx.qdu.edu.cn/studyLearn/courseDirectoryProcess?courseOpenId="
	addQueueListURL    = "http://api.jxjyzx.qdu.edu.cn/process/addedQuestionList"
	leaveCellLogURL    = "http://api.jxjyzx.qdu.edu.cn/studyLearn/leaveCellLog"
	cellDetailURL      = "http://api.jxjyzx.qdu.edu.cn/studyLearn/cellDetail?cellId="
)

// ReqCourseList 请求课程列表
func ReqCourseList(listCourseRB *ListCourseReqBody) (map[string]interface{}, error) {
	body, err := json.Marshal(listCourseRB)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", "parse struct ListCourseReqBody error.")
		return nil, err
	}
	fmt.Printf("Try request for cource list!\n")
	request, err := http.NewRequest("POST", listCourseURL, bytes.NewBuffer(body))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", "build post request error.")
		return nil, err
	}
	setHeaders(request)
	return sendReqAndParse(request)
}

// ReqCourseDirectory 获取课程的详细目录
func ReqCourseDirectory(courseOpenId string) (map[string]interface{}, error) {
	request, err := http.NewRequest("GET", courseDirectoryURL+courseOpenId, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(request)
	return sendReqAndParse(request)
}

func ReqAddQueueList(body *AddQueueReqBody) (map[string]interface{}, error) {
	marshal, err := json.Marshal(body)
	if err != nil {
		log.Logger.Errorln("parse AddQueueReqBody error.")
		return nil, err
	}
	request, err := http.NewRequest("POST", addQueueListURL, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	setHeaders(request)
	return sendReqAndParse(request)
}

func ReqCellDetail(cellId string) (map[string]interface{}, error) {
	request, err := http.NewRequest("GET", cellDetailURL+cellId, nil)
	if err != nil {
		return nil, err
	}
	return sendReqAndParse(request)
}

func ReqLeaveCellLog(body *LeaveCellLogReqBody) (map[string]interface{}, error) {
	marshal, err := json.Marshal(body)
	if err != nil {
		log.Logger.Errorln("parse ReqLeaveCellLog error.")
		return nil, err
	}
	request, err := http.NewRequest("POST", leaveCellLogURL, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	setHeaders(request)
	return sendReqAndParse(request)
}

func GetAccessToken() *string {
	return &accessToken
}

func SetAccessToken(token *string) {
	accessToken = *token
}

func setHeaders(req *http.Request) {
	req.Header.Set(uaKey, uaValEdge)
	req.Header.Set(contentTypeKey, applicationJsonUtf8)
	req.Header.Set(referKey, refererVal)
	req.Header.Set(originKey, originVal)
	req.Header.Set(accessTokenKey, *GetAccessToken())
}

func sendReqAndParse(request *http.Request) (map[string]interface{}, error) {
	response, err := HttpClient.Do(request)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", "send requests error!")
		return nil, err
	}
	if response.StatusCode == 200 {
		result := make(map[string]interface{})
		all, err := io.ReadAll(response.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "read response body error!")
			return nil, err
		}
		if err := json.Unmarshal(all, &result); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "reqCourses 解析响应体为JSON失败]\n")
			return nil, err
		}
		code, ok := result["code"]
		if ok {
			if code.(string) == "200" {
				return result, nil
			} else if code.(string) == "401" {
				log.Logger.Warnln("TOKEN 验证失败,可能需要重新登录.")
				return nil, fmt.Errorf("%d", 401)
			}
		}
		log.Logger.Errorf("错误的响应信息 -> \n%s\n", util.ToJsonString(&result))
		return nil, fmt.Errorf("%s", "invalid response json format.")
	}
	return nil, fmt.Errorf("invalid response code %d\n", response.StatusCode)
}
