package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const file = "models/res/err_code.json"

type ErrMap map[string]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
}
