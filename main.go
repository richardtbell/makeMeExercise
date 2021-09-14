package main

import (
	"math/rand"
	"time"
)

func main() {
	es := getExercises()
	e := es.choseRandomExercise()
	e.printExercise(getRandomReps())
}

func getRandomReps() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(50)
}
