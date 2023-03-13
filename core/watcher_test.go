package core

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNumConvert(t *testing.T) {
	atoi, err := strconv.Atoi("00")
	if err != nil {
		logger.Errorln(err)
		return
	}
	logger.Infoln(atoi)
}

func TestMap(t *testing.T) {

	myMap := make(map[string]any, 16)
	fmt.Println(len(myMap))

}
