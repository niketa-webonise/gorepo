package main

import("fmt"
		"./schedular")

func f1(){

		fmt.Println("execution of f1()")
		a := 10
		b := 20
		fmt.Println("operation -> Addition")
		fmt.Println("Addition of two numbers",a+b)
		fmt.Println()
}

func f2(){

		fmt.Println("execution of f2()")
		a := 10
		b := 20
		fmt.Println("operation -> Multiplication")
		fmt.Println("Multiplication of two numbers",a*b)
		fmt.Println()
}

func f3(){

	fmt.Println("execution of f3()")
	a := 20
	b := 10
	fmt.Println("operation -> Subtraction")
	fmt.Println("Subtraction of two numbers",a-b)
	fmt.Println()
}


func main() {

	 go schedular.ExecuteAtIntervals(1,f1)
	 go schedular.ExecuteAtIntervals(2,f2)
	 go schedular.ExecuteAtIntervals(3,f3)

	 for{}
	
}

