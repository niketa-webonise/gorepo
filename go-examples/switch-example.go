package main

import "fmt"
import "time"

func main() {
i:=2
switch i{
case 1:
       fmt.Println("print one")
case 2:
	fmt.Println("print two")
case 3:
	fmt.Println("print three");
}
t1:=time.Now()
fmt.Println(t1)
t:=time.Now().Weekday()
fmt.Println(t)
switch t{
case time.Sunday,time.Saturday:
			fmt.Println("weekends")
default:
	fmt.Println("weekdays")
}
}
