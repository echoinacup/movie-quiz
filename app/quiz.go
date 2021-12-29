package main

import (
	"embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"math/rand"
	"time"
)

// UI containers initialization BLUE color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff}
// TODO RED FOR WRONG color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
// TODO GREEN FOR color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x3f}
// TODO Constants for colors

func initMainWindow(currentApp fyne.App) fyne.Window {
	mainWindow := currentApp.NewWindow("Movie quiz for Mini Seal Pup")
	mainWindow.Resize(fyne.NewSize(700, 700))
	return mainWindow
}

var currentMovieHint string
var moviesCount int

//go:embed images/layout/harp.jpeg
//go:embed images/movies/*
//go:embed movie-meta/movies.json
var resourcesFiles embed.FS

func main() {
	//Make rand generation les deterministic
	rand.Seed(time.Now().UnixNano())

	//Main app initialization
	mainApp := app.New()

	moviesMeta, err := resourcesFiles.ReadFile("movie-meta/movies.json")
	if err != nil {
		log.Fatal(err)
	}
	movies := UnmarshalMovies(moviesMeta)
	moviesCount = len(movies)

	// TODO to fix somehow or leave as is
	mainScreeHarpBytes, err := resourcesFiles.ReadFile("images/layout/harp.jpeg")
	var mainRes = &fyne.StaticResource{
		StaticName:    "harp.jpeg",
		StaticContent: mainScreeHarpBytes,
	}

	// UI elements initialization
	mainWindow := initMainWindow(mainApp)
	questionButtons := initQuestionButtons()
	fistRow, secondRow := createQuestionsRows(questionButtons)
	mainScreenImg := canvas.NewImageFromResource(mainRes)
	nextBtn := addNextButton(movies, mainScreenImg, questionButtons)
	startBtn := addStartButton(movies, mainScreenImg, questionButtons, fistRow, secondRow, nextBtn)
	hintButton := createHintButton()

	// init UI containers
	startBtnContainer := createStartButtonContainer(startBtn)
	titleContainer := createTitleHorizontalContainerForHeader()
	questionsContainer := container.NewVBox(fistRow, secondRow)
	topContainer := container.NewVBox(
		titleContainer,
		startBtnContainer,
		questionsContainer)
	nextButtonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), nextBtn)
	buttonsContainer := container.NewVBox(
		nextButtonContainer,
		hintButton)

	// Set UI elements and containers to main layout container
	content := container.New(layout.NewBorderLayout(topContainer, buttonsContainer, nil, nil),
		topContainer, buttonsContainer, mainScreenImg)
	// final setup and launch of the app
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

func createStartButtonContainer(startBtn *widget.Button) *fyne.Container {
	startBtnColoredContainer := container.NewMax(canvas.NewRectangle(color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff}), startBtn)
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), startBtnColoredContainer, layout.NewSpacer())
}

func createHintButton() *widget.Button {
	hintButton := widget.NewButton("Help me Master!", func() {})
	hintButton.OnTapped = func() {
		updateButtonLabel(hintButton, currentMovieHint)
	}
	return hintButton
}

func addNextButton(movies []Movie,
	mainScreenImg *canvas.Image,
	questionButtons []*widget.Button) *widget.Button {
	nextBtn := widget.NewButton("Next movie", func() {})
	nextBtn.Disable()
	nextBtn.OnTapped = func() {
		initRound(movies, mainScreenImg, questionButtons, nextBtn)
	}
	return nextBtn
}

func addStartButton(movies []Movie, mainScreenImg *canvas.Image, questionButtons []*widget.Button, firstRow *fyne.Container, secondRow *fyne.Container, nextBtn *widget.Button) *widget.Button {
	startBtn := widget.NewButton("Start!", func() {})
	startBtn.OnTapped = func() {
		startBtn.Hide()
		initRound(movies, mainScreenImg, questionButtons, nextBtn)
		// rows for questions are not visible from the beginning
		firstRow.Show()
		secondRow.Show()
	}
	return startBtn
}

func initRound(
	movies []Movie,
	mainScreenInitImg *canvas.Image,
	questionButtons []*widget.Button,
	nextButton *widget.Button) {
	//TODO
	nextButton.Disable()

	currentRoundMovies := GenerateMovieOptions(movies)
	currentMovie := currentRoundMovies[0]
	currentMovieImgPath := currentMovie.ImgPath
	currentMovieHint = currentMovie.Hint
	updateImg(mainScreenInitImg, currentMovieImgPath)
	questions := GenerateQuestions(currentRoundMovies)
	ShuffleQuestions(questions)
	assignQuestionsToButtons(questionButtons, questions, nextButton)
}

func assignQuestionsToButtons(buttons []*widget.Button, questions []Question, nextButton *widget.Button) {
	for i, q := range questions {
		updateButtonLabel(buttons[i], q.Title)
		setQuestionButtonTapped(buttons[i], questions, nextButton)
	}
}

func createQuestionsRows(buttons []*widget.Button) (*fyne.Container, *fyne.Container) {
	fistRow := container.New(layout.NewHBoxLayout(), buttons[0], layout.NewSpacer(), buttons[1])
	secondRow := container.New(layout.NewHBoxLayout(), buttons[2], layout.NewSpacer(), buttons[3])
	fistRow.Hide()
	secondRow.Hide()
	return fistRow, secondRow
}

func initQuestionButtons() []*widget.Button {
	return []*widget.Button{
		widget.NewButton("", func() {}),
		widget.NewButton("", func() {}),
		widget.NewButton("", func() {}),
		widget.NewButton("", func() {})}
}

func setQuestionButtonTapped(button *widget.Button, currentRoundQuestion []Question, nextButton *widget.Button) {
	button.OnTapped = func() {
		if isCorrectAnswer(button.Text, currentRoundQuestion) {
			button.Importance = 1
			nextButton.Enable()
			button.Refresh()
		}
	}
}

func isCorrectAnswer(title string, questions []Question) bool {
	for _, q := range questions {
		if title == q.Title {
			return q.IsCorrect
		}
	}
	return false
}

func createTitleHorizontalContainerForHeader() *fyne.Container {
	topLabel := canvas.NewText("Let's start our Game, Pup!", color.White)
	titleContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), topLabel, layout.NewSpacer())
	return titleContainer
}

func updateImg(image *canvas.Image, imageName string) {
	imgBytes, _ := resourcesFiles.ReadFile("images/movies/" + imageName)
	image.Resource = &fyne.StaticResource{
		StaticName:    imageName,
		StaticContent: imgBytes,
	}
	image.Refresh()
}

func updateButtonLabel(button *widget.Button, newLabel string) {
	button.Text = newLabel
	button.Importance = 2
	button.Refresh()
}
