package main

import "fmt"

func main() {
	// -- create empty key and value map --
	// var colors := make(map[string]string) 
	colors := map[string]string{
		"red" : "#ff0000",
		"green" : "#4bf754",
		"white" : "#ffffff",
	}

	colors["white"] = "#ffffff"
	// delete(colors,10) //delete key "10" and its value 

	printMap(colors)
}

func printMap(c map[string]string) {
	for color,hex := range c {
		fmt.Println(color + "-" + hex)
	}
}