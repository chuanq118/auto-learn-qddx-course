package api

import (
	"bytes"
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

const (
	listCourseURL      = "http://api.jxjyzx.qdu.edu.cn/LearningSpace/list"
	courseDirectoryURL = "http://api.jxjyzx.qdu.edu.cn/studyLearn/courseDirectoryProcess?courseOpenId="
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

func GetAccessToken() string {
	return "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxMTA2NTIwMjAyMzA0NjMxIiwiaWQiOiJCNzY1MkZDRS03RkM1LTQyQTgtQTI4MS1DQjJCOTk3QUEyN0QiLCJleHAiOjE2NzgzNTk1OTMsImNyZWF0ZWQiOjE2NzgyNzMxOTMwNTR9.YaUmNdOCcrSsW7O-YmeKq5MCXKxVX3jhoKtWp0NHRur2wKqHq2SRP9xAF2rD_rq-oXOHBiigQO3FvUmmlRVnzQ"
}

func setHeaders(req *http.Request) {
	req.Header.Set(uaKey, uaValEdge)
	req.Header.Set(contentTypeKey, applicationJsonUtf8)
	req.Header.Set(referKey, refererVal)
	req.Header.Set(originKey, originVal)
	req.Header.Set(accessTokenKey, GetAccessToken())
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
		return result, nil
	}
	return nil, fmt.Errorf("invalid response code %d\n", response.StatusCode)
}
