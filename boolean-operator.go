package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80

	var lulusUjian = ujian > 70
	var lulusAbsensi = absensi >= 80

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

}
