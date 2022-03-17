package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Elen"
	names[1] = "Ardi"
	names[2] = "Yanto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90, 80, 70,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(names))
	fmt.Println(len(values))
}
