package main

import "fmt"

func rec(n int)int{
if n==0 {
return 1
}
return n*rec(n-1)
}

func main(){
result:=rec(5)
fmt.Println(result)
}
