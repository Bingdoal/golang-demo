package util

import "encoding/json"

func ObjConvert(from interface{}, to interface{}) {
	str, _ := json.Marshal(from)
	json.Unmarshal(str, &to)
}
