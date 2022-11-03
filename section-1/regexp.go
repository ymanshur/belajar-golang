package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z]o)")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eDo"))

	// search := regex.FindAllString("eko eka, edo eto eyo", 1)
	search := regex.FindAllString("eko eka, edo eto eyo", -1)
	fmt.Println(search)
}
