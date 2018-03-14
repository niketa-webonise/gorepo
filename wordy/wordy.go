package wordy

import ("math"
"regexp"
"strings"
"strconv"

)

func Wordy(input string) int {



re := regexp.MustCompile("[0-9]+")
operands := re.FindAllString(input, -1)

op1, _ := strconv.Atoi(operands[0])
op2, _ := strconv.Atoi(operands[1])


   if len(operands) == 2 {

      if strings.Contains(input,"plus") {

      evalval := op1 + op2
   	return evalval

   } else if strings.Contains(input,"minus"){

   	minus_ans := op1 - op2
   	return minus_ans

   } else if strings.Contains(input,"multiplied") {
   	
      multiply_ans := op1 * op2 
   	return multiply_ans

   } else if strings.Contains(input,"divided") {
   	

      divide_ans := op1 / op2
   	return divide_ans

  
   } else if strings.Contains(input,"power") {
   	

      raised_to := math.Pow(float64(op1),float64(op2))
   	return int(raised_to)


   }
} else if len(operands) == 3 {

   op3, _ := strconv.Atoi(operands[2])

   var final_eval int


   if strings.Contains(input,"plus") {

         input = strings.Replace(input,"plus","",1)
         final_eval = op1 + op2


   } else if strings.Contains(input,"minus"){


   	input = strings.Replace(input,"minus","",1)
   	final_eval = op1-op2


   } else if strings.Contains(input,"multiplied") {
   	

      input = strings.Replace(input,"multiplied","",1)
   	final_eval = op1 * op2 


   } else if strings.Contains(input,"divided") {
   	

      input = strings.Replace(input,"divided","",1)
   	final_eval = op1 / op2


   }
   
   if strings.Contains(input,"plus") {
      
      final_eval = final_eval + op3

   } else if strings.Contains(input,"minus"){
   	
      final_eval = final_eval - op3

   } else if strings.Contains(input,"multiplied") {
   	
      final_eval = final_eval * op3 

   } else if strings.Contains(input,"divided") {
   	
      final_eval = final_eval / op3
   }

   return final_eval
}

return 0

}
