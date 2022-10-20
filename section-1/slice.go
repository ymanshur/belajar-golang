package main

import "fmt"

func main() {
	var months = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	slice1[0] = "Diubah lagi"
	fmt.Println(months)

	// var slice2 = months[10:]
	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)

	// slice3[1] = "Bukan Desember"
	slice3[1] = "Bukan April"

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice)) // [Eko Kurniawan]
	// copySlice := make([]string, 1, cap(newSlice)) // [Eko]
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// iniArray := [5]int{1, 2, 3, 4, 5}
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
