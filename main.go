// A file exploring some of the built in data types in Go and some simple use cases

package main

import "fmt"

func main() {

	// Initializing a sctuct instance
	p := Player{LastName: "Celebrini", Team: "Sharks", Salary: 1.7, Goals: 45, Assists: 67}
	fmt.Println(p)

	// Accessing fields from the struct
	fmt.Println(p.LastName)

	// Modifying fields from a struct
	p.Salary = 18.5
	fmt.Println("\nUpdated Salary:", p)

	// Creating a pointer to reference a struct
	s := &p
	s.Goals = s.Goals + 1 // Addition allowed because p.Goals refers to an int
	fmt.Println("\nUpdated Goals:", s.Goals)

	// Creating a slice
	var dailyRevenue []int = []int{250, 349, 580, 390, 207, 873}
	dailyRevenue = append(dailyRevenue, 500)
	
	if(len(dailyRevenue) == 7) {
		fullWeek := true
		fmt.Println("\nA full week of data was recorded: ", fullWeek)
	}
	
	// Conversion of Numeric data types
	var pts int = p.Goals + p.Assists
	fmt.Println("\nPoints:", pts)
	ptsFloat := float32(pts) // Int must be converted to float for division to work
	var PPG float32 = ptsFloat/82
	fmt.Println("Points per game:", PPG)
}

// Defining the struct 
type Player struct {
	LastName string // Declaring the type of attributes
	Team string
	Salary float64
	Goals int
	Assists int
}
