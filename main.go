package main

import (
	"go-demo/api"
)

var rest *api.Rest

func main() {
	Init()
	rest = api.SetUpRoute()
	rest.Run()
}
