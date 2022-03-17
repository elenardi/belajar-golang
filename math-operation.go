package main

import "fmt"

func main() {
	var result = 20 + 30
	fmt.Println(result)

	var a = 30
	var b = 20
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i+10
	fmt.Println(i)

	i++ // i = i+1
	fmt.Println(i)

	//positive and negative value
	var positive = +100
	var negative = -100
	fmt.Println(positive)
	fmt.Println(negative)
}
