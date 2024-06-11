package main

import (
	"fmt"
	"os"
)

// Hàm dược khai báo rõ ràng
func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

// Truyền vào slice list các giá trị
func sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

// Đặt tên tham số trả về
func find(m map[int]int, k int) (v int, status bool) {
	v, status = m[k]
	return
}

// Truyen array vao function, noi dung cua array khong thay doi
func once(x [3]int) {
	for i := range x {
		x[i] *= 2
	}
}

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

func main() {

	// Hàm ẩn danh
	var addFunc = func(a, b int) int {
		return a + b
	}

	fmt.Println(addFunc(1, 2))

	// Một hàm trong go có thể có nhiều tham số, trả về nhiều giá trị. Các tham số và return value
	// sử dụng baằng các truyền tham trị

	// Truyền slice vào hàm sum

	// unpack giá trị truyền vào
	var a = []interface{}{123, 234}
	fmt.Println(a...)
	fmt.Println(a)

	// Đặt tên tham số trả về, chỉ cần gọi return, không cần gọi giá trị cụ thể

	// defer tạm hoãn công viêệc bên trong phạm vi defer cho đến khi công việc bên ngoài hàm thực hiện xong

	defer fmt.Println("Sau")
	fmt.Println("Truoc")

	// use case defer
	// lời gọi defer push vao stack, thuc thi theo thu tu
	// close file, giai phong tai nguye

	f, err := os.Create("file.txt")
	if err != nil {
		panic("cannot create file")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Cannot close file!")
		}
	}(f)

	_, err = fmt.Fprintf(f, "hello world!")
	if err != nil {
		return
	}

	// slice trong function

	// truyen array vao function, ket qua array khong bi thay doi, do truyen gia tri
	x := [3]int{1, 2, 3}
	fmt.Println("x: ", x)
	once(x)
	fmt.Println("x: ", x)

	// truyen slice vao function, ket qua slice bi thay doi, do thuc chat truyen con tro vao function
	x1 := []int{1, 2, 3}
	fmt.Println("x1: ", x1)
	twice(x1)
	fmt.Println("x1: ", x1)

}
