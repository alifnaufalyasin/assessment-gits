package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Masukkan Email : ")
	var email string
	fmt.Scanln(&email)

	if strings.Contains(email, "@") {
		arrEmail := strings.Split(email, "@")
		// fmt.Println(arrEmail)
		if len(arrEmail[0]) <= 20 {

		}
		if strings.Contains(arrEmail[1], ".") {
			if strings.Contains(arrEmail[1], ".co.id") || strings.Contains(arrEmail[1], ".id") {
				fmt.Println("Format Email sudah benar")
			} else {
				fmt.Println("domain tidak diperbolehkan")
			}
		} else {
			fmt.Println("Tidak terdapat titik setalah @")
		}
	} else {
		fmt.Println("Tidak terdapat @")
	}

}
