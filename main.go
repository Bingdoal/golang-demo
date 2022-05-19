package main

import (
	"go-demo/api"
	"go-demo/internal/util"
)

var rest *api.Rest

func init() {
	util.Init()
	rest = api.SetUpRoute()
}

func main() {
	rest.Run()
}
