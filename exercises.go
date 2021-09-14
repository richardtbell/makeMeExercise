package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type exercise struct {
	name        string
	region      string
	description string
}

type exercises []exercise

func getExercises() exercises {
	es := exercises{}
	contents, err := os.ReadFile("exerciseList.txt")
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
			es = append(es, exercise{name: string(e), region: region})
		}
	}
	return es
}

func (es exercises) choseRandomExercise() exercise {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return es[r.Intn(len(es)-1)]
}

func (e exercise) printExercise(r int) {
	fmt.Println("Region:", e.region)
	fmt.Println("Name:", e.name)
	fmt.Println("Reps:", r)
}
