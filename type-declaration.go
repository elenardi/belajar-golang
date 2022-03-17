package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var myNIK NoKTP = "1246154278917"
	var marriedStatus Married = false

	fmt.Println("NIK = ", myNIK)
	fmt.Println("Married = ", marriedStatus)
}
