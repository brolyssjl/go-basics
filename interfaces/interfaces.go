package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct {}

// implementing the Write function from Writer interface
func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {
	// http get
	response, err := http.Get("http://www.google.com")

	//how to handle the error
	if err != nil {
		fmt.Println("estamos jodidos con google")
	}
	//fmt.Println(response.Body)
	// webWriter struct
	e := webWriter{}
	io.Copy(e, response.Body)
}
