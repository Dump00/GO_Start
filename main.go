package main

import "fmt"

func main() {

	// strings

	var nameOne string = "cisco"
	var nameTwo = "golang"
	var nameThree string
	nameThree = "Hello"
	// following shorthand method can not be used outside the functions
	nameFour := "World"

	fmt.Println(nameOne)
	fmt.Println(nameTwo)
	fmt.Println(nameThree)
	fmt.Println(nameFour)

	// ints

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne)
	fmt.Println(ageTwo)
	fmt.Println(ageThree)

	// bits & memory

	var numOne int8 = 25
	var numTwo uint = 255
	fmt.Println(numOne)
	fmt.Println(numTwo)

	// float

	var scoreOne float32 = 1.34
	var scoreTwo float64 = 25787367783789.9879
	scoreThree := 1.5
	fmt.Println(scoreOne)
	fmt.Println(scoreTwo)
	fmt.Println(scoreThree)

}
