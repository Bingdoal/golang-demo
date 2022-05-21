package tests

import (
	"bytes"
	"encoding/json"
	"go-demo/api"
	"go-demo/config"
	"go-demo/config/db"
	"go-demo/internal/dto"
	"go-demo/internal/model/dao"
	"go-demo/internal/service"
	"go-demo/internal/util/logger"
	"net/http"
)

var rest *api.Rest

func init() {
	initialization()
	rest = api.SetUpRoute()
}

func initialization() {
	config.InitConfig("./_assets")
	logger.InitLogger()
	db.InitDB()
	dao.InitDao()
	service.InitService()
	api.InitApiInstance()
}

func loginGetToken() string {
	bodyStr, _ := json.Marshal(dto.LoginDto{
		Username: "admin",
		Password: "admin",
	})
	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewReader(bodyStr))
	res := rest.TestApi(req)
	var resBody map[string]string
	json.Unmarshal(res.Body.Bytes(), &resBody)
	return resBody["token"]
}
