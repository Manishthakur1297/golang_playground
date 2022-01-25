//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var color string = "blue"
	fmt.Println("----Color----", color)

	//* Store your birth year and age (in years) in two variables using compound assignment
	birthYear, age := 1997, 24
	fmt.Println("----BirthYear---", birthYear, " ----Age----", age)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "M"
		lastInitial  = "T"
	)

	fmt.Println("----firstInitial---", firstInitial, " ----lastInitial----", lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	var birthDays int
	fmt.Println("----BirthDays W/o Intialization---", birthDays)

	//  then assign it on the next line by multiplying 365 with the age variable created earlier

	birthDays = age * 365
	fmt.Println("----BirthDays With Intialization---", birthDays)

}
