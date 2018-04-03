package main

import ("fmt"
		"strings")


func main() {
	
	read := "what is 1 plus 2 plus 3?"

	fmt.Println(strings.ContainsAny(read,"plus & plus"))
}