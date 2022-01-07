package exercise

import "fmt"

func (e Exercise) PrintExercise() {
	fmt.Println("--------------------")
	fmt.Println("Region:", e.Region)
	fmt.Println("Name:", e.Name)
	fmt.Println("Reps:", e.Reps)
	fmt.Println(e.Description)
	fmt.Println("--------------------")
}
