package model

import (
	"github.com/ederj98/hex-movies-microservice/domain/validator"
)

type Movie struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Stars    string `json:"stars"`
}

func (movie *Movie) Create(name string, director string, writer string, stars string) (Movie, error) {
	if err := validator.ValidateRequired(name, "Name should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validator.ValidateRequired(director, "Director should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validator.ValidateRequired(writer, "Writer should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validator.ValidateRequired(stars, "Stars should have some value"); err != nil {
		return Movie{}, err
	}

	return Movie{
		Name:     name,
		Director: director,
		Writer:   writer,
		Stars:    stars,
	}, nil
}

func (movie *Movie) CreateWithId(id int64, name string, director string, writer string, stars string) (Movie, error) {
	if err := validator.ValidateRequired(name, "Name should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validator.ValidateRequired(director, "Director should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validator.ValidateRequired(writer, "Writer should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validator.ValidateRequired(stars, "Stars should have some value"); err != nil {
		return Movie{}, err
	}

	return Movie{
		Id:       id,
		Name:     name,
		Director: director,
		Writer:   writer,
		Stars:    stars,
	}, nil
}
