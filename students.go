package main

import (
	"fmt"
)

func main() {
	// Initialize variables to store user data
	var name, regNo, homeLocation string
	var gender int

	// Get user data
	fmt.Print("Enter student's name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter student's registration number: ")
	fmt.Scanln(&regNo)

	fmt.Print(" gender: \n 1. Male \n 2. Female \n Enter: ")
	fmt.Scanln(&gender)

	// Validate gender
	if gender != 1 && gender != 2 {
		fmt.Println("Invalid gender. Please enter either 1 for Male or 2 for Female.")
		return
	}

	fmt.Print("Enter home location: ")
	fmt.Scanln(&homeLocation)

	// Print collected user data
	fmt.Printf("\nCollected data:\n")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Registration Number: %s\n", regNo)
	fmt.Printf("Gender: %s\n", genderStr(gender))
	fmt.Printf("Home Location: %s\n", homeLocation)
}

func genderStr(g int) string {
	switch g {
	case 1:
		return "Male"
	case 2:
		return "Female"
	default:
		return "Invalid gender"
	}
}
