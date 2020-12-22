package usecase

import "github.com/ederj98/hex-movies-microservice/domain/service"

type DeleteMovieUseCase interface {
	Handler(id int64) error
}

type UseCaseDeleteMovie struct {
	movieService service.MovieService
}

func (useCaseDeleteMovie *UseCaseDeleteMovie) Handler(id int64) error {
	_, getMovieError := useCaseDeleteMovie.movieService.Find(id)
	if getMovieError != nil {
		return getMovieError
	}
	err := useCaseDeleteMovie.movieService.Delete(id)
	return err
}
