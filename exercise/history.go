package exercise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
