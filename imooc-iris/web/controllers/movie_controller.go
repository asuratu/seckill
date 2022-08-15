package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"imooc/repositories"
	"imooc/services"
)

type MovieController struct {
}

func (m *MovieController) Get() mvc.View {
	movieRepo := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepo)
	name := movieService.ShowName()
	return mvc.View{
		Name: "movie/index",
		Data: iris.Map{
			"name": name,
		},
	}
}
