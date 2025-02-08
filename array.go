package main

import (
	"fmt"
)

func main() {
	angka := [4]int{1, 2, 3, 4}
	// fmt.Println(angka)
	// fmt.Println(angka[0])
	// fmt.Println(angka[3])
	// fmt.Println(len(angka))
	fmt.Println(angka[len(angka)-1])

	angka[0] = 9
	fmt.Println(angka)

	number := [5]int{1: 89, 3: 88}

	fmt.Println(number)

}
