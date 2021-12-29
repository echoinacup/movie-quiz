package main

import (
	"io/ioutil"
	"os"
	"quiz-gui/test/utls"
	"reflect"
	"testing"
)
import "github.com/google/go-cmp/cmp"

func TestUnmarshalMovies(t *testing.T) {
	got := UnmarshalMovies(utls.FetchFileContent("./movie-meta/movies.json"))
	dirInfo, _ := ioutil.ReadDir("./images/movies/")
	wantSize := len(dirInfo)

	if got == nil || len(got) != wantSize {
		t.Errorf("got %q, wanted %q", got, wantSize)
	}

	for _, movie := range got {
		path := "./images/movies/" + movie.ImgPath
		if _, err := os.Stat(path); err != nil {
			t.Errorf("Path to  movie: %s  imageis not correct or cannot be read", movie.Title)
		}
		checkImg := utls.FetchFileContent(path)

		if checkImg == nil || len(checkImg) == 0 {
			t.Errorf("Image for movie: %s is not correct or cannot be read", movie.Title)
		}
	}
}

func TestMovieOptionGeneration(t *testing.T) {
	movies := UnmarshalMovies(utls.FetchFileContent("../test/resources/movies.json"))
	gotOptions := GenerateMovieOptions(movies)
	wantSize := 4
	if gotOptions == nil || len(gotOptions) != wantSize {
		t.Errorf("got wrong size of options slice %d, wanted %q", len(gotOptions), wantSize)
	}
}

func TestMovie(t *testing.T) {

	movie1 := Movie{
		Id:      1,
		Title:   "The Revenant",
		Hint:    "What if you use a horse like Airbnb?",
		ImgPath: "/images/movies/revenant.jpeg"}
	movie2 := Movie{
		Id:      1,
		Title:   "The Revenant",
		Hint:    "What if you use a horse like Airbnb?",
		ImgPath: "/images/movies/revenant.jpeg"}
	gotCmp := cmp.Equal(movie1, movie2)
	gotDeepEqls := reflect.DeepEqual(movie1, movie2)

	if gotCmp != true || gotDeepEqls != true {
		t.Errorf("Equality of movies is not working for cmp %t, and for deepEqual %t", gotCmp, gotDeepEqls)
	}
	//println(cmp.Equal(movie1, movie2, cmpopts.IgnoreFields(person{}, "ID")))
}
