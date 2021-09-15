package main

import "testing"

func TestGetExercises(t *testing.T) {
	es := getExercises()
	if len(es) != 78 {
		t.Errorf("Expected length of 78, but got %v", len(es))
	}
	e := es[0]
	if e.name != "Weighted Crunch" {
		t.Errorf("Expected first item to have name 'Weighted Crunch', but got '%v'", e.name)
	}
	if e.region != "Abdominals" {
		t.Errorf("Expected first item to have region 'Adominals', but got '%v'", e.name)
	}
}

func TestChooseRandomExercise(t *testing.T) {
	es := getExercises()
	r1 := es.chooseRandomExercise()
	r2 := es.chooseRandomExercise()
	if r1 == r2 {
		t.Errorf("Expected %v and %v to be different exercises", r1, r2)
	}
}

func TestChooseRandomExerciseForRegion(t *testing.T) {
	es := getExercises()
	r1 := es.chooseRandomExerciseForRegion("Shoulders")
	if r1.region != "Shoulders" {
		t.Errorf("Expected region to be 'Shoulders', but got %v", r1)
	}
	r2 := es.chooseRandomExerciseForRegion("Shoulders")

	if r1 == r2 {
		t.Errorf("Expected %v and %v to be different exercises", r1, r2)
	}
}
