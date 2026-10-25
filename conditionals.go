// A file using various conditional statements to examine values in a struct
package main

import "fmt"

type Player struct {
	LastName    string
	Age         int
	Captaincy   bool
	Points      int
	AwardsWon   []string
	GamesPlayed int
	Team        string
}

func main() {

	// instantiating a Player struct
	p := Player{
		LastName:    "McDavid",
		Age:         30,
		Captaincy:   true,
		Points:      153,
		AwardsWon:   []string{"Hart", "Ted Lindsay", "Art Ross", "Rocket Richard"},
		GamesPlayed: 82,
		Team:        "Oilers"}
	fmt.Println(p, "\\n")

	// one-condition if/else statement
	if p.GamesPlayed == 82 {
		fmt.Println("The season is complete! This is an end of season player analysis.\\n")
	} else {
		fmt.Print("These stats are not finalized!\\n")
	}

	// if/else/esle if statement
	if len(p.AwardsWon) == 0 {
		fmt.Println(p.LastName, "didn't win any awards.")
	} else if len(p.AwardsWon) == 1 {
		fmt.Print(p.LastName, "won 1 award. Not easy to do!")
	} else if len(p.AwardsWon) == 2 {
		fmt.Print(p.LastName, "won 2 awards! That's very impressive!")
	} else if len(p.AwardsWon) == 3 {
		fmt.Print(p.LastName, "won 3 awards! Three?? Must be an outstanding player!")
	} else {
		fmt.Println(p.LastName, "won more tha 3 awards!! Hall of famer in the making.")
	}

	// multi-condition if statement
	if p.Age > 29 || p.Captaincy { // notice the short-circut evaluation
		fmt.Println(p.LastName, "is an important team leader")
	}

	// switch statement
	switch p.Team {
	case "Golden Knights":
		fmt.Println(p.LastName, "won the Stanley Cup!")
	case "Panthers":
		fmt.Println(p.LastName, "lost in the Stanley Cup Final...")
	default:
		fmt.Println(p.LastName, "plays for", p.Team, "and was not in the Stanley Cup Final")
	}

}
