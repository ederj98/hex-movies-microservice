package usecase

import (
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/port"
)

type GetAllMovieUseCase interface {
	Handler() ([]model.Movie, error)
}

type UseCaseGetAllMovie struct {
	MovieRepository port.MovieRepository
}

func (useCaseGetAllMovie *UseCaseGetAllMovie) Handler() ([]model.Movie, error) {

	movies, err := useCaseGetAllMovie.MovieRepository.FindAll()
	if err != nil {
		return []model.Movie{}, err
	}
	return movies, nil
}
