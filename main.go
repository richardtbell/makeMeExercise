package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	es := getExercises()
	if len(os.Args) == 1 {
		e := es.chooseRandomExercise()
		e.printExercise(getRandomReps())
	}
	if len(os.Args) == 2 {
		e, err := es.chooseRandomExerciseForRegion(os.Args[1])
		if err != nil {
			fmt.Println("No exercises found for region: ", os.Args[1])
			fmt.Println("Please choose a region from the following list, or omit to get a random exercise.")
			rs := es.getAllPossibleRegions()
			rs.printRegions()
			os.Exit(1)
		}
		e.printExercise(getRandomReps())
	}
}

func getRandomReps() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(50)
}

// Scrape description (https://dumbbell-exercises.com/exercises/triceps/) (https://dumbbell-exercises.com/exercises/dumbbell-back-exercises/) (https://dumbbell-exercises.com/exercises/dumbbell-exercises-for-biceps/)
// Optional command line argument to have one exercise for each region
