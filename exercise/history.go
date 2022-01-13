package exercise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type History []CompletedExercise

var HISTORY_FILENAME = "store/history.json"

func getAllHistory() (history History) {
	file, err := ioutil.ReadFile(HISTORY_FILENAME)
	if err != nil {
		fmt.Println(err)
	}

	// Here the magic happens!
	json.Unmarshal(file, &history)
	return
}

func (e Exercise) GetPreviousAttempts() (attempts History) {
	hist := getAllHistory()
	for _, attempt := range hist {
		if attempt.Exercise.Name == e.Name {
			attempts = append(attempts, attempt)
		}
	}
	return
}

func (e Exercise) DisplayPreviousAttempts() *fyne.Container {
	attempts := e.GetPreviousAttempts()
	data := []string{}
	for _, attempt := range attempts {
		fmt.Println(attempt)
		data = append(data, "Reps: "+strconv.Itoa(attempt.Exercise.Reps)+" Weight: "+attempt.Weight+" "+attempt.Difficulty)
	}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Previous Attempts")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	return container.New(layout.NewGridLayout(3), list)
}
