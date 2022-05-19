package main

import (
	"go-demo/api"
	"go-demo/internal/util"
)

var rest *api.Rest

func main() {
	util.Init()
	rest = api.SetUpRoute()
	rest.Run()
}
