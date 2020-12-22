package usecase

import (
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/service"
)

type GetAllMovieUseCase interface {
	Handler() (model.Movie, error)
}

type UseCaseGetAllMovie struct {
	MovieService service.MovieService
}

func (useCaseGetAllMovie *UseCaseGetAllMovie) Handler() ([]model.Movie, error) {

	[]movie, err := useCaseGetMovie.MovieService.FindAll(idd)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}
