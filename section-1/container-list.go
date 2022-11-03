package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Muhammad")
	data.PushBack("Yusuf")
	data.PushBack("Manshur")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Front().Next().Value)
	// fmt.Println(data.Back().Value)

	// fmt.Println(data.Front().Prev())
	// fmt.Println(data.Back().Next())

	// From left to right
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// From right to left
	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}
