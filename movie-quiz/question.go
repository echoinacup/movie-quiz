package main

import (
	"math/rand"
)

type Question struct {
	Title     string
	IsCorrect bool
}

// Assumption from movies slice that first movie is a correct answer.
func GenerateQuestions(movies []Movie) []Question {
	questions := make([]Question, len(movies))
	for currentMovieIndex, currentMovie := range movies {
		if currentMovieIndex == 0 {
			correctAnswer := Question{Title: currentMovie.Title, IsCorrect: true}
			questions[currentMovieIndex] = correctAnswer
		} else {
			questions[currentMovieIndex] = Question{Title: currentMovie.Title}
		}
	}
	return questions
}

func ShuffleQuestions(questions []Question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}

func IsCorrectAnswer(title string, questions []Question) bool {
	for _, q := range questions {
		if title == q.Title {
			return q.IsCorrect
		}
	}
	return false
}
