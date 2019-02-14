package main

import "fmt"

type Income interface {
	// calculates and return the income from source
	calculate() int
	// name of the source
	source() string
}

// FixedBilling is a type of Project
type FixedBilling struct {
	// name of the project
	projectName string
	// the amount that the org has bid for project
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	totalHours  int
	hourlyRate  int
}

// Define the methods on these struct types which
// calculate and return the actual income and
// source of income.

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.totalHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

// Since both FixedBilling and TimeAndMaterial
// structs provide definitions for the calculate()
// and the source() methods of the Income interface
// both structs implement the Income interface.

func calculateNetIncome(ic []Income) {
	var netIncome int
	for _, income := range ic {
		fmt.Printf("Income from %s = £%d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Printf("Net income of organisation = £%d", netIncome)
}

// The calculateNetIncome function accepts a slice
// of Income interfaces as argument. It calculates
// the total income by iterating over the slice
// and calling calculate() method on each of its
// items.

// In the future, if a new type of project is added
// this function will still calculate the total
// income correctly without a single line of code
// change.

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 300}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 230}
	project3 := FixedBilling{projectName: "Project 3", biddedAmount: 190}
	project4 := TimeAndMaterial{
		projectName: "Project 4",
		totalHours:  40,
		hourlyRate:  20,
	}
	project5 := TimeAndMaterial{
		projectName: "Project 5",
		totalHours:  130,
		hourlyRate:  18,
	}
	incomeStreams := []Income{
		project1,
		project2,
		project3,
		project4,
		project5,
	}
	calculateNetIncome(incomeStreams)
}
