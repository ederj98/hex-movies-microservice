package usecase

import "github.com/ederj98/hex-movies-microservice/domain/service"

type DeleteMovieUseCase interface {
	Handler(id int64) error
}

type UseCaseDeleteMovie struct {
	MovieService service.MovieService
}

func (useCaseDeleteMovie *UseCaseDeleteMovie) Handler(id int64) error {
	movie, getMovieError := useCaseDeleteMovie.MovieService.Find(id)
	if getMovieError != nil {
		return getMovieError
	}
	err := useCaseDeleteMovie.MovieService.Delete(id)
	return err
}
