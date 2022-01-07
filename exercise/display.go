package exercise

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (e Exercise) Display() *fyne.Container {
	e.PrintExercise()
	name := canvas.NewText(e.Name, color.White)
	name.Alignment = fyne.TextAlignCenter
	name.TextSize = 24
	reps := canvas.NewText("Reps: "+strconv.Itoa(e.Reps), color.White)
	reps.Alignment = fyne.TextAlignCenter
	reps.TextSize = 18
	region := canvas.NewText("Region: "+e.Region, color.White)
	description := widget.NewLabel(e.Description)
	return container.NewVBox(name, reps, region, description)
}
