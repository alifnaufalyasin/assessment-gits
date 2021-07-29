package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Masukkan Kalimat (xxx di balik xxx): ")
	var kalimat1, kalimat2, kalimat3, kalimat4 string
	fmt.Scan(&kalimat1)
	fmt.Scan(&kalimat2)
	fmt.Scan(&kalimat3)
	fmt.Scan(&kalimat4)

	// fmt.Println(kalimat1, kalimat2, kalimat3, kalimat4)
	arrChar1 := strings.Split(kalimat1, "")
	arrChar4 := strings.Split(kalimat4, "")
	var match int = 0
	for i := 0; i < len(arrChar1); i++ {
		if arrChar1[i] == arrChar4[len(arrChar4)-1-i] {
			match++
		}
		// fmt.Println(arrChar1[i], arrChar4[len(arrChar4)-1-i])
	}
	if match == len(arrChar1) {
		fmt.Println("Result True (karena palindrom)")
	} else {
		fmt.Println("Result False (karena bukan palindrome)")
	}

}
