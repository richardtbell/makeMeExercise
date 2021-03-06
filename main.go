package main

import (
	"makeMeExercise/exercise"
	"makeMeExercise/exercises"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Make me fit")
	es := exercises.Get()
	workout := es.GetRandomFullBodyWorkout()

	tabs := container.NewAppTabs(
		container.NewTabItem("Full body workout", displayFullBodyWorkout(workout)),
	)
	win.Resize(fyne.NewSize(1200, 600))
	win.SetContent(tabs)
	win.ShowAndRun()
}

func displayFullBodyWorkout(w exercises.Exercises) *widget.Accordion {
	a := widget.NewAccordion()
	for _, e := range w {
		a.Append(widget.NewAccordionItem(e.Region, exerciseCardForRegion(e, a)))
	}
	return a
}

func getNewExercise(e exercise.Exercise) exercise.Exercise {
	newExercise, _ := exercises.Get().ChooseRandomExerciseForRegion(e.Region)
	return newExercise
}
func exerciseCardForRegion(e exercise.Exercise, a *widget.Accordion) *fyne.Container {
	card := widget.NewCard("", "", container.NewVBox())
	e.Display(card, getNewExercise, a)
	return container.NewVBox(card)
}
