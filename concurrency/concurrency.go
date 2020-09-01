package main

import (
	"fmt"
	"net/http"
)

// Concurrency example with go routines and channels
func main() {
	//channel and how to initialize
	channel := make(chan string)
	servers := []string{
		"http://pedidosya.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		// go routine
		go checkServer(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		// how to read the channels value
		fmt.Println(<-channel)
	}
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		//fmt.Println(server, "server not found :(")
		//how to set channel's value
		channel <- server + " server not found!!"
	} else {
		//fmt.Println(server, "works!!! yay!!!")
		channel <- server + " works!!"
	}
}
