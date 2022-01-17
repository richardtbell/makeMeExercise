package exercise

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type getNewExercise func(Exercise) Exercise

func (e Exercise) Display(card *widget.Card, getNewExercise getNewExercise) {
	e.PrintExercise()
	name := canvas.NewText(e.Name, color.White)
	name.Alignment = fyne.TextAlignCenter
	name.TextSize = 24
	reps := canvas.NewText("Reps: "+strconv.Itoa(e.Reps), color.White)
	reps.Alignment = fyne.TextAlignCenter
	reps.TextSize = 18
	region := canvas.NewText("Region: "+e.Region, color.White)
	description := widget.NewLabel(e.Description)
	prev := e.DisplayPreviousAttempts()

	newExerciseButton := widget.NewButton("Change "+e.Region+" exercise", func() {
		newExercise := getNewExercise(e)
		newExercise.Display(card, getNewExercise)
	})
	weightInput := widget.NewEntry()
	weightInput.SetPlaceHolder("Enter weight")
	weightInput.Validator = func(s string) (err error) { _, err = strconv.ParseFloat(s, 64); return }
	difficultyInput := widget.NewRadioGroup([]string{"Easy", "Medium", "Hard"}, func(value string) {})
	saveExerciseButton := widget.NewButton("Completed", func() {
		e.Save(weightInput.Text, difficultyInput.Selected)
		newExercise := getNewExercise(e)
		newExercise.Display(card, getNewExercise)
	})

	card.SetContent(container.NewVBox(name, reps, region, description, prev, centerButton(newExerciseButton), saveRow(saveExerciseButton, weightInput, difficultyInput)))
}

func centerButton(button *widget.Button) *fyne.Container {
	return container.NewHBox(layout.NewSpacer(), button, layout.NewSpacer())
}

func saveRow(button *widget.Button, weightInput *widget.Entry, difficultyInput *widget.RadioGroup) *fyne.Container {
	return container.New(layout.NewGridLayout(3), weightInput, difficultyInput, button)
}
