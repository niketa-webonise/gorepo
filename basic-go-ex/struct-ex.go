package main

import "fmt"

type a struct{
	name string
	age int
	istrue bool
}

/*func (a1 a)method1() string {
if a1.age < 20 {
	return a1.name + "hello"
}else{

	return a1.name + "hi!"
}}*/


func main() {

	var1 := a{name:"Niketa",age:10}	
	validateAge(var1)
	fmt.Println("struct1:",var1)
	//fmt.Println(var1.method1())
	//var2 := a{name:"Anish",age:1}
	//fmt.Println(var2.method1())
	
}

func validateAge(a1 a){

	a1.istrue = a1.age<5
}


