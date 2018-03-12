package wordy

import ("math"
"regexp"
"strings"
"strconv"

)

func Wordy(case1 string) int {

//case1 := "What is 5 plus 13?"
//case1 := "What is 53 plus 2?"

re := regexp.MustCompile("[0-9]+")
   operands := re.FindAllString(case1, -1)

op1, _ := strconv.Atoi(operands[0])
op2, _ := strconv.Atoi(operands[1])

   if len(operands) == 2 {

if strings.Contains(case1,"plus") {
evalval := op1 + op2
   	return evalval
   } else if strings.Contains(case1,"minus"){
   	minus_ans := op1 - op2
   	return minus_ans
   } else if strings.Contains(case1,"multiplied") {
   	multiply_ans := op1 * op2 
   	return multiply_ans
   } else if strings.Contains(case1,"divided") {
   	divide_ans := op1 / op2
   	return divide_ans
  
   } else if strings.Contains(case1,"power") {
   	raised_to := math.Pow(float64(op1),float64(op2))
   	return int(raised_to)
   }
} else if len(operands) == 3 {
op3, _ := strconv.Atoi(operands[2])

var final_eval int

//solve left part
if strings.Contains(case1,"plus") {
case1 = strings.Replace(case1,"plus","",1)
final_eval = op1 + op2
   } else if strings.Contains(case1,"minus"){
   	case1 = strings.Replace(case1,"minus","",1)
   	final_eval = op1-op2
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
   return final_eval
}
return 0
}
