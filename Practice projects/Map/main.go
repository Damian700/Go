package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		//all keys are type string and so are the values
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	// detail the type of key and value respectively
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
