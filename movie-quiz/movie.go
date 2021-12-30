package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Movie struct {
	Id      int    `json:"_id"`
	Title   string `json:"title"`
	Hint    string `json:"hint"`
	ImgPath string `json:"image_source"`
}

var moviesAlreadyUsed = make(map[Movie]bool)

/*
   Movies have ids from 1 to N equally to slice indexes.
*/
func UnmarshalMovies(moviesData []byte) []Movie {
	var movies []Movie
	err := json.Unmarshal(moviesData, &movies)

	if err != nil {
		fmt.Printf("error unmurshaling movies: %s \n", err.Error())
	}
	return movies
}

// In Golang, reflect.DeepEqual function is used to compare the equality of struct, slice, and map
func GenerateMovieOptions(movies []Movie) []Movie {
	optionsCount := 4
	movieOptions := make([]Movie, optionsCount)
	movieUsedInQuestion := make(map[Movie]bool)

	notUsedYetMovie := getNotUsedYetMovie(movies)

	movieUsedInQuestion[notUsedYetMovie] = true
	movieOptions[0] = notUsedYetMovie

	count := 1
	for count < optionsCount {
		current := movies[rand.Intn(len(movies))]
		if movieUsedInQuestion[current] != true {
			movieOptions[count] = current
			movieUsedInQuestion[current] = true
			count++
		}
	}
	return movieOptions
}

//TODO calculate from files
func IsFinishedMovieOptions() bool {
	if len(moviesAlreadyUsed) == 40 {
		return true
	}
	return false
}

func getNotUsedYetMovie(movies []Movie) Movie {
	used := false
	var currentMovie Movie
	for used != true {
		currentMovie = movies[rand.Intn(len(movies))]
		if moviesAlreadyUsed[currentMovie] == false {
			moviesAlreadyUsed[currentMovie] = true
			used = true
			return currentMovie
		}
	}
	return Movie{}
}
