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
	r1, _ := es.chooseRandomExerciseForRegion("Shoulders")
	if r1.region != "Shoulders" {
		t.Errorf("Expected region to be 'Shoulders', but got %v", r1)
	}
	r2, _ := es.chooseRandomExerciseForRegion("Shoulders")

	if r1 == r2 {
		// likely to get a conflict here, so getting a different value should hopefully reduce the chance of this being flaky
		r2, _ = es.chooseRandomExerciseForRegion("Shoulders")
		if r1 == r2 {
			t.Errorf("Expected %v and %v to be different exercises", r1, r2)
		}
	}
	r3, err := es.chooseRandomExerciseForRegion("eyeball")
	if err == nil {
		t.Errorf("Expected error response for non existant region, but got %v", r3)
	}
}

func TestGetAllPossibleRegions(t *testing.T) {
	es := getExercises()
	rs := es.getAllPossibleRegions()
	if len(rs) != 8 {
		t.Errorf("Expected length of 8, but got %v", len(rs))

	}
}
