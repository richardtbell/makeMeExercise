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
	// e := es.ChooseRandomExercise()
	workout := es.GetRandomFullBodyWorkout()

	tabs := container.NewAppTabs(
		// container.NewTabItem("Random Exercise", exerciseCard(e)),
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

func getNewExercise(e exercise.Exercise) exercise.Exercise {
	newExercise, _ := exercises.Get().ChooseRandomExerciseForRegion(e.Region)
	return newExercise
}
func exerciseCardForRegion(e exercise.Exercise) *fyne.Container {
	card := widget.NewCard("", "", container.NewVBox())
	e.Display(card, getNewExercise)
	return container.NewVBox(card)
	// displayedExercise := widget.NewCard("", "", e.Display())
	// newExerciseButton := widget.NewButton("Change "+e.Region+" exercise", func() {
	// 	displayNewExercise(&e, displayedExercise)
	// })
	// weightInput := widget.NewEntry()
	// weightInput.SetPlaceHolder("Enter weight")
	// weightInput.Validator = func(s string) (err error) { _, err = strconv.ParseFloat(s, 64); return }
	// difficultyInput := widget.NewRadioGroup([]string{"Easy", "Medium", "Hard"}, func(value string) {
	// 	log.Println("Radio set to", value)
	// })
	// saveExerciseButton := widget.NewButton("Completed", func() {
	// 	displayNewExercise(&e, displayedExercise)
	// 	e.Save(weightInput.Text, difficultyInput.Selected)
	// })
	// prev := e.DisplayPreviousAttempts()
	// return container.NewVBox(displayedExercise, prev, centerButton(newExerciseButton), saveRow(saveExerciseButton, weightInput, difficultyInput))
}

// func exerciseCard(e exercise.Exercise) *fyne.Container {
// 	displayedExercise := widget.NewCard("", "", e.Display())
// 	newExerciseButton := widget.NewButton("New random exercise", func() {
// 		newExercise := exercises.Get().ChooseRandomExercise()
// 		displayedExercise.SetContent(newExercise.Display())
// 	})
// 	return container.NewVBox(displayedExercise, centerButton(newExerciseButton))
// }

// func centerButton(button *widget.Button) *fyne.Container {
// 	return container.NewHBox(layout.NewSpacer(), button, layout.NewSpacer())
// }

// func saveRow(button *widget.Button, weightInput *widget.Entry, difficultyInput *widget.RadioGroup) *fyne.Container {
// 	return container.New(layout.NewGridLayout(3), weightInput, difficultyInput, button)
// 	// weight := container.NewVBox(layout.NewSpacer(), weightInput, layout.NewSpacer())
// 	// b := container.NewVBox(layout.NewSpacer(), button, layout.NewSpacer())
// 	// return container.NewHBox(layout.NewSpacer(), weight, difficultyInput, b, layout.NewSpacer())
// }
