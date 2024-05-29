package main

import "fmt"

func main() {

	s1 := []int{1, 2}
	fmt.Print("Slice: ", s1)

	s1 = append(s1, 3)
	fmt.Println("\nSlice append: ", s1)

	var s2 = s1[0:1]
	fmt.Println("Slice copy: ", s2)

}
