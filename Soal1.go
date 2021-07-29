package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Masukkan Angka : ")
	var angka int
	fmt.Scanln(&angka)

	// Mod dapat menggunakan math.Mod atau %
	if math.Mod(float64(angka), 3) == 0 && math.Mod(float64(angka), 5) == 0 {
		fmt.Println("Hello World")
	} else if math.Mod(float64(angka), 5) == 0 {
		fmt.Println("World")
	} else if math.Mod(float64(angka), 3) == 0 {
		fmt.Println("Hello")
	}
}
