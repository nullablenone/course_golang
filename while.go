package main

import "fmt"

func main() {
	siswa := []string{"yesa", "dimas", "perdi", "yusuf"}

	i := 0
	// kondisi
	for i < len(siswa){
		fmt.Println(siswa[i])
		i++
	}
}
