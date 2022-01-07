package main

import (
	"encoding/json"
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
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Exercise struct {
	Name        string
	Region      string
	Description string
	Reps        int
}

type Exercises []Exercise
type Regions []string
type completedExercise struct {
	Exercise   Exercise
	Weight     float64
	Difficulty string
	Date       string
}

func getExercises() Exercises {
	es := Exercises{}
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
			es = append(es, Exercise{Name: string(e), Region: region})
		}
	}
	return es
}

func (es Exercises) chooseRandomExercise() Exercise {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	e := es[r.Intn(len(es)-1)]
	e.Description = e.getDescription()
	e.Reps = getRandomReps()
	return e
}

func (es Exercises) chooseRandomExerciseForRegion(r string) (Exercise, error) {
	exercisesForRegion := Exercises{}
	for _, e := range es {
		if e.Region == r {
			exercisesForRegion = append(exercisesForRegion, e)
		}
	}
	if len(exercisesForRegion) == 0 {
		return Exercise{}, errors.New("No Exercises found for region '" + r + "'")
	}
	return exercisesForRegion.chooseRandomExercise(), nil
}

func (es Exercises) getAllPossibleRegions() Regions {
	regions := Regions{}
	for _, e := range es {
		hasRegion := false
		for _, r := range regions {
			if e.Region == r {
				hasRegion = true
			}
		}
		if !hasRegion {
			regions = append(regions, e.Region)
		}
	}
	return regions
}

func (es Exercises) getRandomFullBodyWorkout() Exercises {
	rs := es.getAllPossibleRegions()
	fullBodyWorkout := Exercises{}
	for _, r := range rs {
		e, err := es.chooseRandomExerciseForRegion(r)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fullBodyWorkout = append(fullBodyWorkout, e)
	}
	return fullBodyWorkout
}

func (e Exercise) printExercise() {
	fmt.Println("--------------------")
	fmt.Println("Region:", e.Region)
	fmt.Println("Name:", e.Name)
	fmt.Println("Reps:", e.Reps)
	fmt.Println(e.Description)
	fmt.Println("--------------------")
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Exercise) save() {
	filename := "completed.json"
	err := checkFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	data := []completedExercise{}

	// Here the magic happens!
	json.Unmarshal(file, &data)
	currentTime := time.Time.Format(time.Now(), "2006-01-02 15:04")
	toSave := &completedExercise{
		Exercise: e,
		Date:     currentTime,
	}

	data = append(data, *toSave)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func (e Exercise) display() *fyne.Container {
	e.printExercise()
	name := canvas.NewText(e.Name, color.White)
	name.Alignment = fyne.TextAlignCenter
	name.TextSize = 24
	reps := canvas.NewText("Reps: "+strconv.Itoa(e.Reps), color.White)
	reps.Alignment = fyne.TextAlignCenter
	reps.TextSize = 18
	region := canvas.NewText("Region: "+e.Region, color.White)
	description := widget.NewLabel(e.Description)
	return container.NewVBox(name, reps, region, description)
}

func (e Exercise) getDescription() string {
	d := e.getDescriptionFromFile()
	if len(d) == 0 {
		d = e.getDescriptionFromWebsite()
	}
	return d
}

func (e Exercise) getDescriptionFromFile() string {
	contents, err := os.ReadFile("descriptions/" + e.Name)
	if err != nil {
		fmt.Println("Error", err)
	}
	return string(contents)
}

func (e Exercise) getDescriptionFromWebsite() string {
	baseUrl := "https://dumbbell-exercises.com/exercises/"
	url := baseUrl + e.Region
	if e.Region == "Back" {
		url = "https://dumbbell-exercises.com/exercises/dumbbell-back-exercises"
	}
	if e.Region == "Bicep" {
		url = "https://dumbbell-exercises.com/exercises/dumbbell-exercises-for-biceps"
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	sanitisedName := strings.ReplaceAll(e.Name, "(", "\\(")
	sanitisedName = strings.ReplaceAll(sanitisedName, ")", "\\)")
	reg := regexp.MustCompile(`>` + sanitisedName + `<`)
	nameIndex := reg.FindAllStringSubmatchIndex(string(body), -1)
	if len(nameIndex) == 0 {
		fmt.Println("Error: Could not find description for", e.Name)
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
	os.WriteFile("descriptions/"+e.Name, []byte(d), 0666)
	return d
}

func getRandomReps() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(49) + 1
}
