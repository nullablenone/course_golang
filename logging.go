package main

import (
	"errors"
	"log"
)

func cekMinuman(nama string) (string, error) {
	if nama == "air putih" {
		return "Nih, air putihnya!", nil
	} else {
		return "", errors.New("minuman tidak tersedia")
	}
}

func main() {
	log.Println("INFO : aplikasi di mulai")
	log.Println("DEBUG : memulai proses")
	minuman, err := cekMinuman("air putih")
	if err != nil {
		log.Printf("Error: %t", err)
	} else {
		log.Println("INFO : Dapat minuman:", minuman)
	}
}
