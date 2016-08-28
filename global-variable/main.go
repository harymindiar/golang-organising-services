package main

import (
	"github.com/go-martini/martini"
	"github.com/harymindiar/golang-organising-services/global-variable/controller"
	"github.com/harymindiar/golang-organising-services/global-variable/model"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	// we don't use martini "Services" way,
	// We try to use global variable for this example
	model.InitDb()
	m.Get("/", controller.GetNews)
	m.RunOnAddr(":8080")
}
