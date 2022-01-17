package exercise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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

func (e Exercise) DisplayPreviousAttempts() *widget.Label {
	attempts := e.GetPreviousAttempts()
	data := []string{}
	for _, attempt := range attempts {
		fmt.Println(attempt)
		data = append(data, "Reps: "+strconv.Itoa(attempt.Exercise.Reps)+" Weight: "+attempt.Weight+" "+attempt.Difficulty)
	}
	list := widget.NewLabel(strings.Join(data, "\n"))
	return list
}
