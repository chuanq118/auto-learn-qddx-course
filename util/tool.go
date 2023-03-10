package util

import (
	"bytes"
	"cn.lqservice.qddxCourse/log"
	"encoding/json"
)

func ToJsonString(v any) string {
	jBytes, err := json.Marshal(v)
	if err != nil {
		log.Logger.Errorf("Marshal %v failed!", v)
		return ""
	}
	var fmtJson bytes.Buffer
	err = json.Indent(&fmtJson, jBytes, "", "  ")
	if err != nil {
		log.Logger.Errorf("Indent %v failed!", v)
		return ""
	}
	return fmtJson.String()
}
