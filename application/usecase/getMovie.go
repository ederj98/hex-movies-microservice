package usecase

import (
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/port"
)

type GetMovieUseCase interface {
	Handler(id int64) (model.Movie, error)
}

type UseCaseGetMovie struct {
	MovieRepository port.MovieRepository
}

func (useCaseGetMovie *UseCaseGetMovie) Handler(id int64) (model.Movie, error) {

	movie, err := useCaseGetMovie.MovieRepository.Find(id)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}
