package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	var p Person
	p.lastName = "robert"
	p.firstName = "mccall"
	fmt.Println("Person: ", p)

	p1 := Person{
		firstName: "robert",
		lastName:  "mccall",
		age:       26,
	}
	fmt.Println("Person 1: ", p1)

	p2 := Person{"robert", "mccall", 26}
	fmt.Println("Person 2: ", p2)

	p3 := Person{firstName: "robert"}
	fmt.Println("Person 3: ", p3)

}
