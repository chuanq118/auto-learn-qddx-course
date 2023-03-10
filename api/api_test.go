package api

import (
	"cn.lqservice.qddxCourse/util"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReqCourseList(t *testing.T) {
	courses, err := ReqCourseList(&ListCourseReqBody{
		IsPass:     0,
		Order:      "",
		OrderField: "",
		PageNum:    1,
		PageSize:   10,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Err:: %v\n", err)
		return
	}
	fmt.Println(util.ToJsonString(courses))

}

func TestReqCourseDirectory(t *testing.T) {
	dir, err := ReqCourseDirectory("d4eab501-1fe7-4308-81c4-055e965fede6")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(util.ToJsonString(dir))
}
