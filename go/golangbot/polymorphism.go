package main

import "fmt"

type Income interface {
	// calculates and return the income from source
	calculate() int
	// name of the source
	source() string
	serviceCharge(i int)
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

func (fb *FixedBilling) serviceCharge(i int) {
	fb.biddedAmount -= i
}

func (tm TimeAndMaterial) calculate() int {
	return tm.totalHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (tm *TimeAndMaterial) serviceCharge(i int) {
	tm.hourlyRate -= (i % 10)
}

// Since both FixedBilling and TimeAndMaterial
// structs provide definitions for the calculate()
// and the source() methods of the Income interface
// both structs implement the Income interface.

func calculateNetIncome(ic []Income) []Income {
	var netIncome int
	for _, income := range ic {
		income.serviceCharge(30)
		fmt.Printf("Income from %s = £%d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Printf("Net income of organisation = £%d\n", netIncome)
	return ic
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
	fmt.Printf("Project 5 £'s %d\n", project5.hourlyRate)
	incomeStreams := []Income{
		&project1,
		&project2,
		&project3,
		&project4,
		&project5,
	}
	incomeStreams = calculateNetIncome(incomeStreams)
	incomeStreams = calculateNetIncome(incomeStreams)
	fmt.Printf("Project 5 £'s %d\n", project5.hourlyRate)
}
