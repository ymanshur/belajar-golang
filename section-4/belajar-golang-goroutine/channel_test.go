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

func GiveMeResponse(channel chan string) {
	channel <- "Yusuf Manshur"

	time.Sleep(2 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	channel <- "Yusuf Manshur"
	// invalid operation
	// data := <-channel
	time.Sleep(2 * time.Second)
}

func OnlyOut(channel <-chan string) {
	// invalid operation
	// channel <- "Yusuf Manshur"
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
