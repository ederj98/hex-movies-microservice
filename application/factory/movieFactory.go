package factory

import (
	"github.com/ederj98/hex-movies-microservice/application/command"
	"github.com/ederj98/hex-movies-microservice/domain/model"
)

func Create(movieCommand command.MovieCommand) (model.Movie, error) {
	var movie model.Movie
	movie, err := movie.Create(movieCommand.Name, movieCommand.Director, movieCommand.Writer, movieCommand.Stars)
	return movie, err
}
