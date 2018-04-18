package main

import "fmt"


func main() {

var arr_numbers [10] int
fmt.Println(arr_numbers)
arr_numbers[9]=10
fmt.Println(arr_numbers)
fmt.Println(arr_numbers[9])
fmt.Println("length of arr_numbers",len(arr_numbers))


a:=[10]int{10,20,30}
fmt.Println(a)

var twoD[2][2]int
for i:=0;i<2;i++{
for j:=0;j<2;j++{
twoD[i][j] = i + j;
}}

fmt.Println("two dimensional array",twoD)
}
