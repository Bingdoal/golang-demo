package util

import (
	"encoding/json"
)

func ObjConvert(from interface{}, to interface{}) {
	str, _ := json.Marshal(from)
	json.Unmarshal(str, &to)
}

func Filter[E interface{}](slice []E, callback func(E) bool) []E {
	var result []E
	for _, v := range slice {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[E interface{}, R interface{}](slice []E, callback func(E) R) []R {
	var result []R
	for _, v := range slice {
		result = append(result, callback(v))
	}
	return result
}
