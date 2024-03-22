package main

import "fmt"

func main() {
    var name string
    var age int

    // Take user input for name
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    // Take user input for age
    fmt.Print("Enter your age: ")
    fmt.Scanln(&age)

    // Print greeting message
    fmt.Printf("Hello, %s!\n", name)

    // Calculate age after 10 years
    ageAfter10Years := age + 10
    fmt.Printf("In 10 years, you will be %d years old.\n", ageAfter10Years)
}
