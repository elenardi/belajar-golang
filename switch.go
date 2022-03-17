package main

import "fmt"

func main() {
	name := "Elen"

	switch name {
	case "Elen":
		fmt.Println("Halo, Elen")
		fmt.Println("Elen apa kabar ?")
	case "Ardi":
		fmt.Println("Halo, Ardi")
		fmt.Println("Ardi apa kabar ?")
	default:
		fmt.Println("Loh he")
		fmt.Println("Kamu siapa ?")
	}

	// switch length := len(name); length > 4 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println(name)
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 4:
		fmt.Println("Nama lumaya panjang")
	default:
		fmt.Println(name)
	}
}
