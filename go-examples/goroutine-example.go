package main

import "fmt"




func f1(msg string) {
	for i:=0;i<3;i++ {
           fmt.Println(msg,":",i)
	}
 
	}

func main() {

f1("simple")


go func(msg string) {
        fmt.Println(msg)
    }("complex")

go f1("golang")

fmt.Scanln()
    fmt.Println("done")

}


