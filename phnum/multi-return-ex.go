package main

import "fmt"

func multi() (int,int,int) {
	return 10,20,30
}


func main() {
result1,result2,result3:=multi()
fmt.Println("value1:",result1)
fmt.Println("value2:",result2)
fmt.Println("value3:",result3)

_,_, result4:=multi()
fmt.Println("value3=value4:",result4)
}
