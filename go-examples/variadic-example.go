package main

import "fmt"

func sum(nums...int){

total := 0

for _,num := range nums {
total=total+num
}
fmt.Println(total)

}

func main() {
sum(10,20,30)
sum(10,20)
}
