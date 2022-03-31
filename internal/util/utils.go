package util

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ObjConvert(from interface{}, to interface{}) {
	str, _ := json.Marshal(from)
	json.Unmarshal(str, &to)
}

var Validator = validator.New()

func BindJsonAndValid(ctx *gin.Context, obj interface{}) error {
	if err := ctx.BindJSON(&obj); err != nil {
		return err
	}
	return Validator.Struct(obj)
}
