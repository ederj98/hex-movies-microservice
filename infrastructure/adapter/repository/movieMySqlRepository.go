package repository

import (
	"errors"
	"fmt"

	"github.com/ederj98/hex-movies-microservice/domain/model"
	"github.com/ederj98/hex-movies-microservice/infrastructure/adapter/repository/entity"
	"github.com/ederj98/hex-movies-microservice/infrastructure/mapper"
	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/jinzhu/gorm"
)

type MovieMySqlRepository struct {
	Db *gorm.DB
}

func (movieMySqlRepository *MovieMySqlRepository) Create(movie *model.Movie) error {
	var movieEntity entity.MovieEntity
	movieEntity = mapper.MovieToMovieEntity(*movie)
	if err := movieMySqlRepository.Db.Create(&movieEntity).Error; err != nil {
		logger.Error(fmt.Sprintf("can't work with %s", movieEntity.Name), err)
		return errors.New(fmt.Sprintf("can't work with %s", movieEntity.Name))
	}
	movie.Id = movieEntity.Id
	return nil
}

func (movieMySqlRepository *MovieMySqlRepository) Find(id int64) (model.Movie, error) {
	var movieEntity entity.MovieEntity
	if movieMySqlRepository.Db.First(&movieEntity, id).Error != nil {
		return model.Movie{}, errors.New(fmt.Sprintf("movie with id: %d not found", id))
	}
	movie := mapper.MovieEntityToMovie(movieEntity)
	return movie, nil
}

func (movieMySqlRepository *MovieMySqlRepository) FindAll() ([]model.Movie, error) {
	var moviesEntities []entity.MovieEntity
	if movieMySqlRepository.Db.Find(&moviesEntities).Error != nil {
		return nil, errors.New(fmt.Sprintf("no movies found"))
	}
	if len(moviesEntities) <= 0 {
		return nil, errors.New(fmt.Sprintf("no users found"))
	}
	movies := mapper.MoviesEntitiesToMovies(moviesEntities)
	return movies, nil
}

func (movieMySqlRepository *MovieMySqlRepository) Update(movie *model.Movie) error {
	var current entity.MovieEntity
	if movieMySqlRepository.Db.First(&current, movie.Id).RecordNotFound() {
		return errors.New(fmt.Sprintf("error when updated movie %v", movie.Id))
	}
	if movieMySqlRepository.Db.Model(&current).Update(entity.MovieEntity{Name: movie.Name, Director: movie.Director, Writer: movie.Writer, Stars: movie.Stars}).Error != nil {
		return errors.New(fmt.Sprintf("error when updated movie %v", movie.Id))
	}
	//movieUpdated := mapper.MovieEntityToMovie(current)
	return nil
}

func (movieMySqlRepository *MovieMySqlRepository) Delete(id int64) error {
	var current entity.MovieEntity
	current.Id = id
	if movieMySqlRepository.Db.Delete(current).Error != nil {
		return errors.New(fmt.Sprintf("cannot delete movie %v", id))
	}
	return nil
}
