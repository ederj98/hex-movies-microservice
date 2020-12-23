package usecase

import (
	"github.com/ederj98/hex-movies-microservice/application/command"
	"github.com/ederj98/hex-movies-microservice/application/factory"
	"github.com/ederj98/hex-movies-microservice/domain/port"
)

type CreateMoviePort interface {
	Handler(movieCommand command.MovieCommand) error
}

type UseCaseMovieCreate struct {
	MovieRepository port.MovieRepository
}

func (createUseCase *UseCaseMovieCreate) Handler(movieCommand command.MovieCommand) error {

	movie, err := factory.Create(movieCommand)
	if err != nil {
		return err
	}
	createMovieErr := createUseCase.MovieRepository.Create(&movie)
	if createMovieErr != nil {
		return createMovieErr
	}
	return nil
}
