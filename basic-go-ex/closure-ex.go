package main

import "fmt"

func Sequence() func() int {
	i:=0
	return func() int {
        i++;
return i;
}}

func main() {

generate_seq := Sequence()

fmt.Println(generate_seq());
fmt.Println(generate_seq());
fmt.Println(generate_seq());

generate_seq = Sequence()

fmt.Println(generate_seq());
}
