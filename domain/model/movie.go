package model

import "github.com/ederj98/hex-movies-microservice/domain/validators"

type Movie struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Stars    string `json:"stars"`
}

func (movie *Movie) Create(id int64, name string, director string, writer string, stars string) (Movie, error) {
	if err := validators.ValidateRequired(name, "Name should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validators.ValidateRequired(director, "Director should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validators.ValidateRequired(writer, "Writer should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validators.ValidateRequired(stars, "Stars should have some value"); err != nil {
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
