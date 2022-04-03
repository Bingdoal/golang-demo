package main

import (
	"go-demo/api"
	_ "go-demo/config/db/postgres"
)

var rest *api.Rest

func init() {
	rest = api.SetUpRoute()
}

func main() {
	rest.Run()
}
