package main

import "fmt"

func main() {

	siswa := []string{"yesa", "dimas", "agung", "sasuke", "naruto"}

	for key, value := range siswa {
		fmt.Printf("type key : %T, type value : %T \n", key, value)
	}
}
