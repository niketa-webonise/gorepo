package main

import "fmt"

func main() {

map_var:=make(map[string]string)

fmt.Println(map_var)

map_var["key1"]="value1"
map_var["key2"]="value2"

fmt.Println(map_var)

fmt.Println(map_var["key1"])
fmt.Println("length of map:",len(map_var))

delete(map_var,"key1")
fmt.Println(map_var)

_, present:=map_var["key2"]
fmt.Println(present)

other_var:=map[string]int{"k1" : 10 , "k2" : 20}
fmt.Println("new map",other_var)
}
