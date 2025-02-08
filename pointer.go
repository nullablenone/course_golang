package main

import "fmt"

func main() {

	// nilai := 1
	// fmt.Println(nilai)
	
	// nilai2 := 2
	// fmt.Println(nilai2)

	// nilai3 := 3
	// fmt.Println(nilai3)

	nilai := 1
	almt_nilai := &nilai
	fmt.Println(nilai)

	*almt_nilai = 2
	fmt.Println(nilai)

	*almt_nilai = 3
	fmt.Println(nilai)





	
	


	
}