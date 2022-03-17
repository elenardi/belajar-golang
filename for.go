package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke-", counter2)
	}

	slice := []string{"Elen", "Ardi", "Yanto"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Elen"
	person["tittle"] = "Programmer"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
