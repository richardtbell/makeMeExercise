package main

import (
	"errors"
	"fmt"
	"image/color"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type exercise struct {
	name        string
	region      string
	description string
	reps        int
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
	e := es[r.Intn(len(es)-1)]
	e.description = e.getDescription()
	e.reps = getRandomReps()
	return e
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

func (e exercise) printExercise() {
	fmt.Println("--------------------")
	fmt.Println("Region:", e.region)
	fmt.Println("Name:", e.name)
	fmt.Println("Reps:", e.reps)
	fmt.Println(e.description)
	fmt.Println("--------------------")
}

func (e exercise) display() (*canvas.Text, *canvas.Text, *canvas.Text, *widget.Label) {
	e.printExercise()
	name := canvas.NewText(e.name, color.White)
	name.Alignment = fyne.TextAlignCenter
	name.TextSize = 24
	reps := canvas.NewText("Reps: "+strconv.Itoa(e.reps), color.White)
	reps.Alignment = fyne.TextAlignCenter
	reps.TextSize = 18
	region := canvas.NewText("Region: "+e.region, color.White)
	description := widget.NewLabel(e.description)
	return name, reps, region, description
}

func (e exercise) getDescription() string {
	d := e.getDescriptionFromFile()
	if len(d) == 0 {
		d = e.getDescriptionFromWebsite()
	}
	return d
}

func (e exercise) getDescriptionFromFile() string {
	contents, err := os.ReadFile("descriptions/" + e.name)
	if err != nil {
		fmt.Println("Error", err)
	}
	return string(contents)
}

func (e exercise) getDescriptionFromWebsite() string {
	baseUrl := "https://dumbbell-exercises.com/exercises/"
	url := baseUrl + e.region
	if e.region == "Back" {
		url = "https://dumbbell-exercises.com/exercises/dumbbell-back-exercises"
	}
	if e.region == "Bicep" {
		url = "https://dumbbell-exercises.com/exercises/dumbbell-exercises-for-biceps"
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	sanitisedName := strings.ReplaceAll(e.name, "(", "\\(")
	sanitisedName = strings.ReplaceAll(sanitisedName, ")", "\\)")
	reg := regexp.MustCompile(`>` + sanitisedName + `<`)
	nameIndex := reg.FindAllStringSubmatchIndex(string(body), -1)
	if len(nameIndex) == 0 {
		fmt.Println("Error: Could not find description for", e.name)
		return ""
	}
	var ni int
	if len(nameIndex) == 1 {
		ni = nameIndex[0][0]
	}
	if len(nameIndex) > 1 {
		ni = nameIndex[1][0]
	}
	preg := regexp.MustCompile(`<p>.*</p>`)
	d := string(body)[ni:]
	d = d[0:strings.Index(d, "</ul>")]
	d = d[strings.Index(d, "/>"):]
	d = preg.ReplaceAllString(d, "")
	d = strings.ReplaceAll(d, "/>", "")
	d = strings.ReplaceAll(d, "</div>", "")
	d = strings.ReplaceAll(d, "<ul>", "")
	d = strings.ReplaceAll(d, "<li>", "* ")
	d = strings.ReplaceAll(d, "</li>", "")
	d = strings.ReplaceAll(d, "&#8217;", "'")
	d = strings.ReplaceAll(d, "\n", "\n")
	d = strings.TrimSpace(d)
	if d[0] != '*' {
		d = "* " + d
	}
	os.WriteFile("descriptions/"+e.name, []byte(d), 0666)
	return d
}

func getRandomReps() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(50)
}
