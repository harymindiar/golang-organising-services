package controller

import (
	"github.com/go-martini/martini"
	"github.com/harymindiar/golang-organising-services/global-variable/model"
	"github.com/martini-contrib/render"
)

func GetNews(c martini.Context, render render.Render) {
	res, err := model.GetNews()
	if err != nil {
		render.JSON(503, map[string]interface{}{"error": "Data not found"})
	} else {
		render.JSON(200, map[string]interface{}{"data": res})
	}
}
