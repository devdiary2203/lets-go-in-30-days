package main

import "fmt"

func main() {
	var number int
	number = 1

	fmt.Print("Number \n", number)

	var number1 = 1

	fmt.Print("\nNumber 1: ", number1)

	var (
		name string // default '' empty string, call by zero values
		age  int    // default=0
	)
	name = "test"
	age = 1

	fmt.Print("\nName: ", name)
	fmt.Print("\nAge: ", age)

	var isTest bool // zero values is false
	isTest = false
	fmt.Print("\tIs test: ", isTest)

	var runeTest rune
	runeTest = 'á'

	fmt.Printf("Rune test: %T", runeTest)

	var a float32
	var b int
	a = 2.1
	b = 1
	fmt.Print("\nresult: ", a+float32(b))

}
