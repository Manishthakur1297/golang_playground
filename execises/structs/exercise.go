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

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a, b Coordinate
}

func length(rect Rectangle) int {
	return (rect.a.y - rect.b.y)
}

func breadth(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}

func area(rect Rectangle) int {
	return length(rect) * breadth(rect)
}

func perimeter(rect Rectangle) int {
	return 2 * (length(rect) + breadth(rect))
}

func printInfo(rect Rectangle) {
	fmt.Println("Area: ", area(rect))
	fmt.Println("Perimeter: ", perimeter(rect))
}

func main() {
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}

	printInfo(rect)

	rect.a.y *= 2
	rect.b.x *= 2

	printInfo(rect)
}
