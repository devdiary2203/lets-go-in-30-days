package main

import "fmt"

// map
func main() {
	// init by make function
	var exMap1 = make(map[string]int)
	fmt.Print("Map: ", exMap1)
	if exMap1 == nil {
		fmt.Println("ExMap1 nil")
	} else {
		fmt.Println("ExMap1 not nil")
	}

	// init declare type
	var exMap2 map[string]int
	fmt.Print("Map2: ", exMap2)
	if exMap2 == nil {
		fmt.Println("ExMap2 nil")
	} else {
		fmt.Println("ExMap2 not nil")
	}

	// init declare specific value
	exMap3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("ExMap3: ", exMap3)
	exMap3["c"] = 3
	fmt.Println("ExMap3: ", exMap3)

}
