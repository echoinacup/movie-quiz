package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGenerateQuestions(t *testing.T) {
	movie1 := Movie{
		Id:      1,
		Title:   "The Revenant",
		Hint:    "What if you use a horse like Airbnb?",
		ImgPath: "../images/movies/revenant.jpeg"}
	movie2 := Movie{
		Id:      2,
		Title:   "Movie2",
		Hint:    "What if you use a horse like Airbnb?",
		ImgPath: "../images/movies/Movie2.jpeg"}
	movie3 := Movie{
		Id:      3,
		Title:   "Movie3",
		Hint:    "What if you use a horse like Airbnb?",
		ImgPath: "../images/movies/movie3.jpeg"}
	movie4 := Movie{
		Id:      4,
		Title:   "Movie4",
		Hint:    "What if you use a horse like Airbnb?",
		ImgPath: "../images/movies/movie4.jpeg"}

	movies := []Movie{movie1, movie2, movie3, movie4}

	expectedQuestions := make([]Question, 4)
	expectedQuestions[0] = Question{Title: "The Revenant", IsCorrect: true}
	expectedQuestions[1] = Question{Title: "Movie2", IsCorrect: false}
	expectedQuestions[2] = Question{Title: "Movie3", IsCorrect: false}
	expectedQuestions[3] = Question{Title: "Movie4", IsCorrect: false}

	actualQuestions := GenerateQuestions(movies)

	if !cmp.Equal(actualQuestions, expectedQuestions) {
		t.Errorf("Question creation is not working! expected: %v and got: %v", expectedQuestions, actualQuestions)
	}
}
