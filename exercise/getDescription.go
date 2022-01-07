package exercise

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func (e Exercise) GetDescription() string {
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
	body, _ := ioutil.ReadAll(resp.Body)
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
