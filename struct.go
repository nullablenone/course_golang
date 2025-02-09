package main

import "fmt"

type Manusia struct {
	Nama string
	Umur int
	Jk   string
}

func main() {
	var yesa Manusia
	yesa.Nama = "Muhamad Yesa"
	yesa.Umur = 17
	yesa.Jk = "Pria"

	fmt.Println(yesa)
	fmt.Println(yesa.Umur)

	nullablenone := Manusia{
		Nama: "nullablenone",
		Umur: 17,
		Jk:   "Laki",
	}
	fmt.Println(nullablenone)
	fmt.Println(nullablenone.Umur)
	fmt.Printf("Type : %T", nullablenone.Umur)
}
