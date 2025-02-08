package main

import "fmt"

func main() {
	//array
	var angka = [5]int{1,2,3,4,5}
	slice := angka[1:5]
	fmt.Println(slice)

	// langsung
	slice_langsung := []int{1,2,3,4,5}
	fmt.Println(slice_langsung)

	// pake mathod
	slice_mathod := make([]int, 3, 5)
													// len, cap
	fmt.Println(slice_mathod)
	fmt.Println(len(slice_mathod))
	fmt.Println(cap(slice_mathod))
}
