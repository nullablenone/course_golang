package main

import "fmt"

func status (nama string, umur int)(string, int){
	return nama, umur
	
}


func main() {
	fmt.Println(status("muhamad yesa", 17))
}
