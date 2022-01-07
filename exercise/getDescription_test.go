package exercise

import "testing"

func TestGetDescriptionFromWebsite(t *testing.T) {
	d1 := Exercise{Region: "Back", Name: "Back Fly"}.getDescriptionFromWebsite()
	expectedBackDescription := "* Lie down on your chest on the bench and grab two dumbbells with your hands, elbows at 90 degree angles.\n* Raise the dumbbells until your arms are parallel to the ground and lower them back after a short pause.\n* Breathe out when pulling the dumbbells and breathe in when lowering them back."
	if d1 != expectedBackDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBackDescription, d1)
	}

	d2 := Exercise{Region: "Bicep", Name: "Alternated Biceps Curl"}.getDescriptionFromWebsite()
	expectedBicepDescription := "* Stand up and hold one dumbbell with each hand down the side of your body, palms facing each other.\n* Raise one dumbbell until it reaches your shoulder's height and while slowly lowering it back down after a short pause, start raising the other one.\n* Try NOT to jerk your upper body in an effort to help you lift the weights."
	if d2 != expectedBicepDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBicepDescription, d2)
	}

	d3 := Exercise{Region: "Bicep", Name: "Biceps Curl"}.getDescriptionFromWebsite()
	expectedBicepCurlDescription := "* Stand up and hold one dumbbell with each hand down the side of your body, palms facing each other.\n* Raise both dumbbells until they reach your shoulders' height and slowly lower them back down after a short pause.\n* Try NOT to jerk your upper body in an effort to help you lift the weights."
	if d3 != expectedBicepCurlDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBicepCurlDescription, d3)
	}

	d4 := Exercise{Region: "Abdominals", Name: "Weighted Crunch"}.getDescriptionFromWebsite()
	expectedAbsDescription := "* Lie down on your back on a bench and hold a dumbbell on top of your chest.\n* Raise your upper body until your shoulder blades no longer touch the bench and lower yourself back down after a short pause.\n* To avoid pulling your neck with your hands, look straight up instead of looking at your knees."
	if d4 != expectedAbsDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedAbsDescription, d4)
	}

	d5 := Exercise{Region: "Chest", Name: "Bench Press"}.getDescriptionFromWebsite()
	expectedBPDescription := "* Lie down on your back on a bench and hold 2 dumbbells at chest level along your body, palms facing your feet.\n* Raise the dumbbells straight up until your elbows are close to being locked and lower them back slowly after a short pause.\n* Breathe out when raising the dumbbells and breathe in when lowering them back."
	if d5 != expectedBPDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBPDescription, d5)
	}

	d6 := Exercise{Region: "Chest", Name: "Bench Press (Neutral Grip)"}.getDescriptionFromWebsite()
	expectedBPNGDescription := "* Lie down on your back on a bench and hold 2 dumbbells at chest level along your body, palms facing each other.\n* Raise the dumbbells straight up until your elbows are close to being locked and lower them back slowly after a short pause.\n* Breathe out when raising the dumbbells and breathe in when lowering them back."
	if d6 != expectedBPNGDescription {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedBPNGDescription, d6)
	}

	d7 := Exercise{Region: "Bicep", Name: "Concentrated Biceps Curl"}.getDescriptionFromWebsite()
	expectedd7Description := "* Stand behind an inclined bench and rest one arm on the back support while holding a dumbbell, palm facing up.\n* Raise the dumbbell up to your shoulder and lower it back down after a short pause.\n* Only your lower arm should move throughout the exercise."
	if d7 != expectedd7Description {
		t.Errorf("Expected: \n%v\nGot: \n%v", expectedd7Description, d7)
	}
}
