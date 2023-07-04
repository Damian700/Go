package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		/* in golang we have to explicit the http protocol before the actual domain */
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://aklsdjf.com", /* se pone coma al final tambien */
	}

	for _, link := range links {
		checkLink(link)
	}
}
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err, "not runing properly...")
		os.Exit(1)
	}
	fmt.Println(link, "is up!")
}
