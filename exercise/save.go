package exercise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type CompletedExercise struct {
	Exercise   Exercise
	Weight     float64
	Difficulty string
	Date       string
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Exercise) Save() {
	filename := "completed.json"
	err := checkFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	data := []CompletedExercise{}

	// Here the magic happens!
	json.Unmarshal(file, &data)
	currentTime := time.Time.Format(time.Now(), "2006-01-02 15:04")
	toSave := &CompletedExercise{
		Exercise: e,
		Date:     currentTime,
	}

	data = append(data, *toSave)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
