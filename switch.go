package main

import "fmt"

func main() {
	angka := 3

	switch angka {
	case 1:
		fmt.Println("angka 1")
	case 2:
		fmt.Println("angka 2")
	case 3:
		fmt.Println("angka 3")
	case 4:
		fmt.Println("angka 4")
	default:
		fmt.Println("angka tidak valid")
	}

	hari := "Senin"
	switch hari {
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur!")
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari kerja.")
	default:
		fmt.Println("Hari tidak valid.")
	}

  umur := 18
	switch {
	case umur < 18:
		fmt.Println("Kamu belum cukup umur.")
	case umur >= 18 && umur <= 60:
		fmt.Println("Kamu sudah dewasa.")
	default:
		fmt.Println("Kamu sudah tua.")
	}
}
