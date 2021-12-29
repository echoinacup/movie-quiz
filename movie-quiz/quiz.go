package main

import (
	"embed"
	"log"
	"math/rand"
	"time"
)

// UI containers initialization BLUE color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff}
// TODO RED FOR WRONG color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
// TODO GREEN FOR color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x3f}
// TODO Constants for colors

//go:embed images/layout/harp.jpeg
//go:embed images/movies/*
//go:embed movie-meta/movies.json
var resourcesFiles embed.FS

func main() {
	//Make rand generation les deterministic
	rand.Seed(time.Now().UnixNano())

	moviesMeta, err := resourcesFiles.ReadFile("movie-meta/movies.json")
	if err != nil {
		log.Fatal(err)
	}
	movies := UnmarshalMovies(moviesMeta)

	mainScreeHarpBytes, err := resourcesFiles.ReadFile("images/layout/harp.jpeg")

	//Main app initialization
	// TODO change to resources
	mainWindow := InitApp(movies, mainScreeHarpBytes)
	mainWindow.ShowAndRun()

}
