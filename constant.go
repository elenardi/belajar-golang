package main

import "fmt"

func main() {
	const value = 100

	//multi constant
	const (
		firstName string = "Elen"
		lastName         = "Ardiyanto"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
