package usecase

import (
	"github.com/ederj98/hex-movies-microservice/application/command"
	"github.com/ederj98/hex-movies-microservice/application/factory"
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/port"
)

type UpdateMoviePort interface {
	Handler(movieCommand command.MovieCommand) (model.Movie, error)
}

type UseCaseMovieUpdate struct {
	MovieRepository port.MovieRepository
}

func (updateUseCase *UseCaseMovieUpdate) Handler(movieCommand command.MovieCommand) (model.Movie, error) {

	movie, err := factory.Update(movieCommand)
	if err != nil {
		return model.Movie{}, err
	}
	updateMovieErr := updateUseCase.MovieRepository.Update(&movie)
	if updateMovieErr != nil {
		return model.Movie{}, updateMovieErr
	}
	return movie, nil
}
