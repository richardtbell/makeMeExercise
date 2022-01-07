package exercises

import "testing"

func TestChooseRandomExercise(t *testing.T) {
	es := Get()
	r1 := es.ChooseRandomExercise()
	r2 := es.ChooseRandomExercise()
	if r1 == r2 {
		// likely to get a conflict here, so getting a different value should hopefully reduce the chance of this being flaky
		r2 = es.ChooseRandomExercise()
		if r1 == r2 {
			t.Errorf("Expected %v and %v to be different exercises", r1, r2)
		}
	}
}
func TestChooseRandomExerciseForRegion(t *testing.T) {
	es := Get()
	r1, _ := es.ChooseRandomExerciseForRegion("Shoulders")
	if r1.Region != "Shoulders" {
		t.Errorf("Expected region to be 'Shoulders', but got %v", r1)
	}
	r2, _ := es.ChooseRandomExerciseForRegion("Shoulders")

	if r1 == r2 {
		// likely to get a conflict here, so getting a different value should hopefully reduce the chance of this being flaky
		r2, _ = es.ChooseRandomExerciseForRegion("Shoulders")
		if r1 == r2 {
			t.Errorf("Expected %v and %v to be different exercises", r1, r2)
		}
	}
	r3, err := es.ChooseRandomExerciseForRegion("eyeball")
	if err == nil {
		t.Errorf("Expected error response for non existant region, but got %v", r3)
	}
}
