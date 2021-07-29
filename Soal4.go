package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Masukkan Kalimat : ")
	var kalimat string
	fmt.Scanln(&kalimat)
	arrChar := strings.Split(kalimat, "")

	for i := len(arrChar); i > 0; i-- {
		fmt.Print(arrChar[i-1])
	}

}
