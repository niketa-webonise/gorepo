package main

import "fmt"

func main() {

msg:=make(chan string)

go func() {msg<-"new message"}()

messages:=<-msg
fmt.Println(messages)

}
