package main

import "fmt"

func main() {
	user := map[string]string{
		"nama": "yesa",
		"alamat": "warna",
	}

	fmt.Println(user)
	user["nama"] = "yesnot"
	delete(user, "alamat")

	fmt.Println(user)
}