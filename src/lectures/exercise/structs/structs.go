//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x, y float64
}

type Rectangle struct {
	Point1  Coordinate
	Point2  Coordinate
}

func calcPerimeter(r Rectangle) float32 {
	return float32(math.Abs(r.Point1.x - r.Point2.x)*2 + 2*math.Abs(r.Point1.y - r.Point2.y)) 
}

func calcArea(r Rectangle) float32 {
	return float32(math.Abs(r.Point1.x - r.Point2.x)*math.Abs(r.Point1.y - r.Point2.y)) 
}

//  - Print the results to the terminal
func printInfo(rect Rectangle) {
	fmt.Println("Area is", calcArea(rect))
	fmt.Println("Perimeter is", calcPerimeter(rect))
}

func main() {
	rect := Rectangle{Point1: Coordinate{1.0, 5.5}, Point2: Coordinate{10.2, 2.0}}

	printInfo(rect)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.Point1.y *= 2
	rect.Point2.x *= 2
	//  - Print the new results to the terminal
	printInfo(rect)
}
