package exercises

import (
	"errors"
	"makeMeExercise/exercise"
	"math/rand"
	"time"
)

func getRandomReps() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(49) + 1
}

func (es Exercises) ChooseRandomExercise() exercise.Exercise {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	e := es[r.Intn(len(es)-1)]
	e.Description = e.GetDescription()
	e.Reps = getRandomReps()
	return e
}

func (es Exercises) ChooseRandomExerciseForRegion(r string) (exercise.Exercise, error) {
	exercisesForRegion := Exercises{}
	for _, e := range es {
		if e.Region == r {
			exercisesForRegion = append(exercisesForRegion, e)
		}
	}
	if len(exercisesForRegion) == 0 {
		return exercise.Exercise{}, errors.New("No Exercises found for region '" + r + "'")
	}
	return exercisesForRegion.ChooseRandomExercise(), nil
}
