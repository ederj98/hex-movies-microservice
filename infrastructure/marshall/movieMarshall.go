package marshall

import (
	"encoding/json"
	"fmt"

	"github.com/ederj98/hex-movies-microservice/domain/model"
)

type PublicMovie struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type PrivateMovie struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Stars    string `json:"stars"`
}

func Marshall(isPublic bool, movie model.Movie) interface{} {
	if isPublic {
		return PublicMovie{
			Id:   movie.Id,
			Name: movie.Name,
		}
	}
	movieJson, errUn := json.Marshal(movie)
	fmt.Println(errUn)
	fmt.Println(movie)
	var privateMovie PrivateMovie
	_ = json.Unmarshal(movieJson, &privateMovie)
	return privateMovie
}

func MarshallArray(isPublic bool, movies []model.Movie) []interface{} {
	result := make([]interface{}, len(movies))
	for index, movie := range movies {
		result[index] = Marshall(isPublic, movie)
	}
	return result
}
