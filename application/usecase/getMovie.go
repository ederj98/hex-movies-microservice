package usecase

import (
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/service"
)

type GetMovieUseCase interface {
	Handler(id int64) (model.Movie, error)
}

type UseCaseGetMovie struct {
	movieService service.MovieService
}

func (useCaseGetMovie *UseCaseGetMovie) Handler(id int64) (model.Movie, error) {

	movie, err := useCaseGetMovie.movieService.Find(id)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}
