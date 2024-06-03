package main

import "fmt"

// variadic
func addItem(item int, list ...int) {
	list = append(list, item)
	fmt.Print(list)
}

func change(list ...int) {
	list[0] = 999
}

func main() {

	addItem(6, 1, 2, 3, 4, 5)

	var list = []int{1, 2, 3, 4, 5}
	addItem(100, list...)

	change(list...)
	fmt.Print("List after change: ", list)
}
