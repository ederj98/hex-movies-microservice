package usecase

import (
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/service"
)

type GetAllMovieUseCase interface {
	Handler() ([]model.Movie, error)
}

type UseCaseGetAllMovie struct {
	movieService service.MovieService
}

func (useCaseGetAllMovie *UseCaseGetAllMovie) Handler() ([]model.Movie, error) {

	movies, err := useCaseGetAllMovie.movieService.FindAll()
	if err != nil {
		return []model.Movie{}, err
	}
	return movies, nil
}
