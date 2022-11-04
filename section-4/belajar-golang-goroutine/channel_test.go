package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// channel <- "Yusuf"

	// var data string
	// data = <- channel
	// data := <-channel
	// fmt.Println(data)
	// fmt.Println(<-channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Yusuf Manshur"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
