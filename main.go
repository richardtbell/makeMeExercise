package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Make me fit")
	es := getExercises()
	e := es.chooseRandomExercise()
	workout := es.getRandomFullBodyWorkout()

	tabs := container.NewAppTabs(
		container.NewTabItem("Random Exercise", exerciseCard(e)),
		container.NewTabItem("Full body workout", displayFullBodyWorkout(workout)),
	)
	win.Resize(fyne.NewSize(1200, 600))
	win.SetContent(tabs)
	win.ShowAndRun()
}

func displayFullBodyWorkout(w Exercises) *widget.Accordion {
	displayedWorkout := widget.NewAccordion()
	for _, e := range w {
		displayedWorkout.Append(widget.NewAccordionItem(e.Region, exerciseCardForRegion(e)))
	}
	return displayedWorkout
}

func displayNewExercise(e Exercise, displayedExercise *widget.Card) {
	newExercise, _ := getExercises().chooseRandomExerciseForRegion(e.Region)
	displayedExercise.SetContent(newExercise.display())
}
func exerciseCardForRegion(e Exercise) *fyne.Container {
	displayedExercise := widget.NewCard("", "", e.display())
	newExerciseButton := widget.NewButton("New "+e.Region+" exercise", func() {
		displayNewExercise(e, displayedExercise)
	})
	saveExerciseButton := widget.NewButton("Completed", func() {
		displayNewExercise(e, displayedExercise)
		e.save()
	})
	return container.NewVBox(displayedExercise, centerButton(newExerciseButton), centerButton(saveExerciseButton))
}

func exerciseCard(e Exercise) *fyne.Container {
	displayedExercise := widget.NewCard("", "", e.display())
	newExerciseButton := widget.NewButton("New random exercise", func() {
		newExercise := getExercises().chooseRandomExercise()
		displayedExercise.SetContent(newExercise.display())
	})
	return container.NewVBox(displayedExercise, centerButton(newExerciseButton))
}

func centerButton(button *widget.Button) *fyne.Container {
	return container.NewHBox(layout.NewSpacer(), button, layout.NewSpacer())
}
