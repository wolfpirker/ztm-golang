//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func printContents(line []Part) {
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
}



func main() {
	//* Using a slice, create an assembly line that contains type Part
	//  - Create an assembly line having any three parts
	assemblyLine := make([]Part, 3)

	assemblyLine[0] = "Scissors"
	assemblyLine[1] = "Stone"
	assemblyLine[2] = "Sheet"

	printContents(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Pen", "Fork")
	fmt.Println("\nAdded two parts:")
	printContents(assemblyLine)


	//  - Slice the assembly line so it contains only the two new parts
	//  - Print out the contents of the assembly line at each step
	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced:")
	printContents(assemblyLine)

}
