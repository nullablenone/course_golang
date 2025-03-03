package main

import (
	"errors"
	"fmt"
)

// Function buat cek minuman
func cekMinuman(nama string) (string, error) {
	if nama == "air putih" {
		return "Nih, air putihnya!", nil
	} else {
		return "", errors.New("minuman tidak tersedia")
	}
}

func main() {
	// Coba minta air putih
	minuman, err := cekMinuman("air putih")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Dapat minuman:", minuman)
	}

	// Coba minta kopi (yang gak ada)
	minuman, err = cekMinuman("kopi")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Dapat minuman:", minuman)
	}
}
