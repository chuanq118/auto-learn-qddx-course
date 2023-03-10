package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func printFormatJson(v any) {
	jBytes, err := json.Marshal(v)
	if err != nil {
		return
	}
	var fmtJson bytes.Buffer
	err = json.Indent(&fmtJson, jBytes, "", "  ")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", fmtJson.String())
}

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
	printFormatJson(courses)
}

func TestReqCourseDirectory(t *testing.T) {
	dir, err := ReqCourseDirectory("d4eab501-1fe7-4308-81c4-055e965fede6")
	if err != nil {
		log.Fatal(err)
	}
	printFormatJson(dir)
}
