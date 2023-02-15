package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	lop "github.com/samber/lo/parallel"
)

func main() {
	// Example in docs

	example1Fn := lop.Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Println(example1Fn)

	// Practical example
	// Transformation in ETL

	// Closure
	processEmployee := processEmployeeFn

	// Collection
	employees := []Employee{
		{
			Name:                       "Foo",
			Salary:                     42,
			LikesFunctionalProgramming: false,
			StartDate:                  time.Time{},
		},
		{
			Name:                       "Bar",
			Salary:                     42,
			LikesFunctionalProgramming: false,
			StartDate:                  time.Time{},
		},
		{
			Name:                       "Baz",
			Salary:                     42,
			LikesFunctionalProgramming: false,
			StartDate:                  time.Time{},
		},
	}

	fmt.Println("--| Before")
	printEmployees(employees)

	// Map the function over the collection
	// Note: Because of we've eliminated effects in our transformation function
	//  (to the extent at what we can control) we can easily parallelize these operations.
	results := lop.Map(employees, processEmployee) // Evaluation done here

	fmt.Println("--| After")
	printEmployees(results)
}

type Employee struct {
	Name                       string
	Salary                     float64
	LikesFunctionalProgramming bool
	StartDate                  time.Time
}

// x - the value in the collection
// _ - unused variable; the index of the collection
func processEmployeeFn(x Employee, _ int) Employee {
	// Localize TZ
	x.StartDate = x.StartDate.Local()
	// They like FP now
	x.LikesFunctionalProgramming = true
	// Everyone gets a raise
	x.Salary += 42
	return x
}

// Side effects! Ew!
func printEmployees(e []Employee) {
	ejson, _ := json.MarshalIndent(e, "", " ")
	fmt.Printf("%v\n", string(ejson))
}
