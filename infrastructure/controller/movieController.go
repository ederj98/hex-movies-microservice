package controller

import (
	"net/http"
	"strconv"

	"github.com/ederj98/hex-movies-microservice/application/command"
	"github.com/ederj98/hex-movies-microservice/infrastructure/marshall"

	"github.com/ederj98/hex-movies-microservice/application/usecase"
	"github.com/fmcarrero/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

type RedirectMovieHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	CreateUseCase usecase.CreateMoviePort
	GetUseCase    usecase.GetMovieUseCase
	GetAllUseCase usecase.GetAllMovieUseCase
	UpdateUseCase usecase.UpdateMoviePort
	DeleteUseCase usecase.DeleteMovieUseCase
}

func (h *Handler) Create(c *gin.Context) {

	var movieCommand command.MovieCommand
	if err := c.ShouldBindJSON(&movieCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, createMovieErr := h.CreateUseCase.Handler(movieCommand)

	if createMovieErr != nil {
		_ = c.Error(createMovieErr)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, marshall.Marshall(isPublic, result))
}

func (h *Handler) Get(c *gin.Context) {

	id, movieErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if movieErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	movie, errGet := h.GetUseCase.Handler(id)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}

	c.JSON(http.StatusOK, marshall.Marshall(false, movie))

}

func (h *Handler) GetAll(c *gin.Context) {
	movies, err := h.GetAllUseCase.Handler()
	if err != nil {
		restErr := rest_errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, marshall.MarshallArray(false, movies))
}

func (h *Handler) Update(c *gin.Context) {
	id, movieErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if movieErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}

	var movieCommand command.MovieCommand
	if err := c.ShouldBindJSON(&movieCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	movie, updateErr := h.UpdateUseCase.Handler(id, movieCommand)
	if updateErr != nil {
		restErr := rest_errors.NewBadRequestError(updateErr.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, &movie)
}

func (h *Handler) Delete(c *gin.Context) {
	id, movieErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if movieErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	errDelete := h.DeleteUseCase.Handler(id)
	if errDelete != nil {
		restErr := rest_errors.NewBadRequestError(errDelete.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.Status(http.StatusNoContent)
}
