package people

import (
	"fmt"
	"interview/models"
	"interview/pkg/utilities"
	"math/rand"
)

func GeneratePeoples(num int) []models.Peoples {
	var peoples []models.Peoples

	for i := 0; i < num; i++ {
		people := utilities.GeneratePeople()
		people.ID = i + 1
		peoples = append(peoples, people)
	}

	return peoples
}

func GenderCount(peoples []models.Peoples) string {
	var male, female int

	for _, people := range peoples {
		if people.Gender == "male" {
			male++
		} else {
			female++
		}
	}

	return fmt.Sprintf("total %d males and %d females", male, female)
}

func AssignmentAllocation(peoples []models.Peoples) map[int]string {

	assignment_alloted := make(map[int]string)

	for _, people := range peoples {
		assignment_alloted[people.ID] = utilities.Assignments()[rand.Intn((len(utilities.Assignments())))]
	}

	return assignment_alloted

}
