package service

import (
	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/domain/port/repository"
)

type MovieService interface {
	Create(*model.Movie) (*model.Movie, error)
	Find(int) (*model.Movie, error)
	FindAll() ([]model.Movie, error)
	Update(*model.Movie) (*model.Movie, error)
	Delete(int)
}

type Service struct {
	repository repository.MovieRepository
}

func NewService(repository repository.MovieRepository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Create(movie *model.Movie) (*model.Movie, error) {
	return service.repository.Create(movie)
}

func (service *Service) Find(id int) (*model.Movie, error) {
	return service.repository.Find(id)
}

func (service Service) FindAll() ([]model.Movie, error) {
	return service.repository.FindAll()
}

func (service Service) Update(movie *model.Movie) error {
	return service.repository.Update(movie)
}

func (service Service) Delete(id int) error {
	return service.repository.Delete(id)
}
