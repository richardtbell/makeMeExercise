package exercise

import (
	"fmt"
	"os"
)

func (e Exercise) GetDescription() string {
	return e.getDescriptionFromFile()
}

func (e Exercise) getDescriptionFromFile() string {
	contents, err := os.ReadFile("descriptions/" + e.Name)
	if err != nil {
		fmt.Println("Error", err)
	}
	return string(contents)
}
