package main

import "fmt"

func main() {



x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97,9,17,
}
x1:=x[0]

for i:=0;i<len(x);i++ {
	//for j:=0;j<len(x);j++ {
	if x1>x[i] {
		x1=x[i]
	}

	
//}
}
fmt.Println("smallest no:",x1)

}