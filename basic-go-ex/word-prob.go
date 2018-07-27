package main

import("strings"
"fmt"
"regexp"
"strconv"
"math")

func main() {

//case1 := "what is 5 plus 13?"
//case1 := "What is 7 minus 5?"
//case1 := "What is 6 multiplied by 4?"
//case1 := "What is 25 divided by 5?"
//case1 := "What is 5 plus 13 multiplied 6?"
  // case1 := "What is 2 raised to the 5th power?"
   case1 :="what is 15 plus 1 plus 1?"

re := regexp.MustCompile("[0-9]+")
   operands := re.FindAllString(case1, -1)

op1, _ := strconv.Atoi(operands[0])
op2, _ := strconv.Atoi(operands[1])

   if len(operands) == 2 {

if strings.Contains(case1,"plus") {
evalval := op1 + op2
   	fmt.Println(evalval)
   } else if strings.Contains(case1,"minus"){
   	minus_ans := op1 - op2
   	fmt.Println(minus_ans)
   } else if strings.Contains(case1,"multiplied") {
   	multiply_ans := op1 * op2 
   	fmt.Println(multiply_ans)
   } else if strings.Contains(case1,"divided") {
   	divide_ans := op1 / op2
   	fmt.Println(divide_ans)
  
   } else if strings.Contains(case1,"power") {
   	raised_to := math.Pow(float64(op1),float64(op2))
   	fmt.Println(raised_to)
   }
} else if len(operands) == 3 {
op3, _ := strconv.Atoi(operands[2])

var final_eval int

//solve left part
if strings.Contains(case1,"plus  plus") {
   final_eval = op1+op2+op3
 fmt.Println(final_eval)
}
}
}
case1 = strings.Replace(case1,"plus","",1)

final_eval = op1 + op2
   } else if strings.Contains(case1,"minus"){
   	case1 = strings.Replace(case1,"minus","",1)
   	final_eval = op1 - op2
   } else if strings.Contains(case1,"multiplied") {
   	case1 = strings.Replace(case1,"multiplied","",1)
   	final_eval = op1 * op2 
   } else if strings.Contains(case1,"divided") {
   	case1 = strings.Replace(case1,"divided","",1)
   	final_eval = op1 / op2
   }
   
   if strings.Contains(case1,"plus") {
final_eval = final_eval + op3
   } else if strings.Contains(case1,"minus"){
   	final_eval = final_eval - op3
   } else if strings.Contains(case1,"multiplied") {
   	final_eval = final_eval * op3 
   } else if strings.Contains(case1,"divided") {
   	final_eval = final_eval / op3
   }
   //fmt.Println(final_eval)
}
 