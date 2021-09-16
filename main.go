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
		if os.Args[1] == "help" {
			printHelp()
			os.Exit(0)
		}
		if os.Args[1] == "all" {
			workout := es.getRandomFullBodyWorkout()
			for _, e := range workout {
				e.printExercise(getRandomReps())
			}
			os.Exit(0)
		}
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

func printHelp() {
	fmt.Println("Run program without any arguments to get a random exercise")
	fmt.Println("Run program with \"help\" to see this message")
	fmt.Println("Run program with \"all\" to get full body workout")
	fmt.Println("Run program with a region specified to get a random exercise targeting that region")
	fmt.Println("Acceptable regions are:")
	getExercises().getAllPossibleRegions().printRegions()
	fmt.Println("")
	fmt.Println("All exercises have a random number of reps between 1 and 50")
}
