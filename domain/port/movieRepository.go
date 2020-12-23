package port

import "github.com/ederj98/hex-movies-microservice/domain/model"

type MovieRepository interface {
	Create(*model.Movie) error
	Find(int64) (model.Movie, error)
	FindAll() ([]model.Movie, error)
	Update(int64, *model.Movie) error
	Delete(int64) error
}
