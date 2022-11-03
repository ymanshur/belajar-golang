package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	yusuf := Man{"Yusuf"}
	yusuf.Married()

	fmt.Println(yusuf.Name)
}
