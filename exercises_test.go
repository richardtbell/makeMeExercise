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
		// likely to get a conflict here, so getting a different value should hopefully reduce the chance of this being flaky
		r2, _ = es.chooseRandomExerciseForRegion("Shoulders")
		if r1 == r2 {
			t.Errorf("Expected %v and %v to be different exercises", r1, r2)
		}
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

func TestGetFullBodyWorkout(t *testing.T) {
	es := getExercises()
	w := es.getRandomFullBodyWorkout()
	if len(w) != 8 {
		t.Errorf("Expected length of 8, but got %v", len(w))
	}
	rs := w.getAllPossibleRegions()
	if len(rs) != 8 {
		t.Errorf("Expected length of 8, but got %v for list of regions %v", len(rs), rs)
	}
}

func TestGetDescriptionFromWebsite(t *testing.T) {
	d1 := exercise{region: "Back", name: "Back Fly"}.getDescriptionFromWebsite()
	expectedBackDescription := "* Lie down on your chest on the bench and grab two dumbbells with your hands, elbows at 90 degree angles.\n* Raise the dumbbells until your arms are parallel to the ground and lower them back after a short pause.\n* Breathe out when pulling the dumbbells and breathe in when lowering them back."
	if d1 != expectedBackDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBackDescription, d1)
	}

	d2 := exercise{region: "Bicep", name: "Alternated Biceps Curl"}.getDescriptionFromWebsite()
	expectedBicepDescription := "* Stand up and hold one dumbbell with each hand down the side of your body, palms facing each other.\n* Raise one dumbbell until it reaches your shoulder's height and while slowly lowering it back down after a short pause, start raising the other one.\n* Try NOT to jerk your upper body in an effort to help you lift the weights."
	if d2 != expectedBicepDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBicepDescription, d2)
	}

	d3 := exercise{region: "Bicep", name: "Biceps Curl"}.getDescriptionFromWebsite()
	expectedBicepCurlDescription := "* Stand up and hold one dumbbell with each hand down the side of your body, palms facing each other.\n* Raise both dumbbells until they reach your shoulders' height and slowly lower them back down after a short pause.\n* Try NOT to jerk your upper body in an effort to help you lift the weights."
	if d3 != expectedBicepCurlDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBicepCurlDescription, d3)
	}

	d4 := exercise{region: "Abdominals", name: "Weighted Crunch"}.getDescriptionFromWebsite()
	expectedAbsDescription := "* Lie down on your back on a bench and hold a dumbbell on top of your chest.\n* Raise your upper body until your shoulder blades no longer touch the bench and lower yourself back down after a short pause.\n* To avoid pulling your neck with your hands, look straight up instead of looking at your knees."
	if d4 != expectedAbsDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedAbsDescription, d4)
	}

	d5 := exercise{region: "Chest", name: "Bench Press"}.getDescriptionFromWebsite()
	expectedBPDescription := "* Lie down on your back on a bench and hold 2 dumbbells at chest level along your body, palms facing your feet.\n* Raise the dumbbells straight up until your elbows are close to being locked and lower them back slowly after a short pause.\n* Breathe out when raising the dumbbells and breathe in when lowering them back."
	if d5 != expectedBPDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBPDescription, d5)
	}

	d6 := exercise{region: "Chest", name: "Bench Press (Neutral Grip)"}.getDescriptionFromWebsite()
	expectedBPNGDescription := "* Lie down on your back on a bench and hold 2 dumbbells at chest level along your body, palms facing each other.\n* Raise the dumbbells straight up until your elbows are close to being locked and lower them back slowly after a short pause.\n* Breathe out when raising the dumbbells and breathe in when lowering them back."
	if d6 != expectedBPNGDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBPNGDescription, d6)
	}

	d7 := exercise{region: "Bicep", name: "Concentrated Biceps Curl"}.getDescriptionFromWebsite()
	expectedd7Description := "* Stand behind an inclined bench and rest one arm on the back support while holding a dumbbell, palm facing up.\n* Raise the dumbbell up to your shoulder and lower it back down after a short pause.\n* Only your lower arm should move throughout the exercise."
	if d7 != expectedd7Description {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedd7Description, d7)
	}
}
