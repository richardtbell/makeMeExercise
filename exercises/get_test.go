package exercises

import "testing"

func TestGetExercises(t *testing.T) {
	es := Get()
	if len(es) != 78 {
		t.Errorf("Expected length of 78, but got %v", len(es))
	}
	e := es[0]
	if e.Name != "Weighted Crunch" {
		t.Errorf("Expected first item to have name 'Weighted Crunch', but got '%v'", e.Name)
	}
	if e.Region != "Abdominals" {
		t.Errorf("Expected first item to have region 'Adominals', but got '%v'", e.Region)
	}
}
func TestGetFullBodyWorkout(t *testing.T) {
	es := Get()
	w := es.GetRandomFullBodyWorkout()
	if len(w) != 8 {
		t.Errorf("Expected length of 8, but got %v", len(w))
	}
	rs := w.GetAllPossibleRegions()
	if len(rs) != 8 {
		t.Errorf("Expected length of 8, but got %v for list of regions %v", len(rs), rs)
	}
}
