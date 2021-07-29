package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Masukkan Kalimat : ")
	var kalimat string
	fmt.Scan(&kalimat)
	arrChar2 := strings.Split(kalimat, "")
	kalimat = strings.ToLower(kalimat)

	arrChar := strings.Split(kalimat, "")
	match := 0
	for i := 0; i < len(arrChar); i++ {
		if arrChar[i] == arrChar[len(arrChar)-1-i] {
			match++
		}
	}

	if match == len(arrChar) {
		fmt.Print(kalimat, " di balik ")
		for i := 0; i < len(arrChar2); i++ {
			fmt.Print(arrChar2[len(arrChar2)-1-i])
		}
		fmt.Print(" => Result True (karena palindrom)")
	} else {
		fmt.Print("Result False (karena bukan palindrome )")
	}

}
