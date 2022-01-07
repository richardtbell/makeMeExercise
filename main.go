package main

import (
	"makeMeExercise/exercise"
	"makeMeExercise/exercises"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Make me fit")
	es := exercises.Get()
	e := es.ChooseRandomExercise()
	workout := es.GetRandomFullBodyWorkout()

	tabs := container.NewAppTabs(
		container.NewTabItem("Random Exercise", exerciseCard(e)),
		container.NewTabItem("Full body workout", displayFullBodyWorkout(workout)),
	)
	win.Resize(fyne.NewSize(1200, 600))
	win.SetContent(tabs)
	win.ShowAndRun()
}

func displayFullBodyWorkout(w exercises.Exercises) *widget.Accordion {
	displayedWorkout := widget.NewAccordion()
	for _, e := range w {
		displayedWorkout.Append(widget.NewAccordionItem(e.Region, exerciseCardForRegion(e)))
	}
	return displayedWorkout
}

func displayNewExercise(e exercise.Exercise, displayedExercise *widget.Card) {
	newExercise, _ := exercises.Get().ChooseRandomExerciseForRegion(e.Region)
	displayedExercise.SetContent(newExercise.Display())
}
func exerciseCardForRegion(e exercise.Exercise) *fyne.Container {
	displayedExercise := widget.NewCard("", "", e.Display())
	newExerciseButton := widget.NewButton("New "+e.Region+" exercise", func() {
		displayNewExercise(e, displayedExercise)
	})
	saveExerciseButton := widget.NewButton("Completed", func() {
		displayNewExercise(e, displayedExercise)
		e.Save()
	})
	return container.NewVBox(displayedExercise, centerButton(newExerciseButton), centerButton(saveExerciseButton))
}

func exerciseCard(e exercise.Exercise) *fyne.Container {
	displayedExercise := widget.NewCard("", "", e.Display())
	newExerciseButton := widget.NewButton("New random exercise", func() {
		newExercise := exercises.Get().ChooseRandomExercise()
		displayedExercise.SetContent(newExercise.Display())
	})
	return container.NewVBox(displayedExercise, centerButton(newExerciseButton))
}

func centerButton(button *widget.Button) *fyne.Container {
	return container.NewHBox(layout.NewSpacer(), button, layout.NewSpacer())
}
