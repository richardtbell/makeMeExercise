package exercises

import (
	"testing"
)

func TestGetAllPossibleRegions(t *testing.T) {
	es := Get()
	rs := es.GetAllPossibleRegions()
	if len(rs) != 8 {
		t.Errorf("Expected length of 8, but got %v", len(rs))
	}
}
