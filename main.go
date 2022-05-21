package main

import (
	"go-demo/api"
)

var rest *api.Rest

func main() {
	Initialization()
	rest = api.SetUpRoute()
	rest.Run()
}
