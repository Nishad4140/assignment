package main

import (
	"fmt"
	"interview/pkg/people"
)

func main() {

	// generate random peoples
	generatedPeoples := people.GeneratePeoples(15)

	// checking the total number of male and female
	fmt.Println(people.GenderCount(generatedPeoples))

	// assignement allocation
	assignement_allotted := people.AssignmentAllocation(generatedPeoples)
	fmt.Println(assignement_allotted)
	

}
