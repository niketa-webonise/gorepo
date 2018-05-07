 package main
 import ("fmt"
 		"strconv")


 func main() {

 	input := "what is -11 plus -1 ?"

 	re := regexp.MustCompile("[0-9]+")
	operands := re.FindAllString(input, -1)


 op1, _ := strconv.FormatUnit(unit64(operands[0]),16)
 op2,_ := strconv.FormatUnit(unit64(operands[1]),16)

if strings.Contains(input,"plus"){

	result := op1+op2

	fmt.Println(result)
}
    
    }