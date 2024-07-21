package utilities

import (
	"interview/models"
	"math/rand"
)

var (
	gender      = []string{"male", "female"}
	maleName    = []string{"nishad", "akshay", "vishnu"}
	femaleName  = []string{"neha", "arshia", "sana"}
	assignments = []string{"assignment_1", "assignment_2", "assignment_3", "assignment_4", "assignment_5"}
)

func GeneratePeople() models.Peoples {

	var people models.Peoples

	people.Gender = gender[rand.Intn(2)]

	if people.Gender == "female" {
		people.Name = femaleName[rand.Intn(len(femaleName))]
	} else {
		people.Name = maleName[rand.Intn(len(maleName))]
	}

	return people
}

func Assignments() []string {
	return assignments
}
