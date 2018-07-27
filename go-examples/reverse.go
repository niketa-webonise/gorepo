package main

import "fmt"

func main() {
	
	//var rev string
	input := "i am niketa"

	fmt.Println(Reverse(input))

	

	}

	//fmt.Println("reverse string:",rev)

	func Reverse(s string) string {
    var reverse string
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}

