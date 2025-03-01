//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

type color string

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favouriteColor color  = "lilac"
	fmt.Println("my favourite color", favouriteColor)
	
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment

	birthYear, ageInYears := 1986, 35
	fmt.Println("I was born in", birthYear)
	fmt.Println("I am", ageInYears,"old")

	//* Store your first & last initials in two variables using block assignment
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var (
		firstInitial = "W"
		secondInitial = "P"
	)
	fmt.Println("Initials=", firstInitial, secondInitial)
	ageInDays := ageInYears * 365
	fmt.Println("I am", ageInDays, "days old")
}
