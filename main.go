package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Make me fit")
	es := getExercises()
	e := es.chooseRandomExercise()

	displayedExercise := widget.NewCard("", "", container.NewVBox(e.display()))
	newExerciseButton := widget.NewButton("New random exercise", func() {
		newExercise := es.chooseRandomExercise()
		displayedExercise.SetContent(container.NewVBox(newExercise.display()))
	})
	buttonContainer := container.NewHBox(layout.NewSpacer(), newExerciseButton, layout.NewSpacer())
	w.SetContent(container.NewVBox(displayedExercise, buttonContainer))
	w.ShowAndRun()
}
