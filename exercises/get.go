package exercises

import (
	"fmt"
	"makeMeExercise/exercise"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Get() Exercises {
	es := Exercises{}
	contents, err := os.ReadFile("store/exerciseList.txt")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	c := strings.Split(string(contents), "\n\n")
	for _, regionExercises := range c {
		var region string
		for i, e := range strings.Split(regionExercises, "\n") {
			if i == 0 {
				region = string(e)
				continue
			}
			es = append(es, exercise.Exercise{Name: string(e), Region: region})
		}
	}
	return es
}

func (es Exercises) GetRandomFullBodyWorkout() Exercises {
	rs := es.GetAllPossibleRegions()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rs), func(i, j int) { rs[i], rs[j] = rs[j], rs[i] })
	fullBodyWorkout := Exercises{}
	for _, r := range rs {
		e, err := es.ChooseRandomExerciseForRegion(r)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fullBodyWorkout = append(fullBodyWorkout, e)
	}
	return fullBodyWorkout
}
