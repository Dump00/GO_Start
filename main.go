package main

import "fmt"

func main() {

	name := "cisco"
	age := 30

	// print
	fmt.Print("Hello")
	fmt.Print("World \n")
	fmt.Print("newLine \n")

	fmt.Println("Golang")

	fmt.Println("My name is", name, "and my age is", age)

	// %_ format specifier

	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %v and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f \n", 255.55)
	fmt.Printf("you scored %0.1f \n", 255.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)
}
