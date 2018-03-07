package main

import "fmt"

func mul(a,b int)int {
return a*b;
}

func mulmul(a,b,c int)int {
return a*b*c;
}

func main() {
result:=mul(10,20);
fmt.Println(" 10*20 = ",result);

result=mulmul(10,20,30)
fmt.Println(" 10*20*30 = ",result);
}

