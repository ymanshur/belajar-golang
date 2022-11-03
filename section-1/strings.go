package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Yusuf Manshur", "Yusuf"))
	fmt.Println(strings.Contains("Yusuf Manshur", "Eko"))

	fmt.Println(strings.Split("Yusuf Manshur", ""))

	fmt.Println(strings.ToLower("Yusuf Manshur"))
	fmt.Println(strings.ToUpper("Yusuf Manshur"))
	fmt.Println(strings.ToTitle("yusuf manshur"))

	fmt.Println(strings.Trim("        Yusuf Manshur        ", " "))
	fmt.Println(strings.Trim("        Yusuf Manshur       a", " "))

	fmt.Println(strings.ReplaceAll("Yusuf Yusuf Manshur", "Yusuf", "Muhammad"))
	fmt.Println(strings.Replace("Yusuf Yusuf Manshur", "Yusuf", "Muhammad", 1))
}
