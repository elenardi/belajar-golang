package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Elen",
		"address": "Purbalingga",
	}
	person["tittle"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["tittle"])

	var book map[string]string = make(map[string]string)
	book["tittle"] = "Belajar Go-Lang"
	book["auth"] = "Elen Ardiyanto"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
