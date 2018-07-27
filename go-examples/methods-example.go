package main

import "fmt"

type rect struct{
	length,breadth int
}

func(r *rect)area() int {
  return r.length*r.breadth
}

func(r rect)perimeter() int {
   return 2*(r.length+r.breadth)
}


func main() {

var1 := rect{length:10,breadth:20}
fmt.Println("area=10*20 : ",var1.area())
fmt.Println("perimeter=2*(10+20) : ",var1.perimeter())

var2 := &var1
fmt.Println("area=",var2.area())
fmt.Println("perimeter=",var2.perimeter())
}
