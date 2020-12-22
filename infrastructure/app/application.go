package app

import (
	"github.com/ederj98/hex-movies-microservice/application/usecase"
	"github.com/ederj98/hex-movies-microservice/domain/port/repository"
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

	_ = godotenv.Load()
	router.Use(middlewares.ErrorHandler())
	movieRepository := getMovieRepository()
	var handler = createHandler(movieRepository)
	mapUrls(handler)

	logger.Info("about to start the application")
	_ = router.Run(":8081")
}

func createHandler(movieRepository repository.MovieRepository) controller.RedirectMovieHandler {

	return newHandler(newCreatesUseCase(userRepository), newGetUserUseCase(userRepository),
		newUpdateUserUseCase(userRepository), newDeleteUserUseCase(userRepository),
		newFindUsersByStatusUseCase(userRepository))
}

func newCreateUseCase(repository repository.MovieRepository) usecase.CreateMoviePort {
	return &usecase.UseCaseMovieCreate{
		MovieRepository: repository,
	}
}

func newGetMovieUseCase(repository repository.MovieRepository) usecase.GetMovieUseCase {
	return &usecase.UseCaseGetMovie{
		MovieRepository: repository,
	}
}

func newGetAllMoviesUseCase(repository repository.MovieRepository) usecase.GetAllMovieUseCase {
	return &usecase.UseCaseGetAllMovie{
		MovieRepository: repository,
	}
}

func newUpdateMovieUseCase(repository repository.MovieRepository) usecase.UpdateMoviePort {
	return &usecase.UseCaseMovieUpdate{
		MovieRepository: repository,
	}
}

func newDeleteUserUseCase(repository repository.MovieRepository) usecase.DeleteMovieUseCase {
	return &usecase.UseCaseDeleteMovie{
		MovieRepository: repository,
	}
}

func newHandler(createMovie usecase.CreateMoviePort, getMovieUseCase usecase.GetMovieUseCase, getAllMovieUseCase usecase.GetAllMovieUseCase, updateMovieUseCase usecase.UpdateMoviePort,
	deleteMovieUseCase usecase.DeleteMovieUseCase) controller.RedirectMovieHandler {
	return &controller.Handler{CreateUseCase: createMovie, GetMovieUseCase: getMovieUseCase, UpdateMovieUseCase: updateMovieUseCase,
		DeleteMovieUseCase: deleteMovieUseCase,
	}
}

func getMovieRepository() repository.MovieRepository {
	return &movies.MovieMySqlRepository{
		Db: database.GetDatabaseInstance(),
	}
}
