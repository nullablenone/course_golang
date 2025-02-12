package main

import "fmt"

func main() {

	n := 5 // Tinggi segitiga

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
