package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Masukkan Waktu (HH:MM:SS AM/PM) : ")
	var waktu, ampm string
	fmt.Scan(&waktu)
	fmt.Scan(&ampm)
	// fmt.Println(waktu)

	// pisah := strings.Split(waktu, " ")
	// ampm := pisah[1]
	jam := waktu[0:5]

	if ampm == "AM" {
		if jam[0:2] == "12" {
			jam = "00" + jam[2:]
			fmt.Println(jam)
		} else {
			hour := jam[0:2]
			H, err := strconv.Atoi(hour)
			if err != nil {
				// handle error
				fmt.Println(err)
			}
			H = H + 12

			jam := strconv.Itoa(H) + jam[2:]
			fmt.Println(jam)
		}

	} else {
		fmt.Println(jam)
	}
	// fmt.Println(jam, ampm)

}
