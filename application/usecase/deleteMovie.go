package usecase

import "github.com/ederj98/hex-movies-microservice/domain/port"

type DeleteMovieUseCase interface {
	Handler(id int64) error
}

type UseCaseDeleteMovie struct {
	MovieRepository port.MovieRepository
}

func (useCaseDeleteMovie *UseCaseDeleteMovie) Handler(id int64) error {
	_, getMovieError := useCaseDeleteMovie.MovieRepository.Find(id)
	if getMovieError != nil {
		return getMovieError
	}
	err := useCaseDeleteMovie.MovieRepository.Delete(id)
	return err
}
