package main

import "fmt"

func main() {

var x[5] float64
x[0] = 100
x[1] = 200
x[2] = 300
x[3] = 400
x[4] = 500
r:=[2]int{5,10}
fmt.Println(r)

var total float64
var avg float64
for i:=0;i<5;i++ {

	total = total + x[i]
	//fmt.Println("Total:",total)
}

fmt.Println("total:",total)
avg = total / float64(len(x))
fmt.Println("average:",avg)
}