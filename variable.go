package main

import "fmt"

func main() {
	// with var
	// with type data
	var name string

	name = "Elen Ardi"
	fmt.Println(name)

	name = "Elen Ardiyanto"
	fmt.Println(name)

	//with var
	//without type data
	var friendName = "Elbow"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	//without var
	//without type data
	country := "Indonesia"
	fmt.Println(country)

	//multi variable
	var (
		firstName = "Elen"
		lastName  = "Ardiyanto"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
