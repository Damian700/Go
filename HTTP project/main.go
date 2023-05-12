package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	/* 	bs := make([]byte, 99999) //you can put"n" number and its okay
	   	resp.Body.Read(bs)
	   	fmt.Println(string(bs)) */

	lw := myWriter{}

	io.Copy(lw, resp.Body) //first argument, has to be something that implements the writer interface. (File type, which has a function caled Write ([]byte) (int, err))
	//second argument, something that implements the reader interface

}

func (myWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
