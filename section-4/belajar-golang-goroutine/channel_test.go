package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	// channel := make(chan string)
	// channel := make(chan string, 1)
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Yusuf"
		channel <- "Manshur"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		// fmt.Println(<-channel)
	}()

	fmt.Println("Done")
	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	// defer close(channel)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		// program can stuck in here, if close called using defer
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Done")
}
