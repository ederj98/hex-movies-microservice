package app

import (
	"github.com/ederj98/hex-movies-microservice/application/usecase"
	"github.com/ederj98/hex-movies-microservice/domain/port"
	"github.com/ederj98/hex-movies-microservice/infrastructure/adapter/repository"
	"github.com/ederj98/hex-movies-microservice/infrastructure/app/middlewares"
	"github.com/ederj98/hex-movies-microservice/infrastructure/controller"
	"github.com/ederj98/hex-movies-microservice/infrastructure/database"
	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func StartApplication() {

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	router.Use(middlewares.ErrorHandler())
	movieRepository := getMovieRepository()
	var handler = createHandler(movieRepository)
	mapUrls(handler)

	logger.Info("Start the application...")
	_ = router.Run(":8081")
}

func createHandler(movieRepository port.MovieRepository) controller.RedirectMovieHandler {

	return newHandler(newCreateUseCase(movieRepository), newGetMovieUseCase(movieRepository),
		newGetAllMoviesUseCase(movieRepository), newUpdateMovieUseCase(movieRepository),
		newDeleteMovieUseCase(movieRepository))
}

func newCreateUseCase(repository port.MovieRepository) usecase.CreateMoviePort {
	return &usecase.UseCaseMovieCreate{
		MovieRepository: repository,
	}
}

func newGetMovieUseCase(repository port.MovieRepository) usecase.GetMovieUseCase {
	return &usecase.UseCaseGetMovie{
		MovieRepository: repository,
	}
}

func newGetAllMoviesUseCase(repository port.MovieRepository) usecase.GetAllMovieUseCase {
	return &usecase.UseCaseGetAllMovie{
		MovieRepository: repository,
	}
}

func newUpdateMovieUseCase(repository port.MovieRepository) usecase.UpdateMoviePort {
	return &usecase.UseCaseMovieUpdate{
		MovieRepository: repository,
	}
}

func newDeleteMovieUseCase(repository port.MovieRepository) usecase.DeleteMovieUseCase {
	return &usecase.UseCaseDeleteMovie{
		MovieRepository: repository,
	}
}

func newHandler(createMovie usecase.CreateMoviePort, getMovieUseCase usecase.GetMovieUseCase, getAllMovieUseCase usecase.GetAllMovieUseCase, updateMovieUseCase usecase.UpdateMoviePort,
	deleteMovieUseCase usecase.DeleteMovieUseCase) controller.RedirectMovieHandler {
	return &controller.Handler{CreateUseCase: createMovie, GetUseCase: getMovieUseCase, UpdateUseCase: updateMovieUseCase,
		GetAllUseCase: getAllMovieUseCase, DeleteUseCase: deleteMovieUseCase,
	}
}

func getMovieRepository() port.MovieRepository {
	return &repository.MovieMySqlRepository{
		Db: database.GetDatabaseInstance(),
	}
}
