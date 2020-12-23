package usecase

import (
	"github.com/ederj98/hex-movies-microservice/application/command"
	"github.com/ederj98/hex-movies-microservice/application/factory"
	"github.com/ederj98/hex-movies-microservice/domain/port"
)

type UpdateMoviePort interface {
	Handler(id int64, movieCommand command.MovieCommand) error
}

type UseCaseMovieUpdate struct {
	MovieRepository port.MovieRepository
}

func (updateUseCase *UseCaseMovieUpdate) Handler(id int64, movieCommand command.MovieCommand) error {

	movie, err := factory.Create(movieCommand)
	if err != nil {
		return err
	}
	updateMovieErr := updateUseCase.MovieRepository.Update(id, &movie)
	if updateMovieErr != nil {
		return updateMovieErr
	}
	return nil
}
