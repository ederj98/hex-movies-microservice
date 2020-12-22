package usecase

import (
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/service"
)

type GetMovieUseCase interface {
	Handler(id int64) (model.Movie, error)
}

type UseCaseGetMovie struct {
	MovieService service.MovieService
}

func (useCaseGetMovie *UseCaseGetMovie) Handler(id int64) (model.Movie, error) {

	movie, err := useCaseGetMovie.MovieService.Find(userId)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}
