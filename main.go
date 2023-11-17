package main

import (
	"project/config"
	"project/route"
)

func main() {
	config.Init()
	e := route.New()
	e.Logger.Fatal(e.Start(":8142"))
}
