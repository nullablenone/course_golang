package main

import "fmt"

func main() {
	siswa := []string{"yesa", "dimas", "perdi", "yusuf"}

	// inisialisasi; kondisi; increment/decrement
	for i := 0; i < len(siswa); i++ {
		fmt.Println(siswa[i])
	}
}
