package app

import (
	"github.com/ederj98/hex-movies-microservice/infrastructure/controller"
)

func mapUrls(handler controller.RedirectMovieHandler) {

	router.POST("/movies", handler.Create)
	router.PUT("/movies/:id", handler.Update)
	router.DELETE("/movies/:id", handler.Delete)
	router.GET("/movies/:id", handler.Get)
	router.GET("/movies", handler.GetAll)
}
