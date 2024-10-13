package main

import "fmt"

//ISPISIVANJE ZADATAKA
func main() {
	prvi()
	drugi()

}

/*
1. Declare two integers, `firstNumber` and `secondNumber`, assign values 20 and 30 to them.
Swap values of `firstNumber` and `secondNumber` without using third variable
After all, print values of `firstNumber` and `secondNumber`.*/
func prvi() {
	var firstNumber int = 20
	var secondNumber int = 30

	firstNumber, secondNumber = secondNumber, firstNumber

	fmt.Println("Zamijenjeni brojevi su:", "prvi broj-->", firstNumber, "drugi broj-->", secondNumber)
}

/*
2. Declare two variables, `firstName` and `lastName` assign them with wanted values.
Declare constant named `fullname`
Combine constant and both strings into a full name by concatenating strings with a space in between and print them out.*/
func drugi() {
	var firstName string = "Sara"
	var secondName string = "Budimir"

	const fullName string = "Full name: "

	var full string = fullName + firstName + " " + secondName

	fmt.Println(full)
}
