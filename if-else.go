package main

import "fmt"

func main() {

	var umur int

	fmt.Println("berapa umur mu:")
	fmt.Scan(&umur)

	if umur >= 17 {
		fmt.Println("sudah bisa bermain game")
	} else if umur >= 15 {
		fmt.Println("bentar lagi bisa main game")
	} else {
		fmt.Println("masih kecil")
	}
}
