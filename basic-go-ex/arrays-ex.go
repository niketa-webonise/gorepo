package main

import "fmt"

func main() {

x := [...]int{10,20,30}
xsl := x[0:2]
fmt.Println("before modification",x)

for i := range xsl{
	xsl[i]++
}
fmt.Println("after modification",x)

//fmt.Println("x::",x)
//var b[5] int
//fmt.Println("value of b",b)

y :=  x
y[0]=500
fmt.Println("X is",x)
fmt.Println("Y is",y)

var x1[5]string
x1[0]="abc"
x1[1]="def"
x1[2]="gfh"
fmt.Println("sring arrays:",x1)

sl:=x1[0:3]
fmt.Println("slice1:",sl)

map1 :=map[string]int{"thousands":1000}
fmt.Println("map1:",map1)

fmt.Println(map1["thousands"])

fmt.Println(x)
}