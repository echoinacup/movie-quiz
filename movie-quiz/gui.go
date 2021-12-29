package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var hintBtnLbl = "Help me Master!"
var nextBtn *widget.Button
var startBtn *widget.Button
var hintBtn *widget.Button

var currentMovieHint string

func InitApp(movies []Movie, mainImgBytes []byte) fyne.Window {
	mainWindow := InitMainWindow()
	mainWindow.SetContent(initGuiComponents(movies, mainImgBytes))
	return mainWindow
}

func InitMainWindow() fyne.Window {
	mainApp := app.New()
	mainWindow := mainApp.NewWindow("Movie quiz for Mini Seal Pup")
	mainWindow.Resize(fyne.NewSize(700, 700))
	return mainWindow
}

func initGuiComponents(movies []Movie, mainImgBytes []byte) *fyne.Container {
	// UI elements initialization
	questionButtons := initQuestionButtons()
	fistRow, secondRow := createQuestionsRowsOfButtons(questionButtons)
	var mainRes = &fyne.StaticResource{
		StaticName:    "harp.jpeg",
		StaticContent: mainImgBytes,
	}
	mainScreenImg := canvas.NewImageFromResource(mainRes)
	initNextButton(movies, mainScreenImg, questionButtons)
	initStartButton(movies, mainScreenImg, questionButtons, fistRow, secondRow)
	initHintButton()

	// init UI containers
	startBtnContainer := createStartButtonContainer()
	titleContainer := createTitleHorizontalContainerForHeader()
	questionsContainer := container.NewVBox(fistRow, secondRow)
	topContainer := container.NewVBox(
		titleContainer,
		startBtnContainer,
		questionsContainer)
	nextButtonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), nextBtn)
	buttonsContainer := container.NewVBox(
		nextButtonContainer,
		hintBtn)

	// Set UI elements and containers to main layout container
	return container.New(layout.NewBorderLayout(topContainer, buttonsContainer, nil, nil),
		topContainer, buttonsContainer, mainScreenImg)
}

func createQuestionsRowsOfButtons(buttons []*widget.Button) (*fyne.Container, *fyne.Container) {
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

func initNextButton(movies []Movie, mainScreenImg *canvas.Image, questionButtons []*widget.Button) {
	nextBtn = widget.NewButton("Next movie", func() {})
	nextBtn.Disable()
	nextBtn.OnTapped = func() {
		InitRound(movies, mainScreenImg, questionButtons)
	}
}

func initStartButton(movies []Movie, mainScreenImg *canvas.Image, questionButtons []*widget.Button, firstRow *fyne.Container, secondRow *fyne.Container) {
	startBtn = widget.NewButton("Start!", func() {})
	startBtn.OnTapped = func() {
		startBtn.Hide()
		InitRound(movies, mainScreenImg, questionButtons)
		// rows for questions are not visible from the beginning
		firstRow.Show()
		secondRow.Show()
		hintBtn.Show()
	}
}

func initHintButton() {
	hintBtn = widget.NewButton(hintBtnLbl, func() {})
	hintBtn.OnTapped = func() {
		updateButtonLabel(hintBtn, currentMovieHint)
	}
	hintBtn.Hide()
}

func createTitleHorizontalContainerForHeader() *fyne.Container {
	topLabel := canvas.NewText("Let's start our Game, Pup!", color.White)
	titleContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), topLabel, layout.NewSpacer())
	return titleContainer
}

func assignQuestionsToButtons(buttons []*widget.Button, questions []Question) {
	for i, q := range questions {
		updateButtonLabel(buttons[i], q.Title)
		setQuestionButtonTapped(buttons[i], questions)
	}
}

func setQuestionButtonTapped(button *widget.Button, currentRoundQuestion []Question) {
	button.OnTapped = func() {
		if IsCorrectAnswer(button.Text, currentRoundQuestion) {
			button.Importance = 1
			nextBtn.Enable()
			button.Refresh()
		}
	}
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

func InitRound(
	movies []Movie,
	mainScreenInitImg *canvas.Image,
	questionButtons []*widget.Button) {
	updateButtonLabel(hintBtn, hintBtnLbl)
	nextBtn.Disable()

	currentRoundMovies := GenerateMovieOptions(movies)
	currentMovie := currentRoundMovies[0]
	currentMovieImgPath := currentMovie.ImgPath
	currentMovieHint = currentMovie.Hint
	updateImg(mainScreenInitImg, currentMovieImgPath)
	questions := GenerateQuestions(currentRoundMovies)
	ShuffleQuestions(questions)
	assignQuestionsToButtons(questionButtons, questions)
}

func createStartButtonContainer() *fyne.Container {
	startBtnColoredContainer := container.NewMax(canvas.NewRectangle(color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff}), startBtn)
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), startBtnColoredContainer, layout.NewSpacer())
}
