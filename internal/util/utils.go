package util

import (
	"encoding/json"
	"fmt"
)

func ObjConvert(from interface{}, to interface{}) {
	str, _ := json.Marshal(from)
	json.Unmarshal(str, &to)
}

func IfNilPanic(variable interface{}, msg ...string) {
	if variable == nil {
		fmt.Printf("msg: %v\n", msg)
		panic("IfNilPanic!!")
	}
}
