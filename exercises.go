package main

import (
	"errors"
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
type regions []string

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

func (es exercises) chooseRandomExercise() exercise {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return es[r.Intn(len(es)-1)]
}

func (es exercises) chooseRandomExerciseForRegion(r string) (exercise, error) {
	exercisesForRegion := exercises{}
	for _, e := range es {
		if e.region == r {
			exercisesForRegion = append(exercisesForRegion, e)
		}
	}
	if len(exercisesForRegion) == 0 {
		return exercise{}, errors.New("No exercises found for region '" + r + "'")
	}
	return exercisesForRegion.chooseRandomExercise(), nil
}

func (e exercise) printExercise(r int) {
	fmt.Println("--------------------")
	fmt.Println("Region:", e.region)
	fmt.Println("Name:", e.name)
	fmt.Println("Reps:", r)
	fmt.Println("--------------------")
}

func (es exercises) getAllPossibleRegions() regions {
	regions := regions{}
	for _, e := range es {
		hasRegion := false
		for _, r := range regions {
			if e.region == r {
				hasRegion = true
			}
		}
		if !hasRegion {
			regions = append(regions, e.region)
		}
	}
	return regions
}

func (es exercises) getRandomFullBodyWorkout() exercises {
	rs := es.getAllPossibleRegions()
	fullBodyWorkout := exercises{}
	for _, r := range rs {
		e, err := es.chooseRandomExerciseForRegion(r)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fullBodyWorkout = append(fullBodyWorkout, e)
	}
	return fullBodyWorkout
}

func (rs regions) printRegions() {
	for _, r := range rs {
		fmt.Println(r)
	}
}
