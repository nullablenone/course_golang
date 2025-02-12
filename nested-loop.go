package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ { // Loop luar

		for j := 1; j <= 3; j++ { // Loop dalam
			fmt.Printf("%d ", j)
		}
		fmt.Println() // Ganti baris setelah satu loop dalam selesai
		
	}

}
