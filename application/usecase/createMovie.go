package usecase

import (
	"github.com/ederj98/hex-movies-microservice/application/command"
	"github.com/ederj98/hex-movies-microservice/application/factory"
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/port"
)

type CreateMoviePort interface {
	Handler(movieCommand command.MovieCommand) (model.Movie, error)
}

type UseCaseMovieCreate struct {
	MovieRepository port.MovieRepository
}

func (createUseCase *UseCaseMovieCreate) Handler(movieCommand command.MovieCommand) (model.Movie, error) {

	movie, err := factory.Create(movieCommand)
	if err != nil {
		return model.Movie{}, err
	}
	createMovieErr := createUseCase.MovieRepository.Create(&movie)
	if createMovieErr != nil {
		return model.Movie{}, createMovieErr
	}
	return movie, nil
}
