package main

import "fmt"

func main() {

msg := make(chan string,2)

msg <- "hi i am niketa"
msg <- "jain"

fmt.Println(<-msg)
fmt.Println(<-msg)

}

