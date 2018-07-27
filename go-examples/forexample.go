package main

import "fmt"

func main() {
fmt.Println("outer loop")
i := 1
for i <= 3 {
fmt.Println(i)
i=i+1
}

fmt.Println("inner loop")
for j:=1;j<=5;j++ {
fmt.Println(j)
}

for{
fmt.Println("infinite loop")
break
}




}
