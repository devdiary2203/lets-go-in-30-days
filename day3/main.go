package main

import "fmt"

func main() {

	var arr [4]int
	arr[0] = 1
	fmt.Print(arr)

	arr1 := [3]int{1, 2, 3}
	fmt.Print(arr1[2])

	arr2 := [4]string{"at", "b", "c"}
	fmt.Print(arr2)

	fmt.Println(len(arr2))

	arr3 := [...]string{"test", "test1"}
	fmt.Println(arr3)
	fmt.Println(len(arr3))

	arrRoot := [...]int{1, 2, 3} // array is value type
	arrCopy := arrRoot
	arrCopy[0] = 0
	fmt.Println(arrRoot)
	fmt.Println(arrCopy)

	for i := 0; i < len(arrRoot); i++ {
		fmt.Println(arrRoot[i])
	}

	for i, v := range arrRoot { // _ if not used
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	matrix := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
