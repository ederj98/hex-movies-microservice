package repository

import "github.com/ederj98/hex-movies-microservice/domain/model"

type MovieRepository interface {
	Create(*model.Movie) (*model.Movie, error)
	Find(int) (*model.Movie, error)
	FindAll() ([]model.Movie, error)
	Update(*model.Movie) error
	Delete(int) error
}
