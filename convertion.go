package main

import "fmt"

func main() {
	var nilai32 = 129
	var nilai64 int64 = int64(nilai32)

	//missing convertion
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Elen"
	var e = name[1]
	var eString = string(e)

	fmt.Println(e)
	fmt.Println(eString)
}
