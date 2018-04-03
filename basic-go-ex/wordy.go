package wordy

import ("fmt"
		"regexp"
		"strings"
		"strconv")

func Wordy(case1 string) int {

	//ase1 := "What is 5 plus 13?"
    
	re := regexp.MustCompile("[0-9]+")
	operands := re.FindAllString(case1, -1)

	op1,_ := strconv.Atoi(operands[0])
	op2,_ := strconv.Atoi(operands[1])

	if len(operands) == 2 {

	if strings.Contains(case1,"plus") {
	eval := op1 + op2
	   	//fmt.Println(evalval)
	 }
	 	}
return eval
}