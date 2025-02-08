package main

import "fmt"

func main() {
	// assign
	harga_1 := 1000
	harga_1 += 2000
	fmt.Println(harga_1)

	angka := true
	angka_2 := false
	// OR ||
	fmt.Println(angka || angka_2)
	// AND &&
	fmt.Println(angka && angka_2)
	// NOT !
	fmt.Println(!angka_2)

}