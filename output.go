package main

import (
	"fmt"
)

func main() {
	var nama, alamat string = "yesa", "bogor"

	umur := 17

	// fmt.Print(nama, "\n")
	// fmt.Print(alamat)

	fmt.Printf("umur saya %d nama saya %s alamat saya %s \n", umur, nama, alamat)
	fmt.Printf("umur %d type : %T", umur, umur)
}
