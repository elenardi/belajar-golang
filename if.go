package main

import "fmt"

func main() {
	name := "Elen"
	if name == "Elen" {
		fmt.Println("Hallo Elen")
	} else if name == "Ardi" {
		fmt.Println("Hallo Ardi")
	} else {
		fmt.Println("Kamu siapa ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("Hallo", name)
	}
}
