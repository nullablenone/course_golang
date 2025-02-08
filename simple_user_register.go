package main

import "fmt"

func main() {
	var name, password string
	fmt.Print("register : ")
	fmt.Scanln(&name)
	fmt.Print("password : ")
	fmt.Scanln(&password)
	fmt.Printf("username : %s \n", name)
	fmt.Printf("password  yzyzyz%sshshsh \n", password)
}
