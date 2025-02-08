package main

import "fmt"

func main() {
	var nama string
	var umur int

	fmt.Print("masukan nama dan umur mu : ")
	fmt.Scan(&nama, &umur)
	fmt.Printf("nama : %s, umur %d", nama, umur)
}
