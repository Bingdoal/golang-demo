package api

import (
	"encoding/json"
	"fmt"
	"go-demo/api/common"
	"go-demo/config"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	Server       *gin.Engine
	middlerSlice []gin.HandlerFunc
}

func (r *Rest) Add(root string, routes ...common.IApiRoute) *Rest {
	group := r.Server.Group(root)
	group.Use(r.middlerSlice...)
	for _, route := range routes {
		route.AddRoute(group)
	}
	r.middlerSlice = nil
	return r
}

func (r *Rest) Middleware(middlerSlice ...gin.HandlerFunc) *Rest {
	r.middlerSlice = append(r.middlerSlice, middlerSlice...)
	return r
}

func (r *Rest) TestApi(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	r.Server.ServeHTTP(res, req)
	return res
}

func (r *Rest) Run() {
	fmt.Printf("\n============ Start [%s] version:%s on:%s ============\n",
		config.Env.GetString("name"),
		config.Env.GetString("version"),
		config.Env.GetString("server.port"))
	NoRouteRedirect(r.Server)
	r.Server.Run(":" + config.Env.GetString("server.port"))
}

func NoRouteRedirect(server *gin.Engine) {
	server.NoRoute(func(ctx *gin.Context) {
		baseUrl := ""

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		req, err := http.NewRequest(ctx.Request.Method,
			baseUrl+ctx.Request.RequestURI, ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		for key, value := range ctx.Request.Header {
			for _, v := range value {
				req.Header.Add(key, v)
			}
		}

		resp, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		var resBody map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&resBody)
		ctx.JSON(resp.StatusCode, resBody)
		return
	})
}
