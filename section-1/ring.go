package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)
	// data.Value = "Yusuf"

	// var data2 = data.Next()
	// data2.Value = "Manshur"

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// fmt.Println(*data)

	data.Do(func(a any) {
		fmt.Println(a)
	})

}
